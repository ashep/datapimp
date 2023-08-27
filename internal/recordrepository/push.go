package recordrepository

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ashep/go-apperrors"

	"github.com/ashep/ujds/internal/model"
)

//nolint:cyclop // that's ok
func (r *Repository) Push(ctx context.Context, indexID uint, schema []byte, records []model.Record) error {
	var err error

	if indexID == 0 {
		return apperrors.InvalidArgError{Subj: "index id", Reason: "must not be zero"}
	}

	if len(records) == 0 {
		return apperrors.InvalidArgError{Subj: "records", Reason: "must not be empty"}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("db begin failed: %w", err)
	}

	qGetRecord, err := tx.PrepareContext(ctx, `SELECT log_id FROM record WHERE checksum=$1`)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db prepare failed: %w", err)
	}

	defer func() {
		if err := qGetRecord.Close(); err != nil {
			r.l.Error().Err(err).Msg("prepared statement close failed")
		}
	}()

	qInsertLog, err := tx.PrepareContext(ctx, `INSERT INTO record_log (index_id, record_id, data)
		VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db prepare failed: %w", err)
	}

	defer func() {
		if err := qInsertLog.Close(); err != nil {
			r.l.Error().Err(err).Msg("prepared statement close failed")
		}
	}()

	qInsertRecord, err := tx.PrepareContext(ctx, `INSERT INTO record (id, index_id, log_id, checksum)
VALUES ($1, $2, $3, $4) ON CONFLICT (id, index_id) DO UPDATE SET log_id=$3, checksum=$4, updated_at=now()`)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db prepare failed: %w", err)
	}

	defer func() {
		if err := qInsertRecord.Close(); err != nil {
			r.l.Error().Err(err).Msg("prepared statement close failed")
		}
	}()

	for i, rec := range records {
		if err := r.insertRecord(ctx, qGetRecord, qInsertLog, qInsertRecord, indexID, schema, i, rec); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("db commit failed: %w", err)
	}

	return nil
}

func (r *Repository) insertRecord(
	ctx context.Context,
	qGetRecord, qInsertLog, qInsertRecord *sql.Stmt,
	indexID uint,
	schema []byte,
	i int,
	rec model.Record,
) error {
	if rec.ID == "" {
		return apperrors.InvalidArgError{Subj: fmt.Sprintf("record (%d) id", i), Reason: "must not be empty"}
	}

	if rec.Data == "" {
		return apperrors.InvalidArgError{Subj: fmt.Sprintf("record (%d) data", i), Reason: "must not be empty"}
	}

	// Validate data against schema
	recDataB := []byte(rec.Data)
	if err := model.ValidateJSON(schema, recDataB); err != nil {
		return apperrors.InvalidArgError{Subj: fmt.Sprintf("record data (%d)", i), Reason: err.Error()}
	}

	logID := uint64(0)

	sumSrc := append(recDataB, []byte(rec.Index)...) //nolint:gocritic // it's ok
	sumSrc = append(sumSrc, []byte(rec.ID)...)
	sum := sha256.Sum256(sumSrc)

	// Check if we already have such data recorded as latest version
	row := qGetRecord.QueryRowContext(ctx, sum[:])
	if err := row.Scan(&logID); errors.Is(err, sql.ErrNoRows) { //nolint:revive // this is intentionally empty block
		// Ok, continue to insert
	} else if err != nil {
		return fmt.Errorf("db scan failed: %w", err)
	} else {
		// A record with the same data found, skip it
		return nil
	}

	row = qInsertLog.QueryRowContext(ctx, indexID, rec.ID, rec.Data)
	if err := row.Scan(&logID); err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		return fmt.Errorf("db query failed: %w", err)
	}

	if _, err := qInsertRecord.ExecContext(ctx, rec.ID, indexID, logID, sum[:]); err != nil {
		return fmt.Errorf("db query failed: %w", err)
	}

	return nil
}

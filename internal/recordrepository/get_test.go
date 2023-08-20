package recordrepository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ashep/go-apperrors"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/ujds/internal/recordrepository"
)

func TestRepository_Get(tt *testing.T) {
	tt.Parallel()

	tt.Run("EmptyIndexName", func(t *testing.T) {
		t.Parallel()

		db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)

		repo := recordrepository.New(db, zerolog.Nop())

		_, err = repo.Get(context.Background(), "", "theID")
		require.EqualError(t, err, "invalid index name: must not be empty")
	})

	tt.Run("DbNoRows", func(t *testing.T) {
		t.Parallel()

		db, dbm, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)

		dbm.
			ExpectQuery(`SELECT r.log_id, l.data, r.created_at, r.updated_at FROM record r LEFT JOIN record_log l ON r.log_id = l.id LEFT JOIN index i ON r.index_id = i.id WHERE i.name=$1 AND r.id=$2 ORDER BY l.created_at DESC LIMIT 1`).
			WithArgs("theIndex", "theID").
			WillReturnRows(sqlmock.NewRows([]string{}))

		repo := recordrepository.New(db, zerolog.Nop())

		_, err = repo.Get(context.Background(), "theIndex", "theID")
		require.ErrorIs(t, err, apperrors.NotFoundError{Subj: "record"})
	})

	tt.Run("DbRowScanError", func(t *testing.T) {
		t.Parallel()

		db, dbm, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)

		dbm.
			ExpectQuery(`SELECT r.log_id, l.data, r.created_at, r.updated_at FROM record r LEFT JOIN record_log l ON r.log_id = l.id LEFT JOIN index i ON r.index_id = i.id WHERE i.name=$1 AND r.id=$2 ORDER BY l.created_at DESC LIMIT 1`).
			WithArgs("theIndex", "theID").
			WillReturnError(errors.New("theSQLError"))

		repo := recordrepository.New(db, zerolog.Nop())

		_, err = repo.Get(context.Background(), "theIndex", "theID")
		require.EqualError(t, err, "db scan failed: theSQLError")
	})

	tt.Run("Ok", func(t *testing.T) {
		t.Parallel()

		db, dbm, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)

		rows := sqlmock.NewRows([]string{"r.log_id", "l.data", "r.created_at", "r.updated_at"})
		rows.AddRow(uint64(123), "theData", time.Unix(234, 345), time.Unix(456, 678))
		dbm.
			ExpectQuery(`SELECT r.log_id, l.data, r.created_at, r.updated_at FROM record r LEFT JOIN record_log l ON r.log_id = l.id LEFT JOIN index i ON r.index_id = i.id WHERE i.name=$1 AND r.id=$2 ORDER BY l.created_at DESC LIMIT 1`).
			WithArgs("theIndex", "theID").
			WillReturnRows(rows)

		repo := recordrepository.New(db, zerolog.Nop())

		rec, err := repo.Get(context.Background(), "theIndex", "theID")
		require.NoError(t, err)

		assert.Equal(t, "theID", rec.ID)
		assert.Equal(t, "theIndex", rec.Index)
		assert.Equal(t, uint64(123), rec.Rev)
		assert.Equal(t, "theData", rec.Data)
		assert.Equal(t, time.Unix(234, 345), rec.CreatedAt)
		assert.Equal(t, time.Unix(456, 678), rec.UpdatedAt)
	})
}

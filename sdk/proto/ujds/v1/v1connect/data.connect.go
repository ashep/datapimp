// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/ujds/v1/data.proto

package v1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/ashep/ujds/sdk/proto/ujds/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// DataServiceName is the fully-qualified name of the DataService service.
	DataServiceName = "ujds.v1.DataService"
)

// DataServiceClient is a client for the ujds.v1.DataService service.
type DataServiceClient interface {
	PushSchema(context.Context, *connect_go.Request[v1.PushSchemaRequest]) (*connect_go.Response[v1.PushSchemaResponse], error)
	GetSchema(context.Context, *connect_go.Request[v1.GetSchemaRequest]) (*connect_go.Response[v1.GetSchemaResponse], error)
	PushRecords(context.Context, *connect_go.Request[v1.PushRecordsRequest]) (*connect_go.Response[v1.PushRecordsResponse], error)
	GetRecords(context.Context, *connect_go.Request[v1.GetRecordsRequest]) (*connect_go.Response[v1.GetRecordsResponse], error)
}

// NewDataServiceClient constructs a client for the ujds.v1.DataService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDataServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DataServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dataServiceClient{
		pushSchema: connect_go.NewClient[v1.PushSchemaRequest, v1.PushSchemaResponse](
			httpClient,
			baseURL+"/ujds.v1.DataService/PushSchema",
			opts...,
		),
		getSchema: connect_go.NewClient[v1.GetSchemaRequest, v1.GetSchemaResponse](
			httpClient,
			baseURL+"/ujds.v1.DataService/GetSchema",
			opts...,
		),
		pushRecords: connect_go.NewClient[v1.PushRecordsRequest, v1.PushRecordsResponse](
			httpClient,
			baseURL+"/ujds.v1.DataService/PushRecords",
			opts...,
		),
		getRecords: connect_go.NewClient[v1.GetRecordsRequest, v1.GetRecordsResponse](
			httpClient,
			baseURL+"/ujds.v1.DataService/GetRecords",
			opts...,
		),
	}
}

// dataServiceClient implements DataServiceClient.
type dataServiceClient struct {
	pushSchema  *connect_go.Client[v1.PushSchemaRequest, v1.PushSchemaResponse]
	getSchema   *connect_go.Client[v1.GetSchemaRequest, v1.GetSchemaResponse]
	pushRecords *connect_go.Client[v1.PushRecordsRequest, v1.PushRecordsResponse]
	getRecords  *connect_go.Client[v1.GetRecordsRequest, v1.GetRecordsResponse]
}

// PushSchema calls ujds.v1.DataService.PushSchema.
func (c *dataServiceClient) PushSchema(ctx context.Context, req *connect_go.Request[v1.PushSchemaRequest]) (*connect_go.Response[v1.PushSchemaResponse], error) {
	return c.pushSchema.CallUnary(ctx, req)
}

// GetSchema calls ujds.v1.DataService.GetSchema.
func (c *dataServiceClient) GetSchema(ctx context.Context, req *connect_go.Request[v1.GetSchemaRequest]) (*connect_go.Response[v1.GetSchemaResponse], error) {
	return c.getSchema.CallUnary(ctx, req)
}

// PushRecords calls ujds.v1.DataService.PushRecords.
func (c *dataServiceClient) PushRecords(ctx context.Context, req *connect_go.Request[v1.PushRecordsRequest]) (*connect_go.Response[v1.PushRecordsResponse], error) {
	return c.pushRecords.CallUnary(ctx, req)
}

// GetRecords calls ujds.v1.DataService.GetRecords.
func (c *dataServiceClient) GetRecords(ctx context.Context, req *connect_go.Request[v1.GetRecordsRequest]) (*connect_go.Response[v1.GetRecordsResponse], error) {
	return c.getRecords.CallUnary(ctx, req)
}

// DataServiceHandler is an implementation of the ujds.v1.DataService service.
type DataServiceHandler interface {
	PushSchema(context.Context, *connect_go.Request[v1.PushSchemaRequest]) (*connect_go.Response[v1.PushSchemaResponse], error)
	GetSchema(context.Context, *connect_go.Request[v1.GetSchemaRequest]) (*connect_go.Response[v1.GetSchemaResponse], error)
	PushRecords(context.Context, *connect_go.Request[v1.PushRecordsRequest]) (*connect_go.Response[v1.PushRecordsResponse], error)
	GetRecords(context.Context, *connect_go.Request[v1.GetRecordsRequest]) (*connect_go.Response[v1.GetRecordsResponse], error)
}

// NewDataServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDataServiceHandler(svc DataServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/ujds.v1.DataService/PushSchema", connect_go.NewUnaryHandler(
		"/ujds.v1.DataService/PushSchema",
		svc.PushSchema,
		opts...,
	))
	mux.Handle("/ujds.v1.DataService/GetSchema", connect_go.NewUnaryHandler(
		"/ujds.v1.DataService/GetSchema",
		svc.GetSchema,
		opts...,
	))
	mux.Handle("/ujds.v1.DataService/PushRecords", connect_go.NewUnaryHandler(
		"/ujds.v1.DataService/PushRecords",
		svc.PushRecords,
		opts...,
	))
	mux.Handle("/ujds.v1.DataService/GetRecords", connect_go.NewUnaryHandler(
		"/ujds.v1.DataService/GetRecords",
		svc.GetRecords,
		opts...,
	))
	return "/ujds.v1.DataService/", mux
}

// UnimplementedDataServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDataServiceHandler struct{}

func (UnimplementedDataServiceHandler) PushSchema(context.Context, *connect_go.Request[v1.PushSchemaRequest]) (*connect_go.Response[v1.PushSchemaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.v1.DataService.PushSchema is not implemented"))
}

func (UnimplementedDataServiceHandler) GetSchema(context.Context, *connect_go.Request[v1.GetSchemaRequest]) (*connect_go.Response[v1.GetSchemaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.v1.DataService.GetSchema is not implemented"))
}

func (UnimplementedDataServiceHandler) PushRecords(context.Context, *connect_go.Request[v1.PushRecordsRequest]) (*connect_go.Response[v1.PushRecordsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.v1.DataService.PushRecords is not implemented"))
}

func (UnimplementedDataServiceHandler) GetRecords(context.Context, *connect_go.Request[v1.GetRecordsRequest]) (*connect_go.Response[v1.GetRecordsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.v1.DataService.GetRecords is not implemented"))
}

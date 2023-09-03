// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/ujds/index/v1/index.proto

package v1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/ashep/ujds/sdk/proto/ujds/index/v1"
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
	// IndexServiceName is the fully-qualified name of the IndexService service.
	IndexServiceName = "ujds.index.v1.IndexService"
)

// IndexServiceClient is a client for the ujds.index.v1.IndexService service.
type IndexServiceClient interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Clear(context.Context, *connect_go.Request[v1.ClearRequest]) (*connect_go.Response[v1.ClearResponse], error)
}

// NewIndexServiceClient constructs a client for the ujds.index.v1.IndexService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIndexServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IndexServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &indexServiceClient{
		push: connect_go.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+"/ujds.index.v1.IndexService/Push",
			opts...,
		),
		get: connect_go.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+"/ujds.index.v1.IndexService/Get",
			opts...,
		),
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/ujds.index.v1.IndexService/List",
			opts...,
		),
		clear: connect_go.NewClient[v1.ClearRequest, v1.ClearResponse](
			httpClient,
			baseURL+"/ujds.index.v1.IndexService/Clear",
			opts...,
		),
	}
}

// indexServiceClient implements IndexServiceClient.
type indexServiceClient struct {
	push  *connect_go.Client[v1.PushRequest, v1.PushResponse]
	get   *connect_go.Client[v1.GetRequest, v1.GetResponse]
	list  *connect_go.Client[v1.ListRequest, v1.ListResponse]
	clear *connect_go.Client[v1.ClearRequest, v1.ClearResponse]
}

// Push calls ujds.index.v1.IndexService.Push.
func (c *indexServiceClient) Push(ctx context.Context, req *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// Get calls ujds.index.v1.IndexService.Get.
func (c *indexServiceClient) Get(ctx context.Context, req *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls ujds.index.v1.IndexService.List.
func (c *indexServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Clear calls ujds.index.v1.IndexService.Clear.
func (c *indexServiceClient) Clear(ctx context.Context, req *connect_go.Request[v1.ClearRequest]) (*connect_go.Response[v1.ClearResponse], error) {
	return c.clear.CallUnary(ctx, req)
}

// IndexServiceHandler is an implementation of the ujds.index.v1.IndexService service.
type IndexServiceHandler interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	Clear(context.Context, *connect_go.Request[v1.ClearRequest]) (*connect_go.Response[v1.ClearResponse], error)
}

// NewIndexServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIndexServiceHandler(svc IndexServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/ujds.index.v1.IndexService/Push", connect_go.NewUnaryHandler(
		"/ujds.index.v1.IndexService/Push",
		svc.Push,
		opts...,
	))
	mux.Handle("/ujds.index.v1.IndexService/Get", connect_go.NewUnaryHandler(
		"/ujds.index.v1.IndexService/Get",
		svc.Get,
		opts...,
	))
	mux.Handle("/ujds.index.v1.IndexService/List", connect_go.NewUnaryHandler(
		"/ujds.index.v1.IndexService/List",
		svc.List,
		opts...,
	))
	mux.Handle("/ujds.index.v1.IndexService/Clear", connect_go.NewUnaryHandler(
		"/ujds.index.v1.IndexService/Clear",
		svc.Clear,
		opts...,
	))
	return "/ujds.index.v1.IndexService/", mux
}

// UnimplementedIndexServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIndexServiceHandler struct{}

func (UnimplementedIndexServiceHandler) Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Push is not implemented"))
}

func (UnimplementedIndexServiceHandler) Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Get is not implemented"))
}

func (UnimplementedIndexServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.List is not implemented"))
}

func (UnimplementedIndexServiceHandler) Clear(context.Context, *connect_go.Request[v1.ClearRequest]) (*connect_go.Response[v1.ClearResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Clear is not implemented"))
}

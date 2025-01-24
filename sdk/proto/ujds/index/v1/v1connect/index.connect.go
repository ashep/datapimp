// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ujds/index/v1/index.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ashep/ujds/sdk/proto/ujds/index/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// IndexServiceName is the fully-qualified name of the IndexService service.
	IndexServiceName = "ujds.index.v1.IndexService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IndexServicePushProcedure is the fully-qualified name of the IndexService's Push RPC.
	IndexServicePushProcedure = "/ujds.index.v1.IndexService/Push"
	// IndexServiceGetProcedure is the fully-qualified name of the IndexService's Get RPC.
	IndexServiceGetProcedure = "/ujds.index.v1.IndexService/Get"
	// IndexServiceListProcedure is the fully-qualified name of the IndexService's List RPC.
	IndexServiceListProcedure = "/ujds.index.v1.IndexService/List"
	// IndexServiceClearProcedure is the fully-qualified name of the IndexService's Clear RPC.
	IndexServiceClearProcedure = "/ujds.index.v1.IndexService/Clear"
)

// IndexServiceClient is a client for the ujds.index.v1.IndexService service.
type IndexServiceClient interface {
	Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
	Clear(context.Context, *connect.Request[v1.ClearRequest]) (*connect.Response[v1.ClearResponse], error)
}

// NewIndexServiceClient constructs a client for the ujds.index.v1.IndexService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIndexServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IndexServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	indexServiceMethods := v1.File_ujds_index_v1_index_proto.Services().ByName("IndexService").Methods()
	return &indexServiceClient{
		push: connect.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+IndexServicePushProcedure,
			connect.WithSchema(indexServiceMethods.ByName("Push")),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+IndexServiceGetProcedure,
			connect.WithSchema(indexServiceMethods.ByName("Get")),
			connect.WithClientOptions(opts...),
		),
		list: connect.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+IndexServiceListProcedure,
			connect.WithSchema(indexServiceMethods.ByName("List")),
			connect.WithClientOptions(opts...),
		),
		clear: connect.NewClient[v1.ClearRequest, v1.ClearResponse](
			httpClient,
			baseURL+IndexServiceClearProcedure,
			connect.WithSchema(indexServiceMethods.ByName("Clear")),
			connect.WithClientOptions(opts...),
		),
	}
}

// indexServiceClient implements IndexServiceClient.
type indexServiceClient struct {
	push  *connect.Client[v1.PushRequest, v1.PushResponse]
	get   *connect.Client[v1.GetRequest, v1.GetResponse]
	list  *connect.Client[v1.ListRequest, v1.ListResponse]
	clear *connect.Client[v1.ClearRequest, v1.ClearResponse]
}

// Push calls ujds.index.v1.IndexService.Push.
func (c *indexServiceClient) Push(ctx context.Context, req *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// Get calls ujds.index.v1.IndexService.Get.
func (c *indexServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// List calls ujds.index.v1.IndexService.List.
func (c *indexServiceClient) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Clear calls ujds.index.v1.IndexService.Clear.
func (c *indexServiceClient) Clear(ctx context.Context, req *connect.Request[v1.ClearRequest]) (*connect.Response[v1.ClearResponse], error) {
	return c.clear.CallUnary(ctx, req)
}

// IndexServiceHandler is an implementation of the ujds.index.v1.IndexService service.
type IndexServiceHandler interface {
	Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
	List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error)
	Clear(context.Context, *connect.Request[v1.ClearRequest]) (*connect.Response[v1.ClearResponse], error)
}

// NewIndexServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIndexServiceHandler(svc IndexServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	indexServiceMethods := v1.File_ujds_index_v1_index_proto.Services().ByName("IndexService").Methods()
	indexServicePushHandler := connect.NewUnaryHandler(
		IndexServicePushProcedure,
		svc.Push,
		connect.WithSchema(indexServiceMethods.ByName("Push")),
		connect.WithHandlerOptions(opts...),
	)
	indexServiceGetHandler := connect.NewUnaryHandler(
		IndexServiceGetProcedure,
		svc.Get,
		connect.WithSchema(indexServiceMethods.ByName("Get")),
		connect.WithHandlerOptions(opts...),
	)
	indexServiceListHandler := connect.NewUnaryHandler(
		IndexServiceListProcedure,
		svc.List,
		connect.WithSchema(indexServiceMethods.ByName("List")),
		connect.WithHandlerOptions(opts...),
	)
	indexServiceClearHandler := connect.NewUnaryHandler(
		IndexServiceClearProcedure,
		svc.Clear,
		connect.WithSchema(indexServiceMethods.ByName("Clear")),
		connect.WithHandlerOptions(opts...),
	)
	return "/ujds.index.v1.IndexService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IndexServicePushProcedure:
			indexServicePushHandler.ServeHTTP(w, r)
		case IndexServiceGetProcedure:
			indexServiceGetHandler.ServeHTTP(w, r)
		case IndexServiceListProcedure:
			indexServiceListHandler.ServeHTTP(w, r)
		case IndexServiceClearProcedure:
			indexServiceClearHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIndexServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIndexServiceHandler struct{}

func (UnimplementedIndexServiceHandler) Push(context.Context, *connect.Request[v1.PushRequest]) (*connect.Response[v1.PushResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Push is not implemented"))
}

func (UnimplementedIndexServiceHandler) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Get is not implemented"))
}

func (UnimplementedIndexServiceHandler) List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.List is not implemented"))
}

func (UnimplementedIndexServiceHandler) Clear(context.Context, *connect.Request[v1.ClearRequest]) (*connect.Response[v1.ClearResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ujds.index.v1.IndexService.Clear is not implemented"))
}

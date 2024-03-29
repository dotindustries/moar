// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: moarpb/v1/moar.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/dotindustries/moar/moarpb/v1"
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
	// ModuleRegistryServiceName is the fully-qualified name of the ModuleRegistryService service.
	ModuleRegistryServiceName = "moarpb.v1.ModuleRegistryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ModuleRegistryServiceGetUrlProcedure is the fully-qualified name of the ModuleRegistryService's
	// GetUrl RPC.
	ModuleRegistryServiceGetUrlProcedure = "/moarpb.v1.ModuleRegistryService/GetUrl"
	// ModuleRegistryServiceCreateModuleProcedure is the fully-qualified name of the
	// ModuleRegistryService's CreateModule RPC.
	ModuleRegistryServiceCreateModuleProcedure = "/moarpb.v1.ModuleRegistryService/CreateModule"
	// ModuleRegistryServiceGetModuleProcedure is the fully-qualified name of the
	// ModuleRegistryService's GetModule RPC.
	ModuleRegistryServiceGetModuleProcedure = "/moarpb.v1.ModuleRegistryService/GetModule"
	// ModuleRegistryServiceDeleteModuleProcedure is the fully-qualified name of the
	// ModuleRegistryService's DeleteModule RPC.
	ModuleRegistryServiceDeleteModuleProcedure = "/moarpb.v1.ModuleRegistryService/DeleteModule"
	// ModuleRegistryServiceUploadVersionProcedure is the fully-qualified name of the
	// ModuleRegistryService's UploadVersion RPC.
	ModuleRegistryServiceUploadVersionProcedure = "/moarpb.v1.ModuleRegistryService/UploadVersion"
	// ModuleRegistryServiceDeleteVersionProcedure is the fully-qualified name of the
	// ModuleRegistryService's DeleteVersion RPC.
	ModuleRegistryServiceDeleteVersionProcedure = "/moarpb.v1.ModuleRegistryService/DeleteVersion"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	moduleRegistryServiceServiceDescriptor             = v1.File_moarpb_v1_moar_proto.Services().ByName("ModuleRegistryService")
	moduleRegistryServiceGetUrlMethodDescriptor        = moduleRegistryServiceServiceDescriptor.Methods().ByName("GetUrl")
	moduleRegistryServiceCreateModuleMethodDescriptor  = moduleRegistryServiceServiceDescriptor.Methods().ByName("CreateModule")
	moduleRegistryServiceGetModuleMethodDescriptor     = moduleRegistryServiceServiceDescriptor.Methods().ByName("GetModule")
	moduleRegistryServiceDeleteModuleMethodDescriptor  = moduleRegistryServiceServiceDescriptor.Methods().ByName("DeleteModule")
	moduleRegistryServiceUploadVersionMethodDescriptor = moduleRegistryServiceServiceDescriptor.Methods().ByName("UploadVersion")
	moduleRegistryServiceDeleteVersionMethodDescriptor = moduleRegistryServiceServiceDescriptor.Methods().ByName("DeleteVersion")
)

// ModuleRegistryServiceClient is a client for the moarpb.v1.ModuleRegistryService service.
type ModuleRegistryServiceClient interface {
	GetUrl(context.Context, *connect.Request[v1.GetUrlRequest]) (*connect.Response[v1.GetUrlResponse], error)
	CreateModule(context.Context, *connect.Request[v1.CreateModuleRequest]) (*connect.Response[v1.CreateModuleResponse], error)
	GetModule(context.Context, *connect.Request[v1.GetModuleRequest]) (*connect.Response[v1.GetModuleResponse], error)
	DeleteModule(context.Context, *connect.Request[v1.DeleteModuleRequest]) (*connect.Response[v1.DeleteModuleResponse], error)
	UploadVersion(context.Context, *connect.Request[v1.UploadVersionRequest]) (*connect.Response[v1.UploadVersionResponse], error)
	DeleteVersion(context.Context, *connect.Request[v1.DeleteVersionRequest]) (*connect.Response[v1.DeleteVersionResponse], error)
}

// NewModuleRegistryServiceClient constructs a client for the moarpb.v1.ModuleRegistryService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewModuleRegistryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ModuleRegistryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &moduleRegistryServiceClient{
		getUrl: connect.NewClient[v1.GetUrlRequest, v1.GetUrlResponse](
			httpClient,
			baseURL+ModuleRegistryServiceGetUrlProcedure,
			connect.WithSchema(moduleRegistryServiceGetUrlMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createModule: connect.NewClient[v1.CreateModuleRequest, v1.CreateModuleResponse](
			httpClient,
			baseURL+ModuleRegistryServiceCreateModuleProcedure,
			connect.WithSchema(moduleRegistryServiceCreateModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getModule: connect.NewClient[v1.GetModuleRequest, v1.GetModuleResponse](
			httpClient,
			baseURL+ModuleRegistryServiceGetModuleProcedure,
			connect.WithSchema(moduleRegistryServiceGetModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteModule: connect.NewClient[v1.DeleteModuleRequest, v1.DeleteModuleResponse](
			httpClient,
			baseURL+ModuleRegistryServiceDeleteModuleProcedure,
			connect.WithSchema(moduleRegistryServiceDeleteModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		uploadVersion: connect.NewClient[v1.UploadVersionRequest, v1.UploadVersionResponse](
			httpClient,
			baseURL+ModuleRegistryServiceUploadVersionProcedure,
			connect.WithSchema(moduleRegistryServiceUploadVersionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteVersion: connect.NewClient[v1.DeleteVersionRequest, v1.DeleteVersionResponse](
			httpClient,
			baseURL+ModuleRegistryServiceDeleteVersionProcedure,
			connect.WithSchema(moduleRegistryServiceDeleteVersionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// moduleRegistryServiceClient implements ModuleRegistryServiceClient.
type moduleRegistryServiceClient struct {
	getUrl        *connect.Client[v1.GetUrlRequest, v1.GetUrlResponse]
	createModule  *connect.Client[v1.CreateModuleRequest, v1.CreateModuleResponse]
	getModule     *connect.Client[v1.GetModuleRequest, v1.GetModuleResponse]
	deleteModule  *connect.Client[v1.DeleteModuleRequest, v1.DeleteModuleResponse]
	uploadVersion *connect.Client[v1.UploadVersionRequest, v1.UploadVersionResponse]
	deleteVersion *connect.Client[v1.DeleteVersionRequest, v1.DeleteVersionResponse]
}

// GetUrl calls moarpb.v1.ModuleRegistryService.GetUrl.
func (c *moduleRegistryServiceClient) GetUrl(ctx context.Context, req *connect.Request[v1.GetUrlRequest]) (*connect.Response[v1.GetUrlResponse], error) {
	return c.getUrl.CallUnary(ctx, req)
}

// CreateModule calls moarpb.v1.ModuleRegistryService.CreateModule.
func (c *moduleRegistryServiceClient) CreateModule(ctx context.Context, req *connect.Request[v1.CreateModuleRequest]) (*connect.Response[v1.CreateModuleResponse], error) {
	return c.createModule.CallUnary(ctx, req)
}

// GetModule calls moarpb.v1.ModuleRegistryService.GetModule.
func (c *moduleRegistryServiceClient) GetModule(ctx context.Context, req *connect.Request[v1.GetModuleRequest]) (*connect.Response[v1.GetModuleResponse], error) {
	return c.getModule.CallUnary(ctx, req)
}

// DeleteModule calls moarpb.v1.ModuleRegistryService.DeleteModule.
func (c *moduleRegistryServiceClient) DeleteModule(ctx context.Context, req *connect.Request[v1.DeleteModuleRequest]) (*connect.Response[v1.DeleteModuleResponse], error) {
	return c.deleteModule.CallUnary(ctx, req)
}

// UploadVersion calls moarpb.v1.ModuleRegistryService.UploadVersion.
func (c *moduleRegistryServiceClient) UploadVersion(ctx context.Context, req *connect.Request[v1.UploadVersionRequest]) (*connect.Response[v1.UploadVersionResponse], error) {
	return c.uploadVersion.CallUnary(ctx, req)
}

// DeleteVersion calls moarpb.v1.ModuleRegistryService.DeleteVersion.
func (c *moduleRegistryServiceClient) DeleteVersion(ctx context.Context, req *connect.Request[v1.DeleteVersionRequest]) (*connect.Response[v1.DeleteVersionResponse], error) {
	return c.deleteVersion.CallUnary(ctx, req)
}

// ModuleRegistryServiceHandler is an implementation of the moarpb.v1.ModuleRegistryService service.
type ModuleRegistryServiceHandler interface {
	GetUrl(context.Context, *connect.Request[v1.GetUrlRequest]) (*connect.Response[v1.GetUrlResponse], error)
	CreateModule(context.Context, *connect.Request[v1.CreateModuleRequest]) (*connect.Response[v1.CreateModuleResponse], error)
	GetModule(context.Context, *connect.Request[v1.GetModuleRequest]) (*connect.Response[v1.GetModuleResponse], error)
	DeleteModule(context.Context, *connect.Request[v1.DeleteModuleRequest]) (*connect.Response[v1.DeleteModuleResponse], error)
	UploadVersion(context.Context, *connect.Request[v1.UploadVersionRequest]) (*connect.Response[v1.UploadVersionResponse], error)
	DeleteVersion(context.Context, *connect.Request[v1.DeleteVersionRequest]) (*connect.Response[v1.DeleteVersionResponse], error)
}

// NewModuleRegistryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewModuleRegistryServiceHandler(svc ModuleRegistryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	moduleRegistryServiceGetUrlHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceGetUrlProcedure,
		svc.GetUrl,
		connect.WithSchema(moduleRegistryServiceGetUrlMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moduleRegistryServiceCreateModuleHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceCreateModuleProcedure,
		svc.CreateModule,
		connect.WithSchema(moduleRegistryServiceCreateModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moduleRegistryServiceGetModuleHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceGetModuleProcedure,
		svc.GetModule,
		connect.WithSchema(moduleRegistryServiceGetModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moduleRegistryServiceDeleteModuleHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceDeleteModuleProcedure,
		svc.DeleteModule,
		connect.WithSchema(moduleRegistryServiceDeleteModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moduleRegistryServiceUploadVersionHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceUploadVersionProcedure,
		svc.UploadVersion,
		connect.WithSchema(moduleRegistryServiceUploadVersionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	moduleRegistryServiceDeleteVersionHandler := connect.NewUnaryHandler(
		ModuleRegistryServiceDeleteVersionProcedure,
		svc.DeleteVersion,
		connect.WithSchema(moduleRegistryServiceDeleteVersionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/moarpb.v1.ModuleRegistryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ModuleRegistryServiceGetUrlProcedure:
			moduleRegistryServiceGetUrlHandler.ServeHTTP(w, r)
		case ModuleRegistryServiceCreateModuleProcedure:
			moduleRegistryServiceCreateModuleHandler.ServeHTTP(w, r)
		case ModuleRegistryServiceGetModuleProcedure:
			moduleRegistryServiceGetModuleHandler.ServeHTTP(w, r)
		case ModuleRegistryServiceDeleteModuleProcedure:
			moduleRegistryServiceDeleteModuleHandler.ServeHTTP(w, r)
		case ModuleRegistryServiceUploadVersionProcedure:
			moduleRegistryServiceUploadVersionHandler.ServeHTTP(w, r)
		case ModuleRegistryServiceDeleteVersionProcedure:
			moduleRegistryServiceDeleteVersionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedModuleRegistryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedModuleRegistryServiceHandler struct{}

func (UnimplementedModuleRegistryServiceHandler) GetUrl(context.Context, *connect.Request[v1.GetUrlRequest]) (*connect.Response[v1.GetUrlResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.GetUrl is not implemented"))
}

func (UnimplementedModuleRegistryServiceHandler) CreateModule(context.Context, *connect.Request[v1.CreateModuleRequest]) (*connect.Response[v1.CreateModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.CreateModule is not implemented"))
}

func (UnimplementedModuleRegistryServiceHandler) GetModule(context.Context, *connect.Request[v1.GetModuleRequest]) (*connect.Response[v1.GetModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.GetModule is not implemented"))
}

func (UnimplementedModuleRegistryServiceHandler) DeleteModule(context.Context, *connect.Request[v1.DeleteModuleRequest]) (*connect.Response[v1.DeleteModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.DeleteModule is not implemented"))
}

func (UnimplementedModuleRegistryServiceHandler) UploadVersion(context.Context, *connect.Request[v1.UploadVersionRequest]) (*connect.Response[v1.UploadVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.UploadVersion is not implemented"))
}

func (UnimplementedModuleRegistryServiceHandler) DeleteVersion(context.Context, *connect.Request[v1.DeleteVersionRequest]) (*connect.Response[v1.DeleteVersionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("moarpb.v1.ModuleRegistryService.DeleteVersion is not implemented"))
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package viz

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	StatSummary(ctx context.Context, in *StatSummaryRequest, opts ...grpc.CallOption) (*StatSummaryResponse, error)
	Edges(ctx context.Context, in *EdgesRequest, opts ...grpc.CallOption) (*EdgesResponse, error)
	Gateways(ctx context.Context, in *GatewaysRequest, opts ...grpc.CallOption) (*GatewaysResponse, error)
	TopRoutes(ctx context.Context, in *TopRoutesRequest, opts ...grpc.CallOption) (*TopRoutesResponse, error)
	ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	SelfCheck(ctx context.Context, in *SelfCheckRequest, opts ...grpc.CallOption) (*SelfCheckResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) StatSummary(ctx context.Context, in *StatSummaryRequest, opts ...grpc.CallOption) (*StatSummaryResponse, error) {
	out := new(StatSummaryResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/StatSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Edges(ctx context.Context, in *EdgesRequest, opts ...grpc.CallOption) (*EdgesResponse, error) {
	out := new(EdgesResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/Edges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Gateways(ctx context.Context, in *GatewaysRequest, opts ...grpc.CallOption) (*GatewaysResponse, error) {
	out := new(GatewaysResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/Gateways", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) TopRoutes(ctx context.Context, in *TopRoutesRequest, opts ...grpc.CallOption) (*TopRoutesResponse, error) {
	out := new(TopRoutesResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/TopRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SelfCheck(ctx context.Context, in *SelfCheckRequest, opts ...grpc.CallOption) (*SelfCheckResponse, error) {
	out := new(SelfCheckResponse)
	err := c.cc.Invoke(ctx, "/linkerd2.viz.Api/SelfCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	StatSummary(context.Context, *StatSummaryRequest) (*StatSummaryResponse, error)
	Edges(context.Context, *EdgesRequest) (*EdgesResponse, error)
	Gateways(context.Context, *GatewaysRequest) (*GatewaysResponse, error)
	TopRoutes(context.Context, *TopRoutesRequest) (*TopRoutesResponse, error)
	ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	SelfCheck(context.Context, *SelfCheckRequest) (*SelfCheckResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) StatSummary(context.Context, *StatSummaryRequest) (*StatSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatSummary not implemented")
}
func (UnimplementedApiServer) Edges(context.Context, *EdgesRequest) (*EdgesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edges not implemented")
}
func (UnimplementedApiServer) Gateways(context.Context, *GatewaysRequest) (*GatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gateways not implemented")
}
func (UnimplementedApiServer) TopRoutes(context.Context, *TopRoutesRequest) (*TopRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopRoutes not implemented")
}
func (UnimplementedApiServer) ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (UnimplementedApiServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedApiServer) SelfCheck(context.Context, *SelfCheckRequest) (*SelfCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfCheck not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_StatSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).StatSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/StatSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).StatSummary(ctx, req.(*StatSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Edges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Edges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/Edges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Edges(ctx, req.(*EdgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Gateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Gateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/Gateways",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Gateways(ctx, req.(*GatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_TopRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).TopRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/TopRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).TopRoutes(ctx, req.(*TopRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListPods(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SelfCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SelfCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.viz.Api/SelfCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SelfCheck(ctx, req.(*SelfCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkerd2.viz.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatSummary",
			Handler:    _Api_StatSummary_Handler,
		},
		{
			MethodName: "Edges",
			Handler:    _Api_Edges_Handler,
		},
		{
			MethodName: "Gateways",
			Handler:    _Api_Gateways_Handler,
		},
		{
			MethodName: "TopRoutes",
			Handler:    _Api_TopRoutes_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _Api_ListPods_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Api_ListServices_Handler,
		},
		{
			MethodName: "SelfCheck",
			Handler:    _Api_SelfCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "viz.proto",
}
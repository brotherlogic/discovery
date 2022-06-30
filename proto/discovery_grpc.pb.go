// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	RegisterService(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (*DiscoverResponse, error)
	ListAllServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
}

type discoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceClient(cc grpc.ClientConnInterface) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) RegisterService(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryService/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (*DiscoverResponse, error) {
	out := new(DiscoverResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryService/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) ListAllServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryService/ListAllServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryService/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
// All implementations should embed UnimplementedDiscoveryServiceServer
// for forward compatibility
type DiscoveryServiceServer interface {
	RegisterService(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Discover(context.Context, *DiscoverRequest) (*DiscoverResponse, error)
	ListAllServices(context.Context, *ListRequest) (*ListResponse, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
}

// UnimplementedDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceServer struct {
}

func (UnimplementedDiscoveryServiceServer) RegisterService(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedDiscoveryServiceServer) Discover(context.Context, *DiscoverRequest) (*DiscoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedDiscoveryServiceServer) ListAllServices(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllServices not implemented")
}
func (UnimplementedDiscoveryServiceServer) State(context.Context, *StateRequest) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}

// UnsafeDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServiceServer will
// result in compilation errors.
type UnsafeDiscoveryServiceServer interface {
	mustEmbedUnimplementedDiscoveryServiceServer()
}

func RegisterDiscoveryServiceServer(s grpc.ServiceRegistrar, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryService/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).RegisterService(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryService/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Discover(ctx, req.(*DiscoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_ListAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).ListAllServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryService/ListAllServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).ListAllServices(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryService/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _DiscoveryService_RegisterService_Handler,
		},
		{
			MethodName: "Discover",
			Handler:    _DiscoveryService_Discover_Handler,
		},
		{
			MethodName: "ListAllServices",
			Handler:    _DiscoveryService_ListAllServices_Handler,
		},
		{
			MethodName: "State",
			Handler:    _DiscoveryService_State_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

// DiscoveryServiceV2Client is the client API for DiscoveryServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryServiceV2Client interface {
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	RegisterV2(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error)
	MasterElect(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*MasterResponse, error)
	GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error)
	GetInternalState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type discoveryServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceV2Client(cc grpc.ClientConnInterface) DiscoveryServiceV2Client {
	return &discoveryServiceV2Client{cc}
}

func (c *discoveryServiceV2Client) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) RegisterV2(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/RegisterV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error) {
	out := new(UnregisterResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) MasterElect(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*MasterResponse, error) {
	out := new(MasterResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/MasterElect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error) {
	out := new(GetFriendsResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) GetInternalState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/GetInternalState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceV2Client) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/discovery.DiscoveryServiceV2/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceV2Server is the server API for DiscoveryServiceV2 service.
// All implementations should embed UnimplementedDiscoveryServiceV2Server
// for forward compatibility
type DiscoveryServiceV2Server interface {
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	RegisterV2(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error)
	MasterElect(context.Context, *MasterRequest) (*MasterResponse, error)
	GetFriends(context.Context, *GetFriendsRequest) (*GetFriendsResponse, error)
	GetInternalState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
}

// UnimplementedDiscoveryServiceV2Server should be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceV2Server struct {
}

func (UnimplementedDiscoveryServiceV2Server) Lock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) RegisterV2(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterV2 not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) MasterElect(context.Context, *MasterRequest) (*MasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MasterElect not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) GetFriends(context.Context, *GetFriendsRequest) (*GetFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) GetInternalState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInternalState not implemented")
}
func (UnimplementedDiscoveryServiceV2Server) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}

// UnsafeDiscoveryServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServiceV2Server will
// result in compilation errors.
type UnsafeDiscoveryServiceV2Server interface {
	mustEmbedUnimplementedDiscoveryServiceV2Server()
}

func RegisterDiscoveryServiceV2Server(s grpc.ServiceRegistrar, srv DiscoveryServiceV2Server) {
	s.RegisterService(&_DiscoveryServiceV2_serviceDesc, srv)
}

func _DiscoveryServiceV2_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_RegisterV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).RegisterV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/RegisterV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).RegisterV2(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_MasterElect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).MasterElect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/MasterElect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).MasterElect(ctx, req.(*MasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).GetFriends(ctx, req.(*GetFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_GetInternalState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).GetInternalState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/GetInternalState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).GetInternalState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryServiceV2_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceV2Server).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.DiscoveryServiceV2/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceV2Server).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.DiscoveryServiceV2",
	HandlerType: (*DiscoveryServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _DiscoveryServiceV2_Lock_Handler,
		},
		{
			MethodName: "RegisterV2",
			Handler:    _DiscoveryServiceV2_RegisterV2_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DiscoveryServiceV2_Get_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _DiscoveryServiceV2_Unregister_Handler,
		},
		{
			MethodName: "MasterElect",
			Handler:    _DiscoveryServiceV2_MasterElect_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _DiscoveryServiceV2_GetFriends_Handler,
		},
		{
			MethodName: "GetInternalState",
			Handler:    _DiscoveryServiceV2_GetInternalState_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _DiscoveryServiceV2_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

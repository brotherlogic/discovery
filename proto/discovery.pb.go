// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery.proto

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	discovery.proto

It has these top-level messages:
	RegistryEntry
	ServiceList
	Empty
	StateResponse
	StateRequest
*/
package discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegistryEntry struct {
	// The ip address associated with this entry
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	// The port number assigned / requested for this entry
	Port int32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	// The name of this service
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Is this an external facing port
	ExternalPort bool `protobuf:"varint,4,opt,name=external_port,json=externalPort" json:"external_port,omitempty"`
	// This is the machine identifier
	Identifier string `protobuf:"bytes,5,opt,name=identifier" json:"identifier,omitempty"`
	// Boolean to show we're master/slave
	Master bool `protobuf:"varint,6,opt,name=master" json:"master,omitempty"`
	// The time at which this binary was registered
	RegisterTime int64 `protobuf:"varint,7,opt,name=register_time,json=registerTime" json:"register_time,omitempty"`
	// The time at which this binary should be cleaned
	TimeToClean int64 `protobuf:"varint,8,opt,name=time_to_clean,json=timeToClean" json:"time_to_clean,omitempty"`
	// The time at which this binary was last seen
	LastSeenTime int64 `protobuf:"varint,9,opt,name=last_seen_time,json=lastSeenTime" json:"last_seen_time,omitempty"`
	// We are never going to be master
	IgnoresMaster bool `protobuf:"varint,10,opt,name=ignores_master,json=ignoresMaster" json:"ignores_master,omitempty"`
}

func (m *RegistryEntry) Reset()                    { *m = RegistryEntry{} }
func (m *RegistryEntry) String() string            { return proto.CompactTextString(m) }
func (*RegistryEntry) ProtoMessage()               {}
func (*RegistryEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistryEntry) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegistryEntry) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegistryEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegistryEntry) GetExternalPort() bool {
	if m != nil {
		return m.ExternalPort
	}
	return false
}

func (m *RegistryEntry) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *RegistryEntry) GetMaster() bool {
	if m != nil {
		return m.Master
	}
	return false
}

func (m *RegistryEntry) GetRegisterTime() int64 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *RegistryEntry) GetTimeToClean() int64 {
	if m != nil {
		return m.TimeToClean
	}
	return 0
}

func (m *RegistryEntry) GetLastSeenTime() int64 {
	if m != nil {
		return m.LastSeenTime
	}
	return 0
}

func (m *RegistryEntry) GetIgnoresMaster() bool {
	if m != nil {
		return m.IgnoresMaster
	}
	return false
}

type ServiceList struct {
	Services []*RegistryEntry `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ServiceList) Reset()                    { *m = ServiceList{} }
func (m *ServiceList) String() string            { return proto.CompactTextString(m) }
func (*ServiceList) ProtoMessage()               {}
func (*ServiceList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceList) GetServices() []*RegistryEntry {
	if m != nil {
		return m.Services
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StateResponse struct {
	Counts string `protobuf:"bytes,1,opt,name=counts" json:"counts,omitempty"`
	Len    int32  `protobuf:"varint,2,opt,name=len" json:"len,omitempty"`
}

func (m *StateResponse) Reset()                    { *m = StateResponse{} }
func (m *StateResponse) String() string            { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()               {}
func (*StateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StateResponse) GetCounts() string {
	if m != nil {
		return m.Counts
	}
	return ""
}

func (m *StateResponse) GetLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}

type StateRequest struct {
}

func (m *StateRequest) Reset()                    { *m = StateRequest{} }
func (m *StateRequest) String() string            { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()               {}
func (*StateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*RegistryEntry)(nil), "discovery.RegistryEntry")
	proto.RegisterType((*ServiceList)(nil), "discovery.ServiceList")
	proto.RegisterType((*Empty)(nil), "discovery.Empty")
	proto.RegisterType((*StateResponse)(nil), "discovery.StateResponse")
	proto.RegisterType((*StateRequest)(nil), "discovery.StateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DiscoveryService service

type DiscoveryServiceClient interface {
	RegisterService(ctx context.Context, in *RegistryEntry, opts ...grpc.CallOption) (*RegistryEntry, error)
	Discover(ctx context.Context, in *RegistryEntry, opts ...grpc.CallOption) (*RegistryEntry, error)
	ListAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceList, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
}

type discoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceClient(cc *grpc.ClientConn) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) RegisterService(ctx context.Context, in *RegistryEntry, opts ...grpc.CallOption) (*RegistryEntry, error) {
	out := new(RegistryEntry)
	err := grpc.Invoke(ctx, "/discovery.DiscoveryService/RegisterService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) Discover(ctx context.Context, in *RegistryEntry, opts ...grpc.CallOption) (*RegistryEntry, error) {
	out := new(RegistryEntry)
	err := grpc.Invoke(ctx, "/discovery.DiscoveryService/Discover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) ListAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceList, error) {
	out := new(ServiceList)
	err := grpc.Invoke(ctx, "/discovery.DiscoveryService/ListAllServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := grpc.Invoke(ctx, "/discovery.DiscoveryService/State", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiscoveryService service

type DiscoveryServiceServer interface {
	RegisterService(context.Context, *RegistryEntry) (*RegistryEntry, error)
	Discover(context.Context, *RegistryEntry) (*RegistryEntry, error)
	ListAllServices(context.Context, *Empty) (*ServiceList, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryEntry)
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
		return srv.(DiscoveryServiceServer).RegisterService(ctx, req.(*RegistryEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryEntry)
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
		return srv.(DiscoveryServiceServer).Discover(ctx, req.(*RegistryEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_ListAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(DiscoveryServiceServer).ListAllServices(ctx, req.(*Empty))
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

func init() { proto.RegisterFile("discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x74, 0xed, 0xda, 0xd7, 0xa6, 0xad, 0x7c, 0x18, 0xd6, 0x0e, 0x28, 0x32, 0x20,
	0xe5, 0xb4, 0xc3, 0xe0, 0x82, 0x04, 0x12, 0x68, 0x4c, 0x5c, 0x40, 0x42, 0xee, 0xee, 0x51, 0xc8,
	0x1e, 0x93, 0xa5, 0xc4, 0x0e, 0xf6, 0xdb, 0x44, 0xbe, 0x0a, 0x1f, 0x85, 0x4f, 0x87, 0xec, 0x3a,
	0x23, 0x3d, 0xec, 0xc2, 0xed, 0xbd, 0x9f, 0xdf, 0xfb, 0xfb, 0x9f, 0xbf, 0x15, 0xd8, 0xde, 0x2a,
	0x57, 0x9b, 0x07, 0xb4, 0xfd, 0x45, 0x67, 0x0d, 0x19, 0xb6, 0x7c, 0x04, 0xe2, 0x4f, 0x0a, 0x99,
	0xc4, 0x3b, 0xe5, 0xc8, 0xf6, 0xd7, 0x9a, 0x6c, 0xcf, 0x36, 0x90, 0xaa, 0x8e, 0x27, 0x79, 0x52,
	0x2c, 0x65, 0xaa, 0x3a, 0xc6, 0xe0, 0xa4, 0x33, 0x96, 0x78, 0x9a, 0x27, 0xc5, 0x4c, 0x86, 0xda,
	0x33, 0x5d, 0xb5, 0xc8, 0xa7, 0x61, 0x2a, 0xd4, 0xec, 0x05, 0x64, 0xf8, 0x8b, 0xd0, 0xea, 0xaa,
	0x29, 0xc3, 0xc2, 0x49, 0x9e, 0x14, 0x0b, 0xb9, 0x1e, 0xe0, 0x37, 0xbf, 0xf8, 0x1c, 0x40, 0xdd,
	0xa2, 0x26, 0xf5, 0x43, 0xa1, 0xe5, 0xb3, 0xb0, 0x3e, 0x22, 0xec, 0x0c, 0xe6, 0x6d, 0xe5, 0x08,
	0x2d, 0x9f, 0x87, 0xed, 0xd8, 0x79, 0x71, 0x1b, 0x5c, 0xa2, 0x2d, 0x49, 0xb5, 0xc8, 0x4f, 0xf3,
	0xa4, 0x98, 0xca, 0xf5, 0x00, 0x6f, 0x54, 0x8b, 0x4c, 0x40, 0xe6, 0xcf, 0x4a, 0x32, 0x65, 0xdd,
	0x60, 0xa5, 0xf9, 0x22, 0x0c, 0xad, 0x3c, 0xbc, 0x31, 0x57, 0x1e, 0xb1, 0x97, 0xb0, 0x69, 0x2a,
	0x47, 0xa5, 0x43, 0xd4, 0x07, 0xa5, 0xe5, 0x41, 0xc9, 0xd3, 0x3d, 0xa2, 0x0e, 0x4a, 0xaf, 0x60,
	0xa3, 0xee, 0xb4, 0xb1, 0xe8, 0xca, 0x68, 0x07, 0x82, 0x9d, 0x2c, 0xd2, 0xaf, 0x01, 0x8a, 0x2b,
	0x58, 0xed, 0xd1, 0x3e, 0xa8, 0x1a, 0xbf, 0x28, 0x47, 0xec, 0x0d, 0x2c, 0xdc, 0xa1, 0x75, 0x3c,
	0xc9, 0xa7, 0xc5, 0xea, 0x92, 0x5f, 0xfc, 0x8b, 0xfe, 0x28, 0x65, 0xf9, 0x38, 0x29, 0x4e, 0x61,
	0x76, 0xdd, 0x76, 0xd4, 0x8b, 0xb7, 0x90, 0xed, 0xa9, 0x22, 0x94, 0xe8, 0x3a, 0xa3, 0x1d, 0xfa,
	0x30, 0x6a, 0x73, 0xaf, 0xc9, 0xc5, 0xd7, 0x88, 0x1d, 0xdb, 0xc1, 0xb4, 0x41, 0x1d, 0x1f, 0xc4,
	0x97, 0x62, 0x03, 0xeb, 0xb8, 0xfa, 0xf3, 0x1e, 0x1d, 0x5d, 0xfe, 0x4e, 0x61, 0xf7, 0x69, 0xb8,
	0x39, 0x5a, 0x64, 0x9f, 0x61, 0x2b, 0x63, 0x5c, 0x03, 0x7a, 0xd2, 0xdf, 0xf9, 0x93, 0x27, 0x62,
	0xc2, 0x3e, 0xc0, 0x62, 0x10, 0xff, 0x4f, 0x85, 0xf7, 0xb0, 0xf5, 0x89, 0x7d, 0x6c, 0x9a, 0xe8,
	0xc4, 0xb1, 0xdd, 0x68, 0x3c, 0xe4, 0x71, 0x7e, 0x36, 0x22, 0xa3, 0x98, 0xc5, 0x84, 0xbd, 0x83,
	0x59, 0xf8, 0x5c, 0xf6, 0x6c, 0x3c, 0x32, 0x0a, 0xe0, 0xe8, 0xf2, 0xa3, 0x50, 0xc5, 0xe4, 0xfb,
	0x3c, 0xfc, 0x04, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x46, 0xbf, 0xd8, 0x62, 0x17, 0x03,
	0x00, 0x00,
}

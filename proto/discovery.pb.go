// Code generated by protoc-gen-go.
// source: discovery.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	discovery.proto

It has these top-level messages:
	RegistryEntry
	ServiceList
	Empty
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
}

func (m *RegistryEntry) Reset()                    { *m = RegistryEntry{} }
func (m *RegistryEntry) String() string            { return proto.CompactTextString(m) }
func (*RegistryEntry) ProtoMessage()               {}
func (*RegistryEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func init() {
	proto.RegisterType((*RegistryEntry)(nil), "discovery.RegistryEntry")
	proto.RegisterType((*ServiceList)(nil), "discovery.ServiceList")
	proto.RegisterType((*Empty)(nil), "discovery.Empty")
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

// Server API for DiscoveryService service

type DiscoveryServiceServer interface {
	RegisterService(context.Context, *RegistryEntry) (*RegistryEntry, error)
	Discover(context.Context, *RegistryEntry) (*RegistryEntry, error)
	ListAllServices(context.Context, *Empty) (*ServiceList, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xb4, 0xf9, 0xff, 0xf4, 0x96, 0x92, 0xca, 0x03, 0xb2, 0x3a, 0xa0, 0xc8, 0x2c,
	0x99, 0x3a, 0x14, 0x56, 0x24, 0x10, 0x54, 0x2c, 0x0c, 0xc8, 0x3c, 0x00, 0x2a, 0xed, 0x05, 0x5d,
	0x29, 0x4d, 0xac, 0x6b, 0xab, 0x22, 0xef, 0xc0, 0x23, 0xf2, 0x30, 0x28, 0x6e, 0x53, 0xc2, 0xd0,
	0x85, 0xed, 0xfa, 0xf3, 0x39, 0x47, 0xf7, 0xd8, 0x90, 0xae, 0xc9, 0xad, 0xaa, 0x2d, 0x72, 0x3d,
	0xb3, 0x5c, 0xf9, 0x4a, 0x0e, 0x0f, 0x40, 0x7f, 0x0a, 0x18, 0x1b, 0x7c, 0x27, 0xe7, 0xb9, 0x5e,
	0x94, 0x9e, 0x6b, 0x79, 0x0a, 0x11, 0x59, 0x25, 0x32, 0x91, 0x0f, 0x4d, 0x44, 0x56, 0x4a, 0x18,
	0xd8, 0x8a, 0xbd, 0x8a, 0x32, 0x91, 0xc7, 0x26, 0xcc, 0x0d, 0x2b, 0x97, 0x1b, 0x54, 0xfd, 0xa0,
	0x0a, 0xb3, 0xbc, 0x80, 0x31, 0x7e, 0x78, 0xe4, 0x72, 0x59, 0xbc, 0x04, 0xc3, 0x20, 0x13, 0x79,
	0x62, 0x4e, 0x5a, 0xf8, 0xd4, 0x18, 0xcf, 0x01, 0x68, 0x8d, 0xa5, 0xa7, 0x37, 0x42, 0x56, 0x71,
	0xb0, 0x77, 0x88, 0xbe, 0x83, 0xd1, 0x33, 0xf2, 0x96, 0x56, 0xf8, 0x48, 0xce, 0xcb, 0x2b, 0x48,
	0xdc, 0xee, 0xe8, 0x94, 0xc8, 0xfa, 0xf9, 0x68, 0xae, 0x66, 0x3f, 0x65, 0x7e, 0xed, 0x6d, 0x0e,
	0x4a, 0xfd, 0x1f, 0xe2, 0xc5, 0xc6, 0xfa, 0x7a, 0xfe, 0x25, 0x60, 0x72, 0xdf, 0xca, 0xf7, 0xb9,
	0xf2, 0x01, 0xd2, 0x9d, 0x11, 0xb9, 0x45, 0x47, 0x43, 0xa7, 0x47, 0x6f, 0x74, 0x4f, 0xde, 0x40,
	0xd2, 0x86, 0xff, 0x31, 0xe1, 0x1a, 0xd2, 0xa6, 0xe6, 0x6d, 0x51, 0xec, 0x37, 0x71, 0x72, 0xd2,
	0x91, 0x87, 0x12, 0xd3, 0xb3, 0x0e, 0xe9, 0xbc, 0x8d, 0xee, 0xbd, 0xfe, 0x0b, 0xbf, 0x79, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x34, 0x3b, 0x42, 0xe0, 0x01, 0x00, 0x00,
}

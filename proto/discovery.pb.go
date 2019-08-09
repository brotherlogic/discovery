// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery.proto

package discovery

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type RegistryEntry_Version int32

const (
	RegistryEntry_V1 RegistryEntry_Version = 0
	RegistryEntry_V2 RegistryEntry_Version = 1
)

var RegistryEntry_Version_name = map[int32]string{
	0: "V1",
	1: "V2",
}

var RegistryEntry_Version_value = map[string]int32{
	"V1": 0,
	"V2": 1,
}

func (x RegistryEntry_Version) String() string {
	return proto.EnumName(RegistryEntry_Version_name, int32(x))
}

func (RegistryEntry_Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0, 0}
}

type RegistryEntry struct {
	// The ip address associated with this entry
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// The port number assigned / requested for this entry
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// The name of this service
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Is this an external facing port
	ExternalPort bool `protobuf:"varint,4,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	// This is the machine identifier
	Identifier string `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Boolean to show we're master/slave
	Master     bool `protobuf:"varint,6,opt,name=master,proto3" json:"master,omitempty"`
	WeakMaster bool `protobuf:"varint,12,opt,name=weak_master,json=weakMaster,proto3" json:"weak_master,omitempty"`
	// The time at which this binary was registered
	RegisterTime int64 `protobuf:"varint,7,opt,name=register_time,json=registerTime,proto3" json:"register_time,omitempty"`
	// The time at which this binary should be cleaned
	TimeToClean int64 `protobuf:"varint,8,opt,name=time_to_clean,json=timeToClean,proto3" json:"time_to_clean,omitempty"`
	// The time at which this binary was last seen
	LastSeenTime int64 `protobuf:"varint,9,opt,name=last_seen_time,json=lastSeenTime,proto3" json:"last_seen_time,omitempty"`
	// We are never going to be master
	IgnoresMaster bool `protobuf:"varint,10,opt,name=ignores_master,json=ignoresMaster,proto3" json:"ignores_master,omitempty"`
	// The time at which we were set master
	MasterTime           int64                 `protobuf:"varint,11,opt,name=master_time,json=masterTime,proto3" json:"master_time,omitempty"`
	Version              RegistryEntry_Version `protobuf:"varint,13,opt,name=version,proto3,enum=discovery.RegistryEntry_Version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegistryEntry) Reset()         { *m = RegistryEntry{} }
func (m *RegistryEntry) String() string { return proto.CompactTextString(m) }
func (*RegistryEntry) ProtoMessage()    {}
func (*RegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}

func (m *RegistryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryEntry.Unmarshal(m, b)
}
func (m *RegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryEntry.Marshal(b, m, deterministic)
}
func (m *RegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryEntry.Merge(m, src)
}
func (m *RegistryEntry) XXX_Size() int {
	return xxx_messageInfo_RegistryEntry.Size(m)
}
func (m *RegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryEntry proto.InternalMessageInfo

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

func (m *RegistryEntry) GetWeakMaster() bool {
	if m != nil {
		return m.WeakMaster
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

func (m *RegistryEntry) GetMasterTime() int64 {
	if m != nil {
		return m.MasterTime
	}
	return 0
}

func (m *RegistryEntry) GetVersion() RegistryEntry_Version {
	if m != nil {
		return m.Version
	}
	return RegistryEntry_V1
}

type ServiceList struct {
	Services             []*RegistryEntry `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ServiceList) Reset()         { *m = ServiceList{} }
func (m *ServiceList) String() string { return proto.CompactTextString(m) }
func (*ServiceList) ProtoMessage()    {}
func (*ServiceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{1}
}

func (m *ServiceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceList.Unmarshal(m, b)
}
func (m *ServiceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceList.Marshal(b, m, deterministic)
}
func (m *ServiceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceList.Merge(m, src)
}
func (m *ServiceList) XXX_Size() int {
	return xxx_messageInfo_ServiceList.Size(m)
}
func (m *ServiceList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceList.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceList proto.InternalMessageInfo

func (m *ServiceList) GetServices() []*RegistryEntry {
	if m != nil {
		return m.Services
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type StateResponse struct {
	LongestCall          int64    `protobuf:"varint,1,opt,name=longest_call,json=longestCall,proto3" json:"longest_call,omitempty"`
	MostFrequent         string   `protobuf:"bytes,2,opt,name=most_frequent,json=mostFrequent,proto3" json:"most_frequent,omitempty"`
	Frequency            int32    `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Count                string   `protobuf:"bytes,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{3}
}

func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetLongestCall() int64 {
	if m != nil {
		return m.LongestCall
	}
	return 0
}

func (m *StateResponse) GetMostFrequent() string {
	if m != nil {
		return m.MostFrequent
	}
	return ""
}

func (m *StateResponse) GetFrequency() int32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *StateResponse) GetCount() string {
	if m != nil {
		return m.Count
	}
	return ""
}

type StateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRequest) Reset()         { *m = StateRequest{} }
func (m *StateRequest) String() string { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()    {}
func (*StateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{4}
}

func (m *StateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRequest.Unmarshal(m, b)
}
func (m *StateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRequest.Marshal(b, m, deterministic)
}
func (m *StateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequest.Merge(m, src)
}
func (m *StateRequest) XXX_Size() int {
	return xxx_messageInfo_StateRequest.Size(m)
}
func (m *StateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequest proto.InternalMessageInfo

type RegisterRequest struct {
	Service              *RegistryEntry `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Caller               string         `protobuf:"bytes,2,opt,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{5}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetService() *RegistryEntry {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *RegisterRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

type RegisterResponse struct {
	Service              *RegistryEntry `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{6}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetService() *RegistryEntry {
	if m != nil {
		return m.Service
	}
	return nil
}

type DiscoverRequest struct {
	Request              *RegistryEntry `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Caller               string         `protobuf:"bytes,2,opt,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoverRequest) Reset()         { *m = DiscoverRequest{} }
func (m *DiscoverRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoverRequest) ProtoMessage()    {}
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{7}
}

func (m *DiscoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRequest.Unmarshal(m, b)
}
func (m *DiscoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRequest.Marshal(b, m, deterministic)
}
func (m *DiscoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRequest.Merge(m, src)
}
func (m *DiscoverRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoverRequest.Size(m)
}
func (m *DiscoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRequest proto.InternalMessageInfo

func (m *DiscoverRequest) GetRequest() *RegistryEntry {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *DiscoverRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

type DiscoverResponse struct {
	Service              *RegistryEntry `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoverResponse) Reset()         { *m = DiscoverResponse{} }
func (m *DiscoverResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoverResponse) ProtoMessage()    {}
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{8}
}

func (m *DiscoverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverResponse.Unmarshal(m, b)
}
func (m *DiscoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverResponse.Marshal(b, m, deterministic)
}
func (m *DiscoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverResponse.Merge(m, src)
}
func (m *DiscoverResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoverResponse.Size(m)
}
func (m *DiscoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverResponse proto.InternalMessageInfo

func (m *DiscoverResponse) GetService() *RegistryEntry {
	if m != nil {
		return m.Service
	}
	return nil
}

type ListRequest struct {
	Caller               string   `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{9}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

type ListResponse struct {
	Services             *ServiceList `protobuf:"bytes,1,opt,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{10}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() *ServiceList {
	if m != nil {
		return m.Services
	}
	return nil
}

type GetRequest struct {
	Job                  string   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{11}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *GetRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

type GetResponse struct {
	Services             []*RegistryEntry `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{12}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetServices() []*RegistryEntry {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterEnum("discovery.RegistryEntry_Version", RegistryEntry_Version_name, RegistryEntry_Version_value)
	proto.RegisterType((*RegistryEntry)(nil), "discovery.RegistryEntry")
	proto.RegisterType((*ServiceList)(nil), "discovery.ServiceList")
	proto.RegisterType((*Empty)(nil), "discovery.Empty")
	proto.RegisterType((*StateResponse)(nil), "discovery.StateResponse")
	proto.RegisterType((*StateRequest)(nil), "discovery.StateRequest")
	proto.RegisterType((*RegisterRequest)(nil), "discovery.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "discovery.RegisterResponse")
	proto.RegisterType((*DiscoverRequest)(nil), "discovery.DiscoverRequest")
	proto.RegisterType((*DiscoverResponse)(nil), "discovery.DiscoverResponse")
	proto.RegisterType((*ListRequest)(nil), "discovery.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "discovery.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "discovery.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "discovery.GetResponse")
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor_1e7ff60feb39c8d0) }

var fileDescriptor_1e7ff60feb39c8d0 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xae, 0x93, 0xe6, 0x34, 0xb6, 0xd3, 0x68, 0xf5, 0xff, 0xc5, 0x04, 0x04, 0xc6, 0xa5, 0x52,
	0xae, 0x2a, 0x61, 0x50, 0x2f, 0x10, 0x37, 0xd0, 0xd3, 0x4d, 0x2b, 0xa1, 0x6d, 0x95, 0x3b, 0x64,
	0xb9, 0xe9, 0xb4, 0x32, 0x38, 0x76, 0xba, 0xbb, 0x2d, 0xe4, 0x19, 0xb8, 0xe6, 0x29, 0x78, 0x1c,
	0x5e, 0x08, 0xed, 0xc1, 0x87, 0x24, 0x50, 0xa4, 0xf6, 0x2a, 0xbb, 0xdf, 0xcc, 0x7c, 0xf3, 0xcd,
	0xce, 0x67, 0x05, 0x36, 0x2e, 0x12, 0x3e, 0xc9, 0x6f, 0x91, 0xcd, 0x77, 0x66, 0x2c, 0x17, 0x39,
	0xe9, 0x95, 0x40, 0xf0, 0xab, 0x09, 0x2e, 0xc5, 0xab, 0x84, 0x0b, 0x36, 0x3f, 0xc8, 0x04, 0x9b,
	0x93, 0x3e, 0x34, 0x92, 0x99, 0x67, 0xf9, 0xd6, 0xa8, 0x47, 0x1b, 0xc9, 0x8c, 0x10, 0x58, 0x9f,
	0xe5, 0x4c, 0x78, 0x0d, 0xdf, 0x1a, 0xb5, 0xa8, 0x3a, 0x4b, 0x2c, 0x8b, 0xa7, 0xe8, 0x35, 0x55,
	0x96, 0x3a, 0x93, 0x2d, 0x70, 0xf1, 0x9b, 0x40, 0x96, 0xc5, 0x69, 0xa4, 0x0a, 0xd6, 0x7d, 0x6b,
	0xd4, 0xa5, 0x4e, 0x01, 0x7e, 0x94, 0x85, 0xcf, 0x00, 0x92, 0x0b, 0xcc, 0x44, 0x72, 0x99, 0x20,
	0xf3, 0x5a, 0xaa, 0xbc, 0x86, 0x90, 0x4d, 0x68, 0x4f, 0x63, 0x2e, 0x90, 0x79, 0x6d, 0x55, 0x6d,
	0x6e, 0xe4, 0x39, 0xd8, 0x5f, 0x31, 0xfe, 0x12, 0x99, 0xa0, 0xa3, 0x82, 0x20, 0xa1, 0x13, 0x9d,
	0xb0, 0x05, 0x2e, 0x53, 0x63, 0x20, 0x8b, 0x44, 0x32, 0x45, 0xaf, 0xe3, 0x5b, 0xa3, 0x26, 0x75,
	0x0a, 0xf0, 0x2c, 0x99, 0x22, 0x09, 0xc0, 0x95, 0xb1, 0x48, 0xe4, 0xd1, 0x24, 0xc5, 0x38, 0xf3,
	0xba, 0x2a, 0xc9, 0x96, 0xe0, 0x59, 0xbe, 0x27, 0x21, 0xf2, 0x12, 0xfa, 0x69, 0xcc, 0x45, 0xc4,
	0x11, 0x33, 0xcd, 0xd4, 0xd3, 0x4c, 0x12, 0x3d, 0x45, 0xcc, 0x14, 0xd3, 0x36, 0xf4, 0x93, 0xab,
	0x2c, 0x67, 0xc8, 0x0b, 0x49, 0xa0, 0x24, 0xb9, 0x06, 0x3d, 0x29, 0x65, 0xeb, 0xb0, 0x66, 0xb2,
	0x15, 0x13, 0x68, 0x48, 0xf1, 0xbc, 0x85, 0xce, 0x2d, 0x32, 0x9e, 0xe4, 0x99, 0xe7, 0xfa, 0xd6,
	0xa8, 0x1f, 0xfa, 0x3b, 0xd5, 0xb2, 0x16, 0xf6, 0xb2, 0x33, 0xd6, 0x79, 0xb4, 0x28, 0x08, 0x1e,
	0x43, 0xc7, 0x60, 0xa4, 0x0d, 0x8d, 0xf1, 0xab, 0xc1, 0x9a, 0xfa, 0x0d, 0x07, 0x56, 0xb0, 0x07,
	0xf6, 0x29, 0xb2, 0xdb, 0x64, 0x82, 0xc7, 0x09, 0x17, 0xe4, 0x0d, 0x74, 0xb9, 0xbe, 0x72, 0xcf,
	0xf2, 0x9b, 0x23, 0x3b, 0xf4, 0xfe, 0xd6, 0x86, 0x96, 0x99, 0x41, 0x07, 0x5a, 0x07, 0xd3, 0x99,
	0x98, 0x07, 0xdf, 0x2d, 0x70, 0x4f, 0x45, 0x2c, 0x90, 0x22, 0x9f, 0xe5, 0x19, 0x47, 0xf2, 0x02,
	0x9c, 0x34, 0xcf, 0xae, 0x90, 0x8b, 0x68, 0x12, 0xa7, 0xa9, 0x72, 0x4b, 0x93, 0xda, 0x06, 0xdb,
	0x8b, 0xd3, 0x54, 0x2e, 0x64, 0x9a, 0x73, 0x11, 0x5d, 0x32, 0xbc, 0xbe, 0xc1, 0x4c, 0xfb, 0xa7,
	0x47, 0x1d, 0x09, 0x1e, 0x1a, 0x8c, 0x3c, 0x85, 0x9e, 0x89, 0x4f, 0xe6, 0xca, 0x4c, 0x2d, 0x5a,
	0x01, 0xe4, 0x3f, 0x68, 0x4d, 0xf2, 0x9b, 0x4c, 0x3b, 0xa9, 0x47, 0xf5, 0x25, 0xe8, 0x83, 0x63,
	0xc4, 0x5c, 0xdf, 0x20, 0x17, 0xc1, 0x27, 0xd8, 0xa0, 0x66, 0xc9, 0x06, 0x22, 0x21, 0x74, 0xcc,
	0x14, 0x4a, 0xd9, 0x5d, 0xe3, 0x16, 0x89, 0xd2, 0x79, 0x72, 0x14, 0x64, 0x46, 0xa8, 0xb9, 0x05,
	0x87, 0x30, 0xa8, 0xe8, 0xcd, 0xf8, 0xf7, 0xe0, 0x97, 0x32, 0xf7, 0x4d, 0x4e, 0x4d, 0x26, 0xd3,
	0xc7, 0x7f, 0xd3, 0x98, 0xc4, 0xbb, 0x64, 0x56, 0xf4, 0x0f, 0x90, 0xb9, 0x0d, 0xb6, 0xb4, 0x0c,
	0x5d, 0x69, 0x67, 0x2d, 0xb4, 0xfb, 0x00, 0x8e, 0x4e, 0x2b, 0x5b, 0xd5, 0x1d, 0x26, 0x7b, 0x6d,
	0xd6, 0x7a, 0xd5, 0xbc, 0x58, 0xf3, 0xd7, 0x2e, 0xc0, 0x11, 0x96, 0x9d, 0x06, 0xd0, 0xfc, 0x9c,
	0x9f, 0x9b, 0x36, 0xf2, 0x28, 0x7b, 0xcb, 0xdc, 0x6a, 0x54, 0x7d, 0x93, 0xe6, 0x56, 0x75, 0xa6,
	0xf5, 0xbd, 0xcc, 0x1d, 0xfe, 0x6c, 0x54, 0x0f, 0x36, 0x37, 0xfa, 0xc8, 0x71, 0x65, 0xa5, 0x02,
	0x1a, 0xae, 0x70, 0x95, 0xfb, 0x1b, 0x3e, 0xf9, 0x63, 0x4c, 0xcb, 0x0a, 0xd6, 0xc8, 0x01, 0x74,
	0x8b, 0x0e, 0x0b, 0x34, 0x4b, 0x36, 0x58, 0xa0, 0x59, 0xde, 0x61, 0xb0, 0x46, 0xf6, 0x61, 0x43,
	0x3e, 0xdc, 0xfb, 0x34, 0x35, 0x9a, 0x38, 0xa9, 0xbf, 0x6d, 0x6d, 0x5b, 0xc3, 0x47, 0x2b, 0x78,
	0xc9, 0xf2, 0x0e, 0x5a, 0xea, 0xab, 0x21, 0xf5, 0x9c, 0xfa, 0x77, 0x34, 0xf4, 0x56, 0x03, 0x45,
	0x75, 0xf8, 0xc3, 0x02, 0xb2, 0xfc, 0x5a, 0xe3, 0x90, 0x1c, 0x01, 0x14, 0x73, 0x8f, 0xc3, 0x87,
	0x3c, 0xd5, 0x2e, 0x34, 0x8f, 0x50, 0x90, 0xff, 0x6b, 0x59, 0x95, 0x35, 0x86, 0x9b, 0xcb, 0x70,
	0x51, 0x77, 0xde, 0x56, 0xff, 0x67, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x8d, 0x19,
	0x87, 0xe2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	RegisterService(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (*DiscoverResponse, error)
	ListAllServices(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
}

type discoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceClient(cc *grpc.ClientConn) DiscoveryServiceClient {
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
type DiscoveryServiceServer interface {
	RegisterService(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Discover(context.Context, *DiscoverRequest) (*DiscoverResponse, error)
	ListAllServices(context.Context, *ListRequest) (*ListResponse, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
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
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryServiceV2Client interface {
	RegisterV2(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type discoveryServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceV2Client(cc *grpc.ClientConn) DiscoveryServiceV2Client {
	return &discoveryServiceV2Client{cc}
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

// DiscoveryServiceV2Server is the server API for DiscoveryServiceV2 service.
type DiscoveryServiceV2Server interface {
	RegisterV2(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterDiscoveryServiceV2Server(s *grpc.Server, srv DiscoveryServiceV2Server) {
	s.RegisterService(&_DiscoveryServiceV2_serviceDesc, srv)
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

var _DiscoveryServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.DiscoveryServiceV2",
	HandlerType: (*DiscoveryServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterV2",
			Handler:    _DiscoveryServiceV2_RegisterV2_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DiscoveryServiceV2_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adjacency.proto

package adjacency

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Adjacency event type
type AdjacencyEventType int32

const (
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_NO_EVENT      AdjacencyEventType = 0
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_ADD           AdjacencyEventType = 1
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_DELETE        AdjacencyEventType = 2
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_UPDATE        AdjacencyEventType = 3
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_DOWNLOAD      AdjacencyEventType = 4
	AdjacencyEventType_ADJACENCY_EVENT_TYPE_DOWNLOAD_DONE AdjacencyEventType = 5
)

var AdjacencyEventType_name = map[int32]string{
	0: "ADJACENCY_EVENT_TYPE_NO_EVENT",
	1: "ADJACENCY_EVENT_TYPE_ADD",
	2: "ADJACENCY_EVENT_TYPE_DELETE",
	3: "ADJACENCY_EVENT_TYPE_UPDATE",
	4: "ADJACENCY_EVENT_TYPE_DOWNLOAD",
	5: "ADJACENCY_EVENT_TYPE_DOWNLOAD_DONE",
}

var AdjacencyEventType_value = map[string]int32{
	"ADJACENCY_EVENT_TYPE_NO_EVENT":      0,
	"ADJACENCY_EVENT_TYPE_ADD":           1,
	"ADJACENCY_EVENT_TYPE_DELETE":        2,
	"ADJACENCY_EVENT_TYPE_UPDATE":        3,
	"ADJACENCY_EVENT_TYPE_DOWNLOAD":      4,
	"ADJACENCY_EVENT_TYPE_DOWNLOAD_DONE": 5,
}

func (x AdjacencyEventType) String() string {
	return proto.EnumName(AdjacencyEventType_name, int32(x))
}

func (AdjacencyEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2f009b605f0ca8, []int{0}
}

// Adjacency address family
type AdjacencyAddressFamily int32

const (
	AdjacencyAddressFamily_ADJACENCY_AF_IPV4 AdjacencyAddressFamily = 0
	AdjacencyAddressFamily_ADJACENCY_AF_IPV6 AdjacencyAddressFamily = 1
)

var AdjacencyAddressFamily_name = map[int32]string{
	0: "ADJACENCY_AF_IPV4",
	1: "ADJACENCY_AF_IPV6",
}

var AdjacencyAddressFamily_value = map[string]int32{
	"ADJACENCY_AF_IPV4": 0,
	"ADJACENCY_AF_IPV6": 1,
}

func (x AdjacencyAddressFamily) String() string {
	return proto.EnumName(AdjacencyAddressFamily_name, int32(x))
}

func (AdjacencyAddressFamily) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2f009b605f0ca8, []int{1}
}

// Adjacency message
type NxAdjacencyProto struct {
	// IP address
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// MAC address
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Interface name
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Physical interface name
	PhysicalInterfaceName string `protobuf:"bytes,4,opt,name=physical_interface_name,json=physicalInterfaceName,proto3" json:"physical_interface_name,omitempty"`
	// vrf name
	VrfName string `protobuf:"bytes,5,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	// Preference
	Preference uint32 `protobuf:"varint,6,opt,name=preference,proto3" json:"preference,omitempty"`
	// source for the adjacency
	Source string `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	// Address family for the adjacency
	AddressFamily AdjacencyAddressFamily `protobuf:"varint,8,opt,name=address_family,json=addressFamily,proto3,enum=AdjacencyAddressFamily" json:"address_family,omitempty"`
	// Adjacency event type
	EventType AdjacencyEventType `protobuf:"varint,9,opt,name=event_type,json=eventType,proto3,enum=AdjacencyEventType" json:"event_type,omitempty"`
	Timestamp uint64             `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Mac associated with multiple ip.
	Addrlist             []string `protobuf:"bytes,11,rep,name=addrlist,proto3" json:"addrlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NxAdjacencyProto) Reset()         { *m = NxAdjacencyProto{} }
func (m *NxAdjacencyProto) String() string { return proto.CompactTextString(m) }
func (*NxAdjacencyProto) ProtoMessage()    {}
func (*NxAdjacencyProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2f009b605f0ca8, []int{0}
}

func (m *NxAdjacencyProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NxAdjacencyProto.Unmarshal(m, b)
}
func (m *NxAdjacencyProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NxAdjacencyProto.Marshal(b, m, deterministic)
}
func (m *NxAdjacencyProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NxAdjacencyProto.Merge(m, src)
}
func (m *NxAdjacencyProto) XXX_Size() int {
	return xxx_messageInfo_NxAdjacencyProto.Size(m)
}
func (m *NxAdjacencyProto) XXX_DiscardUnknown() {
	xxx_messageInfo_NxAdjacencyProto.DiscardUnknown(m)
}

var xxx_messageInfo_NxAdjacencyProto proto.InternalMessageInfo

func (m *NxAdjacencyProto) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *NxAdjacencyProto) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *NxAdjacencyProto) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *NxAdjacencyProto) GetPhysicalInterfaceName() string {
	if m != nil {
		return m.PhysicalInterfaceName
	}
	return ""
}

func (m *NxAdjacencyProto) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *NxAdjacencyProto) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *NxAdjacencyProto) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *NxAdjacencyProto) GetAddressFamily() AdjacencyAddressFamily {
	if m != nil {
		return m.AddressFamily
	}
	return AdjacencyAddressFamily_ADJACENCY_AF_IPV4
}

func (m *NxAdjacencyProto) GetEventType() AdjacencyEventType {
	if m != nil {
		return m.EventType
	}
	return AdjacencyEventType_ADJACENCY_EVENT_TYPE_NO_EVENT
}

func (m *NxAdjacencyProto) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *NxAdjacencyProto) GetAddrlist() []string {
	if m != nil {
		return m.Addrlist
	}
	return nil
}

func init() {
	proto.RegisterEnum("AdjacencyEventType", AdjacencyEventType_name, AdjacencyEventType_value)
	proto.RegisterEnum("AdjacencyAddressFamily", AdjacencyAddressFamily_name, AdjacencyAddressFamily_value)
	proto.RegisterType((*NxAdjacencyProto)(nil), "NxAdjacencyProto")
}

func init() { proto.RegisterFile("adjacency.proto", fileDescriptor_7b2f009b605f0ca8) }

var fileDescriptor_7b2f009b605f0ca8 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x14, 0x85, 0xe7, 0xb5, 0xeb, 0x9a, 0x3b, 0xb5, 0x18, 0xa3, 0x6d, 0x06, 0x36, 0x16, 0x26, 0x81,
	0xa2, 0x3d, 0xec, 0x61, 0xa0, 0x3d, 0x22, 0x85, 0xc5, 0x95, 0x86, 0xa6, 0xb4, 0x8a, 0xc2, 0xd0,
	0x78, 0xb1, 0x4c, 0x7a, 0x23, 0x82, 0x9a, 0x34, 0x72, 0x42, 0x45, 0xfe, 0x16, 0x3f, 0x89, 0x5f,
	0xc1, 0x23, 0xaa, 0xdb, 0xa6, 0x03, 0xd2, 0xbd, 0xe5, 0x9c, 0xf3, 0x1d, 0xc7, 0xbe, 0xba, 0xf0,
	0x48, 0x8d, 0xbf, 0xa9, 0x08, 0xb3, 0xa8, 0x3a, 0xcf, 0xf5, 0xb4, 0x9c, 0x9e, 0xfe, 0x6c, 0x01,
	0xf5, 0x7f, 0xb8, 0x2b, 0x77, 0x34, 0x37, 0xd9, 0x31, 0x40, 0x92, 0x4b, 0x35, 0x1e, 0x6b, 0x2c,
	0x0a, 0x4e, 0x6c, 0xe2, 0x58, 0x81, 0x95, 0xe4, 0xee, 0xc2, 0x60, 0x27, 0xb0, 0x97, 0xaa, 0xa8,
	0xce, 0xb7, 0x4d, 0x0e, 0xa9, 0x8a, 0x56, 0xc0, 0x2b, 0xe8, 0x27, 0x59, 0x89, 0x3a, 0x56, 0x11,
	0xca, 0x4c, 0xa5, 0xc8, 0x5b, 0x86, 0xe9, 0xd5, 0xae, 0xaf, 0x52, 0x64, 0x97, 0x70, 0x98, 0x7f,
	0xad, 0x8a, 0x24, 0x52, 0x13, 0xf9, 0x0f, 0xdf, 0x36, 0xfc, 0xfe, 0x2a, 0xbe, 0xfe, 0xab, 0xf7,
	0x14, 0xba, 0x33, 0x1d, 0x2f, 0xc0, 0x1d, 0x03, 0xee, 0xce, 0x74, 0x6c, 0xa2, 0x17, 0x00, 0xb9,
	0xc6, 0x18, 0x35, 0x66, 0x11, 0xf2, 0x8e, 0x4d, 0x9c, 0x5e, 0x70, 0xcf, 0x61, 0x07, 0xd0, 0x29,
	0xa6, 0xdf, 0x75, 0x84, 0x7c, 0xd7, 0x14, 0x97, 0x8a, 0xbd, 0x83, 0xfe, 0xf2, 0x39, 0x32, 0x56,
	0x69, 0x32, 0xa9, 0x78, 0xd7, 0x26, 0x4e, 0xff, 0xe2, 0xf0, 0xbc, 0x1e, 0xcd, 0xf2, 0x71, 0x03,
	0x13, 0x07, 0x3d, 0x75, 0x5f, 0xb2, 0x0b, 0x00, 0x9c, 0x61, 0x56, 0xca, 0xb2, 0xca, 0x91, 0x5b,
	0xa6, 0xfb, 0x64, 0xdd, 0x15, 0xf3, 0x2c, 0xac, 0x72, 0x0c, 0x2c, 0x5c, 0x7d, 0xb2, 0x23, 0xb0,
	0xca, 0x24, 0xc5, 0xa2, 0x54, 0x69, 0xce, 0xc1, 0x26, 0x4e, 0x3b, 0x58, 0x1b, 0xec, 0x19, 0x74,
	0xe7, 0xbf, 0x98, 0x24, 0x45, 0xc9, 0xf7, 0xec, 0x96, 0x63, 0x05, 0xb5, 0x3e, 0xfb, 0x45, 0x80,
	0xfd, 0x7f, 0x36, 0x7b, 0x09, 0xc7, 0xae, 0xf7, 0xc1, 0xbd, 0x12, 0xfe, 0xd5, 0x9d, 0x14, 0xb7,
	0xc2, 0x0f, 0x65, 0x78, 0x37, 0x12, 0xd2, 0x1f, 0x2e, 0x14, 0xdd, 0x62, 0x47, 0xc0, 0x1b, 0x11,
	0xd7, 0xf3, 0x28, 0x61, 0x27, 0xf0, 0xbc, 0x31, 0xf5, 0xc4, 0x8d, 0x08, 0x05, 0xdd, 0xde, 0x08,
	0x7c, 0x1c, 0x79, 0x6e, 0x28, 0x68, 0x6b, 0xe3, 0x15, 0xbc, 0xe1, 0x27, 0xff, 0x66, 0xe8, 0x7a,
	0xb4, 0xcd, 0x5e, 0xc3, 0xe9, 0x83, 0x88, 0xf4, 0x86, 0xbe, 0xa0, 0x3b, 0x67, 0x03, 0x38, 0x68,
	0x9e, 0x3d, 0xdb, 0x87, 0xc7, 0xeb, 0x13, 0xdc, 0x81, 0xbc, 0x1e, 0xdd, 0xbe, 0xa5, 0x5b, 0x4d,
	0xf6, 0x25, 0x25, 0xef, 0xfb, 0x9f, 0xad, 0x7a, 0xe9, 0x7f, 0x13, 0xf2, 0xa5, 0x63, 0x16, 0xff,
	0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd2, 0xb4, 0xb2, 0x0b, 0x03, 0x00, 0x00,
}

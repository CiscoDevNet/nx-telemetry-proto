// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mac_all.proto

package mac_all

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

// MAC event types
type MacAllEventType int32

const (
	MacAllEventType_MAC_ALL_EVENT_TYPE_NO_EVENT      MacAllEventType = 0
	MacAllEventType_MAC_ALL_EVENT_TYPE_ADD           MacAllEventType = 1
	MacAllEventType_MAC_ALL_EVENT_TYPE_DELETE        MacAllEventType = 2
	MacAllEventType_MAC_ALL_EVENT_TYPE_UPDATE        MacAllEventType = 3
	MacAllEventType_MAC_ALL_EVENT_TYPE_DOWNLOAD      MacAllEventType = 4
	MacAllEventType_MAC_ALL_EVENT_TYPE_DOWNLOAD_DONE MacAllEventType = 5
)

var MacAllEventType_name = map[int32]string{
	0: "MAC_ALL_EVENT_TYPE_NO_EVENT",
	1: "MAC_ALL_EVENT_TYPE_ADD",
	2: "MAC_ALL_EVENT_TYPE_DELETE",
	3: "MAC_ALL_EVENT_TYPE_UPDATE",
	4: "MAC_ALL_EVENT_TYPE_DOWNLOAD",
	5: "MAC_ALL_EVENT_TYPE_DOWNLOAD_DONE",
}

var MacAllEventType_value = map[string]int32{
	"MAC_ALL_EVENT_TYPE_NO_EVENT":      0,
	"MAC_ALL_EVENT_TYPE_ADD":           1,
	"MAC_ALL_EVENT_TYPE_DELETE":        2,
	"MAC_ALL_EVENT_TYPE_UPDATE":        3,
	"MAC_ALL_EVENT_TYPE_DOWNLOAD":      4,
	"MAC_ALL_EVENT_TYPE_DOWNLOAD_DONE": 5,
}

func (x MacAllEventType) String() string {
	return proto.EnumName(MacAllEventType_name, int32(x))
}

func (MacAllEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{0}
}

// MAC address types
type Type int32

const (
	Type_MAC_ALL_ADDRESS_TYPE_NO_TYPE Type = 0
	Type_MAC_ALL_ADDRESS_TYPE_STATIC  Type = 1
	Type_MAC_ALL_ADDRESS_TYPE_DYNAMIC Type = 2
)

var Type_name = map[int32]string{
	0: "MAC_ALL_ADDRESS_TYPE_NO_TYPE",
	1: "MAC_ALL_ADDRESS_TYPE_STATIC",
	2: "MAC_ALL_ADDRESS_TYPE_DYNAMIC",
}

var Type_value = map[string]int32{
	"MAC_ALL_ADDRESS_TYPE_NO_TYPE": 0,
	"MAC_ALL_ADDRESS_TYPE_STATIC":  1,
	"MAC_ALL_ADDRESS_TYPE_DYNAMIC": 2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{1}
}

type MacL2Type int32

const (
	MacL2Type_MAC_ALL_MAC_L2_TYPE_UNKNOWN          MacL2Type = 0
	MacL2Type_MAC_ALL_MAC_L2_TYPE_PRIMARY          MacL2Type = 1
	MacL2Type_MAC_ALL_MAC_L2_TYPE_GATEWAY          MacL2Type = 2
	MacL2Type_MAC_ALL_MAC_L2_TYPE_OVERLAY          MacL2Type = 4
	MacL2Type_MAC_ALL_MAC_L2_TYPE_PRIMARY_VPC_PEER MacL2Type = 5
	MacL2Type_MAC_ALL_MAC_L2_TYPE_CONTROL_PLANE    MacL2Type = 6
	MacL2Type_MAC_ALL_MAC_L2_TYPE_VSAN             MacL2Type = 7
)

var MacL2Type_name = map[int32]string{
	0: "MAC_ALL_MAC_L2_TYPE_UNKNOWN",
	1: "MAC_ALL_MAC_L2_TYPE_PRIMARY",
	2: "MAC_ALL_MAC_L2_TYPE_GATEWAY",
	4: "MAC_ALL_MAC_L2_TYPE_OVERLAY",
	5: "MAC_ALL_MAC_L2_TYPE_PRIMARY_VPC_PEER",
	6: "MAC_ALL_MAC_L2_TYPE_CONTROL_PLANE",
	7: "MAC_ALL_MAC_L2_TYPE_VSAN",
}

var MacL2Type_value = map[string]int32{
	"MAC_ALL_MAC_L2_TYPE_UNKNOWN":          0,
	"MAC_ALL_MAC_L2_TYPE_PRIMARY":          1,
	"MAC_ALL_MAC_L2_TYPE_GATEWAY":          2,
	"MAC_ALL_MAC_L2_TYPE_OVERLAY":          4,
	"MAC_ALL_MAC_L2_TYPE_PRIMARY_VPC_PEER": 5,
	"MAC_ALL_MAC_L2_TYPE_CONTROL_PLANE":    6,
	"MAC_ALL_MAC_L2_TYPE_VSAN":             7,
}

func (x MacL2Type) String() string {
	return proto.EnumName(MacL2Type_name, int32(x))
}

func (MacL2Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{2}
}

type MacInfo int32

const (
	MacInfo_MAC_ALL_INFO_STANDARD         MacInfo = 0
	MacInfo_MAC_ALL_INFO_DROP             MacInfo = 1
	MacInfo_MAC_ALL_INFO_SVI_DOWN_FLOOD   MacInfo = 2
	MacInfo_MAC_ALL_INFO_SUP_INBAND_CFSOE MacInfo = 3
	MacInfo_MAC_ALL_INFO_VPC_PEER_LINK    MacInfo = 4
	MacInfo_MAC_ALL_INFO_NVE              MacInfo = 5
	MacInfo_MAC_ALL_INFO_SUP_ETH          MacInfo = 6
)

var MacInfo_name = map[int32]string{
	0: "MAC_ALL_INFO_STANDARD",
	1: "MAC_ALL_INFO_DROP",
	2: "MAC_ALL_INFO_SVI_DOWN_FLOOD",
	3: "MAC_ALL_INFO_SUP_INBAND_CFSOE",
	4: "MAC_ALL_INFO_VPC_PEER_LINK",
	5: "MAC_ALL_INFO_NVE",
	6: "MAC_ALL_INFO_SUP_ETH",
}

var MacInfo_value = map[string]int32{
	"MAC_ALL_INFO_STANDARD":         0,
	"MAC_ALL_INFO_DROP":             1,
	"MAC_ALL_INFO_SVI_DOWN_FLOOD":   2,
	"MAC_ALL_INFO_SUP_INBAND_CFSOE": 3,
	"MAC_ALL_INFO_VPC_PEER_LINK":    4,
	"MAC_ALL_INFO_NVE":              5,
	"MAC_ALL_INFO_SUP_ETH":          6,
}

func (x MacInfo) String() string {
	return proto.EnumName(MacInfo_name, int32(x))
}

func (MacInfo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{3}
}

// MAC message
type Mac struct {
	// Age - seconds since last seen
	Age uint32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	// Additional Information on the MAC
	MacInfo MacInfo `protobuf:"varint,2,opt,name=mac_info,json=macInfo,proto3,enum=mac_all.MacInfo" json:"mac_info,omitempty"`
	// NTFY
	Ntfy bool `protobuf:"varint,3,opt,name=ntfy,proto3" json:"ntfy,omitempty"`
	// Port -
	//Must match first field in the output of show intf brief.
	//Example: Eth1/1 or Vlan100
	Port string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	// Routed or not
	Routed bool `protobuf:"varint,5,opt,name=routed,proto3" json:"routed,omitempty"`
	// Secure or not
	Secure bool `protobuf:"varint,6,opt,name=secure,proto3" json:"secure,omitempty"`
	// Static or Dynamic
	MacType Type `protobuf:"varint,7,opt,name=mac_type,json=macType,proto3,enum=mac_all.Type" json:"mac_type,omitempty"`
	// MAC L2 Type
	L2Type MacL2Type `protobuf:"varint,8,opt,name=l2_type,json=l2Type,proto3,enum=mac_all.MacL2Type" json:"l2_type,omitempty"`
	// MAC address
	MacAddress string `protobuf:"bytes,9,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// VLAN number
	Vlan uint32 `protobuf:"varint,10,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// Event types
	EventType            MacAllEventType `protobuf:"varint,11,opt,name=event_type,json=eventType,proto3,enum=mac_all.MacAllEventType" json:"event_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Mac) Reset()         { *m = Mac{} }
func (m *Mac) String() string { return proto.CompactTextString(m) }
func (*Mac) ProtoMessage()    {}
func (*Mac) Descriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{0}
}

func (m *Mac) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mac.Unmarshal(m, b)
}
func (m *Mac) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mac.Marshal(b, m, deterministic)
}
func (m *Mac) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mac.Merge(m, src)
}
func (m *Mac) XXX_Size() int {
	return xxx_messageInfo_Mac.Size(m)
}
func (m *Mac) XXX_DiscardUnknown() {
	xxx_messageInfo_Mac.DiscardUnknown(m)
}

var xxx_messageInfo_Mac proto.InternalMessageInfo

func (m *Mac) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Mac) GetMacInfo() MacInfo {
	if m != nil {
		return m.MacInfo
	}
	return MacInfo_MAC_ALL_INFO_STANDARD
}

func (m *Mac) GetNtfy() bool {
	if m != nil {
		return m.Ntfy
	}
	return false
}

func (m *Mac) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Mac) GetRouted() bool {
	if m != nil {
		return m.Routed
	}
	return false
}

func (m *Mac) GetSecure() bool {
	if m != nil {
		return m.Secure
	}
	return false
}

func (m *Mac) GetMacType() Type {
	if m != nil {
		return m.MacType
	}
	return Type_MAC_ALL_ADDRESS_TYPE_NO_TYPE
}

func (m *Mac) GetL2Type() MacL2Type {
	if m != nil {
		return m.L2Type
	}
	return MacL2Type_MAC_ALL_MAC_L2_TYPE_UNKNOWN
}

func (m *Mac) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Mac) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *Mac) GetEventType() MacAllEventType {
	if m != nil {
		return m.EventType
	}
	return MacAllEventType_MAC_ALL_EVENT_TYPE_NO_EVENT
}

type Macall struct {
	List                 []*MacallList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Macall) Reset()         { *m = Macall{} }
func (m *Macall) String() string { return proto.CompactTextString(m) }
func (*Macall) ProtoMessage()    {}
func (*Macall) Descriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{1}
}

func (m *Macall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macall.Unmarshal(m, b)
}
func (m *Macall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macall.Marshal(b, m, deterministic)
}
func (m *Macall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macall.Merge(m, src)
}
func (m *Macall) XXX_Size() int {
	return xxx_messageInfo_Macall.Size(m)
}
func (m *Macall) XXX_DiscardUnknown() {
	xxx_messageInfo_Macall.DiscardUnknown(m)
}

var xxx_messageInfo_Macall proto.InternalMessageInfo

func (m *Macall) GetList() []*MacallList {
	if m != nil {
		return m.List
	}
	return nil
}

type MacallList struct {
	VlanId               uint32   `protobuf:"varint,1,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	Mac                  string   `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`
	Value                *Mac     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MacallList) Reset()         { *m = MacallList{} }
func (m *MacallList) String() string { return proto.CompactTextString(m) }
func (*MacallList) ProtoMessage()    {}
func (*MacallList) Descriptor() ([]byte, []int) {
	return fileDescriptor_965b3496b765a6cc, []int{2}
}

func (m *MacallList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MacallList.Unmarshal(m, b)
}
func (m *MacallList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MacallList.Marshal(b, m, deterministic)
}
func (m *MacallList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MacallList.Merge(m, src)
}
func (m *MacallList) XXX_Size() int {
	return xxx_messageInfo_MacallList.Size(m)
}
func (m *MacallList) XXX_DiscardUnknown() {
	xxx_messageInfo_MacallList.DiscardUnknown(m)
}

var xxx_messageInfo_MacallList proto.InternalMessageInfo

func (m *MacallList) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *MacallList) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *MacallList) GetValue() *Mac {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("mac_all.MacAllEventType", MacAllEventType_name, MacAllEventType_value)
	proto.RegisterEnum("mac_all.Type", Type_name, Type_value)
	proto.RegisterEnum("mac_all.MacL2Type", MacL2Type_name, MacL2Type_value)
	proto.RegisterEnum("mac_all.MacInfo", MacInfo_name, MacInfo_value)
	proto.RegisterType((*Mac)(nil), "mac_all.Mac")
	proto.RegisterType((*Macall)(nil), "mac_all.Macall")
	proto.RegisterType((*MacallList)(nil), "mac_all.MacallList")
}

func init() { proto.RegisterFile("mac_all.proto", fileDescriptor_965b3496b765a6cc) }

var fileDescriptor_965b3496b765a6cc = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdf, 0x4e, 0xdb, 0x4a,
	0x10, 0xc6, 0xd9, 0xfc, 0x71, 0xc8, 0xe4, 0x70, 0xce, 0x9e, 0x2d, 0x50, 0x43, 0xa1, 0x84, 0x88,
	0xaa, 0x51, 0x90, 0x90, 0x9a, 0x5e, 0xf4, 0x7a, 0xc9, 0x2e, 0xad, 0x85, 0xb3, 0xb6, 0x36, 0x26,
	0x28, 0xea, 0xc5, 0xca, 0x4d, 0x4c, 0x15, 0xd5, 0x24, 0x51, 0x62, 0x90, 0x78, 0xca, 0x5e, 0xf4,
	0x09, 0xfa, 0x16, 0xbd, 0xaa, 0xaa, 0x5d, 0x27, 0x34, 0x6e, 0x13, 0xae, 0x76, 0x76, 0xe6, 0xe7,
	0xf9, 0xbe, 0xf1, 0x58, 0x86, 0xad, 0xdb, 0xb0, 0xaf, 0xc2, 0x38, 0x3e, 0x9b, 0x4c, 0xc7, 0xc9,
	0x98, 0x94, 0xe6, 0xd7, 0xda, 0xf7, 0x1c, 0xe4, 0xdb, 0x61, 0x9f, 0x60, 0xc8, 0x87, 0x9f, 0x23,
	0x1b, 0x55, 0x51, 0x7d, 0x4b, 0xea, 0x90, 0x9c, 0xc2, 0xa6, 0x86, 0x86, 0xa3, 0x9b, 0xb1, 0x9d,
	0xab, 0xa2, 0xfa, 0xbf, 0x4d, 0x7c, 0xb6, 0x68, 0xd2, 0x0e, 0xfb, 0xce, 0xe8, 0x66, 0x2c, 0x75,
	0x1b, 0x1d, 0x10, 0x02, 0x85, 0x51, 0x72, 0xf3, 0x60, 0xe7, 0xab, 0xa8, 0xbe, 0x29, 0x4d, 0xac,
	0x73, 0x93, 0xf1, 0x34, 0xb1, 0x0b, 0x55, 0x54, 0x2f, 0x4b, 0x13, 0x93, 0x5d, 0xb0, 0xa6, 0xe3,
	0xbb, 0x24, 0x1a, 0xd8, 0x45, 0x43, 0xce, 0x6f, 0x3a, 0x3f, 0x8b, 0xfa, 0x77, 0xd3, 0xc8, 0xb6,
	0xd2, 0x7c, 0x7a, 0x23, 0xf5, 0xd4, 0x44, 0xf2, 0x30, 0x89, 0xec, 0x92, 0x31, 0xb1, 0xf5, 0x68,
	0x22, 0x78, 0x98, 0x44, 0xc6, 0x81, 0x0e, 0xc8, 0x29, 0x94, 0xe2, 0x66, 0x0a, 0x6e, 0x1a, 0x90,
	0x2c, 0xbb, 0x75, 0x9b, 0x86, 0xb6, 0x62, 0x73, 0x92, 0x23, 0xa8, 0x98, 0xe2, 0x60, 0x30, 0x8d,
	0x66, 0x33, 0xbb, 0x6c, 0x1c, 0xc2, 0x6d, 0xd8, 0xa7, 0x69, 0x46, 0x7b, 0xbf, 0x8f, 0xc3, 0x91,
	0x0d, 0xe6, 0x7d, 0x98, 0x98, 0xbc, 0x03, 0x88, 0xee, 0xa3, 0x51, 0x92, 0x8a, 0x54, 0x8c, 0x88,
	0xbd, 0x2c, 0x42, 0xe3, 0x98, 0x6b, 0xc0, 0x48, 0x95, 0xa3, 0x45, 0x58, 0x7b, 0x03, 0x56, 0x3b,
	0xec, 0x87, 0x71, 0x4c, 0x5e, 0x43, 0x21, 0x1e, 0xce, 0x12, 0x1b, 0x55, 0xf3, 0xf5, 0x4a, 0xf3,
	0xd9, 0xf2, 0xc3, 0x61, 0x1c, 0xbb, 0xc3, 0x59, 0x22, 0x0d, 0x50, 0xfb, 0x08, 0xf0, 0x3b, 0x47,
	0x9e, 0x43, 0x49, 0x3b, 0x50, 0xc3, 0xc1, 0x7c, 0x41, 0x96, 0xbe, 0x3a, 0x03, 0xbd, 0xb5, 0xdb,
	0xb0, 0x6f, 0xd6, 0x53, 0x96, 0x3a, 0x24, 0x35, 0x28, 0xde, 0x87, 0xf1, 0x5d, 0x64, 0x36, 0x51,
	0x69, 0xfe, 0xb3, 0x2c, 0x21, 0xd3, 0x52, 0xe3, 0x1b, 0x82, 0xff, 0xfe, 0xb0, 0x4b, 0x8e, 0xe0,
	0x45, 0x9b, 0xb6, 0x14, 0x75, 0x5d, 0xc5, 0xbb, 0x5c, 0x04, 0x2a, 0xe8, 0xf9, 0x5c, 0x09, 0x2f,
	0xbd, 0xe1, 0x0d, 0xb2, 0x0f, 0xbb, 0x2b, 0x00, 0xca, 0x18, 0x46, 0xe4, 0x10, 0xf6, 0x56, 0xd4,
	0x18, 0x77, 0x79, 0xc0, 0x71, 0x6e, 0x4d, 0xf9, 0xca, 0x67, 0x34, 0xe0, 0x38, 0xbf, 0x46, 0x9a,
	0x79, 0xd7, 0xc2, 0xf5, 0x28, 0xc3, 0x05, 0x72, 0x02, 0xd5, 0x27, 0x00, 0xc5, 0x3c, 0xc1, 0x71,
	0xb1, 0xf1, 0x05, 0x0a, 0x66, 0x92, 0x2a, 0x1c, 0x2c, 0x68, 0xca, 0x98, 0xe4, 0x9d, 0xce, 0xe3,
	0x2c, 0xfa, 0xc4, 0x1b, 0xcb, 0x82, 0x19, 0xa2, 0x13, 0xd0, 0xc0, 0x69, 0x61, 0xb4, 0xb6, 0x05,
	0xeb, 0x09, 0xda, 0x76, 0x5a, 0x38, 0xd7, 0xf8, 0x89, 0xa0, 0xfc, 0xf8, 0x59, 0x2d, 0x37, 0xd4,
	0xa7, 0xdb, 0x9c, 0x4f, 0x28, 0x2e, 0x85, 0x77, 0x2d, 0xb2, 0x8a, 0xcb, 0x80, 0x2f, 0x9d, 0x36,
	0x95, 0x3d, 0x8c, 0xd6, 0x01, 0xef, 0x69, 0xc0, 0xaf, 0x69, 0x0f, 0xe7, 0xd6, 0x01, 0x5e, 0x97,
	0x4b, 0x97, 0xf6, 0x70, 0x81, 0xd4, 0xe1, 0xe4, 0x09, 0x09, 0xd5, 0xf5, 0x5b, 0xca, 0xe7, 0x5c,
	0xe2, 0x22, 0x79, 0x05, 0xc7, 0xab, 0xc8, 0x96, 0x27, 0x02, 0xe9, 0xb9, 0xca, 0x77, 0xa9, 0xe0,
	0xd8, 0x22, 0x07, 0x60, 0xaf, 0xc2, 0xba, 0x1d, 0x2a, 0x70, 0xa9, 0xf1, 0x15, 0x41, 0x69, 0xfe,
	0x17, 0x20, 0x7b, 0xb0, 0xb3, 0x20, 0x1d, 0x71, 0xe1, 0xe9, 0xf7, 0x28, 0x18, 0x95, 0x0c, 0x6f,
	0x90, 0x1d, 0xf8, 0x3f, 0x53, 0x62, 0xd2, 0xf3, 0xb3, 0xe3, 0xa6, 0x4f, 0x74, 0x1d, 0xb3, 0x4f,
	0x75, 0xe1, 0x7a, 0x1e, 0xc3, 0x39, 0x72, 0x0c, 0x87, 0x59, 0xe0, 0xca, 0x57, 0x8e, 0x38, 0xa7,
	0x82, 0xa9, 0xd6, 0x45, 0xc7, 0xd3, 0x9f, 0xcd, 0x4b, 0xd8, 0xcf, 0x20, 0x8b, 0x09, 0x95, 0xeb,
	0x88, 0x4b, 0x5c, 0x20, 0xdb, 0x80, 0x33, 0x75, 0xd1, 0xe5, 0xb8, 0x48, 0x6c, 0xd8, 0xfe, 0xab,
	0x31, 0x0f, 0x3e, 0x60, 0xeb, 0x3c, 0xff, 0x03, 0xa1, 0x4f, 0x96, 0xf9, 0x3d, 0xbe, 0xfd, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x63, 0x07, 0xab, 0x2f, 0x05, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alliancecontract.proto

package tfc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AllianceState int32

const (
	AllianceState_ACTIVE    AllianceState = 0
	AllianceState_COMPLETED AllianceState = 1
	AllianceState_FAILED    AllianceState = 2
)

var AllianceState_name = map[int32]string{
	0: "ACTIVE",
	1: "COMPLETED",
	2: "FAILED",
}
var AllianceState_value = map[string]int32{
	"ACTIVE":    0,
	"COMPLETED": 1,
	"FAILED":    2,
}

func (x AllianceState) String() string {
	return proto.EnumName(AllianceState_name, int32(x))
}
func (AllianceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_alliancecontract_25dcd2a5c6697941, []int{0}
}

type AllianceTrxType int32

const (
	AllianceTrxType_INIT   AllianceTrxType = 0
	AllianceTrxType_INVOKE AllianceTrxType = 1
)

var AllianceTrxType_name = map[int32]string{
	0: "INIT",
	1: "INVOKE",
}
var AllianceTrxType_value = map[string]int32{
	"INIT":   0,
	"INVOKE": 1,
}

func (x AllianceTrxType) String() string {
	return proto.EnumName(AllianceTrxType_name, int32(x))
}
func (AllianceTrxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_alliancecontract_25dcd2a5c6697941, []int{1}
}

type AllianceData struct {
	Terms                []*GameContractTrxArgs `protobuf:"bytes,1,rep,name=terms,proto3" json:"terms,omitempty"`
	Lifespan             int32                  `protobuf:"varint,2,opt,name=lifespan,proto3" json:"lifespan,omitempty"`
	StartGameState       GameState              `protobuf:"varint,3,opt,name=startGameState,proto3,enum=tfc.GameState" json:"startGameState,omitempty"`
	State                AllianceState          `protobuf:"varint,4,opt,name=state,proto3,enum=tfc.AllianceState" json:"state,omitempty"`
	Allies               []Player               `protobuf:"varint,5,rep,packed,name=allies,proto3,enum=tfc.Player" json:"allies,omitempty"`
	LastUpdated          *timestamp.Timestamp   `protobuf:"bytes,9,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	ContractID           uint32                 `protobuf:"varint,10,opt,name=contractID,proto3" json:"contractID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AllianceData) Reset()         { *m = AllianceData{} }
func (m *AllianceData) String() string { return proto.CompactTextString(m) }
func (*AllianceData) ProtoMessage()    {}
func (*AllianceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_alliancecontract_25dcd2a5c6697941, []int{0}
}
func (m *AllianceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllianceData.Unmarshal(m, b)
}
func (m *AllianceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllianceData.Marshal(b, m, deterministic)
}
func (dst *AllianceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllianceData.Merge(dst, src)
}
func (m *AllianceData) XXX_Size() int {
	return xxx_messageInfo_AllianceData.Size(m)
}
func (m *AllianceData) XXX_DiscardUnknown() {
	xxx_messageInfo_AllianceData.DiscardUnknown(m)
}

var xxx_messageInfo_AllianceData proto.InternalMessageInfo

func (m *AllianceData) GetTerms() []*GameContractTrxArgs {
	if m != nil {
		return m.Terms
	}
	return nil
}

func (m *AllianceData) GetLifespan() int32 {
	if m != nil {
		return m.Lifespan
	}
	return 0
}

func (m *AllianceData) GetStartGameState() GameState {
	if m != nil {
		return m.StartGameState
	}
	return GameState_JOINING
}

func (m *AllianceData) GetState() AllianceState {
	if m != nil {
		return m.State
	}
	return AllianceState_ACTIVE
}

func (m *AllianceData) GetAllies() []Player {
	if m != nil {
		return m.Allies
	}
	return nil
}

func (m *AllianceData) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *AllianceData) GetContractID() uint32 {
	if m != nil {
		return m.ContractID
	}
	return 0
}

type AllianceTrxArgs struct {
	Type                 AllianceTrxType `protobuf:"varint,1,opt,name=type,proto3,enum=tfc.AllianceTrxType" json:"type,omitempty"`
	Data                 *AllianceData   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AllianceTrxArgs) Reset()         { *m = AllianceTrxArgs{} }
func (m *AllianceTrxArgs) String() string { return proto.CompactTextString(m) }
func (*AllianceTrxArgs) ProtoMessage()    {}
func (*AllianceTrxArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_alliancecontract_25dcd2a5c6697941, []int{1}
}
func (m *AllianceTrxArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllianceTrxArgs.Unmarshal(m, b)
}
func (m *AllianceTrxArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllianceTrxArgs.Marshal(b, m, deterministic)
}
func (dst *AllianceTrxArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllianceTrxArgs.Merge(dst, src)
}
func (m *AllianceTrxArgs) XXX_Size() int {
	return xxx_messageInfo_AllianceTrxArgs.Size(m)
}
func (m *AllianceTrxArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AllianceTrxArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AllianceTrxArgs proto.InternalMessageInfo

func (m *AllianceTrxArgs) GetType() AllianceTrxType {
	if m != nil {
		return m.Type
	}
	return AllianceTrxType_INIT
}

func (m *AllianceTrxArgs) GetData() *AllianceData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*AllianceData)(nil), "tfc.AllianceData")
	proto.RegisterType((*AllianceTrxArgs)(nil), "tfc.AllianceTrxArgs")
	proto.RegisterEnum("tfc.AllianceState", AllianceState_name, AllianceState_value)
	proto.RegisterEnum("tfc.AllianceTrxType", AllianceTrxType_name, AllianceTrxType_value)
}

func init() {
	proto.RegisterFile("alliancecontract.proto", fileDescriptor_alliancecontract_25dcd2a5c6697941)
}

var fileDescriptor_alliancecontract_25dcd2a5c6697941 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0x41, 0xef, 0x9a, 0x40,
	0x10, 0xc5, 0x8b, 0x82, 0xf9, 0xff, 0x07, 0xa5, 0x76, 0xd3, 0x34, 0x84, 0x43, 0x6b, 0x6c, 0x9a,
	0x1a, 0x0f, 0x98, 0xd0, 0xc4, 0x5b, 0x0f, 0x44, 0x69, 0x43, 0x6a, 0xd5, 0x6c, 0xa9, 0xd7, 0x66,
	0xc5, 0x95, 0x98, 0x80, 0x10, 0x58, 0x93, 0xfa, 0xed, 0xfa, 0xd1, 0x3a, 0xbb, 0x80, 0x29, 0xbd,
	0x31, 0x33, 0xbf, 0x37, 0xf3, 0x78, 0x0b, 0x6f, 0x58, 0x9a, 0x5e, 0xd8, 0x35, 0xe6, 0x71, 0x7e,
	0x15, 0x25, 0x8b, 0x85, 0x5b, 0x94, 0xb9, 0xc8, 0x49, 0x5f, 0x9c, 0x63, 0xe7, 0x5d, 0x92, 0xe7,
	0x49, 0xca, 0x17, 0xaa, 0x75, 0xbc, 0x9d, 0x17, 0xe2, 0x92, 0xf1, 0x4a, 0xb0, 0xac, 0xa8, 0x29,
	0xc7, 0x4a, 0x58, 0xc6, 0x4f, 0x4c, 0xb0, 0xa6, 0x26, 0xb2, 0xee, 0x6e, 0x9a, 0xfe, 0xe9, 0xc1,
	0xd0, 0x6f, 0x8e, 0xac, 0x11, 0x25, 0x2e, 0x18, 0x82, 0x97, 0x59, 0x65, 0x6b, 0x93, 0xfe, 0xcc,
	0xf4, 0x6c, 0x17, 0x4f, 0xb9, 0x5f, 0x51, 0xb8, 0x6a, 0x84, 0x51, 0xf9, 0xdb, 0x2f, 0x93, 0x8a,
	0xd6, 0x18, 0x71, 0xe0, 0x29, 0xbd, 0x9c, 0x79, 0x55, 0xb0, 0xab, 0xdd, 0x9b, 0x68, 0x33, 0x83,
	0x3e, 0x6a, 0xb2, 0x04, 0x0b, 0xfd, 0x94, 0x42, 0xca, 0x7f, 0x08, 0x26, 0xb8, 0xdd, 0x47, 0xc2,
	0xf2, 0xac, 0xc7, 0x52, 0xd5, 0xa5, 0xff, 0x51, 0x64, 0x06, 0x46, 0xa5, 0x70, 0x5d, 0xe1, 0x44,
	0xe1, 0xad, 0xcb, 0x5a, 0x52, 0x03, 0xe4, 0x3d, 0x0c, 0x64, 0x44, 0xbc, 0xb2, 0x0d, 0xb4, 0x6b,
	0x79, 0xa6, 0x42, 0xf7, 0x29, 0xbb, 0xf3, 0x92, 0x36, 0x23, 0xf2, 0x19, 0x86, 0x29, 0xab, 0xc4,
	0xaf, 0x5b, 0x81, 0x61, 0xf0, 0x93, 0xfd, 0x8c, 0x5b, 0x4d, 0xcf, 0x71, 0xeb, 0xfc, 0xdc, 0x36,
	0x3f, 0x37, 0x6a, 0xf3, 0xa3, 0xa6, 0xe4, 0x7f, 0xd6, 0x38, 0x79, 0x0b, 0xd0, 0x86, 0x16, 0xae,
	0x6d, 0x40, 0xf1, 0x88, 0xfe, 0xd3, 0x99, 0x1e, 0xe1, 0x65, 0xeb, 0xad, 0xc9, 0x06, 0x7f, 0x40,
	0x17, 0xf7, 0x82, 0x63, 0x86, 0xd2, 0xff, 0xeb, 0x8e, 0x7f, 0x64, 0x22, 0x9c, 0x51, 0x45, 0x90,
	0x0f, 0xa0, 0xcb, 0x17, 0x52, 0xd1, 0x99, 0xde, 0xab, 0x0e, 0x29, 0xdf, 0x83, 0xaa, 0xf1, 0x7c,
	0x09, 0xa3, 0xce, 0xff, 0x13, 0x80, 0x81, 0xbf, 0x8a, 0xc2, 0x43, 0x30, 0x7e, 0x41, 0x46, 0xf0,
	0xbc, 0xda, 0x7d, 0xdf, 0x6f, 0x82, 0x28, 0x58, 0x8f, 0x35, 0x39, 0xfa, 0xe2, 0x87, 0x1b, 0xfc,
	0xee, 0xcd, 0x3f, 0x76, 0xbc, 0xc9, 0xbb, 0xe4, 0x09, 0xf4, 0x70, 0x1b, 0x46, 0xa8, 0x43, 0x30,
	0xdc, 0x1e, 0x76, 0xdf, 0x82, 0xb1, 0x76, 0x1c, 0xa8, 0x14, 0x3e, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0xe4, 0x98, 0x83, 0x72, 0x02, 0x00, 0x00,
}

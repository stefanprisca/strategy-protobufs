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
	return fileDescriptor_alliancecontract_4c30c3d2364c8189, []int{0}
}

type AllianceData struct {
	Terms                []*GameContractTrxArgs `protobuf:"bytes,1,rep,name=terms,proto3" json:"terms,omitempty"`
	Lifespan             int32                  `protobuf:"varint,2,opt,name=lifespan,proto3" json:"lifespan,omitempty"`
	StartGameState       GameState              `protobuf:"varint,3,opt,name=startGameState,proto3,enum=tfc.GameState" json:"startGameState,omitempty"`
	State                AllianceState          `protobuf:"varint,4,opt,name=state,proto3,enum=tfc.AllianceState" json:"state,omitempty"`
	LastUpdated          *timestamp.Timestamp   `protobuf:"bytes,9,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AllianceData) Reset()         { *m = AllianceData{} }
func (m *AllianceData) String() string { return proto.CompactTextString(m) }
func (*AllianceData) ProtoMessage()    {}
func (*AllianceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_alliancecontract_4c30c3d2364c8189, []int{0}
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

func (m *AllianceData) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func init() {
	proto.RegisterType((*AllianceData)(nil), "tfc.AllianceData")
	proto.RegisterEnum("tfc.AllianceState", AllianceState_name, AllianceState_value)
}

func init() {
	proto.RegisterFile("alliancecontract.proto", fileDescriptor_alliancecontract_4c30c3d2364c8189)
}

var fileDescriptor_alliancecontract_4c30c3d2364c8189 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x6b, 0x83, 0x40,
	0x10, 0xc5, 0x6b, 0xac, 0xa1, 0x59, 0x13, 0x09, 0x7b, 0x28, 0xe2, 0xa5, 0xa1, 0x27, 0xe9, 0x61,
	0x03, 0x16, 0x72, 0xeb, 0x41, 0x8c, 0x2d, 0x81, 0x94, 0x16, 0x6b, 0x7b, 0x2d, 0x1b, 0xb3, 0x8a,
	0xe0, 0x3f, 0x76, 0x27, 0xd0, 0xef, 0xde, 0x4b, 0xd7, 0x5d, 0x0d, 0xa4, 0x37, 0xdf, 0xf8, 0x7b,
	0x6f, 0x66, 0x1f, 0xba, 0xa5, 0x55, 0x55, 0xd2, 0x26, 0x63, 0x59, 0xdb, 0x00, 0xa7, 0x19, 0x90,
	0x8e, 0xb7, 0xd0, 0x62, 0x13, 0xf2, 0xcc, 0xbb, 0x2b, 0xda, 0xb6, 0xa8, 0xd8, 0x5a, 0x8d, 0x0e,
	0xa7, 0x7c, 0x0d, 0x65, 0xcd, 0x04, 0xd0, 0xba, 0xd3, 0x94, 0xe7, 0x14, 0xb4, 0x66, 0x47, 0x0a,
	0x74, 0xd0, 0xb8, 0xd7, 0x97, 0x49, 0xf7, 0xbf, 0x06, 0x9a, 0x87, 0xc3, 0x92, 0xad, 0x44, 0x31,
	0x41, 0x16, 0x30, 0x5e, 0x0b, 0xd7, 0x58, 0x99, 0xbe, 0x1d, 0xb8, 0x44, 0xae, 0x22, 0x2f, 0xd2,
	0x18, 0x0d, 0xc6, 0x94, 0xff, 0x84, 0xbc, 0x10, 0x89, 0xc6, 0xb0, 0x87, 0x6e, 0xaa, 0x32, 0x67,
	0xa2, 0xa3, 0x8d, 0x3b, 0x59, 0x19, 0xbe, 0x95, 0x9c, 0x35, 0xde, 0x20, 0x47, 0xde, 0xc3, 0xa1,
	0xb7, 0x7f, 0x00, 0x05, 0xe6, 0x9a, 0x92, 0x70, 0x02, 0xe7, 0x1c, 0xaa, 0xa6, 0xc9, 0x3f, 0x0a,
	0xfb, 0xc8, 0x12, 0x0a, 0xbf, 0x56, 0x38, 0x56, 0xf8, 0x78, 0xa5, 0xb6, 0x68, 0x00, 0x3f, 0xa1,
	0x79, 0x45, 0x05, 0x7c, 0x9f, 0x3a, 0xf9, 0x4e, 0x76, 0x74, 0x67, 0xd2, 0x60, 0x07, 0x1e, 0xd1,
	0xd5, 0x90, 0xb1, 0x1a, 0x92, 0x8e, 0xd5, 0x24, 0x76, 0xcf, 0x7f, 0x6a, 0xfc, 0x61, 0x83, 0x16,
	0x17, 0xb1, 0x18, 0xa1, 0x69, 0x18, 0xa5, 0xbb, 0xaf, 0x78, 0x79, 0x85, 0x17, 0x68, 0x16, 0xbd,
	0xbd, 0xbe, 0xef, 0xe3, 0x34, 0xde, 0x2e, 0x8d, 0xfe, 0xd7, 0x73, 0xb8, 0xdb, 0xcb, 0xef, 0xc9,
	0x61, 0xaa, 0x82, 0x1f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x81, 0x5b, 0xb1, 0x23, 0xa0, 0x01,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tictactoe.proto

package tictactoe

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

type Mark int32

const (
	Mark_X Mark = 0
	Mark_O Mark = 1
	Mark_E Mark = 2
)

var Mark_name = map[int32]string{
	0: "X",
	1: "O",
	2: "E",
}
var Mark_value = map[string]int32{
	"X": 0,
	"O": 1,
	"E": 2,
}

func (x Mark) String() string {
	return proto.EnumName(Mark_name, int32(x))
}
func (Mark) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{0}
}

type TrxType int32

const (
	TrxType_MOVE TrxType = 0
)

var TrxType_name = map[int32]string{
	0: "MOVE",
}
var TrxType_value = map[string]int32{
	"MOVE": 0,
}

func (x TrxType) String() string {
	return proto.EnumName(TrxType_name, int32(x))
}
func (TrxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{1}
}

type TttContract_State int32

const (
	TttContract_XTURN TttContract_State = 0
	TttContract_OTURN TttContract_State = 1
	TttContract_XWON  TttContract_State = 2
	TttContract_OWON  TttContract_State = 3
	TttContract_TIE   TttContract_State = 4
)

var TttContract_State_name = map[int32]string{
	0: "XTURN",
	1: "OTURN",
	2: "XWON",
	3: "OWON",
	4: "TIE",
}
var TttContract_State_value = map[string]int32{
	"XTURN": 0,
	"OTURN": 1,
	"XWON":  2,
	"OWON":  3,
	"TIE":   4,
}

func (x TttContract_State) String() string {
	return proto.EnumName(TttContract_State_name, int32(x))
}
func (TttContract_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{0, 0}
}

type TttContract struct {
	Positions            []Mark               `protobuf:"varint,1,rep,packed,name=positions,proto3,enum=tictactoe.Mark" json:"positions,omitempty"`
	State                TttContract_State    `protobuf:"varint,2,opt,name=state,proto3,enum=tictactoe.TttContract_State" json:"state,omitempty"`
	XPlayer              string               `protobuf:"bytes,3,opt,name=x_player,json=xPlayer,proto3" json:"x_player,omitempty"`
	OPlayer              string               `protobuf:"bytes,4,opt,name=o_player,json=oPlayer,proto3" json:"o_player,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TttContract) Reset()         { *m = TttContract{} }
func (m *TttContract) String() string { return proto.CompactTextString(m) }
func (*TttContract) ProtoMessage()    {}
func (*TttContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{0}
}
func (m *TttContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TttContract.Unmarshal(m, b)
}
func (m *TttContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TttContract.Marshal(b, m, deterministic)
}
func (dst *TttContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TttContract.Merge(dst, src)
}
func (m *TttContract) XXX_Size() int {
	return xxx_messageInfo_TttContract.Size(m)
}
func (m *TttContract) XXX_DiscardUnknown() {
	xxx_messageInfo_TttContract.DiscardUnknown(m)
}

var xxx_messageInfo_TttContract proto.InternalMessageInfo

func (m *TttContract) GetPositions() []Mark {
	if m != nil {
		return m.Positions
	}
	return nil
}

func (m *TttContract) GetState() TttContract_State {
	if m != nil {
		return m.State
	}
	return TttContract_XTURN
}

func (m *TttContract) GetXPlayer() string {
	if m != nil {
		return m.XPlayer
	}
	return ""
}

func (m *TttContract) GetOPlayer() string {
	if m != nil {
		return m.OPlayer
	}
	return ""
}

func (m *TttContract) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type InitTrxArgs struct {
	Player1Mark          Mark     `protobuf:"varint,1,opt,name=player1_mark,json=player1Mark,proto3,enum=tictactoe.Mark" json:"player1_mark,omitempty"`
	Player2Mark          Mark     `protobuf:"varint,2,opt,name=player2_mark,json=player2Mark,proto3,enum=tictactoe.Mark" json:"player2_mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitTrxArgs) Reset()         { *m = InitTrxArgs{} }
func (m *InitTrxArgs) String() string { return proto.CompactTextString(m) }
func (*InitTrxArgs) ProtoMessage()    {}
func (*InitTrxArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{1}
}
func (m *InitTrxArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitTrxArgs.Unmarshal(m, b)
}
func (m *InitTrxArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitTrxArgs.Marshal(b, m, deterministic)
}
func (dst *InitTrxArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitTrxArgs.Merge(dst, src)
}
func (m *InitTrxArgs) XXX_Size() int {
	return xxx_messageInfo_InitTrxArgs.Size(m)
}
func (m *InitTrxArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_InitTrxArgs.DiscardUnknown(m)
}

var xxx_messageInfo_InitTrxArgs proto.InternalMessageInfo

func (m *InitTrxArgs) GetPlayer1Mark() Mark {
	if m != nil {
		return m.Player1Mark
	}
	return Mark_X
}

func (m *InitTrxArgs) GetPlayer2Mark() Mark {
	if m != nil {
		return m.Player2Mark
	}
	return Mark_X
}

type TrxArgs struct {
	Type                 TrxType              `protobuf:"varint,1,opt,name=type,proto3,enum=tictactoe.TrxType" json:"type,omitempty"`
	MovePayload          *MoveTrxPayload      `protobuf:"bytes,2,opt,name=movePayload,proto3" json:"movePayload,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TrxArgs) Reset()         { *m = TrxArgs{} }
func (m *TrxArgs) String() string { return proto.CompactTextString(m) }
func (*TrxArgs) ProtoMessage()    {}
func (*TrxArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{2}
}
func (m *TrxArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrxArgs.Unmarshal(m, b)
}
func (m *TrxArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrxArgs.Marshal(b, m, deterministic)
}
func (dst *TrxArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrxArgs.Merge(dst, src)
}
func (m *TrxArgs) XXX_Size() int {
	return xxx_messageInfo_TrxArgs.Size(m)
}
func (m *TrxArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_TrxArgs.DiscardUnknown(m)
}

var xxx_messageInfo_TrxArgs proto.InternalMessageInfo

func (m *TrxArgs) GetType() TrxType {
	if m != nil {
		return m.Type
	}
	return TrxType_MOVE
}

func (m *TrxArgs) GetMovePayload() *MoveTrxPayload {
	if m != nil {
		return m.MovePayload
	}
	return nil
}

func (m *TrxArgs) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type MoveTrxPayload struct {
	Position             int32    `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Mark                 Mark     `protobuf:"varint,2,opt,name=mark,proto3,enum=tictactoe.Mark" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveTrxPayload) Reset()         { *m = MoveTrxPayload{} }
func (m *MoveTrxPayload) String() string { return proto.CompactTextString(m) }
func (*MoveTrxPayload) ProtoMessage()    {}
func (*MoveTrxPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_tictactoe_984f32c8a064e548, []int{3}
}
func (m *MoveTrxPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveTrxPayload.Unmarshal(m, b)
}
func (m *MoveTrxPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveTrxPayload.Marshal(b, m, deterministic)
}
func (dst *MoveTrxPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveTrxPayload.Merge(dst, src)
}
func (m *MoveTrxPayload) XXX_Size() int {
	return xxx_messageInfo_MoveTrxPayload.Size(m)
}
func (m *MoveTrxPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveTrxPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MoveTrxPayload proto.InternalMessageInfo

func (m *MoveTrxPayload) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *MoveTrxPayload) GetMark() Mark {
	if m != nil {
		return m.Mark
	}
	return Mark_X
}

func init() {
	proto.RegisterType((*TttContract)(nil), "tictactoe.TttContract")
	proto.RegisterType((*InitTrxArgs)(nil), "tictactoe.InitTrxArgs")
	proto.RegisterType((*TrxArgs)(nil), "tictactoe.TrxArgs")
	proto.RegisterType((*MoveTrxPayload)(nil), "tictactoe.MoveTrxPayload")
	proto.RegisterEnum("tictactoe.Mark", Mark_name, Mark_value)
	proto.RegisterEnum("tictactoe.TrxType", TrxType_name, TrxType_value)
	proto.RegisterEnum("tictactoe.TttContract_State", TttContract_State_name, TttContract_State_value)
}

func init() { proto.RegisterFile("tictactoe.proto", fileDescriptor_tictactoe_984f32c8a064e548) }

var fileDescriptor_tictactoe_984f32c8a064e548 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0xaf, 0x9a, 0x40,
	0x14, 0x95, 0xaf, 0x2a, 0x17, 0xa3, 0x64, 0xba, 0x41, 0xdb, 0xa4, 0x86, 0x26, 0x4d, 0x63, 0x52,
	0x4c, 0xe9, 0xae, 0x4d, 0x17, 0x4d, 0xe3, 0xc2, 0x85, 0x62, 0xa7, 0xd8, 0xba, 0x33, 0xa8, 0xd4,
	0x90, 0x82, 0x43, 0x60, 0x7c, 0xc1, 0x9f, 0xf3, 0x7e, 0xc0, 0xfb, 0x8f, 0x6f, 0x66, 0x00, 0xe5,
	0x25, 0x2f, 0x6e, 0xde, 0x6a, 0xee, 0xe5, 0x9c, 0x73, 0x3f, 0xce, 0x05, 0xfa, 0x34, 0xda, 0xd1,
	0x60, 0x47, 0x49, 0xe8, 0xa4, 0x19, 0xa1, 0x04, 0xe9, 0x97, 0x0f, 0xc3, 0x77, 0x07, 0x42, 0x0e,
	0x71, 0x38, 0x11, 0xc0, 0xf6, 0xf4, 0x6f, 0x42, 0xa3, 0x24, 0xcc, 0x69, 0x90, 0xa4, 0x25, 0xd7,
	0xbe, 0x97, 0xc1, 0xf0, 0x29, 0xfd, 0x49, 0x8e, 0x34, 0x63, 0x12, 0xf4, 0x09, 0xf4, 0x94, 0xe4,
	0x11, 0x8d, 0xc8, 0x31, 0xb7, 0xa4, 0x91, 0xf2, 0xb1, 0xe7, 0xf6, 0x9d, 0x6b, 0x83, 0x79, 0x90,
	0xfd, 0xc7, 0x57, 0x06, 0x72, 0x41, 0x63, 0xd5, 0x68, 0x68, 0xc9, 0x23, 0x89, 0x51, 0xdf, 0x36,
	0xa8, 0x8d, 0xaa, 0xce, 0x6f, 0xce, 0xc1, 0x25, 0x15, 0x0d, 0xa0, 0x53, 0x6c, 0xd2, 0x38, 0x38,
	0x87, 0x99, 0xa5, 0x30, 0x99, 0x8e, 0xdb, 0xc5, 0x52, 0xa4, 0x1c, 0x22, 0x35, 0xa4, 0x96, 0x10,
	0xa9, 0xa0, 0xef, 0xd0, 0x8d, 0x83, 0x9c, 0x6e, 0x4e, 0xe9, 0x9e, 0x15, 0xd9, 0x5b, 0x3a, 0x83,
	0x0d, 0x77, 0xe8, 0x94, 0x0b, 0x3a, 0xf5, 0x82, 0x8e, 0x5f, 0x2f, 0x88, 0x0d, 0xce, 0x5f, 0x95,
	0x74, 0xfb, 0x2b, 0x68, 0x62, 0x08, 0xa4, 0x83, 0xb6, 0xf6, 0x57, 0x78, 0x61, 0xb6, 0x78, 0xe8,
	0x89, 0x50, 0x42, 0x1d, 0x50, 0xd7, 0x7f, 0xbd, 0x85, 0x29, 0xf3, 0xc8, 0xe3, 0x91, 0x82, 0xda,
	0xa0, 0xf8, 0xb3, 0xa9, 0xa9, 0xda, 0x27, 0x30, 0x66, 0xc7, 0x88, 0xfa, 0x59, 0xf1, 0x23, 0x3b,
	0xf0, 0x9d, 0xbb, 0xe5, 0x88, 0x9f, 0x37, 0x09, 0xb3, 0x83, 0xb9, 0x24, 0x3d, 0xe7, 0x92, 0x51,
	0x91, 0x78, 0x72, 0xd5, 0xb8, 0xa5, 0x46, 0xbe, 0xa9, 0x71, 0x79, 0x62, 0x3f, 0x48, 0xd0, 0xae,
	0x7b, 0x7e, 0x00, 0x95, 0x9e, 0xd3, 0xb0, 0xea, 0x85, 0x9a, 0x36, 0x67, 0x85, 0xcf, 0x10, 0x2c,
	0x70, 0xf4, 0x0d, 0x8c, 0x84, 0xdc, 0x85, 0xcb, 0xe0, 0x1c, 0x93, 0x60, 0x2f, 0xda, 0x18, 0xee,
	0xa0, 0xd9, 0x86, 0xa1, 0x4c, 0x52, 0x11, 0x70, 0x93, 0xfd, 0x52, 0x8b, 0x7f, 0x41, 0xef, 0x69,
	0x75, 0x34, 0x84, 0x4e, 0xfd, 0xab, 0x88, 0xc9, 0x35, 0x7c, 0xc9, 0xd1, 0x7b, 0x50, 0x6f, 0x39,
	0x21, 0xc0, 0xf1, 0x1b, 0x50, 0x85, 0x7d, 0x1a, 0x48, 0x6b, 0x76, 0x30, 0xf6, 0x78, 0xec, 0x58,
	0xec, 0x99, 0x9a, 0xf2, 0xf8, 0xb5, 0xb0, 0x87, 0x2f, 0xcf, 0x8f, 0x36, 0xf7, 0xfe, 0x4c, 0xcd,
	0xd6, 0xf6, 0x95, 0x98, 0xf2, 0xcb, 0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x09, 0x0e, 0x08,
	0x15, 0x03, 0x00, 0x00,
}

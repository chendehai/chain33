// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package types // import "github.com/33cn/chain33/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Reply struct {
	IsOk                 bool     `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Msg                  []byte   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{0}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *Reply) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type ReqString struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqString) Reset()         { *m = ReqString{} }
func (m *ReqString) String() string { return proto.CompactTextString(m) }
func (*ReqString) ProtoMessage()    {}
func (*ReqString) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{1}
}
func (m *ReqString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqString.Unmarshal(m, b)
}
func (m *ReqString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqString.Marshal(b, m, deterministic)
}
func (dst *ReqString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqString.Merge(dst, src)
}
func (m *ReqString) XXX_Size() int {
	return xxx_messageInfo_ReqString.Size(m)
}
func (m *ReqString) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqString.DiscardUnknown(m)
}

var xxx_messageInfo_ReqString proto.InternalMessageInfo

func (m *ReqString) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReplyString struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyString) Reset()         { *m = ReplyString{} }
func (m *ReplyString) String() string { return proto.CompactTextString(m) }
func (*ReplyString) ProtoMessage()    {}
func (*ReplyString) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{2}
}
func (m *ReplyString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyString.Unmarshal(m, b)
}
func (m *ReplyString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyString.Marshal(b, m, deterministic)
}
func (dst *ReplyString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyString.Merge(dst, src)
}
func (m *ReplyString) XXX_Size() int {
	return xxx_messageInfo_ReplyString.Size(m)
}
func (m *ReplyString) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyString.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyString proto.InternalMessageInfo

func (m *ReplyString) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReplyStrings struct {
	Datas                []string `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStrings) Reset()         { *m = ReplyStrings{} }
func (m *ReplyStrings) String() string { return proto.CompactTextString(m) }
func (*ReplyStrings) ProtoMessage()    {}
func (*ReplyStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{3}
}
func (m *ReplyStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStrings.Unmarshal(m, b)
}
func (m *ReplyStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStrings.Marshal(b, m, deterministic)
}
func (dst *ReplyStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStrings.Merge(dst, src)
}
func (m *ReplyStrings) XXX_Size() int {
	return xxx_messageInfo_ReplyStrings.Size(m)
}
func (m *ReplyStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStrings.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStrings proto.InternalMessageInfo

func (m *ReplyStrings) GetDatas() []string {
	if m != nil {
		return m.Datas
	}
	return nil
}

type ReqInt struct {
	Height               int64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqInt) Reset()         { *m = ReqInt{} }
func (m *ReqInt) String() string { return proto.CompactTextString(m) }
func (*ReqInt) ProtoMessage()    {}
func (*ReqInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{4}
}
func (m *ReqInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqInt.Unmarshal(m, b)
}
func (m *ReqInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqInt.Marshal(b, m, deterministic)
}
func (dst *ReqInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqInt.Merge(dst, src)
}
func (m *ReqInt) XXX_Size() int {
	return xxx_messageInfo_ReqInt.Size(m)
}
func (m *ReqInt) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqInt.DiscardUnknown(m)
}

var xxx_messageInfo_ReqInt proto.InternalMessageInfo

func (m *ReqInt) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Int64 struct {
	Data                 int64    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{5}
}
func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (dst *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(dst, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetData() int64 {
	if m != nil {
		return m.Data
	}
	return 0
}

type ReqHash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHash) Reset()         { *m = ReqHash{} }
func (m *ReqHash) String() string { return proto.CompactTextString(m) }
func (*ReqHash) ProtoMessage()    {}
func (*ReqHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{6}
}
func (m *ReqHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHash.Unmarshal(m, b)
}
func (m *ReqHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHash.Marshal(b, m, deterministic)
}
func (dst *ReqHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHash.Merge(dst, src)
}
func (m *ReqHash) XXX_Size() int {
	return xxx_messageInfo_ReqHash.Size(m)
}
func (m *ReqHash) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHash.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHash proto.InternalMessageInfo

func (m *ReqHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ReplyHash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyHash) Reset()         { *m = ReplyHash{} }
func (m *ReplyHash) String() string { return proto.CompactTextString(m) }
func (*ReplyHash) ProtoMessage()    {}
func (*ReplyHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{7}
}
func (m *ReplyHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHash.Unmarshal(m, b)
}
func (m *ReplyHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHash.Marshal(b, m, deterministic)
}
func (dst *ReplyHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHash.Merge(dst, src)
}
func (m *ReplyHash) XXX_Size() int {
	return xxx_messageInfo_ReplyHash.Size(m)
}
func (m *ReplyHash) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHash.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHash proto.InternalMessageInfo

func (m *ReplyHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type ReqNil struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNil) Reset()         { *m = ReqNil{} }
func (m *ReqNil) String() string { return proto.CompactTextString(m) }
func (*ReqNil) ProtoMessage()    {}
func (*ReqNil) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{8}
}
func (m *ReqNil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNil.Unmarshal(m, b)
}
func (m *ReqNil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNil.Marshal(b, m, deterministic)
}
func (dst *ReqNil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNil.Merge(dst, src)
}
func (m *ReqNil) XXX_Size() int {
	return xxx_messageInfo_ReqNil.Size(m)
}
func (m *ReqNil) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNil.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNil proto.InternalMessageInfo

type ReqHashes struct {
	Hashes               [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHashes) Reset()         { *m = ReqHashes{} }
func (m *ReqHashes) String() string { return proto.CompactTextString(m) }
func (*ReqHashes) ProtoMessage()    {}
func (*ReqHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{9}
}
func (m *ReqHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHashes.Unmarshal(m, b)
}
func (m *ReqHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHashes.Marshal(b, m, deterministic)
}
func (dst *ReqHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHashes.Merge(dst, src)
}
func (m *ReqHashes) XXX_Size() int {
	return xxx_messageInfo_ReqHashes.Size(m)
}
func (m *ReqHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHashes.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHashes proto.InternalMessageInfo

func (m *ReqHashes) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type ReplyHashes struct {
	Hashes               [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyHashes) Reset()         { *m = ReplyHashes{} }
func (m *ReplyHashes) String() string { return proto.CompactTextString(m) }
func (*ReplyHashes) ProtoMessage()    {}
func (*ReplyHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{10}
}
func (m *ReplyHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyHashes.Unmarshal(m, b)
}
func (m *ReplyHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyHashes.Marshal(b, m, deterministic)
}
func (dst *ReplyHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyHashes.Merge(dst, src)
}
func (m *ReplyHashes) XXX_Size() int {
	return xxx_messageInfo_ReplyHashes.Size(m)
}
func (m *ReplyHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyHashes.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyHashes proto.InternalMessageInfo

func (m *ReplyHashes) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type KeyValue struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{11}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (dst *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(dst, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type TxHash struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxHash) Reset()         { *m = TxHash{} }
func (m *TxHash) String() string { return proto.CompactTextString(m) }
func (*TxHash) ProtoMessage()    {}
func (*TxHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{12}
}
func (m *TxHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxHash.Unmarshal(m, b)
}
func (m *TxHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxHash.Marshal(b, m, deterministic)
}
func (dst *TxHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxHash.Merge(dst, src)
}
func (m *TxHash) XXX_Size() int {
	return xxx_messageInfo_TxHash.Size(m)
}
func (m *TxHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TxHash.DiscardUnknown(m)
}

var xxx_messageInfo_TxHash proto.InternalMessageInfo

func (m *TxHash) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type TimeStatus struct {
	NtpTime              string   `protobuf:"bytes,1,opt,name=ntpTime,proto3" json:"ntpTime,omitempty"`
	LocalTime            string   `protobuf:"bytes,2,opt,name=localTime,proto3" json:"localTime,omitempty"`
	Diff                 int64    `protobuf:"varint,3,opt,name=diff,proto3" json:"diff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeStatus) Reset()         { *m = TimeStatus{} }
func (m *TimeStatus) String() string { return proto.CompactTextString(m) }
func (*TimeStatus) ProtoMessage()    {}
func (*TimeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{13}
}
func (m *TimeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeStatus.Unmarshal(m, b)
}
func (m *TimeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeStatus.Marshal(b, m, deterministic)
}
func (dst *TimeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeStatus.Merge(dst, src)
}
func (m *TimeStatus) XXX_Size() int {
	return xxx_messageInfo_TimeStatus.Size(m)
}
func (m *TimeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TimeStatus proto.InternalMessageInfo

func (m *TimeStatus) GetNtpTime() string {
	if m != nil {
		return m.NtpTime
	}
	return ""
}

func (m *TimeStatus) GetLocalTime() string {
	if m != nil {
		return m.LocalTime
	}
	return ""
}

func (m *TimeStatus) GetDiff() int64 {
	if m != nil {
		return m.Diff
	}
	return 0
}

type ReqKey struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqKey) Reset()         { *m = ReqKey{} }
func (m *ReqKey) String() string { return proto.CompactTextString(m) }
func (*ReqKey) ProtoMessage()    {}
func (*ReqKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{14}
}
func (m *ReqKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqKey.Unmarshal(m, b)
}
func (m *ReqKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqKey.Marshal(b, m, deterministic)
}
func (dst *ReqKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqKey.Merge(dst, src)
}
func (m *ReqKey) XXX_Size() int {
	return xxx_messageInfo_ReqKey.Size(m)
}
func (m *ReqKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqKey.DiscardUnknown(m)
}

var xxx_messageInfo_ReqKey proto.InternalMessageInfo

func (m *ReqKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type ReqRandHash struct {
	ExecName             string   `protobuf:"bytes,1,opt,name=execName,proto3" json:"execName,omitempty"`
	Height               int64    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	BlockNum             int64    `protobuf:"varint,3,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRandHash) Reset()         { *m = ReqRandHash{} }
func (m *ReqRandHash) String() string { return proto.CompactTextString(m) }
func (*ReqRandHash) ProtoMessage()    {}
func (*ReqRandHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c365cf7167dd3614, []int{15}
}
func (m *ReqRandHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRandHash.Unmarshal(m, b)
}
func (m *ReqRandHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRandHash.Marshal(b, m, deterministic)
}
func (dst *ReqRandHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRandHash.Merge(dst, src)
}
func (m *ReqRandHash) XXX_Size() int {
	return xxx_messageInfo_ReqRandHash.Size(m)
}
func (m *ReqRandHash) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRandHash.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRandHash proto.InternalMessageInfo

func (m *ReqRandHash) GetExecName() string {
	if m != nil {
		return m.ExecName
	}
	return ""
}

func (m *ReqRandHash) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqRandHash) GetBlockNum() int64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Reply)(nil), "types.Reply")
	proto.RegisterType((*ReqString)(nil), "types.ReqString")
	proto.RegisterType((*ReplyString)(nil), "types.ReplyString")
	proto.RegisterType((*ReplyStrings)(nil), "types.ReplyStrings")
	proto.RegisterType((*ReqInt)(nil), "types.ReqInt")
	proto.RegisterType((*Int64)(nil), "types.Int64")
	proto.RegisterType((*ReqHash)(nil), "types.ReqHash")
	proto.RegisterType((*ReplyHash)(nil), "types.ReplyHash")
	proto.RegisterType((*ReqNil)(nil), "types.ReqNil")
	proto.RegisterType((*ReqHashes)(nil), "types.ReqHashes")
	proto.RegisterType((*ReplyHashes)(nil), "types.ReplyHashes")
	proto.RegisterType((*KeyValue)(nil), "types.KeyValue")
	proto.RegisterType((*TxHash)(nil), "types.TxHash")
	proto.RegisterType((*TimeStatus)(nil), "types.TimeStatus")
	proto.RegisterType((*ReqKey)(nil), "types.ReqKey")
	proto.RegisterType((*ReqRandHash)(nil), "types.ReqRandHash")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_c365cf7167dd3614) }

var fileDescriptor_common_c365cf7167dd3614 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xa5, 0x53, 0xdb, 0x6d, 0xaf, 0x7d, 0x90, 0x20, 0x52, 0xd6, 0x5d, 0xb6, 0x46, 0x85, 0x79,
	0x71, 0x07, 0xac, 0xf8, 0x01, 0x3e, 0xb9, 0x2c, 0x8c, 0x90, 0x1d, 0x44, 0x04, 0x1f, 0x32, 0x9d,
	0xb4, 0x0d, 0xd3, 0xa6, 0xed, 0x24, 0x95, 0xe9, 0xdf, 0x4b, 0xee, 0xc4, 0x71, 0x64, 0xd0, 0x7d,
	0xbb, 0x27, 0xf7, 0xdc, 0x9c, 0x73, 0x4f, 0x02, 0x49, 0xd1, 0xb5, 0x6d, 0xa7, 0x6e, 0xfb, 0x5d,
	0x67, 0x3a, 0x12, 0x98, 0xa9, 0x17, 0x9a, 0xbe, 0x83, 0x80, 0x89, 0xbe, 0x99, 0x08, 0x81, 0x27,
	0x52, 0x7f, 0xd9, 0xa6, 0x5e, 0xe6, 0xcd, 0x23, 0x86, 0x35, 0x79, 0x06, 0x7e, 0xab, 0xab, 0x74,
	0x96, 0x79, 0xf3, 0x84, 0xd9, 0x92, 0xde, 0x40, 0xcc, 0xc4, 0xf0, 0x60, 0x76, 0x52, 0x55, 0x76,
	0x64, 0xc3, 0x0d, 0xc7, 0x91, 0x98, 0x61, 0x4d, 0x5f, 0xc1, 0x53, 0xbc, 0xef, 0x3f, 0x94, 0x37,
	0x90, 0x9c, 0x50, 0x34, 0x79, 0x0e, 0x81, 0x3d, 0xd7, 0xa9, 0x97, 0xf9, 0xf3, 0x98, 0x1d, 0x00,
	0xcd, 0x20, 0x64, 0x62, 0xb8, 0x53, 0x86, 0xbc, 0x80, 0xb0, 0x16, 0xb2, 0xaa, 0x0d, 0xde, 0xe2,
	0x33, 0x87, 0xe8, 0x4b, 0x08, 0xee, 0x94, 0xf9, 0xf8, 0xe1, 0x2f, 0x11, 0xdf, 0x89, 0x5c, 0xc3,
	0x05, 0x13, 0xc3, 0x67, 0xae, 0x6b, 0xdb, 0xae, 0xb9, 0xae, 0xb1, 0x9d, 0x30, 0xac, 0x0f, 0x7b,
	0xf4, 0xcd, 0xf4, 0x4f, 0x42, 0x84, 0xf2, 0x4b, 0xd9, 0xd0, 0xd7, 0xb8, 0xb2, 0x25, 0x0a, 0x8d,
	0x5e, 0xb0, 0x42, 0xb3, 0x09, 0x73, 0x88, 0xbe, 0x75, 0x6b, 0x3f, 0x42, 0x7b, 0x0f, 0xd1, 0xbd,
	0x98, 0xbe, 0xf2, 0x66, 0x14, 0x36, 0xdc, 0xad, 0x98, 0x9c, 0xa8, 0x2d, 0x6d, 0x10, 0x3f, 0x6d,
	0xcb, 0x05, 0x7e, 0x00, 0xf4, 0x0a, 0xc2, 0xd5, 0xfe, 0xcc, 0x67, 0xec, 0x7c, 0x7e, 0x03, 0x58,
	0xc9, 0x56, 0x3c, 0x18, 0x6e, 0x46, 0x4d, 0x52, 0xb8, 0x50, 0xa6, 0xb7, 0x07, 0x8e, 0xf4, 0x1b,
	0x92, 0x2b, 0x88, 0x9b, 0xae, 0xe0, 0x0d, 0xf6, 0x66, 0xd8, 0xfb, 0x73, 0x80, 0x09, 0xca, 0xb2,
	0x4c, 0x7d, 0x97, 0xa0, 0x2c, 0x4b, 0x7a, 0x89, 0x09, 0xdc, 0x8b, 0xe9, 0xdc, 0x29, 0xfd, 0x61,
	0xd7, 0x1d, 0x18, 0x57, 0x1b, 0x34, 0x76, 0x09, 0x91, 0xd8, 0x8b, 0x62, 0xc9, 0x8f, 0xba, 0x47,
	0x7c, 0xf2, 0x7a, 0xb3, 0xd3, 0xd7, 0xb3, 0x33, 0xeb, 0xa6, 0x2b, 0xb6, 0xcb, 0xb1, 0x75, 0xb2,
	0x47, 0xfc, 0xe9, 0xe6, 0xfb, 0x75, 0x25, 0x4d, 0x3d, 0xae, 0x6f, 0x8b, 0xae, 0x5d, 0xe4, 0x79,
	0xa1, 0x16, 0x45, 0xcd, 0xa5, 0xca, 0xf3, 0x05, 0xfe, 0xda, 0x75, 0x88, 0x7f, 0x38, 0xff, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0xfb, 0x3a, 0x44, 0xd3, 0x02, 0x00, 0x00,
}

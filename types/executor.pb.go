// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: executor.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Genesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isrun bool `protobuf:"varint,1,opt,name=isrun,proto3" json:"isrun,omitempty"`
}

func (x *Genesis) Reset() {
	*x = Genesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genesis) ProtoMessage() {}

func (x *Genesis) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genesis.ProtoReflect.Descriptor instead.
func (*Genesis) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{0}
}

func (x *Genesis) GetIsrun() bool {
	if x != nil {
		return x.Isrun
	}
	return false
}

type ExecTxList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateHash  []byte         `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	ParentHash []byte         `protobuf:"bytes,7,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	MainHash   []byte         `protobuf:"bytes,8,opt,name=mainHash,proto3" json:"mainHash,omitempty"`
	MainHeight int64          `protobuf:"varint,9,opt,name=mainHeight,proto3" json:"mainHeight,omitempty"`
	BlockTime  int64          `protobuf:"varint,3,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	Height     int64          `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Difficulty uint64         `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	IsMempool  bool           `protobuf:"varint,6,opt,name=isMempool,proto3" json:"isMempool,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *ExecTxList) Reset() {
	*x = ExecTxList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecTxList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecTxList) ProtoMessage() {}

func (x *ExecTxList) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecTxList.ProtoReflect.Descriptor instead.
func (*ExecTxList) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{1}
}

func (x *ExecTxList) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *ExecTxList) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *ExecTxList) GetMainHash() []byte {
	if x != nil {
		return x.MainHash
	}
	return nil
}

func (x *ExecTxList) GetMainHeight() int64 {
	if x != nil {
		return x.MainHeight
	}
	return 0
}

func (x *ExecTxList) GetBlockTime() int64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *ExecTxList) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ExecTxList) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *ExecTxList) GetIsMempool() bool {
	if x != nil {
		return x.IsMempool
	}
	return false
}

func (x *ExecTxList) GetTxs() []*Transaction {
	if x != nil {
		return x.Txs
	}
	return nil
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Execer   []byte `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	FuncName string `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Payload  []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetExecer() []byte {
	if x != nil {
		return x.Execer
	}
	return nil
}

func (x *Query) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *Query) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateTxIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Execer     []byte `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	ActionName string `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Payload    []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateTxIn) Reset() {
	*x = CreateTxIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxIn) ProtoMessage() {}

func (x *CreateTxIn) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxIn.ProtoReflect.Descriptor instead.
func (*CreateTxIn) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTxIn) GetExecer() []byte {
	if x != nil {
		return x.Execer
	}
	return nil
}

func (x *CreateTxIn) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *CreateTxIn) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// 配置修改部分
type ArrayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ArrayConfig) Reset() {
	*x = ArrayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayConfig) ProtoMessage() {}

func (x *ArrayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayConfig.ProtoReflect.Descriptor instead.
func (*ArrayConfig) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{4}
}

func (x *ArrayConfig) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type StringConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringConfig) Reset() {
	*x = StringConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringConfig) ProtoMessage() {}

func (x *StringConfig) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringConfig.ProtoReflect.Descriptor instead.
func (*StringConfig) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{5}
}

func (x *StringConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Int32Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Int32Config) Reset() {
	*x = Int32Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Config) ProtoMessage() {}

func (x *Int32Config) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Config.ProtoReflect.Descriptor instead.
func (*Int32Config) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{6}
}

func (x *Int32Config) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ConfigItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// Types that are assignable to Value:
	//	*ConfigItem_Arr
	//	*ConfigItem_Str
	//	*ConfigItem_Int
	Value isConfigItem_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,11,opt,name=Ty,proto3" json:"Ty,omitempty"`
}

func (x *ConfigItem) Reset() {
	*x = ConfigItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigItem) ProtoMessage() {}

func (x *ConfigItem) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigItem.ProtoReflect.Descriptor instead.
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ConfigItem) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (m *ConfigItem) GetValue() isConfigItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ConfigItem) GetArr() *ArrayConfig {
	if x, ok := x.GetValue().(*ConfigItem_Arr); ok {
		return x.Arr
	}
	return nil
}

func (x *ConfigItem) GetStr() *StringConfig {
	if x, ok := x.GetValue().(*ConfigItem_Str); ok {
		return x.Str
	}
	return nil
}

func (x *ConfigItem) GetInt() *Int32Config {
	if x, ok := x.GetValue().(*ConfigItem_Int); ok {
		return x.Int
	}
	return nil
}

func (x *ConfigItem) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isConfigItem_Value interface {
	isConfigItem_Value()
}

type ConfigItem_Arr struct {
	Arr *ArrayConfig `protobuf:"bytes,3,opt,name=arr,proto3,oneof"`
}

type ConfigItem_Str struct {
	Str *StringConfig `protobuf:"bytes,4,opt,name=str,proto3,oneof"`
}

type ConfigItem_Int struct {
	Int *Int32Config `protobuf:"bytes,5,opt,name=int,proto3,oneof"`
}

func (*ConfigItem_Arr) isConfigItem_Value() {}

func (*ConfigItem_Str) isConfigItem_Value() {}

func (*ConfigItem_Int) isConfigItem_Value() {}

type ModifyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Op    string `protobuf:"bytes,3,opt,name=op,proto3" json:"op,omitempty"`
	Addr  string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *ModifyConfig) Reset() {
	*x = ModifyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyConfig) ProtoMessage() {}

func (x *ModifyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyConfig.ProtoReflect.Descriptor instead.
func (*ModifyConfig) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{8}
}

func (x *ModifyConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ModifyConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ModifyConfig) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *ModifyConfig) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type ReceiptConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prev    *ConfigItem `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current *ConfigItem `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *ReceiptConfig) Reset() {
	*x = ReceiptConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptConfig) ProtoMessage() {}

func (x *ReceiptConfig) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptConfig.ProtoReflect.Descriptor instead.
func (*ReceiptConfig) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{9}
}

func (x *ReceiptConfig) GetPrev() *ConfigItem {
	if x != nil {
		return x.Prev
	}
	return nil
}

func (x *ReceiptConfig) GetCurrent() *ConfigItem {
	if x != nil {
		return x.Current
	}
	return nil
}

type ReplyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReplyConfig) Reset() {
	*x = ReplyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConfig) ProtoMessage() {}

func (x *ReplyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConfig.ProtoReflect.Descriptor instead.
func (*ReplyConfig) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{10}
}

func (x *ReplyConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReplyConfig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type HistoryCertStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rootcerts         [][]byte `protobuf:"bytes,1,rep,name=rootcerts,proto3" json:"rootcerts,omitempty"`
	IntermediateCerts [][]byte `protobuf:"bytes,2,rep,name=intermediateCerts,proto3" json:"intermediateCerts,omitempty"`
	RevocationList    [][]byte `protobuf:"bytes,3,rep,name=revocationList,proto3" json:"revocationList,omitempty"`
	CurHeigth         int64    `protobuf:"varint,4,opt,name=curHeigth,proto3" json:"curHeigth,omitempty"`
	NxtHeight         int64    `protobuf:"varint,5,opt,name=nxtHeight,proto3" json:"nxtHeight,omitempty"`
}

func (x *HistoryCertStore) Reset() {
	*x = HistoryCertStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryCertStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryCertStore) ProtoMessage() {}

func (x *HistoryCertStore) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryCertStore.ProtoReflect.Descriptor instead.
func (*HistoryCertStore) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{11}
}

func (x *HistoryCertStore) GetRootcerts() [][]byte {
	if x != nil {
		return x.Rootcerts
	}
	return nil
}

func (x *HistoryCertStore) GetIntermediateCerts() [][]byte {
	if x != nil {
		return x.IntermediateCerts
	}
	return nil
}

func (x *HistoryCertStore) GetRevocationList() [][]byte {
	if x != nil {
		return x.RevocationList
	}
	return nil
}

func (x *HistoryCertStore) GetCurHeigth() int64 {
	if x != nil {
		return x.CurHeigth
	}
	return 0
}

func (x *HistoryCertStore) GetNxtHeight() int64 {
	if x != nil {
		return x.NxtHeight
	}
	return 0
}

var File_executor_proto protoreflect.FileDescriptor

var file_executor_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x47, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x72, 0x75, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x72, 0x75, 0x6e, 0x22, 0xa0, 0x02, 0x0a, 0x0a,
	0x45, 0x78, 0x65, 0x63, 0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x4d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x4d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x55,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x65, 0x63, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x78, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x65, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x78, 0x65, 0x63, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x23, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x72,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x03, 0x61,
	0x72, 0x72, 0x12, 0x27, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x26, 0x0a, 0x03, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x03,
	0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x54, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x63, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x72, 0x65,
	0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76,
	0x12, 0x2b, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a,
	0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x6f, 0x6f,
	0x74, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x6f,
	0x6f, 0x74, 0x63, 0x65, 0x72, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x72,
	0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x75, 0x72, 0x48, 0x65, 0x69, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x75, 0x72, 0x48, 0x65, 0x69, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x78, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6e, 0x78, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x33, 0x63, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x33, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_executor_proto_rawDescOnce sync.Once
	file_executor_proto_rawDescData = file_executor_proto_rawDesc
)

func file_executor_proto_rawDescGZIP() []byte {
	file_executor_proto_rawDescOnce.Do(func() {
		file_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_executor_proto_rawDescData)
	})
	return file_executor_proto_rawDescData
}

var file_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_executor_proto_goTypes = []interface{}{
	(*Genesis)(nil),          // 0: types.Genesis
	(*ExecTxList)(nil),       // 1: types.ExecTxList
	(*Query)(nil),            // 2: types.Query
	(*CreateTxIn)(nil),       // 3: types.CreateTxIn
	(*ArrayConfig)(nil),      // 4: types.ArrayConfig
	(*StringConfig)(nil),     // 5: types.StringConfig
	(*Int32Config)(nil),      // 6: types.Int32Config
	(*ConfigItem)(nil),       // 7: types.ConfigItem
	(*ModifyConfig)(nil),     // 8: types.ModifyConfig
	(*ReceiptConfig)(nil),    // 9: types.ReceiptConfig
	(*ReplyConfig)(nil),      // 10: types.ReplyConfig
	(*HistoryCertStore)(nil), // 11: types.HistoryCertStore
	(*Transaction)(nil),      // 12: types.Transaction
}
var file_executor_proto_depIdxs = []int32{
	12, // 0: types.ExecTxList.txs:type_name -> types.Transaction
	4,  // 1: types.ConfigItem.arr:type_name -> types.ArrayConfig
	5,  // 2: types.ConfigItem.str:type_name -> types.StringConfig
	6,  // 3: types.ConfigItem.int:type_name -> types.Int32Config
	7,  // 4: types.ReceiptConfig.prev:type_name -> types.ConfigItem
	7,  // 5: types.ReceiptConfig.current:type_name -> types.ConfigItem
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_executor_proto_init() }
func file_executor_proto_init() {
	if File_executor_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genesis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecTxList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_executor_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryCertStore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_executor_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ConfigItem_Arr)(nil),
		(*ConfigItem_Str)(nil),
		(*ConfigItem_Int)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_executor_proto_goTypes,
		DependencyIndexes: file_executor_proto_depIdxs,
		MessageInfos:      file_executor_proto_msgTypes,
	}.Build()
	File_executor_proto = out.File
	file_executor_proto_rawDesc = nil
	file_executor_proto_goTypes = nil
	file_executor_proto_depIdxs = nil
}

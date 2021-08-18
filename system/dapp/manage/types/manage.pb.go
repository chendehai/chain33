// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: manage.proto

package types

import (
	reflect "reflect"
	sync "sync"

	types "github.com/33cn/chain33/types"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ManageAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*ManageAction_Modify
	Value isManageAction_Value `protobuf_oneof:"value"`
	Ty    int32                `protobuf:"varint,2,opt,name=Ty,proto3" json:"Ty,omitempty"`
}

func (x *ManageAction) Reset() {
	*x = ManageAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageAction) ProtoMessage() {}

func (x *ManageAction) ProtoReflect() protoreflect.Message {
	mi := &file_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageAction.ProtoReflect.Descriptor instead.
func (*ManageAction) Descriptor() ([]byte, []int) {
	return file_manage_proto_rawDescGZIP(), []int{0}
}

func (m *ManageAction) GetValue() isManageAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ManageAction) GetModify() *types.ModifyConfig {
	if x, ok := x.GetValue().(*ManageAction_Modify); ok {
		return x.Modify
	}
	return nil
}

func (x *ManageAction) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isManageAction_Value interface {
	isManageAction_Value()
}

type ManageAction_Modify struct {
	Modify *types.ModifyConfig `protobuf:"bytes,1,opt,name=modify,proto3,oneof"`
}

func (*ManageAction_Modify) isManageAction_Value() {}

var File_manage_proto protoreflect.FileDescriptor

var file_manage_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x54, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manage_proto_rawDescOnce sync.Once
	file_manage_proto_rawDescData = file_manage_proto_rawDesc
)

func file_manage_proto_rawDescGZIP() []byte {
	file_manage_proto_rawDescOnce.Do(func() {
		file_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_manage_proto_rawDescData)
	})
	return file_manage_proto_rawDescData
}

var file_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_manage_proto_goTypes = []interface{}{
	(*ManageAction)(nil),       // 0: types.ManageAction
	(*types.ModifyConfig)(nil), // 1: types.ModifyConfig
}
var file_manage_proto_depIdxs = []int32{
	1, // 0: types.ManageAction.modify:type_name -> types.ModifyConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manage_proto_init() }
func file_manage_proto_init() {
	if File_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageAction); i {
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
	file_manage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ManageAction_Modify)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manage_proto_goTypes,
		DependencyIndexes: file_manage_proto_depIdxs,
		MessageInfos:      file_manage_proto_msgTypes,
	}.Build()
	File_manage_proto = out.File
	file_manage_proto_rawDesc = nil
	file_manage_proto_goTypes = nil
	file_manage_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/kv.proto

package pb

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

type KVOperation_Type int32

const (
	KVOperation_UNSET  KVOperation_Type = 0 // Protobuf default should not be used, this is used so that the consume can ensure that the value was actually specified
	KVOperation_SET    KVOperation_Type = 1
	KVOperation_DELETE KVOperation_Type = 2
)

// Enum value maps for KVOperation_Type.
var (
	KVOperation_Type_name = map[int32]string{
		0: "UNSET",
		1: "SET",
		2: "DELETE",
	}
	KVOperation_Type_value = map[string]int32{
		"UNSET":  0,
		"SET":    1,
		"DELETE": 2,
	}
)

func (x KVOperation_Type) Enum() *KVOperation_Type {
	p := new(KVOperation_Type)
	*p = x
	return p
}

func (x KVOperation_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KVOperation_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_kv_proto_enumTypes[0].Descriptor()
}

func (KVOperation_Type) Type() protoreflect.EnumType {
	return &file_proto_kv_proto_enumTypes[0]
}

func (x KVOperation_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KVOperation_Type.Descriptor instead.
func (KVOperation_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_kv_proto_rawDescGZIP(), []int{1, 0}
}

type KVOperations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*KVOperation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *KVOperations) Reset() {
	*x = KVOperations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVOperations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVOperations) ProtoMessage() {}

func (x *KVOperations) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVOperations.ProtoReflect.Descriptor instead.
func (*KVOperations) Descriptor() ([]byte, []int) {
	return file_proto_kv_proto_rawDescGZIP(), []int{0}
}

func (x *KVOperations) GetOperations() []*KVOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type KVOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ordinal uint64           `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Type    KVOperation_Type `protobuf:"varint,4,opt,name=type,proto3,enum=sf.substreams.sink.kv.v1.KVOperation_Type" json:"type,omitempty"`
}

func (x *KVOperation) Reset() {
	*x = KVOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVOperation) ProtoMessage() {}

func (x *KVOperation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVOperation.ProtoReflect.Descriptor instead.
func (*KVOperation) Descriptor() ([]byte, []int) {
	return file_proto_kv_proto_rawDescGZIP(), []int{1}
}

func (x *KVOperation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVOperation) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KVOperation) GetOrdinal() uint64 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *KVOperation) GetType() KVOperation_Type {
	if x != nil {
		return x.Type
	}
	return KVOperation_UNSET
}

var File_proto_kv_proto protoreflect.FileDescriptor

var file_proto_kv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x22, 0x55, 0x0a, 0x0c, 0x4b, 0x56,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x4b, 0x56, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x26, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2d, 0x73, 0x69, 0x6e, 0x6b, 0x2d, 0x6b, 0x76, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kv_proto_rawDescOnce sync.Once
	file_proto_kv_proto_rawDescData = file_proto_kv_proto_rawDesc
)

func file_proto_kv_proto_rawDescGZIP() []byte {
	file_proto_kv_proto_rawDescOnce.Do(func() {
		file_proto_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kv_proto_rawDescData)
	})
	return file_proto_kv_proto_rawDescData
}

var file_proto_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_kv_proto_goTypes = []interface{}{
	(KVOperation_Type)(0), // 0: sf.substreams.sink.kv.v1.KVOperation.Type
	(*KVOperations)(nil),  // 1: sf.substreams.sink.kv.v1.KVOperations
	(*KVOperation)(nil),   // 2: sf.substreams.sink.kv.v1.KVOperation
}
var file_proto_kv_proto_depIdxs = []int32{
	2, // 0: sf.substreams.sink.kv.v1.KVOperations.operations:type_name -> sf.substreams.sink.kv.v1.KVOperation
	0, // 1: sf.substreams.sink.kv.v1.KVOperation.type:type_name -> sf.substreams.sink.kv.v1.KVOperation.Type
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_kv_proto_init() }
func file_proto_kv_proto_init() {
	if File_proto_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVOperations); i {
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
		file_proto_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVOperation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_kv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_kv_proto_goTypes,
		DependencyIndexes: file_proto_kv_proto_depIdxs,
		EnumInfos:         file_proto_kv_proto_enumTypes,
		MessageInfos:      file_proto_kv_proto_msgTypes,
	}.Build()
	File_proto_kv_proto = out.File
	file_proto_kv_proto_rawDesc = nil
	file_proto_kv_proto_goTypes = nil
	file_proto_kv_proto_depIdxs = nil
}

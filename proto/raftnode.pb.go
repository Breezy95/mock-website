// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: raftnode.proto

package proto

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

type IncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val uint64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IncRequest) Reset() {
	*x = IncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncRequest) ProtoMessage() {}

func (x *IncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raftnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncRequest.ProtoReflect.Descriptor instead.
func (*IncRequest) Descriptor() ([]byte, []int) {
	return file_raftnode_proto_rawDescGZIP(), []int{0}
}

func (x *IncRequest) GetVal() uint64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type IncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncCounter uint64 `protobuf:"varint,1,opt,name=incCounter,proto3" json:"incCounter,omitempty"`
}

func (x *IncResponse) Reset() {
	*x = IncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncResponse) ProtoMessage() {}

func (x *IncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raftnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncResponse.ProtoReflect.Descriptor instead.
func (*IncResponse) Descriptor() ([]byte, []int) {
	return file_raftnode_proto_rawDescGZIP(), []int{1}
}

func (x *IncResponse) GetIncCounter() uint64 {
	if x != nil {
		return x.IncCounter
	}
	return 0
}

type DecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DecVal int32 `protobuf:"zigzag32,1,opt,name=decVal,proto3" json:"decVal,omitempty"`
}

func (x *DecRequest) Reset() {
	*x = DecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecRequest) ProtoMessage() {}

func (x *DecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raftnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecRequest.ProtoReflect.Descriptor instead.
func (*DecRequest) Descriptor() ([]byte, []int) {
	return file_raftnode_proto_rawDescGZIP(), []int{2}
}

func (x *DecRequest) GetDecVal() int32 {
	if x != nil {
		return x.DecVal
	}
	return 0
}

type DecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DecCounter int32 `protobuf:"zigzag32,1,opt,name=decCounter,proto3" json:"decCounter,omitempty"`
}

func (x *DecResponse) Reset() {
	*x = DecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecResponse) ProtoMessage() {}

func (x *DecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raftnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecResponse.ProtoReflect.Descriptor instead.
func (*DecResponse) Descriptor() ([]byte, []int) {
	return file_raftnode_proto_rawDescGZIP(), []int{3}
}

func (x *DecResponse) GetDecCounter() int32 {
	if x != nil {
		return x.DecCounter
	}
	return 0
}

var File_raftnode_proto protoreflect.FileDescriptor

var file_raftnode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x61, 0x66, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x2d, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22,
	0x24, 0x0a, 0x0a, 0x44, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x64,
	0x65, 0x63, 0x56, 0x61, 0x6c, 0x22, 0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x32, 0x5d, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x49,
	0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x49, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x09, 0x64, 0x65, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x66, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raftnode_proto_rawDescOnce sync.Once
	file_raftnode_proto_rawDescData = file_raftnode_proto_rawDesc
)

func file_raftnode_proto_rawDescGZIP() []byte {
	file_raftnode_proto_rawDescOnce.Do(func() {
		file_raftnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_raftnode_proto_rawDescData)
	})
	return file_raftnode_proto_rawDescData
}

var file_raftnode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_raftnode_proto_goTypes = []interface{}{
	(*IncRequest)(nil),  // 0: IncRequest
	(*IncResponse)(nil), // 1: IncResponse
	(*DecRequest)(nil),  // 2: DecRequest
	(*DecResponse)(nil), // 3: DecResponse
}
var file_raftnode_proto_depIdxs = []int32{
	0, // 0: Example.increment:input_type -> IncRequest
	2, // 1: Example.decrement:input_type -> DecRequest
	1, // 2: Example.increment:output_type -> IncResponse
	3, // 3: Example.decrement:output_type -> DecResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_raftnode_proto_init() }
func file_raftnode_proto_init() {
	if File_raftnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raftnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncRequest); i {
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
		file_raftnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncResponse); i {
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
		file_raftnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecRequest); i {
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
		file_raftnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecResponse); i {
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
			RawDescriptor: file_raftnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raftnode_proto_goTypes,
		DependencyIndexes: file_raftnode_proto_depIdxs,
		MessageInfos:      file_raftnode_proto_msgTypes,
	}.Build()
	File_raftnode_proto = out.File
	file_raftnode_proto_rawDesc = nil
	file_raftnode_proto_goTypes = nil
	file_raftnode_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/alloy_analyzer.proto

package protos

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

type ModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *ModelRequest) Reset() {
	*x = ModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_alloy_analyzer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelRequest) ProtoMessage() {}

func (x *ModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_alloy_analyzer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelRequest.ProtoReflect.Descriptor instead.
func (*ModelRequest) Descriptor() ([]byte, []int) {
	return file_protos_alloy_analyzer_proto_rawDescGZIP(), []int{0}
}

func (x *ModelRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type ModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ModelResponse) Reset() {
	*x = ModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_alloy_analyzer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelResponse) ProtoMessage() {}

func (x *ModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_alloy_analyzer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelResponse.ProtoReflect.Descriptor instead.
func (*ModelResponse) Descriptor() ([]byte, []int) {
	return file_protos_alloy_analyzer_proto_rawDescGZIP(), []int{1}
}

func (x *ModelResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_protos_alloy_analyzer_proto protoreflect.FileDescriptor

var file_protos_alloy_analyzer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x6c, 0x6c, 0x6f, 0x79, 0x22, 0x2b, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x4c, 0x0a, 0x0d, 0x41, 0x6c,
	0x6c, 0x6f, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x61, 0x6c,
	0x6c, 0x6f, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x63, 0x6d, 0x35, 0x33, 0x34, 0x33, 0x2f, 0x63,
	0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x2d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_alloy_analyzer_proto_rawDescOnce sync.Once
	file_protos_alloy_analyzer_proto_rawDescData = file_protos_alloy_analyzer_proto_rawDesc
)

func file_protos_alloy_analyzer_proto_rawDescGZIP() []byte {
	file_protos_alloy_analyzer_proto_rawDescOnce.Do(func() {
		file_protos_alloy_analyzer_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_alloy_analyzer_proto_rawDescData)
	})
	return file_protos_alloy_analyzer_proto_rawDescData
}

var file_protos_alloy_analyzer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_alloy_analyzer_proto_goTypes = []any{
	(*ModelRequest)(nil),  // 0: alloy.ModelRequest
	(*ModelResponse)(nil), // 1: alloy.ModelResponse
}
var file_protos_alloy_analyzer_proto_depIdxs = []int32{
	0, // 0: alloy.AlloyAnalyzer.analyzeModel:input_type -> alloy.ModelRequest
	1, // 1: alloy.AlloyAnalyzer.analyzeModel:output_type -> alloy.ModelResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_alloy_analyzer_proto_init() }
func file_protos_alloy_analyzer_proto_init() {
	if File_protos_alloy_analyzer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_alloy_analyzer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ModelRequest); i {
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
		file_protos_alloy_analyzer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ModelResponse); i {
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
			RawDescriptor: file_protos_alloy_analyzer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_alloy_analyzer_proto_goTypes,
		DependencyIndexes: file_protos_alloy_analyzer_proto_depIdxs,
		MessageInfos:      file_protos_alloy_analyzer_proto_msgTypes,
	}.Build()
	File_protos_alloy_analyzer_proto = out.File
	file_protos_alloy_analyzer_proto_rawDesc = nil
	file_protos_alloy_analyzer_proto_goTypes = nil
	file_protos_alloy_analyzer_proto_depIdxs = nil
}

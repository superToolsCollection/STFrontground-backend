// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: tools.proto

package tools

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MorseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *MorseReq) Reset() {
	*x = MorseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorseReq) ProtoMessage() {}

func (x *MorseReq) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorseReq.ProtoReflect.Descriptor instead.
func (*MorseReq) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{0}
}

func (x *MorseReq) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type MorseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MorseStr string `protobuf:"bytes,1,opt,name=morseStr,proto3" json:"morseStr,omitempty"`
}

func (x *MorseResp) Reset() {
	*x = MorseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MorseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MorseResp) ProtoMessage() {}

func (x *MorseResp) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MorseResp.ProtoReflect.Descriptor instead.
func (*MorseResp) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{1}
}

func (x *MorseResp) GetMorseStr() string {
	if x != nil {
		return x.MorseStr
	}
	return ""
}

type QrCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *QrCodeReq) Reset() {
	*x = QrCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrCodeReq) ProtoMessage() {}

func (x *QrCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrCodeReq.ProtoReflect.Descriptor instead.
func (*QrCodeReq) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{2}
}

func (x *QrCodeReq) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type QrCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrCodeStr string `protobuf:"bytes,1,opt,name=qrCodeStr,proto3" json:"qrCodeStr,omitempty"`
}

func (x *QrCodeResp) Reset() {
	*x = QrCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QrCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QrCodeResp) ProtoMessage() {}

func (x *QrCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QrCodeResp.ProtoReflect.Descriptor instead.
func (*QrCodeResp) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{3}
}

func (x *QrCodeResp) GetQrCodeStr() string {
	if x != nil {
		return x.QrCodeStr
	}
	return ""
}

type Rgb2HexReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *Rgb2HexReq) Reset() {
	*x = Rgb2HexReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rgb2HexReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rgb2HexReq) ProtoMessage() {}

func (x *Rgb2HexReq) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rgb2HexReq.ProtoReflect.Descriptor instead.
func (*Rgb2HexReq) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{4}
}

func (x *Rgb2HexReq) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type Rgb2HexResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rgb2HexStr string `protobuf:"bytes,1,opt,name=rgb2HexStr,proto3" json:"rgb2HexStr,omitempty"`
}

func (x *Rgb2HexResp) Reset() {
	*x = Rgb2HexResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rgb2HexResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rgb2HexResp) ProtoMessage() {}

func (x *Rgb2HexResp) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rgb2HexResp.ProtoReflect.Descriptor instead.
func (*Rgb2HexResp) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{5}
}

func (x *Rgb2HexResp) GetRgb2HexStr() string {
	if x != nil {
		return x.Rgb2HexStr
	}
	return ""
}

type GetStoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStoryReq) Reset() {
	*x = GetStoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoryReq) ProtoMessage() {}

func (x *GetStoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoryReq.ProtoReflect.Descriptor instead.
func (*GetStoryReq) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{6}
}

func (x *GetStoryReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetStoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Story  string `protobuf:"bytes,1,opt,name=story,proto3" json:"story,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetStoryResp) Reset() {
	*x = GetStoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoryResp) ProtoMessage() {}

func (x *GetStoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_tools_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoryResp.ProtoReflect.Descriptor instead.
func (*GetStoryResp) Descriptor() ([]byte, []int) {
	return file_tools_proto_rawDescGZIP(), []int{7}
}

func (x *GetStoryResp) GetStory() string {
	if x != nil {
		return x.Story
	}
	return ""
}

func (x *GetStoryResp) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_tools_proto protoreflect.FileDescriptor

var file_tools_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x1c, 0x0a, 0x08, 0x4d, 0x6f, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x22, 0x27, 0x0a, 0x09, 0x4d, 0x6f, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x72, 0x73, 0x65, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x6f, 0x72, 0x73, 0x65, 0x53, 0x74, 0x72, 0x22, 0x1d, 0x0a, 0x09, 0x51,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x22, 0x2a, 0x0a, 0x0a, 0x51, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x72, 0x22, 0x1e, 0x0a, 0x0a, 0x52, 0x67, 0x62, 0x32, 0x48, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x22, 0x2d, 0x0a, 0x0b, 0x52, 0x67, 0x62, 0x32, 0x48, 0x65,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x67, 0x62, 0x32, 0x48, 0x65, 0x78,
	0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x67, 0x62, 0x32, 0x48,
	0x65, 0x78, 0x53, 0x74, 0x72, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x32, 0xc6, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x05,
	0x4d, 0x6f, 0x72, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x4d, 0x6f,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x4d,
	0x6f, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x51, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x51, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x51, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x07, 0x52, 0x67, 0x62, 0x32, 0x48,
	0x65, 0x78, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x52, 0x67, 0x62, 0x32, 0x48,
	0x65, 0x78, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x52, 0x67,
	0x62, 0x32, 0x48, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tools_proto_rawDescOnce sync.Once
	file_tools_proto_rawDescData = file_tools_proto_rawDesc
)

func file_tools_proto_rawDescGZIP() []byte {
	file_tools_proto_rawDescOnce.Do(func() {
		file_tools_proto_rawDescData = protoimpl.X.CompressGZIP(file_tools_proto_rawDescData)
	})
	return file_tools_proto_rawDescData
}

var file_tools_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tools_proto_goTypes = []interface{}{
	(*MorseReq)(nil),     // 0: tools.MorseReq
	(*MorseResp)(nil),    // 1: tools.MorseResp
	(*QrCodeReq)(nil),    // 2: tools.QrCodeReq
	(*QrCodeResp)(nil),   // 3: tools.QrCodeResp
	(*Rgb2HexReq)(nil),   // 4: tools.Rgb2HexReq
	(*Rgb2HexResp)(nil),  // 5: tools.Rgb2HexResp
	(*GetStoryReq)(nil),  // 6: tools.GetStoryReq
	(*GetStoryResp)(nil), // 7: tools.GetStoryResp
}
var file_tools_proto_depIdxs = []int32{
	0, // 0: tools.Tools.Morse:input_type -> tools.MorseReq
	2, // 1: tools.Tools.QrCode:input_type -> tools.QrCodeReq
	4, // 2: tools.Tools.Rgb2Hex:input_type -> tools.Rgb2HexReq
	6, // 3: tools.Tools.Story:input_type -> tools.GetStoryReq
	1, // 4: tools.Tools.Morse:output_type -> tools.MorseResp
	3, // 5: tools.Tools.QrCode:output_type -> tools.QrCodeResp
	5, // 6: tools.Tools.Rgb2Hex:output_type -> tools.Rgb2HexResp
	7, // 7: tools.Tools.Story:output_type -> tools.GetStoryResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tools_proto_init() }
func file_tools_proto_init() {
	if File_tools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tools_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorseReq); i {
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
		file_tools_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MorseResp); i {
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
		file_tools_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrCodeReq); i {
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
		file_tools_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QrCodeResp); i {
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
		file_tools_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rgb2HexReq); i {
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
		file_tools_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rgb2HexResp); i {
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
		file_tools_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoryReq); i {
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
		file_tools_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoryResp); i {
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
			RawDescriptor: file_tools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tools_proto_goTypes,
		DependencyIndexes: file_tools_proto_depIdxs,
		MessageInfos:      file_tools_proto_msgTypes,
	}.Build()
	File_tools_proto = out.File
	file_tools_proto_rawDesc = nil
	file_tools_proto_goTypes = nil
	file_tools_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToolsClient is the client API for Tools service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToolsClient interface {
	Morse(ctx context.Context, in *MorseReq, opts ...grpc.CallOption) (*MorseResp, error)
	QrCode(ctx context.Context, in *QrCodeReq, opts ...grpc.CallOption) (*QrCodeResp, error)
	Rgb2Hex(ctx context.Context, in *Rgb2HexReq, opts ...grpc.CallOption) (*Rgb2HexResp, error)
	Story(ctx context.Context, in *GetStoryReq, opts ...grpc.CallOption) (*GetStoryResp, error)
}

type toolsClient struct {
	cc grpc.ClientConnInterface
}

func NewToolsClient(cc grpc.ClientConnInterface) ToolsClient {
	return &toolsClient{cc}
}

func (c *toolsClient) Morse(ctx context.Context, in *MorseReq, opts ...grpc.CallOption) (*MorseResp, error) {
	out := new(MorseResp)
	err := c.cc.Invoke(ctx, "/tools.Tools/Morse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) QrCode(ctx context.Context, in *QrCodeReq, opts ...grpc.CallOption) (*QrCodeResp, error) {
	out := new(QrCodeResp)
	err := c.cc.Invoke(ctx, "/tools.Tools/QrCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) Rgb2Hex(ctx context.Context, in *Rgb2HexReq, opts ...grpc.CallOption) (*Rgb2HexResp, error) {
	out := new(Rgb2HexResp)
	err := c.cc.Invoke(ctx, "/tools.Tools/Rgb2Hex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) Story(ctx context.Context, in *GetStoryReq, opts ...grpc.CallOption) (*GetStoryResp, error) {
	out := new(GetStoryResp)
	err := c.cc.Invoke(ctx, "/tools.Tools/Story", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToolsServer is the server API for Tools service.
type ToolsServer interface {
	Morse(context.Context, *MorseReq) (*MorseResp, error)
	QrCode(context.Context, *QrCodeReq) (*QrCodeResp, error)
	Rgb2Hex(context.Context, *Rgb2HexReq) (*Rgb2HexResp, error)
	Story(context.Context, *GetStoryReq) (*GetStoryResp, error)
}

// UnimplementedToolsServer can be embedded to have forward compatible implementations.
type UnimplementedToolsServer struct {
}

func (*UnimplementedToolsServer) Morse(context.Context, *MorseReq) (*MorseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Morse not implemented")
}
func (*UnimplementedToolsServer) QrCode(context.Context, *QrCodeReq) (*QrCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QrCode not implemented")
}
func (*UnimplementedToolsServer) Rgb2Hex(context.Context, *Rgb2HexReq) (*Rgb2HexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rgb2Hex not implemented")
}
func (*UnimplementedToolsServer) Story(context.Context, *GetStoryReq) (*GetStoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Story not implemented")
}

func RegisterToolsServer(s *grpc.Server, srv ToolsServer) {
	s.RegisterService(&_Tools_serviceDesc, srv)
}

func _Tools_Morse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MorseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).Morse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tools.Tools/Morse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).Morse(ctx, req.(*MorseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_QrCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QrCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).QrCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tools.Tools/QrCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).QrCode(ctx, req.(*QrCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_Rgb2Hex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rgb2HexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).Rgb2Hex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tools.Tools/Rgb2Hex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).Rgb2Hex(ctx, req.(*Rgb2HexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_Story_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).Story(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tools.Tools/Story",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).Story(ctx, req.(*GetStoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tools_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tools.Tools",
	HandlerType: (*ToolsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Morse",
			Handler:    _Tools_Morse_Handler,
		},
		{
			MethodName: "QrCode",
			Handler:    _Tools_QrCode_Handler,
		},
		{
			MethodName: "Rgb2Hex",
			Handler:    _Tools_Rgb2Hex_Handler,
		},
		{
			MethodName: "Story",
			Handler:    _Tools_Story_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tools.proto",
}

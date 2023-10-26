// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: proto/multiplayer.proto

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

type RequestType int32

const (
	RequestType_JOIN RequestType = 0
	RequestType_GET  RequestType = 1
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "JOIN",
		1: "GET",
	}
	RequestType_value = map[string]int32{
		"JOIN": 0,
		"GET":  1,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_multiplayer_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_proto_multiplayer_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type RequestType   `protobuf:"varint,1,opt,name=type,proto3,enum=multiplayer.RequestType" json:"type,omitempty"`
	Join *Request_JOIN `protobuf:"bytes,2,opt,name=join,proto3,oneof" json:"join,omitempty"`
	Get  *Request_GET  `protobuf:"bytes,3,opt,name=get,proto3,oneof" json:"get,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetType() RequestType {
	if x != nil {
		return x.Type
	}
	return RequestType_JOIN
}

func (x *Request) GetJoin() *Request_JOIN {
	if x != nil {
		return x.Join
	}
	return nil
}

func (x *Request) GetGet() *Request_GET {
	if x != nil {
		return x.Get
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Join *Response_JOIN `protobuf:"bytes,2,opt,name=join,proto3,oneof" json:"join,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetJoin() *Response_JOIN {
	if x != nil {
		return x.Join
	}
	return nil
}

type Request_JOIN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login  string  `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	MapId  int64   `protobuf:"varint,2,opt,name=mapId,proto3" json:"mapId,omitempty"`
	Health int64   `protobuf:"varint,3,opt,name=health,proto3" json:"health,omitempty"`
	StartX float64 `protobuf:"fixed64,4,opt,name=startX,proto3" json:"startX,omitempty"`
	StartY float64 `protobuf:"fixed64,5,opt,name=startY,proto3" json:"startY,omitempty"`
}

func (x *Request_JOIN) Reset() {
	*x = Request_JOIN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_JOIN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_JOIN) ProtoMessage() {}

func (x *Request_JOIN) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_JOIN.ProtoReflect.Descriptor instead.
func (*Request_JOIN) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_JOIN) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Request_JOIN) GetMapId() int64 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *Request_JOIN) GetHealth() int64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Request_JOIN) GetStartX() float64 {
	if x != nil {
		return x.StartX
	}
	return 0
}

func (x *Request_JOIN) GetStartY() float64 {
	if x != nil {
		return x.StartY
	}
	return 0
}

type Request_GET struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Health int64   `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
	X      float64 `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y      float64 `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
	Dx     float64 `protobuf:"fixed64,5,opt,name=dx,proto3" json:"dx,omitempty"`
	Dy     float64 `protobuf:"fixed64,6,opt,name=dy,proto3" json:"dy,omitempty"`
}

func (x *Request_GET) Reset() {
	*x = Request_GET{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_GET) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_GET) ProtoMessage() {}

func (x *Request_GET) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_GET.ProtoReflect.Descriptor instead.
func (*Request_GET) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Request_GET) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Request_GET) GetHealth() int64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Request_GET) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Request_GET) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Request_GET) GetDx() float64 {
	if x != nil {
		return x.Dx
	}
	return 0
}

func (x *Request_GET) GetDy() float64 {
	if x != nil {
		return x.Dy
	}
	return 0
}

type Response_JOIN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok   bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Response_JOIN) Reset() {
	*x = Response_JOIN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_multiplayer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_JOIN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_JOIN) ProtoMessage() {}

func (x *Response_JOIN) ProtoReflect() protoreflect.Message {
	mi := &file_proto_multiplayer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_JOIN.ProtoReflect.Descriptor instead.
func (*Response_JOIN) Descriptor() ([]byte, []int) {
	return file_proto_multiplayer_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_JOIN) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Response_JOIN) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_proto_multiplayer_proto protoreflect.FileDescriptor

var file_proto_multiplayer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x98, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4a, 0x4f, 0x49, 0x4e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x45, 0x54, 0x48, 0x01, 0x52, 0x03, 0x67,
	0x65, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x7a, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x58, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74, 0x58, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x59, 0x1a, 0x6d, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x64, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x64, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x64, 0x79,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x67, 0x65,
	0x74, 0x22, 0x74, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4a, 0x4f, 0x49, 0x4e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x88,
	0x01, 0x01, 0x1a, 0x2a, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x2a, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_multiplayer_proto_rawDescOnce sync.Once
	file_proto_multiplayer_proto_rawDescData = file_proto_multiplayer_proto_rawDesc
)

func file_proto_multiplayer_proto_rawDescGZIP() []byte {
	file_proto_multiplayer_proto_rawDescOnce.Do(func() {
		file_proto_multiplayer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_multiplayer_proto_rawDescData)
	})
	return file_proto_multiplayer_proto_rawDescData
}

var file_proto_multiplayer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_multiplayer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_multiplayer_proto_goTypes = []interface{}{
	(RequestType)(0),      // 0: multiplayer.RequestType
	(*Request)(nil),       // 1: multiplayer.Request
	(*Response)(nil),      // 2: multiplayer.Response
	(*Request_JOIN)(nil),  // 3: multiplayer.Request.JOIN
	(*Request_GET)(nil),   // 4: multiplayer.Request.GET
	(*Response_JOIN)(nil), // 5: multiplayer.Response.JOIN
}
var file_proto_multiplayer_proto_depIdxs = []int32{
	0, // 0: multiplayer.Request.type:type_name -> multiplayer.RequestType
	3, // 1: multiplayer.Request.join:type_name -> multiplayer.Request.JOIN
	4, // 2: multiplayer.Request.get:type_name -> multiplayer.Request.GET
	5, // 3: multiplayer.Response.join:type_name -> multiplayer.Response.JOIN
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_multiplayer_proto_init() }
func file_proto_multiplayer_proto_init() {
	if File_proto_multiplayer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_multiplayer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_multiplayer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_multiplayer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_JOIN); i {
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
		file_proto_multiplayer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_GET); i {
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
		file_proto_multiplayer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_JOIN); i {
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
	file_proto_multiplayer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_multiplayer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_multiplayer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_multiplayer_proto_goTypes,
		DependencyIndexes: file_proto_multiplayer_proto_depIdxs,
		EnumInfos:         file_proto_multiplayer_proto_enumTypes,
		MessageInfos:      file_proto_multiplayer_proto_msgTypes,
	}.Build()
	File_proto_multiplayer_proto = out.File
	file_proto_multiplayer_proto_rawDesc = nil
	file_proto_multiplayer_proto_goTypes = nil
	file_proto_multiplayer_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: v1/health/health.proto

package health

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RetrieveResponse is the response for health service.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v1_health_health_proto_msgTypes[0]
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
	return file_v1_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Response) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type ResPatient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ResPatient) Reset() {
	*x = ResPatient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_health_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPatient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPatient) ProtoMessage() {}

func (x *ResPatient) ProtoReflect() protoreflect.Message {
	mi := &file_v1_health_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPatient.ProtoReflect.Descriptor instead.
func (*ResPatient) Descriptor() ([]byte, []int) {
	return file_v1_health_health_proto_rawDescGZIP(), []int{1}
}

func (x *ResPatient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResPatient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResAllPatients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllPatient []*ResPatient `protobuf:"bytes,1,rep,name=allPatient,proto3" json:"allPatient,omitempty"`
}

func (x *ResAllPatients) Reset() {
	*x = ResAllPatients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_health_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResAllPatients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResAllPatients) ProtoMessage() {}

func (x *ResAllPatients) ProtoReflect() protoreflect.Message {
	mi := &file_v1_health_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResAllPatients.ProtoReflect.Descriptor instead.
func (*ResAllPatients) Descriptor() ([]byte, []int) {
	return file_v1_health_health_proto_rawDescGZIP(), []int{2}
}

func (x *ResAllPatients) GetAllPatient() []*ResPatient {
	if x != nil {
		return x.AllPatient
	}
	return nil
}

type ReqPatient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReqPatient) Reset() {
	*x = ReqPatient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_health_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPatient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPatient) ProtoMessage() {}

func (x *ReqPatient) ProtoReflect() protoreflect.Message {
	mi := &file_v1_health_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPatient.ProtoReflect.Descriptor instead.
func (*ReqPatient) Descriptor() ([]byte, []int) {
	return file_v1_health_health_proto_rawDescGZIP(), []int{3}
}

func (x *ReqPatient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqPatient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_v1_health_health_proto protoreflect.FileDescriptor

var file_v1_health_health_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x22, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x30, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xeb, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x65,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x71, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x62, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x73, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_health_health_proto_rawDescOnce sync.Once
	file_v1_health_health_proto_rawDescData = file_v1_health_health_proto_rawDesc
)

func file_v1_health_health_proto_rawDescGZIP() []byte {
	file_v1_health_health_proto_rawDescOnce.Do(func() {
		file_v1_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_health_health_proto_rawDescData)
	})
	return file_v1_health_health_proto_rawDescData
}

var file_v1_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_health_health_proto_goTypes = []interface{}{
	(*Response)(nil),       // 0: api.learnProto.bri.v1.health.Response
	(*ResPatient)(nil),     // 1: api.learnProto.bri.v1.health.ResPatient
	(*ResAllPatients)(nil), // 2: api.learnProto.bri.v1.health.ResAllPatients
	(*ReqPatient)(nil),     // 3: api.learnProto.bri.v1.health.ReqPatient
	(*emptypb.Empty)(nil),  // 4: google.protobuf.Empty
}
var file_v1_health_health_proto_depIdxs = []int32{
	1, // 0: api.learnProto.bri.v1.health.ResAllPatients.allPatient:type_name -> api.learnProto.bri.v1.health.ResPatient
	4, // 1: api.learnProto.bri.v1.health.HealthService.Get:input_type -> google.protobuf.Empty
	4, // 2: api.learnProto.bri.v1.health.HealthService.GetPatient:input_type -> google.protobuf.Empty
	3, // 3: api.learnProto.bri.v1.health.HealthService.CreatePatient:input_type -> api.learnProto.bri.v1.health.ReqPatient
	4, // 4: api.learnProto.bri.v1.health.HealthService.GetAllPatients:input_type -> google.protobuf.Empty
	0, // 5: api.learnProto.bri.v1.health.HealthService.Get:output_type -> api.learnProto.bri.v1.health.Response
	1, // 6: api.learnProto.bri.v1.health.HealthService.GetPatient:output_type -> api.learnProto.bri.v1.health.ResPatient
	1, // 7: api.learnProto.bri.v1.health.HealthService.CreatePatient:output_type -> api.learnProto.bri.v1.health.ResPatient
	2, // 8: api.learnProto.bri.v1.health.HealthService.GetAllPatients:output_type -> api.learnProto.bri.v1.health.ResAllPatients
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_health_health_proto_init() }
func file_v1_health_health_proto_init() {
	if File_v1_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_health_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPatient); i {
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
		file_v1_health_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResAllPatients); i {
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
		file_v1_health_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPatient); i {
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
			RawDescriptor: file_v1_health_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_health_health_proto_goTypes,
		DependencyIndexes: file_v1_health_health_proto_depIdxs,
		MessageInfos:      file_v1_health_health_proto_msgTypes,
	}.Build()
	File_v1_health_health_proto = out.File
	file_v1_health_health_proto_rawDesc = nil
	file_v1_health_health_proto_goTypes = nil
	file_v1_health_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error)
	GetPatient(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResPatient, error)
	CreatePatient(ctx context.Context, in *ReqPatient, opts ...grpc.CallOption) (*ResPatient, error)
	GetAllPatients(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResAllPatients, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.learnProto.bri.v1.health.HealthService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetPatient(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResPatient, error) {
	out := new(ResPatient)
	err := c.cc.Invoke(ctx, "/api.learnProto.bri.v1.health.HealthService/GetPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) CreatePatient(ctx context.Context, in *ReqPatient, opts ...grpc.CallOption) (*ResPatient, error) {
	out := new(ResPatient)
	err := c.cc.Invoke(ctx, "/api.learnProto.bri.v1.health.HealthService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetAllPatients(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResAllPatients, error) {
	out := new(ResAllPatients)
	err := c.cc.Invoke(ctx, "/api.learnProto.bri.v1.health.HealthService/GetAllPatients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	Get(context.Context, *emptypb.Empty) (*Response, error)
	GetPatient(context.Context, *emptypb.Empty) (*ResPatient, error)
	CreatePatient(context.Context, *ReqPatient) (*ResPatient, error)
	GetAllPatients(context.Context, *emptypb.Empty) (*ResAllPatients, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) Get(context.Context, *emptypb.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedHealthServiceServer) GetPatient(context.Context, *emptypb.Empty) (*ResPatient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (*UnimplementedHealthServiceServer) CreatePatient(context.Context, *ReqPatient) (*ResPatient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (*UnimplementedHealthServiceServer) GetAllPatients(context.Context, *emptypb.Empty) (*ResAllPatients, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPatients not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.learnProto.bri.v1.health.HealthService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.learnProto.bri.v1.health.HealthService/GetPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetPatient(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPatient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.learnProto.bri.v1.health.HealthService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).CreatePatient(ctx, req.(*ReqPatient))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetAllPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetAllPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.learnProto.bri.v1.health.HealthService/GetAllPatients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetAllPatients(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.learnProto.bri.v1.health.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _HealthService_Get_Handler,
		},
		{
			MethodName: "GetPatient",
			Handler:    _HealthService_GetPatient_Handler,
		},
		{
			MethodName: "CreatePatient",
			Handler:    _HealthService_CreatePatient_Handler,
		},
		{
			MethodName: "GetAllPatients",
			Handler:    _HealthService_GetAllPatients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/health/health.proto",
}
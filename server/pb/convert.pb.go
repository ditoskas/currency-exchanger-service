// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: convert.proto

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

type ConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	FromCurrency string  `protobuf:"bytes,2,opt,name=fromCurrency,proto3" json:"fromCurrency,omitempty"`
	ToCurrency   string  `protobuf:"bytes,3,opt,name=toCurrency,proto3" json:"toCurrency,omitempty"`
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_convert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_convert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_convert_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ConvertRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ConvertRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

type ConvertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountTo       float64 `protobuf:"fixed64,1,opt,name=amountTo,proto3" json:"amountTo,omitempty"`
	ToCurrency     string  `protobuf:"bytes,2,opt,name=toCurrency,proto3" json:"toCurrency,omitempty"`
	AmountOriginal float64 `protobuf:"fixed64,3,opt,name=amountOriginal,proto3" json:"amountOriginal,omitempty"`
	FromCurrency   string  `protobuf:"bytes,4,opt,name=fromCurrency,proto3" json:"fromCurrency,omitempty"`
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_convert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_convert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_convert_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertResponse) GetAmountTo() float64 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *ConvertResponse) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ConvertResponse) GetAmountOriginal() float64 {
	if x != nil {
		return x.AmountOriginal
	}
	return 0
}

func (x *ConvertResponse) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

var File_convert_proto protoreflect.FileDescriptor

var file_convert_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x72, 0x22, 0x6c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f,
	0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x32, 0x64, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x74, 0x6f, 0x73, 0x6b, 0x61, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_convert_proto_rawDescOnce sync.Once
	file_convert_proto_rawDescData = file_convert_proto_rawDesc
)

func file_convert_proto_rawDescGZIP() []byte {
	file_convert_proto_rawDescOnce.Do(func() {
		file_convert_proto_rawDescData = protoimpl.X.CompressGZIP(file_convert_proto_rawDescData)
	})
	return file_convert_proto_rawDescData
}

var file_convert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_convert_proto_goTypes = []interface{}{
	(*ConvertRequest)(nil),  // 0: currency_exchanger.ConvertRequest
	(*ConvertResponse)(nil), // 1: currency_exchanger.ConvertResponse
}
var file_convert_proto_depIdxs = []int32{
	0, // 0: currency_exchanger.ConvertService.Convert:input_type -> currency_exchanger.ConvertRequest
	1, // 1: currency_exchanger.ConvertService.Convert:output_type -> currency_exchanger.ConvertResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_convert_proto_init() }
func file_convert_proto_init() {
	if File_convert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_convert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertRequest); i {
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
		file_convert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertResponse); i {
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
			RawDescriptor: file_convert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_convert_proto_goTypes,
		DependencyIndexes: file_convert_proto_depIdxs,
		MessageInfos:      file_convert_proto_msgTypes,
	}.Build()
	File_convert_proto = out.File
	file_convert_proto_rawDesc = nil
	file_convert_proto_goTypes = nil
	file_convert_proto_depIdxs = nil
}

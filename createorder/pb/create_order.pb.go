// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: create_order.proto

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

type OrderPlaced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity    int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *OrderPlaced) Reset() {
	*x = OrderPlaced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_create_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPlaced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPlaced) ProtoMessage() {}

func (x *OrderPlaced) ProtoReflect() protoreflect.Message {
	mi := &file_create_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPlaced.ProtoReflect.Descriptor instead.
func (*OrderPlaced) Descriptor() ([]byte, []int) {
	return file_create_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderPlaced) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OrderPlaced) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrderPlaced) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderPlaced) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_create_order_proto protoreflect.FileDescriptor

var file_create_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x77, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_create_order_proto_rawDescOnce sync.Once
	file_create_order_proto_rawDescData = file_create_order_proto_rawDesc
)

func file_create_order_proto_rawDescGZIP() []byte {
	file_create_order_proto_rawDescOnce.Do(func() {
		file_create_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_create_order_proto_rawDescData)
	})
	return file_create_order_proto_rawDescData
}

var file_create_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_create_order_proto_goTypes = []interface{}{
	(*OrderPlaced)(nil), // 0: createorder.OrderPlaced
}
var file_create_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_create_order_proto_init() }
func file_create_order_proto_init() {
	if File_create_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_create_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPlaced); i {
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
			RawDescriptor: file_create_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_create_order_proto_goTypes,
		DependencyIndexes: file_create_order_proto_depIdxs,
		MessageInfos:      file_create_order_proto_msgTypes,
	}.Build()
	File_create_order_proto = out.File
	file_create_order_proto_rawDesc = nil
	file_create_order_proto_goTypes = nil
	file_create_order_proto_depIdxs = nil
}

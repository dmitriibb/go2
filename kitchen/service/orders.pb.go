// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: kitchen/service/orders.proto

package service

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

type OrderEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *OrderEvent) Reset() {
	*x = OrderEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_service_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEvent) ProtoMessage() {}

func (x *OrderEvent) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_service_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEvent.ProtoReflect.Descriptor instead.
func (*OrderEvent) Descriptor() ([]byte, []int) {
	return file_kitchen_service_orders_proto_rawDescGZIP(), []int{0}
}

func (x *OrderEvent) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type OrderEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderEventResponse) Reset() {
	*x = OrderEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_service_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderEventResponse) ProtoMessage() {}

func (x *OrderEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_service_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderEventResponse.ProtoReflect.Descriptor instead.
func (*OrderEventResponse) Descriptor() ([]byte, []int) {
	return file_kitchen_service_orders_proto_rawDescGZIP(), []int{1}
}

func (x *OrderEventResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderEventResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_kitchen_service_orders_proto protoreflect.FileDescriptor

var file_kitchen_service_orders_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x32, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x48, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x6d, 0x0a, 0x0e, 0x4b, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x32, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e,
	0x67, 0x6f, 0x32, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x64, 0x6d, 0x62, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x32, 0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitchen_service_orders_proto_rawDescOnce sync.Once
	file_kitchen_service_orders_proto_rawDescData = file_kitchen_service_orders_proto_rawDesc
)

func file_kitchen_service_orders_proto_rawDescGZIP() []byte {
	file_kitchen_service_orders_proto_rawDescOnce.Do(func() {
		file_kitchen_service_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitchen_service_orders_proto_rawDescData)
	})
	return file_kitchen_service_orders_proto_rawDescData
}

var file_kitchen_service_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kitchen_service_orders_proto_goTypes = []interface{}{
	(*OrderEvent)(nil),         // 0: go2.kitchen.service.OrderEvent
	(*OrderEventResponse)(nil), // 1: go2.kitchen.service.OrderEventResponse
}
var file_kitchen_service_orders_proto_depIdxs = []int32{
	0, // 0: go2.kitchen.service.KitchenService.NewOrderEvent:input_type -> go2.kitchen.service.OrderEvent
	1, // 1: go2.kitchen.service.KitchenService.NewOrderEvent:output_type -> go2.kitchen.service.OrderEventResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kitchen_service_orders_proto_init() }
func file_kitchen_service_orders_proto_init() {
	if File_kitchen_service_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitchen_service_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEvent); i {
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
		file_kitchen_service_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderEventResponse); i {
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
			RawDescriptor: file_kitchen_service_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kitchen_service_orders_proto_goTypes,
		DependencyIndexes: file_kitchen_service_orders_proto_depIdxs,
		MessageInfos:      file_kitchen_service_orders_proto_msgTypes,
	}.Build()
	File_kitchen_service_orders_proto = out.File
	file_kitchen_service_orders_proto_rawDesc = nil
	file_kitchen_service_orders_proto_goTypes = nil
	file_kitchen_service_orders_proto_depIdxs = nil
}

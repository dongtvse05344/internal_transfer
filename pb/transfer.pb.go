// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.20.3
// source: transfer.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountId int64   `protobuf:"varint,1,opt,name=from_account_id,json=fromAccountId,proto3" json:"from_account_id,omitempty"`
	ToAccountId   int64   `protobuf:"varint,2,opt,name=to_account_id,json=toAccountId,proto3" json:"to_account_id,omitempty"`
	Amount        float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetFromAccountId() int64 {
	if x != nil {
		return x.FromAccountId
	}
	return 0
}

func (x *TransferRequest) GetToAccountId() int64 {
	if x != nil {
		return x.ToAccountId
	}
	return 0
}

func (x *TransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       *int64  `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message    *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	TransferId *int64  `protobuf:"varint,3,opt,name=transfer_id,json=transferId,proto3,oneof" json:"transfer_id,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetCode() int64 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *TransferResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *TransferResponse) GetTransferId() int64 {
	if x != nil && x.TransferId != nil {
		return *x.TransferId
	}
	return 0
}

type GetTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTransferRequest) Reset() {
	*x = GetTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferRequest) ProtoMessage() {}

func (x *GetTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferRequest.ProtoReflect.Descriptor instead.
func (*GetTransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransferRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     *int64    `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message  *string   `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Transfer *Transfer `protobuf:"bytes,3,opt,name=transfer,proto3,oneof" json:"transfer,omitempty"`
}

func (x *GetTransferResponse) Reset() {
	*x = GetTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferResponse) ProtoMessage() {}

func (x *GetTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferResponse.ProtoReflect.Descriptor instead.
func (*GetTransferResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransferResponse) GetCode() int64 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *GetTransferResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetTransferResponse) GetTransfer() *Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId    *int64   `protobuf:"varint,1,opt,name=transfer_id,json=transferId,proto3,oneof" json:"transfer_id,omitempty"`
	FromAccountId *int64   `protobuf:"varint,2,opt,name=from_account_id,json=fromAccountId,proto3,oneof" json:"from_account_id,omitempty"`
	ToAccountId   *int64   `protobuf:"varint,3,opt,name=to_account_id,json=toAccountId,proto3,oneof" json:"to_account_id,omitempty"`
	Amount        *float64 `protobuf:"fixed64,4,opt,name=amount,proto3,oneof" json:"amount,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *Transfer) GetTransferId() int64 {
	if x != nil && x.TransferId != nil {
		return *x.TransferId
	}
	return 0
}

func (x *Transfer) GetFromAccountId() int64 {
	if x != nil && x.FromAccountId != nil {
		return *x.FromAccountId
	}
	return 0
}

func (x *Transfer) GetToAccountId() int64 {
	if x != nil && x.ToAccountId != nil {
		return *x.ToAccountId
	}
	return 0
}

func (x *Transfer) GetAmount() float64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x75, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x02, 0x52, 0x08, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x22, 0xe4, 0x01, 0x0a, 0x08, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a,
	0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x74, 0x6f,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transfer_proto_goTypes = []interface{}{
	(*TransferRequest)(nil),     // 0: pb.TransferRequest
	(*TransferResponse)(nil),    // 1: pb.TransferResponse
	(*GetTransferRequest)(nil),  // 2: pb.GetTransferRequest
	(*GetTransferResponse)(nil), // 3: pb.GetTransferResponse
	(*Transfer)(nil),            // 4: pb.Transfer
}
var file_transfer_proto_depIdxs = []int32{
	4, // 0: pb.GetTransferResponse.transfer:type_name -> pb.Transfer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
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
		file_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransferRequest); i {
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
		file_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransferResponse); i {
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
		file_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
	file_transfer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_transfer_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_transfer_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}

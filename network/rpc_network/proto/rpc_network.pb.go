// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v3.12.4
// source: rpc_network.proto

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

type GetBlockchainMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId string `protobuf:"bytes,1,opt,name=peerId,proto3" json:"peerId,omitempty"`
}

func (x *GetBlockchainMessage) Reset() {
	*x = GetBlockchainMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainMessage) ProtoMessage() {}

func (x *GetBlockchainMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockchainMessage.ProtoReflect.Descriptor instead.
func (*GetBlockchainMessage) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockchainMessage) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PreviousHash string `protobuf:"bytes,2,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	Timestamp    int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce        int64  `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Hash         string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Block) GetPreviousHash() string {
	if x != nil {
		return x.PreviousHash
	}
	return ""
}

func (x *Block) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Block) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Blockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *Blockchain) Reset() {
	*x = Blockchain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blockchain) ProtoMessage() {}

func (x *Blockchain) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blockchain.ProtoReflect.Descriptor instead.
func (*Blockchain) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{2}
}

func (x *Blockchain) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type SendBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *SendBlockResponse) Reset() {
	*x = SendBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlockResponse) ProtoMessage() {}

func (x *SendBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlockResponse.ProtoReflect.Descriptor instead.
func (*SendBlockResponse) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{3}
}

func (x *SendBlockResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type SendBlockchainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *SendBlockchainResponse) Reset() {
	*x = SendBlockchainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBlockchainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlockchainResponse) ProtoMessage() {}

func (x *SendBlockchainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlockchainResponse.ProtoReflect.Descriptor instead.
func (*SendBlockchainResponse) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{4}
}

func (x *SendBlockchainResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type GetBlockchainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlockchainResponse) Reset() {
	*x = GetBlockchainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainResponse) ProtoMessage() {}

func (x *GetBlockchainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockchainResponse.ProtoReflect.Descriptor instead.
func (*GetBlockchainResponse) Descriptor() ([]byte, []int) {
	return file_rpc_network_proto_rawDescGZIP(), []int{5}
}

var File_rpc_network_proto protoreflect.FileDescriptor

var file_rpc_network_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x87, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x38, 0x0a, 0x0a, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x22, 0x2f, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf5, 0x01, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a,
	0x23, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x21, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_network_proto_rawDescOnce sync.Once
	file_rpc_network_proto_rawDescData = file_rpc_network_proto_rawDesc
)

func file_rpc_network_proto_rawDescGZIP() []byte {
	file_rpc_network_proto_rawDescOnce.Do(func() {
		file_rpc_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_network_proto_rawDescData)
	})
	return file_rpc_network_proto_rawDescData
}

var file_rpc_network_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_network_proto_goTypes = []interface{}{
	(*GetBlockchainMessage)(nil),   // 0: rpc_network.GetBlockchainMessage
	(*Block)(nil),                  // 1: rpc_network.Block
	(*Blockchain)(nil),             // 2: rpc_network.Blockchain
	(*SendBlockResponse)(nil),      // 3: rpc_network.SendBlockResponse
	(*SendBlockchainResponse)(nil), // 4: rpc_network.SendBlockchainResponse
	(*GetBlockchainResponse)(nil),  // 5: rpc_network.GetBlockchainResponse
}
var file_rpc_network_proto_depIdxs = []int32{
	1, // 0: rpc_network.Blockchain.blocks:type_name -> rpc_network.Block
	1, // 1: rpc_network.RpcNetwork.SendBlock:input_type -> rpc_network.Block
	2, // 2: rpc_network.RpcNetwork.SendBlockchain:input_type -> rpc_network.Blockchain
	0, // 3: rpc_network.RpcNetwork.GetBlockchain:input_type -> rpc_network.GetBlockchainMessage
	3, // 4: rpc_network.RpcNetwork.SendBlock:output_type -> rpc_network.SendBlockResponse
	4, // 5: rpc_network.RpcNetwork.SendBlockchain:output_type -> rpc_network.SendBlockchainResponse
	5, // 6: rpc_network.RpcNetwork.GetBlockchain:output_type -> rpc_network.GetBlockchainResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_network_proto_init() }
func file_rpc_network_proto_init() {
	if File_rpc_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockchainMessage); i {
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
		file_rpc_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_rpc_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blockchain); i {
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
		file_rpc_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBlockResponse); i {
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
		file_rpc_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBlockchainResponse); i {
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
		file_rpc_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockchainResponse); i {
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
			RawDescriptor: file_rpc_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_network_proto_goTypes,
		DependencyIndexes: file_rpc_network_proto_depIdxs,
		MessageInfos:      file_rpc_network_proto_msgTypes,
	}.Build()
	File_rpc_network_proto = out.File
	file_rpc_network_proto_rawDesc = nil
	file_rpc_network_proto_goTypes = nil
	file_rpc_network_proto_depIdxs = nil
}
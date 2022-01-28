// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: src/blockchain_txn_rewards_v2.proto

package transaction

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

type BlockchainTxnRewardV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account []byte `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BlockchainTxnRewardV2) Reset() {
	*x = BlockchainTxnRewardV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_blockchain_txn_rewards_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnRewardV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnRewardV2) ProtoMessage() {}

func (x *BlockchainTxnRewardV2) ProtoReflect() protoreflect.Message {
	mi := &file_src_blockchain_txn_rewards_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnRewardV2.ProtoReflect.Descriptor instead.
func (*BlockchainTxnRewardV2) Descriptor() ([]byte, []int) {
	return file_src_blockchain_txn_rewards_v2_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnRewardV2) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *BlockchainTxnRewardV2) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BlockchainTxnRewardsV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartEpoch uint64                   `protobuf:"varint,1,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch   uint64                   `protobuf:"varint,2,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Rewards    []*BlockchainTxnRewardV2 `protobuf:"bytes,3,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *BlockchainTxnRewardsV2) Reset() {
	*x = BlockchainTxnRewardsV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_blockchain_txn_rewards_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnRewardsV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnRewardsV2) ProtoMessage() {}

func (x *BlockchainTxnRewardsV2) ProtoReflect() protoreflect.Message {
	mi := &file_src_blockchain_txn_rewards_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnRewardsV2.ProtoReflect.Descriptor instead.
func (*BlockchainTxnRewardsV2) Descriptor() ([]byte, []int) {
	return file_src_blockchain_txn_rewards_v2_proto_rawDescGZIP(), []int{1}
}

func (x *BlockchainTxnRewardsV2) GetStartEpoch() uint64 {
	if x != nil {
		return x.StartEpoch
	}
	return 0
}

func (x *BlockchainTxnRewardsV2) GetEndEpoch() uint64 {
	if x != nil {
		return x.EndEpoch
	}
	return 0
}

func (x *BlockchainTxnRewardsV2) GetRewards() []*BlockchainTxnRewardV2 {
	if x != nil {
		return x.Rewards
	}
	return nil
}

var File_src_blockchain_txn_rewards_v2_proto protoreflect.FileDescriptor

var file_src_blockchain_txn_rewards_v2_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x72, 0x63, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x76, 0x32, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68,
	0x22, 0x4c, 0x0a, 0x18, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x78, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x98,
	0x01, 0x0a, 0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78,
	0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x76, 0x32, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x32,
	0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_blockchain_txn_rewards_v2_proto_rawDescOnce sync.Once
	file_src_blockchain_txn_rewards_v2_proto_rawDescData = file_src_blockchain_txn_rewards_v2_proto_rawDesc
)

func file_src_blockchain_txn_rewards_v2_proto_rawDescGZIP() []byte {
	file_src_blockchain_txn_rewards_v2_proto_rawDescOnce.Do(func() {
		file_src_blockchain_txn_rewards_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_blockchain_txn_rewards_v2_proto_rawDescData)
	})
	return file_src_blockchain_txn_rewards_v2_proto_rawDescData
}

var file_src_blockchain_txn_rewards_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_blockchain_txn_rewards_v2_proto_goTypes = []interface{}{
	(*BlockchainTxnRewardV2)(nil),  // 0: smartmesh.blockchain_txn_reward_v2
	(*BlockchainTxnRewardsV2)(nil), // 1: smartmesh.blockchain_txn_rewards_v2
}
var file_src_blockchain_txn_rewards_v2_proto_depIdxs = []int32{
	0, // 0: smartmesh.blockchain_txn_rewards_v2.rewards:type_name -> smartmesh.blockchain_txn_reward_v2
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_blockchain_txn_rewards_v2_proto_init() }
func file_src_blockchain_txn_rewards_v2_proto_init() {
	if File_src_blockchain_txn_rewards_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_blockchain_txn_rewards_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnRewardV2); i {
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
		file_src_blockchain_txn_rewards_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnRewardsV2); i {
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
			RawDescriptor: file_src_blockchain_txn_rewards_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_blockchain_txn_rewards_v2_proto_goTypes,
		DependencyIndexes: file_src_blockchain_txn_rewards_v2_proto_depIdxs,
		MessageInfos:      file_src_blockchain_txn_rewards_v2_proto_msgTypes,
	}.Build()
	File_src_blockchain_txn_rewards_v2_proto = out.File
	file_src_blockchain_txn_rewards_v2_proto_rawDesc = nil
	file_src_blockchain_txn_rewards_v2_proto_goTypes = nil
	file_src_blockchain_txn_rewards_v2_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: src/service/local.proto

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

type PubkeyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PubkeyRes) Reset() {
	*x = PubkeyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubkeyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubkeyRes) ProtoMessage() {}

func (x *PubkeyRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubkeyRes.ProtoReflect.Descriptor instead.
func (*PubkeyRes) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{0}
}

func (x *PubkeyRes) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type PubkeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PubkeyReq) Reset() {
	*x = PubkeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubkeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubkeyReq) ProtoMessage() {}

func (x *PubkeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubkeyReq.ProtoReflect.Descriptor instead.
func (*PubkeyReq) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{1}
}

type SignReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SignReq) Reset() {
	*x = SignReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignReq) ProtoMessage() {}

func (x *SignReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignReq.ProtoReflect.Descriptor instead.
func (*SignReq) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{2}
}

func (x *SignReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SignRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignRes) Reset() {
	*x = SignRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRes) ProtoMessage() {}

func (x *SignRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRes.ProtoReflect.Descriptor instead.
func (*SignRes) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{3}
}

func (x *SignRes) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EcdhReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *EcdhReq) Reset() {
	*x = EcdhReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhReq) ProtoMessage() {}

func (x *EcdhReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhReq.ProtoReflect.Descriptor instead.
func (*EcdhReq) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{4}
}

func (x *EcdhReq) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type EcdhRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret []byte `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *EcdhRes) Reset() {
	*x = EcdhRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcdhRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdhRes) ProtoMessage() {}

func (x *EcdhRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdhRes.ProtoReflect.Descriptor instead.
func (*EcdhRes) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{5}
}

func (x *EcdhRes) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

type ConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigReq) Reset() {
	*x = ConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigReq) ProtoMessage() {}

func (x *ConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigReq.ProtoReflect.Descriptor instead.
func (*ConfigReq) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ConfigRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*ConfigValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ConfigRes) Reset() {
	*x = ConfigRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRes) ProtoMessage() {}

func (x *ConfigRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRes.ProtoReflect.Descriptor instead.
func (*ConfigRes) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigRes) GetValues() []*ConfigValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type ConfigValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigValue) Reset() {
	*x = ConfigValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigValue) ProtoMessage() {}

func (x *ConfigValue) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigValue.ProtoReflect.Descriptor instead.
func (*ConfigValue) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{8}
}

func (x *ConfigValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigValue) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ConfigValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KeyedUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Uri     string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *KeyedUri) Reset() {
	*x = KeyedUri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyedUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyedUri) ProtoMessage() {}

func (x *KeyedUri) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyedUri.ProtoReflect.Descriptor instead.
func (*KeyedUri) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{9}
}

func (x *KeyedUri) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *KeyedUri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type HeightReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeightReq) Reset() {
	*x = HeightReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeightReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeightReq) ProtoMessage() {}

func (x *HeightReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeightReq.ProtoReflect.Descriptor instead.
func (*HeightReq) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{10}
}

type HeightRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height   uint64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	BlockAge uint64    `protobuf:"varint,2,opt,name=block_age,json=blockAge,proto3" json:"block_age,omitempty"`
	Gateway  *KeyedUri `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *HeightRes) Reset() {
	*x = HeightRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_service_local_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeightRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeightRes) ProtoMessage() {}

func (x *HeightRes) ProtoReflect() protoreflect.Message {
	mi := &file_src_service_local_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeightRes.ProtoReflect.Descriptor instead.
func (*HeightRes) Descriptor() ([]byte, []int) {
	return file_src_service_local_proto_rawDescGZIP(), []int{11}
}

func (x *HeightRes) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *HeightRes) GetBlockAge() uint64 {
	if x != nil {
		return x.BlockAge
	}
	return 0
}

func (x *HeightRes) GetGateway() *KeyedUri {
	if x != nil {
		return x.Gateway
	}
	return nil
}

var File_src_service_local_proto protoreflect.FileDescriptor

var file_src_service_local_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x26, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71,
	0x22, 0x1e, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x28, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x65, 0x63,
	0x64, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x22, 0x0a, 0x08, 0x65, 0x63, 0x64, 0x68, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x20, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x72, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x65, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x22, 0x0c, 0x0a, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x22, 0x77, 0x0a, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x41, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x69,
	0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0xcd, 0x02, 0x0a, 0x03, 0x61, 0x70,
	0x69, 0x12, 0x42, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x19, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x65, 0x63, 0x64, 0x68, 0x12, 0x19, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x65, 0x63,
	0x64, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x65, 0x63, 0x64, 0x68, 0x5f, 0x72, 0x65,
	0x73, 0x12, 0x42, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x2e, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_service_local_proto_rawDescOnce sync.Once
	file_src_service_local_proto_rawDescData = file_src_service_local_proto_rawDesc
)

func file_src_service_local_proto_rawDescGZIP() []byte {
	file_src_service_local_proto_rawDescOnce.Do(func() {
		file_src_service_local_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_service_local_proto_rawDescData)
	})
	return file_src_service_local_proto_rawDescData
}

var file_src_service_local_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_src_service_local_proto_goTypes = []interface{}{
	(*PubkeyRes)(nil),   // 0: smartmesh.local.pubkey_res
	(*PubkeyReq)(nil),   // 1: smartmesh.local.pubkey_req
	(*SignReq)(nil),     // 2: smartmesh.local.sign_req
	(*SignRes)(nil),     // 3: smartmesh.local.sign_res
	(*EcdhReq)(nil),     // 4: smartmesh.local.ecdh_req
	(*EcdhRes)(nil),     // 5: smartmesh.local.ecdh_res
	(*ConfigReq)(nil),   // 6: smartmesh.local.config_req
	(*ConfigRes)(nil),   // 7: smartmesh.local.config_res
	(*ConfigValue)(nil), // 8: smartmesh.local.config_value
	(*KeyedUri)(nil),    // 9: smartmesh.local.keyed_uri
	(*HeightReq)(nil),   // 10: smartmesh.local.height_req
	(*HeightRes)(nil),   // 11: smartmesh.local.height_res
}
var file_src_service_local_proto_depIdxs = []int32{
	8,  // 0: smartmesh.local.config_res.values:type_name -> smartmesh.local.config_value
	9,  // 1: smartmesh.local.height_res.gateway:type_name -> smartmesh.local.keyed_uri
	1,  // 2: smartmesh.local.api.pubkey:input_type -> smartmesh.local.pubkey_req
	2,  // 3: smartmesh.local.api.sign:input_type -> smartmesh.local.sign_req
	4,  // 4: smartmesh.local.api.ecdh:input_type -> smartmesh.local.ecdh_req
	6,  // 5: smartmesh.local.api.config:input_type -> smartmesh.local.config_req
	10, // 6: smartmesh.local.api.height:input_type -> smartmesh.local.height_req
	0,  // 7: smartmesh.local.api.pubkey:output_type -> smartmesh.local.pubkey_res
	3,  // 8: smartmesh.local.api.sign:output_type -> smartmesh.local.sign_res
	5,  // 9: smartmesh.local.api.ecdh:output_type -> smartmesh.local.ecdh_res
	7,  // 10: smartmesh.local.api.config:output_type -> smartmesh.local.config_res
	11, // 11: smartmesh.local.api.height:output_type -> smartmesh.local.height_res
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_src_service_local_proto_init() }
func file_src_service_local_proto_init() {
	if File_src_service_local_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_service_local_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubkeyRes); i {
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
		file_src_service_local_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubkeyReq); i {
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
		file_src_service_local_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignReq); i {
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
		file_src_service_local_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRes); i {
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
		file_src_service_local_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhReq); i {
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
		file_src_service_local_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcdhRes); i {
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
		file_src_service_local_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigReq); i {
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
		file_src_service_local_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRes); i {
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
		file_src_service_local_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigValue); i {
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
		file_src_service_local_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyedUri); i {
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
		file_src_service_local_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeightReq); i {
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
		file_src_service_local_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeightRes); i {
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
			RawDescriptor: file_src_service_local_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_service_local_proto_goTypes,
		DependencyIndexes: file_src_service_local_proto_depIdxs,
		MessageInfos:      file_src_service_local_proto_msgTypes,
	}.Build()
	File_src_service_local_proto = out.File
	file_src_service_local_proto_rawDesc = nil
	file_src_service_local_proto_goTypes = nil
	file_src_service_local_proto_depIdxs = nil
}

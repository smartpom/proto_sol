// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: src/longfi.proto

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

type LongFiSpreading int32

const (
	LongFiSpreading_SF_INVALID LongFiSpreading = 0
	LongFiSpreading_SF7        LongFiSpreading = 1
	LongFiSpreading_SF8        LongFiSpreading = 2
	LongFiSpreading_SF9        LongFiSpreading = 3
	LongFiSpreading_SF10       LongFiSpreading = 4
)

// Enum value maps for LongFiSpreading.
var (
	LongFiSpreading_name = map[int32]string{
		0: "SF_INVALID",
		1: "SF7",
		2: "SF8",
		3: "SF9",
		4: "SF10",
	}
	LongFiSpreading_value = map[string]int32{
		"SF_INVALID": 0,
		"SF7":        1,
		"SF8":        2,
		"SF9":        3,
		"SF10":       4,
	}
)

func (x LongFiSpreading) Enum() *LongFiSpreading {
	p := new(LongFiSpreading)
	*p = x
	return p
}

func (x LongFiSpreading) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LongFiSpreading) Descriptor() protoreflect.EnumDescriptor {
	return file_src_longfi_proto_enumTypes[0].Descriptor()
}

func (LongFiSpreading) Type() protoreflect.EnumType {
	return &file_src_longfi_proto_enumTypes[0]
}

func (x LongFiSpreading) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LongFiSpreading.Descriptor instead.
func (LongFiSpreading) EnumDescriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{0}
}

type LongFiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Kind:
	//	*LongFiReq_Tx
	Kind isLongFiReq_Kind `protobuf_oneof:"kind"`
}

func (x *LongFiReq) Reset() {
	*x = LongFiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_longfi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongFiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongFiReq) ProtoMessage() {}

func (x *LongFiReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_longfi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongFiReq.ProtoReflect.Descriptor instead.
func (*LongFiReq) Descriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{0}
}

func (x *LongFiReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *LongFiReq) GetKind() isLongFiReq_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *LongFiReq) GetTx() *LongFiTxPacket {
	if x, ok := x.GetKind().(*LongFiReq_Tx); ok {
		return x.Tx
	}
	return nil
}

type isLongFiReq_Kind interface {
	isLongFiReq_Kind()
}

type LongFiReq_Tx struct {
	Tx *LongFiTxPacket `protobuf:"bytes,2,opt,name=tx,proto3,oneof"`
}

func (*LongFiReq_Tx) isLongFiReq_Kind() {}

type LongFiResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Kind:
	//	*LongFiResp_TxStatus
	//	*LongFiResp_Rx
	//	*LongFiResp_ParseErr
	Kind      isLongFiResp_Kind `protobuf_oneof:"kind"`
	MinerName []byte            `protobuf:"bytes,5,opt,name=miner_name,json=minerName,proto3" json:"miner_name,omitempty"`
}

func (x *LongFiResp) Reset() {
	*x = LongFiResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_longfi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongFiResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongFiResp) ProtoMessage() {}

func (x *LongFiResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_longfi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongFiResp.ProtoReflect.Descriptor instead.
func (*LongFiResp) Descriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{1}
}

func (x *LongFiResp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *LongFiResp) GetKind() isLongFiResp_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *LongFiResp) GetTxStatus() *LongFiTxStatus {
	if x, ok := x.GetKind().(*LongFiResp_TxStatus); ok {
		return x.TxStatus
	}
	return nil
}

func (x *LongFiResp) GetRx() *LongFiRxPacket {
	if x, ok := x.GetKind().(*LongFiResp_Rx); ok {
		return x.Rx
	}
	return nil
}

func (x *LongFiResp) GetParseErr() []byte {
	if x, ok := x.GetKind().(*LongFiResp_ParseErr); ok {
		return x.ParseErr
	}
	return nil
}

func (x *LongFiResp) GetMinerName() []byte {
	if x != nil {
		return x.MinerName
	}
	return nil
}

type isLongFiResp_Kind interface {
	isLongFiResp_Kind()
}

type LongFiResp_TxStatus struct {
	TxStatus *LongFiTxStatus `protobuf:"bytes,2,opt,name=tx_status,json=txStatus,proto3,oneof"`
}

type LongFiResp_Rx struct {
	Rx *LongFiRxPacket `protobuf:"bytes,3,opt,name=rx,proto3,oneof"`
}

type LongFiResp_ParseErr struct {
	ParseErr []byte `protobuf:"bytes,4,opt,name=parse_err,json=parseErr,proto3,oneof"`
}

func (*LongFiResp_TxStatus) isLongFiResp_Kind() {}

func (*LongFiResp_Rx) isLongFiResp_Kind() {}

func (*LongFiResp_ParseErr) isLongFiResp_Kind() {}

type LongFiTxStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LongFiTxStatus) Reset() {
	*x = LongFiTxStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_longfi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongFiTxStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongFiTxStatus) ProtoMessage() {}

func (x *LongFiTxStatus) ProtoReflect() protoreflect.Message {
	mi := &file_src_longfi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongFiTxStatus.ProtoReflect.Descriptor instead.
func (*LongFiTxStatus) Descriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{2}
}

func (x *LongFiTxStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type LongFiRxPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of CRC check.
	CrcCheck bool `protobuf:"varint,1,opt,name=crc_check,json=crcCheck,proto3" json:"crc_check,omitempty"`
	// 1uS-resolution timestamp derived from concentrator's internal counter.
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Average packet RSSI in dB.
	Rssi float32 `protobuf:"fixed32,3,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// Average packet SNR, in dB.
	Snr float32 `protobuf:"fixed32,4,opt,name=snr,proto3" json:"snr,omitempty"`
	// Organization Unique ID
	Oui uint32 `protobuf:"varint,5,opt,name=oui,proto3" json:"oui,omitempty"`
	// Device ID
	DeviceId uint32 `protobuf:"varint,6,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Fingerprint
	Fingerprint uint32 `protobuf:"varint,7,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Sequence
	Sequence uint32 `protobuf:"varint,9,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Spreading to be used
	Spreading LongFiSpreading `protobuf:"varint,10,opt,name=spreading,proto3,enum=smartmesh.longfi.LongFiSpreading" json:"spreading,omitempty"`
	// the fully reassembled payload
	Payload []byte `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
	// De-golayed datagram id and flag bits.
	// NOTE: only the lowest 12 bits are valid.
	TagBits uint32 `protobuf:"varint,12,opt,name=tag_bits,json=tagBits,proto3" json:"tag_bits,omitempty"`
}

func (x *LongFiRxPacket) Reset() {
	*x = LongFiRxPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_longfi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongFiRxPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongFiRxPacket) ProtoMessage() {}

func (x *LongFiRxPacket) ProtoReflect() protoreflect.Message {
	mi := &file_src_longfi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongFiRxPacket.ProtoReflect.Descriptor instead.
func (*LongFiRxPacket) Descriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{3}
}

func (x *LongFiRxPacket) GetCrcCheck() bool {
	if x != nil {
		return x.CrcCheck
	}
	return false
}

func (x *LongFiRxPacket) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LongFiRxPacket) GetRssi() float32 {
	if x != nil {
		return x.Rssi
	}
	return 0
}

func (x *LongFiRxPacket) GetSnr() float32 {
	if x != nil {
		return x.Snr
	}
	return 0
}

func (x *LongFiRxPacket) GetOui() uint32 {
	if x != nil {
		return x.Oui
	}
	return 0
}

func (x *LongFiRxPacket) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *LongFiRxPacket) GetFingerprint() uint32 {
	if x != nil {
		return x.Fingerprint
	}
	return 0
}

func (x *LongFiRxPacket) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *LongFiRxPacket) GetSpreading() LongFiSpreading {
	if x != nil {
		return x.Spreading
	}
	return LongFiSpreading_SF_INVALID
}

func (x *LongFiRxPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *LongFiRxPacket) GetTagBits() uint32 {
	if x != nil {
		return x.TagBits
	}
	return 0
}

type LongFiTxPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// is device receiver (downlink) or is router receiver (uplink)
	// note: when Hotspot is sending Proof of Coverage packet,
	// it should behave as a device and flag this as "uplink"
	Downlink bool `protobuf:"varint,1,opt,name=downlink,proto3" json:"downlink,omitempty"`
	// should the receiver ACK
	ShouldAck bool `protobuf:"varint,2,opt,name=should_ack,json=shouldAck,proto3" json:"should_ack,omitempty"`
	// on uplink, this indicates the device is ready to receive downlink
	Cts bool `protobuf:"varint,3,opt,name=cts,proto3" json:"cts,omitempty"`
	// is the packet urgent
	Priority bool `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// the packet beyond the tag field is encoded with LDPC
	Ldpc bool `protobuf:"varint,5,opt,name=ldpc,proto3" json:"ldpc,omitempty"`
	// Organization Unique ID
	Oui uint32 `protobuf:"varint,6,opt,name=oui,proto3" json:"oui,omitempty"`
	// Device ID
	DeviceId uint32 `protobuf:"varint,7,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Fingerprint
	Fingerpint uint32 `protobuf:"varint,8,opt,name=fingerpint,proto3" json:"fingerpint,omitempty"`
	// Sequence
	Sequence uint32 `protobuf:"varint,9,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Spreading to be used
	Spreading LongFiSpreading `protobuf:"varint,10,opt,name=spreading,proto3,enum=smartmesh.longfi.LongFiSpreading" json:"spreading,omitempty"`
	Payload   []byte          `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *LongFiTxPacket) Reset() {
	*x = LongFiTxPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_longfi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongFiTxPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongFiTxPacket) ProtoMessage() {}

func (x *LongFiTxPacket) ProtoReflect() protoreflect.Message {
	mi := &file_src_longfi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongFiTxPacket.ProtoReflect.Descriptor instead.
func (*LongFiTxPacket) Descriptor() ([]byte, []int) {
	return file_src_longfi_proto_rawDescGZIP(), []int{4}
}

func (x *LongFiTxPacket) GetDownlink() bool {
	if x != nil {
		return x.Downlink
	}
	return false
}

func (x *LongFiTxPacket) GetShouldAck() bool {
	if x != nil {
		return x.ShouldAck
	}
	return false
}

func (x *LongFiTxPacket) GetCts() bool {
	if x != nil {
		return x.Cts
	}
	return false
}

func (x *LongFiTxPacket) GetPriority() bool {
	if x != nil {
		return x.Priority
	}
	return false
}

func (x *LongFiTxPacket) GetLdpc() bool {
	if x != nil {
		return x.Ldpc
	}
	return false
}

func (x *LongFiTxPacket) GetOui() uint32 {
	if x != nil {
		return x.Oui
	}
	return 0
}

func (x *LongFiTxPacket) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *LongFiTxPacket) GetFingerpint() uint32 {
	if x != nil {
		return x.Fingerpint
	}
	return 0
}

func (x *LongFiTxPacket) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *LongFiTxPacket) GetSpreading() LongFiSpreading {
	if x != nil {
		return x.Spreading
	}
	return LongFiSpreading_SF_INVALID
}

func (x *LongFiTxPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_src_longfi_proto protoreflect.FileDescriptor

var file_src_longfi_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x66, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f,
	0x6e, 0x67, 0x66, 0x69, 0x22, 0x57, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x66, 0x69,
	0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x54, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x02, 0x74, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xd7, 0x01,
	0x0a, 0x0a, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x09,
	0x74, 0x78, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x66, 0x69, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x48, 0x00, 0x52, 0x08, 0x74, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x02, 0x72, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x66, 0x69, 0x2e, 0x4c, 0x6f, 0x6e,
	0x67, 0x46, 0x69, 0x52, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x02, 0x72,
	0x78, 0x12, 0x1d, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x4c, 0x6f, 0x6e, 0x67, 0x46,
	0x69, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0xd4, 0x02, 0x0a, 0x0e, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x52, 0x78,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x63, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x72, 0x63, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x73, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x72, 0x73, 0x73, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6e, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x73, 0x6e, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x69, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6f, 0x75, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x66, 0x69, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x46,
	0x69, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x73, 0x70, 0x72, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x74, 0x61, 0x67, 0x42, 0x69, 0x74, 0x73, 0x22, 0xd3, 0x02, 0x0a, 0x0e, 0x4c,
	0x6f, 0x6e, 0x67, 0x46, 0x69, 0x54, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x63, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x64, 0x70, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x64, 0x70, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75,
	0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6f, 0x75, 0x69, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x66, 0x69, 0x2e, 0x4c, 0x6f, 0x6e, 0x67,
	0x46, 0x69, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2a, 0x46, 0x0a, 0x0f, 0x4c, 0x6f, 0x6e, 0x67, 0x46, 0x69, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x46, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x37, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x46, 0x38, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x39, 0x10, 0x03, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x46, 0x31, 0x30, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_longfi_proto_rawDescOnce sync.Once
	file_src_longfi_proto_rawDescData = file_src_longfi_proto_rawDesc
)

func file_src_longfi_proto_rawDescGZIP() []byte {
	file_src_longfi_proto_rawDescOnce.Do(func() {
		file_src_longfi_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_longfi_proto_rawDescData)
	})
	return file_src_longfi_proto_rawDescData
}

var file_src_longfi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_longfi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_src_longfi_proto_goTypes = []interface{}{
	(LongFiSpreading)(0),   // 0: smartmesh.longfi.LongFiSpreading
	(*LongFiReq)(nil),      // 1: smartmesh.longfi.LongFiReq
	(*LongFiResp)(nil),     // 2: smartmesh.longfi.LongFiResp
	(*LongFiTxStatus)(nil), // 3: smartmesh.longfi.LongFiTxStatus
	(*LongFiRxPacket)(nil), // 4: smartmesh.longfi.LongFiRxPacket
	(*LongFiTxPacket)(nil), // 5: smartmesh.longfi.LongFiTxPacket
}
var file_src_longfi_proto_depIdxs = []int32{
	5, // 0: smartmesh.longfi.LongFiReq.tx:type_name -> smartmesh.longfi.LongFiTxPacket
	3, // 1: smartmesh.longfi.LongFiResp.tx_status:type_name -> smartmesh.longfi.LongFiTxStatus
	4, // 2: smartmesh.longfi.LongFiResp.rx:type_name -> smartmesh.longfi.LongFiRxPacket
	0, // 3: smartmesh.longfi.LongFiRxPacket.spreading:type_name -> smartmesh.longfi.LongFiSpreading
	0, // 4: smartmesh.longfi.LongFiTxPacket.spreading:type_name -> smartmesh.longfi.LongFiSpreading
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_src_longfi_proto_init() }
func file_src_longfi_proto_init() {
	if File_src_longfi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_longfi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongFiReq); i {
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
		file_src_longfi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongFiResp); i {
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
		file_src_longfi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongFiTxStatus); i {
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
		file_src_longfi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongFiRxPacket); i {
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
		file_src_longfi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongFiTxPacket); i {
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
	file_src_longfi_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LongFiReq_Tx)(nil),
	}
	file_src_longfi_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LongFiResp_TxStatus)(nil),
		(*LongFiResp_Rx)(nil),
		(*LongFiResp_ParseErr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_longfi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_longfi_proto_goTypes,
		DependencyIndexes: file_src_longfi_proto_depIdxs,
		EnumInfos:         file_src_longfi_proto_enumTypes,
		MessageInfos:      file_src_longfi_proto_msgTypes,
	}.Build()
	File_src_longfi_proto = out.File
	file_src_longfi_proto_rawDesc = nil
	file_src_longfi_proto_goTypes = nil
	file_src_longfi_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/kubemq.proto

package records

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KubeMQ struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Channel              string   `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Value                []byte   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sequence             int64    `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeMQ) Reset()         { *m = KubeMQ{} }
func (m *KubeMQ) String() string { return proto.CompactTextString(m) }
func (*KubeMQ) ProtoMessage()    {}
func (*KubeMQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f60de9d0db475b6, []int{0}
}

func (m *KubeMQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeMQ.Unmarshal(m, b)
}
func (m *KubeMQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeMQ.Marshal(b, m, deterministic)
}
func (m *KubeMQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeMQ.Merge(m, src)
}
func (m *KubeMQ) XXX_Size() int {
	return xxx_messageInfo_KubeMQ.Size(m)
}
func (m *KubeMQ) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeMQ.DiscardUnknown(m)
}

var xxx_messageInfo_KubeMQ proto.InternalMessageInfo

func (m *KubeMQ) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KubeMQ) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *KubeMQ) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *KubeMQ) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KubeMQ) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *KubeMQ) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*KubeMQ)(nil), "protos.records.KubeMQ")
}

func init() { proto.RegisterFile("records/kubemq.proto", fileDescriptor_3f60de9d0db475b6) }

var fileDescriptor_3f60de9d0db475b6 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x46, 0xe5, 0x94, 0x86, 0xc6, 0x42, 0x1d, 0xac, 0x0e, 0x16, 0x30, 0x44, 0x4c, 0x59, 0x88,
	0x07, 0x56, 0xc4, 0xc0, 0x86, 0x10, 0x03, 0x19, 0x59, 0x90, 0x7d, 0x7d, 0xd5, 0x58, 0xd8, 0x71,
	0xea, 0x1f, 0x5e, 0x87, 0x57, 0x45, 0x72, 0x02, 0x9d, 0x3e, 0x9d, 0x73, 0xa6, 0x8f, 0x1e, 0x02,
	0x82, 0x0f, 0x3a, 0x8a, 0xaf, 0xac, 0xd0, 0x9d, 0xfa, 0x39, 0xf8, 0xe4, 0xd9, 0xbe, 0x4c, 0xec,
	0xd7, 0x78, 0xf7, 0x43, 0x68, 0xfd, 0x9a, 0x15, 0xbe, 0xbd, 0xb3, 0x3d, 0xad, 0x8c, 0xe6, 0xa4,
	0x25, 0x5d, 0x33, 0x54, 0x46, 0xb3, 0x1b, 0xda, 0x80, 0x35, 0x38, 0xa5, 0x4f, 0xa3, 0x79, 0x55,
	0xf4, 0x6e, 0x11, 0x2f, 0x9a, 0x71, 0x7a, 0x09, 0xa3, 0x9c, 0x26, 0xb4, 0x7c, 0x53, 0xd2, 0x1f,
	0xb2, 0x03, 0xdd, 0x7e, 0x4b, 0x9b, 0x91, 0x5f, 0xb4, 0xa4, 0xbb, 0x1a, 0x16, 0x60, 0xb7, 0xb4,
	0x49, 0xc6, 0x61, 0x4c, 0xd2, 0xcd, 0x7c, 0xdb, 0x92, 0x6e, 0x33, 0x9c, 0x05, 0xbb, 0xa6, 0xbb,
	0x88, 0xa7, 0x8c, 0x13, 0x20, 0xaf, 0x4b, 0xfc, 0xe7, 0xe7, 0xa7, 0x8f, 0xc7, 0xa3, 0x49, 0x63,
	0x56, 0x3d, 0x78, 0x27, 0x94, 0x4c, 0x30, 0x82, 0x0f, 0xb3, 0x98, 0x6d, 0x76, 0x0a, 0xc3, 0x7d,
	0x84, 0x11, 0x9d, 0x8c, 0x42, 0x65, 0x63, 0xb5, 0x38, 0x7a, 0xb1, 0x3c, 0x14, 0xeb, 0x43, 0x55,
	0x17, 0x7e, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x90, 0x56, 0x7b, 0xd1, 0x10, 0x01, 0x00, 0x00,
}

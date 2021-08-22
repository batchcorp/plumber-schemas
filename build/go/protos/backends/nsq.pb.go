// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nsq.proto

package backends

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

type NSQ struct {
	// Required
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Required for reads
	// Ignored for writes
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NSQ) Reset()         { *m = NSQ{} }
func (m *NSQ) String() string { return proto.CompactTextString(m) }
func (*NSQ) ProtoMessage()    {}
func (*NSQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_f83e7c7ac3c992fd, []int{0}
}

func (m *NSQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSQ.Unmarshal(m, b)
}
func (m *NSQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSQ.Marshal(b, m, deterministic)
}
func (m *NSQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSQ.Merge(m, src)
}
func (m *NSQ) XXX_Size() int {
	return xxx_messageInfo_NSQ.Size(m)
}
func (m *NSQ) XXX_DiscardUnknown() {
	xxx_messageInfo_NSQ.DiscardUnknown(m)
}

var xxx_messageInfo_NSQ proto.InternalMessageInfo

func (m *NSQ) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *NSQ) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*NSQ)(nil), "protos.backends.NSQ")
}

func init() { proto.RegisterFile("nsq.proto", fileDescriptor_f83e7c7ac3c992fd) }

var fileDescriptor_f83e7c7ac3c992fd = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2b, 0x2e, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x07, 0x53, 0xc5, 0x7a, 0x49, 0x89, 0xc9, 0xd9, 0xa9,
	0x79, 0x29, 0xc5, 0x4a, 0xa6, 0x5c, 0xcc, 0x7e, 0xc1, 0x81, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9,
	0x05, 0x99, 0xc9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x04, 0x17, 0x7b,
	0x72, 0x46, 0x62, 0x5e, 0x5e, 0x6a, 0x8e, 0x04, 0x13, 0x58, 0x1c, 0xc6, 0x75, 0xb2, 0x8f, 0xb2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4a, 0x2c, 0x49, 0xce,
	0x48, 0xce, 0x2f, 0x2a, 0xd0, 0x2f, 0xc8, 0x29, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2d, 0x4e, 0xce,
	0x48, 0xcd, 0x4d, 0x2c, 0xd6, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x4f, 0xcf, 0xd7, 0x87, 0xd8,
	0xab, 0x0f, 0xb3, 0x37, 0x89, 0x0d, 0x2c, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x66, 0xb0,
	0xb4, 0x28, 0x9c, 0x00, 0x00, 0x00,
}
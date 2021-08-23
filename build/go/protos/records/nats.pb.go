// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/nats.proto

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

type Nats struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Replay               string   `protobuf:"bytes,2,opt,name=replay,proto3" json:"replay,omitempty"`
	Queue                string   `protobuf:"bytes,4,opt,name=queue,proto3" json:"queue,omitempty"`
	Value                []byte   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nats) Reset()         { *m = Nats{} }
func (m *Nats) String() string { return proto.CompactTextString(m) }
func (*Nats) ProtoMessage()    {}
func (*Nats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba27460c29a2f63a, []int{0}
}

func (m *Nats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nats.Unmarshal(m, b)
}
func (m *Nats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nats.Marshal(b, m, deterministic)
}
func (m *Nats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nats.Merge(m, src)
}
func (m *Nats) XXX_Size() int {
	return xxx_messageInfo_Nats.Size(m)
}
func (m *Nats) XXX_DiscardUnknown() {
	xxx_messageInfo_Nats.DiscardUnknown(m)
}

var xxx_messageInfo_Nats proto.InternalMessageInfo

func (m *Nats) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Nats) GetReplay() string {
	if m != nil {
		return m.Replay
	}
	return ""
}

func (m *Nats) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *Nats) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Nats)(nil), "protos.records.Nats")
}

func init() { proto.RegisterFile("records/nats.proto", fileDescriptor_ba27460c29a2f63a) }

var fileDescriptor_ba27460c29a2f63a = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x05, 0x60, 0xfa, 0xd3, 0xf6, 0xc7, 0x20, 0x0e, 0x41, 0x24, 0x63, 0x71, 0xea, 0x62, 0x33,
	0xb8, 0x8a, 0x83, 0x0f, 0xe0, 0xd0, 0xd1, 0x2d, 0x49, 0x2f, 0x6d, 0x25, 0x35, 0x31, 0x37, 0x57,
	0xf0, 0xed, 0xc5, 0xb4, 0x4e, 0x87, 0xef, 0x9c, 0xe5, 0x30, 0x1e, 0xc0, 0xb8, 0xd0, 0xa1, 0x7c,
	0xa8, 0x88, 0x8d, 0x0f, 0x2e, 0x3a, 0xbe, 0x49, 0x81, 0xcd, 0x32, 0xed, 0x3b, 0x96, 0x5f, 0x55,
	0x44, 0x2e, 0xd8, 0x3f, 0x92, 0xbe, 0x83, 0x89, 0x22, 0xab, 0xb2, 0x7a, 0xd5, 0xfe, 0xc8, 0x77,
	0xac, 0x0c, 0xe0, 0xad, 0x7a, 0x8b, 0xbf, 0x34, 0x2c, 0xe2, 0x5b, 0x56, 0x3c, 0x09, 0x08, 0x44,
	0x9e, 0xea, 0x19, 0xdf, 0xf6, 0xa5, 0x2c, 0x81, 0x28, 0xaa, 0xac, 0x5e, 0xb7, 0x33, 0x2e, 0xe7,
	0xdb, 0xa9, 0x1f, 0xe3, 0x40, 0xba, 0x31, 0x6e, 0x92, 0x5a, 0x45, 0x33, 0x18, 0x17, 0xbc, 0xf4,
	0x96, 0x26, 0x0d, 0xe1, 0x80, 0x66, 0x80, 0x49, 0xa1, 0xd4, 0x34, 0xda, 0x4e, 0xf6, 0x4e, 0xce,
	0x2f, 0xe5, 0xf2, 0x52, 0x97, 0xc9, 0xc7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x50, 0x10,
	0xf6, 0xd2, 0x00, 0x00, 0x00,
}

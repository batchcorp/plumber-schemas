// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_aws_sqs.proto

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

type AWSSQS struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            int64             `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RecipientHandle      string            `protobuf:"bytes,3,opt,name=recipient_handle,json=recipientHandle,proto3" json:"recipient_handle,omitempty"`
	Attributes           map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value                []byte            `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AWSSQS) Reset()         { *m = AWSSQS{} }
func (m *AWSSQS) String() string { return proto.CompactTextString(m) }
func (*AWSSQS) ProtoMessage()    {}
func (*AWSSQS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c74c4ca597ada00, []int{0}
}

func (m *AWSSQS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSQS.Unmarshal(m, b)
}
func (m *AWSSQS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSQS.Marshal(b, m, deterministic)
}
func (m *AWSSQS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSQS.Merge(m, src)
}
func (m *AWSSQS) XXX_Size() int {
	return xxx_messageInfo_AWSSQS.Size(m)
}
func (m *AWSSQS) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSQS.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSQS proto.InternalMessageInfo

func (m *AWSSQS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AWSSQS) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AWSSQS) GetRecipientHandle() string {
	if m != nil {
		return m.RecipientHandle
	}
	return ""
}

func (m *AWSSQS) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *AWSSQS) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*AWSSQS)(nil), "protos.records.AWSSQS")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.AWSSQS.AttributesEntry")
}

func init() { proto.RegisterFile("records/ps_records_aws_sqs.proto", fileDescriptor_5c74c4ca597ada00) }

var fileDescriptor_5c74c4ca597ada00 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x1c, 0x85, 0x49, 0xeb, 0x06, 0x8b, 0xb2, 0x8d, 0xe0, 0xa1, 0x88, 0x87, 0xe2, 0x41, 0xea, 0xc1,
	0x04, 0xf4, 0x22, 0xa2, 0xc2, 0x04, 0xc5, 0xab, 0xed, 0x41, 0xf0, 0x52, 0x92, 0x34, 0xac, 0xc1,
	0xb6, 0x89, 0xc9, 0xaf, 0xca, 0xfe, 0x79, 0x11, 0xd3, 0xb2, 0xa9, 0xa7, 0xbc, 0xf7, 0x92, 0x0f,
	0xbe, 0xe0, 0xd4, 0x29, 0x69, 0x5c, 0xe5, 0x99, 0xf5, 0xe5, 0x18, 0x4b, 0xfe, 0xe9, 0x4b, 0xff,
	0xee, 0xa9, 0x75, 0x06, 0x0c, 0x99, 0x87, 0xc3, 0xd3, 0xf1, 0xf6, 0xe4, 0x0b, 0xe1, 0xe9, 0xea,
	0xa5, 0x28, 0x9e, 0x0b, 0x32, 0xc7, 0x91, 0xae, 0x12, 0x94, 0xa2, 0x6c, 0x96, 0x47, 0xba, 0x22,
	0xc7, 0x78, 0x06, 0xba, 0x55, 0x1e, 0x78, 0x6b, 0x93, 0x28, 0x45, 0x59, 0x9c, 0xef, 0x06, 0x72,
	0x86, 0x97, 0x4e, 0x49, 0x6d, 0xb5, 0xea, 0xa0, 0xac, 0x79, 0x57, 0x35, 0x2a, 0x89, 0x03, 0xbb,
	0xd8, 0xee, 0x4f, 0x61, 0x26, 0x8f, 0x18, 0x73, 0x00, 0xa7, 0x45, 0x0f, 0xca, 0x27, 0x7b, 0x69,
	0x9c, 0xed, 0x5f, 0x9c, 0xd2, 0xbf, 0x22, 0x74, 0x90, 0xa0, 0xab, 0xed, 0xc3, 0x87, 0x0e, 0xdc,
	0x26, 0xff, 0x45, 0x92, 0x43, 0x3c, 0xf9, 0xe0, 0x4d, 0xaf, 0x92, 0x49, 0x8a, 0xb2, 0x83, 0x7c,
	0x28, 0x47, 0xb7, 0x78, 0xf1, 0x0f, 0x22, 0x4b, 0x1c, 0xbf, 0xa9, 0xcd, 0xf8, 0x95, 0x9f, 0xb8,
	0x43, 0xa3, 0xb0, 0x0d, 0xe5, 0x3a, 0xba, 0x42, 0xf7, 0x77, 0xaf, 0x37, 0x6b, 0x0d, 0x75, 0x2f,
	0xa8, 0x34, 0x2d, 0x13, 0x1c, 0x64, 0x2d, 0x8d, 0xb3, 0xcc, 0x36, 0x7d, 0x2b, 0x94, 0x3b, 0xf7,
	0xb2, 0x56, 0x2d, 0xf7, 0x4c, 0xf4, 0xba, 0xa9, 0xd8, 0xda, 0xb0, 0xc1, 0x9b, 0x8d, 0xde, 0x62,
	0x1a, 0xfa, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xbc, 0x5a, 0x64, 0x7b, 0x01, 0x00,
	0x00,
}
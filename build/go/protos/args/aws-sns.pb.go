// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aws-sns.proto

package args

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

type AWSSNSConn struct {
	// @gotags: kong:"en=AWS_DEFAULT_REGION,hidden,required"
	AwsRegion string `protobuf:"bytes,1,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty" kong:"en=AWS_DEFAULT_REGION,hidden,required"`
	// @gotags: kong:"env=AWS_ACCESS_KEY_ID,hidden,required"
	AwsAccessKeyId string `protobuf:"bytes,2,opt,name=aws_access_key_id,json=awsAccessKeyId,proto3" json:"aws_access_key_id,omitempty" kong:"env=AWS_ACCESS_KEY_ID,hidden,required"`
	// @gotags: kong:"env=AWS_SECRET_ACCESS_KEY,hidden,required"
	AwsSecretAccessKey   string   `protobuf:"bytes,3,opt,name=aws_secret_access_key,json=awsSecretAccessKey,proto3" json:"aws_secret_access_key,omitempty" kong:"env=AWS_SECRET_ACCESS_KEY,hidden,required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSNSConn) Reset()         { *m = AWSSNSConn{} }
func (m *AWSSNSConn) String() string { return proto.CompactTextString(m) }
func (*AWSSNSConn) ProtoMessage()    {}
func (*AWSSNSConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac516fe8f86ec289, []int{0}
}

func (m *AWSSNSConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSNSConn.Unmarshal(m, b)
}
func (m *AWSSNSConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSNSConn.Marshal(b, m, deterministic)
}
func (m *AWSSNSConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSNSConn.Merge(m, src)
}
func (m *AWSSNSConn) XXX_Size() int {
	return xxx_messageInfo_AWSSNSConn.Size(m)
}
func (m *AWSSNSConn) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSNSConn.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSNSConn proto.InternalMessageInfo

func (m *AWSSNSConn) GetAwsRegion() string {
	if m != nil {
		return m.AwsRegion
	}
	return ""
}

func (m *AWSSNSConn) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *AWSSNSConn) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

type AWSSNSWriteArgs struct {
	// @gotags: kong:"help='Topic ARN',required"
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic ARN',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSSNSWriteArgs) Reset()         { *m = AWSSNSWriteArgs{} }
func (m *AWSSNSWriteArgs) String() string { return proto.CompactTextString(m) }
func (*AWSSNSWriteArgs) ProtoMessage()    {}
func (*AWSSNSWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac516fe8f86ec289, []int{1}
}

func (m *AWSSNSWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSSNSWriteArgs.Unmarshal(m, b)
}
func (m *AWSSNSWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSSNSWriteArgs.Marshal(b, m, deterministic)
}
func (m *AWSSNSWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSSNSWriteArgs.Merge(m, src)
}
func (m *AWSSNSWriteArgs) XXX_Size() int {
	return xxx_messageInfo_AWSSNSWriteArgs.Size(m)
}
func (m *AWSSNSWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSSNSWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AWSSNSWriteArgs proto.InternalMessageInfo

func (m *AWSSNSWriteArgs) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*AWSSNSConn)(nil), "protos.args.AWSSNSConn")
	proto.RegisterType((*AWSSNSWriteArgs)(nil), "protos.args.AWSSNSWriteArgs")
}

func init() { proto.RegisterFile("aws-sns.proto", fileDescriptor_ac516fe8f86ec289) }

var fileDescriptor_ac516fe8f86ec289 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0xa2, 0xb0, 0x23, 0x2a, 0x06, 0x85, 0x5e, 0x04, 0xd9, 0x8b, 0x7a, 0xd8, 0x06,
	0xf1, 0x24, 0x9e, 0xaa, 0x27, 0x11, 0x3c, 0x6c, 0x0f, 0x0b, 0x5e, 0xca, 0x24, 0x1d, 0xd2, 0xe0,
	0xb6, 0x29, 0x99, 0x94, 0xd0, 0x9f, 0xe0, 0xbf, 0x96, 0xa6, 0x22, 0x9e, 0x86, 0x79, 0xef, 0x7b,
	0xf0, 0x1e, 0x9c, 0x62, 0xe4, 0x0d, 0xf7, 0x5c, 0x0c, 0xde, 0x05, 0x27, 0x4e, 0xd2, 0xe1, 0x02,
	0xbd, 0xe1, 0xf5, 0x77, 0x06, 0x50, 0xee, 0xaa, 0xea, 0xa3, 0x7a, 0x75, 0x7d, 0x2f, 0xae, 0x01,
	0x30, 0x72, 0xed, 0xc9, 0x58, 0xd7, 0xe7, 0xd9, 0x4d, 0x76, 0xb7, 0xda, 0xae, 0x30, 0xf2, 0x36,
	0x09, 0xe2, 0x1e, 0x2e, 0x66, 0x1b, 0xb5, 0x26, 0xe6, 0xfa, 0x8b, 0xa6, 0xda, 0x36, 0xf9, 0x41,
	0xa2, 0xce, 0x30, 0x72, 0x99, 0xf4, 0x77, 0x9a, 0xde, 0x1a, 0xf1, 0x00, 0x57, 0x33, 0xca, 0xa4,
	0x3d, 0x85, 0x7f, 0x89, 0xfc, 0x30, 0xe1, 0x02, 0x23, 0x57, 0xc9, 0xfb, 0x0b, 0xad, 0x6f, 0xe1,
	0x7c, 0xa9, 0xb2, 0xf3, 0x36, 0x50, 0xe9, 0x0d, 0x8b, 0x4b, 0x38, 0x0a, 0x6e, 0xb0, 0xfa, 0xb7,
	0xca, 0xf2, 0xbc, 0x3c, 0x7f, 0x3e, 0x19, 0x1b, 0xda, 0x51, 0x15, 0xda, 0x75, 0x52, 0x61, 0xd0,
	0xad, 0x76, 0x7e, 0x90, 0xc3, 0x7e, 0xec, 0x14, 0xf9, 0x0d, 0xeb, 0x96, 0x3a, 0x64, 0xa9, 0x46,
	0xbb, 0x6f, 0xa4, 0x71, 0x72, 0x59, 0x2c, 0xe7, 0xc5, 0xea, 0x38, 0x3d, 0x8f, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x31, 0x7a, 0x6c, 0xc4, 0x16, 0x01, 0x00, 0x00,
}
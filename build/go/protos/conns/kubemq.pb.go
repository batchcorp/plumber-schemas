// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kubemq.proto

package conns

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
	// Required. Ex: localhost:50000
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Required
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Optional
	// Contents of TLS certificate file
	TlsCertificate []byte `protobuf:"bytes,3,opt,name=tls_certificate,json=tlsCertificate,proto3" json:"tls_certificate,omitempty"`
	// Optional
	// JWT authentication token
	AuthToken            string   `protobuf:"bytes,4,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeMQ) Reset()         { *m = KubeMQ{} }
func (m *KubeMQ) String() string { return proto.CompactTextString(m) }
func (*KubeMQ) ProtoMessage()    {}
func (*KubeMQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_5618575e3d8f6914, []int{0}
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

func (m *KubeMQ) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *KubeMQ) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *KubeMQ) GetTlsCertificate() []byte {
	if m != nil {
		return m.TlsCertificate
	}
	return nil
}

func (m *KubeMQ) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func init() {
	proto.RegisterType((*KubeMQ)(nil), "protos.conns.KubeMQ")
}

func init() { proto.RegisterFile("kubemq.proto", fileDescriptor_5618575e3d8f6914) }

var fileDescriptor_5618575e3d8f6914 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x05, 0x60, 0x56, 0xe5, 0x74, 0xc3, 0xa2, 0x90, 0x2a, 0x20, 0xc2, 0x61, 0xe3, 0x35, 0x6e,
	0x0a, 0x3b, 0xb1, 0xd2, 0x4a, 0xc4, 0xc2, 0xc3, 0xca, 0x66, 0x49, 0x26, 0xe3, 0x6d, 0xb8, 0x6c,
	0xb2, 0x66, 0x26, 0xbf, 0xc1, 0xbf, 0x2d, 0x97, 0x45, 0xae, 0x1a, 0xde, 0x37, 0xf0, 0xe0, 0x89,
	0x6e, 0x5f, 0x2c, 0x4e, 0x3f, 0xfd, 0x9c, 0x13, 0x27, 0xd9, 0xd5, 0x43, 0x3d, 0xa4, 0x18, 0xe9,
	0xf6, 0xb7, 0x11, 0xab, 0xb7, 0x62, 0xf1, 0xfd, 0x43, 0x2a, 0x71, 0x6e, 0x9c, 0xcb, 0x48, 0xa4,
	0x9a, 0x75, 0xb3, 0x69, 0xb7, 0xff, 0x51, 0x5e, 0x8b, 0x16, 0x82, 0xc7, 0xc8, 0x83, 0x77, 0xea,
	0xa4, 0xfe, 0x2e, 0x16, 0x78, 0x75, 0xf2, 0x4e, 0x5c, 0x71, 0xa0, 0x01, 0x30, 0xb3, 0xff, 0xf6,
	0x60, 0x18, 0xd5, 0xe9, 0xba, 0xd9, 0x74, 0xdb, 0x4b, 0x0e, 0xf4, 0x72, 0x54, 0x79, 0x23, 0x84,
	0x29, 0x3c, 0x0e, 0x9c, 0xf6, 0x18, 0xd5, 0x59, 0xad, 0x69, 0x0f, 0xf2, 0x79, 0x80, 0xe7, 0xa7,
	0xaf, 0xc7, 0x9d, 0xe7, 0xb1, 0xd8, 0x1e, 0xd2, 0xa4, 0xad, 0x61, 0x18, 0x21, 0xe5, 0x59, 0xcf,
	0xa1, 0x4c, 0x16, 0xf3, 0x3d, 0xc1, 0x88, 0x93, 0x21, 0x6d, 0x8b, 0x0f, 0x4e, 0xef, 0x92, 0x5e,
	0x76, 0xe8, 0xba, 0xc3, 0xae, 0x6a, 0x7a, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x26, 0x56,
	0x5b, 0xec, 0x00, 0x00, 0x00,
}

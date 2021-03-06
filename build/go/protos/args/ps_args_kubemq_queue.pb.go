// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_kubemq_queue.proto

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

type KubeMQQueueConn struct {
	// @gotags: kong:"help='Dial string for KubeMQ server',env='PLUMBER_RELAY_KUBEMQ_QUEUE_ADDRESS',default='localhost:50000',required"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Dial string for KubeMQ server',env='PLUMBER_RELAY_KUBEMQ_QUEUE_ADDRESS',default='localhost:50000',required"`
	// @gotags: kong:"help='Client JWT authentication token',env='PLUMBER_RELAY_KUBEMQ_QUEUE_AUTH_TOKEN'"
	AuthToken string `protobuf:"bytes,2,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty" kong:"help='Client JWT authentication token',env='PLUMBER_RELAY_KUBEMQ_QUEUE_AUTH_TOKEN'"`
	// @gotags: kong:"help='KubeMQ client cert file',env='PLUMBER_RELAY_KUBEMQ_QUEUE_TLS_CLIENT_CERT'"
	TlsClientCert string `protobuf:"bytes,3,opt,name=tls_client_cert,json=tlsClientCert,proto3" json:"tls_client_cert,omitempty" kong:"help='KubeMQ client cert file',env='PLUMBER_RELAY_KUBEMQ_QUEUE_TLS_CLIENT_CERT'"`
	// @gotags: kong:"help='KubeMQ client ID',env='PLUMBER_RELAY_KUBEMQ_QUEUE_CLIENT_ID',default=plumber"
	ClientId             string   `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" kong:"help='KubeMQ client ID',env='PLUMBER_RELAY_KUBEMQ_QUEUE_CLIENT_ID',default=plumber"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeMQQueueConn) Reset()         { *m = KubeMQQueueConn{} }
func (m *KubeMQQueueConn) String() string { return proto.CompactTextString(m) }
func (*KubeMQQueueConn) ProtoMessage()    {}
func (*KubeMQQueueConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6c5d36c3dc96c32, []int{0}
}

func (m *KubeMQQueueConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeMQQueueConn.Unmarshal(m, b)
}
func (m *KubeMQQueueConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeMQQueueConn.Marshal(b, m, deterministic)
}
func (m *KubeMQQueueConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeMQQueueConn.Merge(m, src)
}
func (m *KubeMQQueueConn) XXX_Size() int {
	return xxx_messageInfo_KubeMQQueueConn.Size(m)
}
func (m *KubeMQQueueConn) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeMQQueueConn.DiscardUnknown(m)
}

var xxx_messageInfo_KubeMQQueueConn proto.InternalMessageInfo

func (m *KubeMQQueueConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *KubeMQQueueConn) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *KubeMQQueueConn) GetTlsClientCert() string {
	if m != nil {
		return m.TlsClientCert
	}
	return ""
}

func (m *KubeMQQueueConn) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type KubeMQQueueReadArgs struct {
	// @gotags: kong:"help='KubeMQ queue name',env='PLUMBER_RELAY_KUBEMQ_QUEUE_QUEUE'"
	QueueName            string   `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='KubeMQ queue name',env='PLUMBER_RELAY_KUBEMQ_QUEUE_QUEUE'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeMQQueueReadArgs) Reset()         { *m = KubeMQQueueReadArgs{} }
func (m *KubeMQQueueReadArgs) String() string { return proto.CompactTextString(m) }
func (*KubeMQQueueReadArgs) ProtoMessage()    {}
func (*KubeMQQueueReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6c5d36c3dc96c32, []int{1}
}

func (m *KubeMQQueueReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeMQQueueReadArgs.Unmarshal(m, b)
}
func (m *KubeMQQueueReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeMQQueueReadArgs.Marshal(b, m, deterministic)
}
func (m *KubeMQQueueReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeMQQueueReadArgs.Merge(m, src)
}
func (m *KubeMQQueueReadArgs) XXX_Size() int {
	return xxx_messageInfo_KubeMQQueueReadArgs.Size(m)
}
func (m *KubeMQQueueReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeMQQueueReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KubeMQQueueReadArgs proto.InternalMessageInfo

func (m *KubeMQQueueReadArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type KubeMQQueueWriteArgs struct {
	// @gotags: kong:"help='KubeMQ queue name'"
	QueueName            string   `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty" kong:"help='KubeMQ queue name'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeMQQueueWriteArgs) Reset()         { *m = KubeMQQueueWriteArgs{} }
func (m *KubeMQQueueWriteArgs) String() string { return proto.CompactTextString(m) }
func (*KubeMQQueueWriteArgs) ProtoMessage()    {}
func (*KubeMQQueueWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6c5d36c3dc96c32, []int{2}
}

func (m *KubeMQQueueWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeMQQueueWriteArgs.Unmarshal(m, b)
}
func (m *KubeMQQueueWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeMQQueueWriteArgs.Marshal(b, m, deterministic)
}
func (m *KubeMQQueueWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeMQQueueWriteArgs.Merge(m, src)
}
func (m *KubeMQQueueWriteArgs) XXX_Size() int {
	return xxx_messageInfo_KubeMQQueueWriteArgs.Size(m)
}
func (m *KubeMQQueueWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeMQQueueWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KubeMQQueueWriteArgs proto.InternalMessageInfo

func (m *KubeMQQueueWriteArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func init() {
	proto.RegisterType((*KubeMQQueueConn)(nil), "protos.args.KubeMQQueueConn")
	proto.RegisterType((*KubeMQQueueReadArgs)(nil), "protos.args.KubeMQQueueReadArgs")
	proto.RegisterType((*KubeMQQueueWriteArgs)(nil), "protos.args.KubeMQQueueWriteArgs")
}

func init() { proto.RegisterFile("ps_args_kubemq_queue.proto", fileDescriptor_c6c5d36c3dc96c32) }

var fileDescriptor_c6c5d36c3dc96c32 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x8a, 0xda, 0x88, 0x2c, 0x54, 0x0f, 0x45, 0x11, 0xa4, 0x07, 0xf1, 0x62, 0x73,
	0x50, 0x0f, 0xe2, 0x49, 0x7b, 0x12, 0x51, 0xd8, 0x45, 0x10, 0xbc, 0x84, 0x7c, 0x0c, 0x6d, 0xd9,
	0x7c, 0x74, 0x93, 0xc9, 0xef, 0xf0, 0x2f, 0x4b, 0xb2, 0x2b, 0xec, 0x71, 0x4f, 0x61, 0x9e, 0x79,
	0x86, 0xcc, 0xbc, 0xe4, 0x62, 0x0a, 0x8c, 0xfb, 0x3e, 0xb0, 0x65, 0x14, 0x60, 0x56, 0x6c, 0x15,
	0x21, 0x42, 0x3b, 0x79, 0x87, 0xae, 0x3a, 0xc9, 0x4f, 0x68, 0x53, 0xbf, 0xf9, 0x2d, 0xc8, 0xec,
	0x3d, 0x0a, 0xf8, 0x98, 0xcf, 0x93, 0xd2, 0x39, 0x6b, 0xab, 0x9a, 0x1c, 0x71, 0xa5, 0x3c, 0x84,
	0x50, 0x17, 0xd7, 0xc5, 0x6d, 0xb9, 0xf8, 0x2f, 0xab, 0x2b, 0x42, 0x78, 0xc4, 0x81, 0xa1, 0x5b,
	0x82, 0xad, 0xf7, 0x72, 0xb3, 0x4c, 0xe4, 0x2b, 0x81, 0xea, 0x86, 0xcc, 0x50, 0x07, 0x26, 0xf5,
	0x08, 0x16, 0x99, 0x04, 0x8f, 0xf5, 0x7e, 0x76, 0x4e, 0x51, 0x87, 0x2e, 0xd3, 0x0e, 0x3c, 0x56,
	0x97, 0xa4, 0xdc, 0x38, 0xa3, 0xaa, 0x0f, 0xb2, 0x71, 0xbc, 0x06, 0x6f, 0xaa, 0x79, 0x20, 0x67,
	0x5b, 0x0b, 0x2d, 0x80, 0xab, 0x17, 0xdf, 0xe7, 0xaf, 0xf3, 0x11, 0xcc, 0x72, 0x03, 0x9b, 0xbd,
	0xca, 0x4c, 0x3e, 0xb9, 0x81, 0xe6, 0x91, 0x9c, 0x6f, 0x4d, 0x7d, 0xfb, 0x11, 0x61, 0x87, 0xb1,
	0xd7, 0xe7, 0x9f, 0xa7, 0x7e, 0xc4, 0x21, 0x8a, 0x56, 0x3a, 0x43, 0x05, 0x47, 0x39, 0x48, 0xe7,
	0x27, 0x3a, 0xe9, 0x68, 0x04, 0xf8, 0xbb, 0x20, 0x07, 0x30, 0x3c, 0x50, 0x11, 0x47, 0xad, 0x68,
	0xef, 0xe8, 0x3a, 0x3b, 0x9a, 0xb2, 0x13, 0x87, 0xb9, 0xb8, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x82, 0x6f, 0x1e, 0x6d, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongo.proto

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

type MongoConn struct {
	// @gotags: kong:"help='Dial string for mongo server (Ex: mongodb://localhost:27017)',env=,default=PLUMBER_RELAY_CDCMONGO_DSN'mongodb://localhost:27017'"
	Dsn                  string   `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty" kong:"help='Dial string for mongo server (Ex: mongodb://localhost:27017)',env=,default=PLUMBER_RELAY_CDCMONGO_DSN'mongodb://localhost:27017'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoConn) Reset()         { *m = MongoConn{} }
func (m *MongoConn) String() string { return proto.CompactTextString(m) }
func (*MongoConn) ProtoMessage()    {}
func (*MongoConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c3854ae3094fec9, []int{0}
}

func (m *MongoConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoConn.Unmarshal(m, b)
}
func (m *MongoConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoConn.Marshal(b, m, deterministic)
}
func (m *MongoConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoConn.Merge(m, src)
}
func (m *MongoConn) XXX_Size() int {
	return xxx_messageInfo_MongoConn.Size(m)
}
func (m *MongoConn) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoConn.DiscardUnknown(m)
}

var xxx_messageInfo_MongoConn proto.InternalMessageInfo

func (m *MongoConn) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

type MongoReadArgs struct {
	// @gotags: kong:"help='Database name',env='PLUMBER_RELAY_CDCMONGO_DATABASE'"
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty" kong:"help='Database name',env='PLUMBER_RELAY_CDCMONGO_DATABASE'"`
	// @gotags: kong:"help='Collection name',env='PLUMBER_RELAY_CDCMONGO_COLLECTION'"
	Collection string `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty" kong:"help='Collection name',env='PLUMBER_RELAY_CDCMONGO_COLLECTION'"`
	// @gotags: kong:"help='Include full document in update in update changes (default - return deltas only)',env='PLUMBER_RELAY_CDCMONGO_INCLUDE_FULL_DOC'"
	IncludeFullDocument  bool     `protobuf:"varint,3,opt,name=include_full_document,json=includeFullDocument,proto3" json:"include_full_document,omitempty" kong:"help='Include full document in update in update changes (default - return deltas only)',env='PLUMBER_RELAY_CDCMONGO_INCLUDE_FULL_DOC'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoReadArgs) Reset()         { *m = MongoReadArgs{} }
func (m *MongoReadArgs) String() string { return proto.CompactTextString(m) }
func (*MongoReadArgs) ProtoMessage()    {}
func (*MongoReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c3854ae3094fec9, []int{1}
}

func (m *MongoReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoReadArgs.Unmarshal(m, b)
}
func (m *MongoReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoReadArgs.Marshal(b, m, deterministic)
}
func (m *MongoReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoReadArgs.Merge(m, src)
}
func (m *MongoReadArgs) XXX_Size() int {
	return xxx_messageInfo_MongoReadArgs.Size(m)
}
func (m *MongoReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_MongoReadArgs proto.InternalMessageInfo

func (m *MongoReadArgs) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *MongoReadArgs) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *MongoReadArgs) GetIncludeFullDocument() bool {
	if m != nil {
		return m.IncludeFullDocument
	}
	return false
}

func init() {
	proto.RegisterType((*MongoConn)(nil), "protos.backends.MongoConn")
	proto.RegisterType((*MongoReadArgs)(nil), "protos.backends.MongoReadArgs")
}

func init() { proto.RegisterFile("mongo.proto", fileDescriptor_5c3854ae3094fec9) }

var fileDescriptor_5c3854ae3094fec9 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0x05, 0x70, 0xd6, 0x03, 0xb9, 0x8b, 0x88, 0x12, 0x11, 0x16, 0x41, 0x39, 0xae, 0xba, 0xc6,
	0x0d, 0x68, 0x2d, 0xe2, 0x1f, 0xec, 0x6c, 0xb6, 0xb4, 0x39, 0x92, 0xc9, 0xb8, 0xbb, 0x38, 0x99,
	0x59, 0x36, 0x49, 0xed, 0x57, 0x97, 0x8b, 0xab, 0x5c, 0x35, 0x33, 0xef, 0x37, 0xcd, 0x53, 0x27,
	0x41, 0xb8, 0x93, 0x66, 0x9c, 0x24, 0x89, 0x3e, 0x2b, 0x23, 0x36, 0xce, 0xc2, 0x17, 0xb2, 0x8f,
	0x9b, 0x6b, 0xb5, 0x7a, 0xdf, 0xfb, 0x8b, 0x30, 0xeb, 0x73, 0xb5, 0xf0, 0x91, 0xeb, 0x6a, 0x5d,
	0x6d, 0x57, 0xed, 0x7e, 0xdd, 0x7c, 0xab, 0xd3, 0xc2, 0x2d, 0x5a, 0xff, 0x34, 0x75, 0x51, 0x5f,
	0xa9, 0xa5, 0xb7, 0xc9, 0x3a, 0x1b, 0x71, 0xfe, 0xfb, 0xbf, 0xf5, 0x8d, 0x52, 0x20, 0x44, 0x08,
	0x69, 0x10, 0xae, 0x8f, 0x8a, 0x1e, 0x24, 0xfa, 0x4e, 0x5d, 0x0e, 0x0c, 0x94, 0x3d, 0xee, 0x3e,
	0x33, 0xd1, 0xce, 0x0b, 0xe4, 0x80, 0x9c, 0xea, 0xc5, 0xba, 0xda, 0x2e, 0xdb, 0x8b, 0x19, 0xdf,
	0x32, 0xd1, 0xeb, 0x4c, 0xcf, 0x8f, 0x1f, 0x0f, 0xdd, 0x90, 0xfa, 0xec, 0x1a, 0x90, 0x60, 0x9c,
	0x4d, 0xd0, 0x83, 0x4c, 0xa3, 0x19, 0x29, 0x07, 0x87, 0xd3, 0x6d, 0x84, 0x1e, 0x83, 0x8d, 0xc6,
	0xe5, 0x81, 0xbc, 0xe9, 0xc4, 0xfc, 0x16, 0x34, 0x7f, 0x05, 0xdd, 0x71, 0x09, 0xee, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x01, 0x36, 0xfa, 0x36, 0x07, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package encoding

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

type EncodeType int32

const (
	EncodeType_ENCODE_TYPE_UNSET  EncodeType = 0
	EncodeType_ENCODE_TYPE_JSONPB EncodeType = 1
	EncodeType_ENCODE_TYPE_AVRO   EncodeType = 2
)

var EncodeType_name = map[int32]string{
	0: "ENCODE_TYPE_UNSET",
	1: "ENCODE_TYPE_JSONPB",
	2: "ENCODE_TYPE_AVRO",
}

var EncodeType_value = map[string]int32{
	"ENCODE_TYPE_UNSET":  0,
	"ENCODE_TYPE_JSONPB": 1,
	"ENCODE_TYPE_AVRO":   2,
}

func (x EncodeType) String() string {
	return proto.EnumName(EncodeType_name, int32(x))
}

func (EncodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

type DecodeType int32

const (
	DecodeType_DECODE_TYPE_UNSET      DecodeType = 0
	DecodeType_DECODE_TYPE_JSONPB     DecodeType = 1
	DecodeType_DECODE_TYPE_PROTOBUF   DecodeType = 2
	DecodeType_DECODE_TYPE_AVRO       DecodeType = 3
	DecodeType_DECODE_TYPE_THRIFT     DecodeType = 4
	DecodeType_DECODE_TYPE_FLATBUFFER DecodeType = 5
)

var DecodeType_name = map[int32]string{
	0: "DECODE_TYPE_UNSET",
	1: "DECODE_TYPE_JSONPB",
	2: "DECODE_TYPE_PROTOBUF",
	3: "DECODE_TYPE_AVRO",
	4: "DECODE_TYPE_THRIFT",
	5: "DECODE_TYPE_FLATBUFFER",
}

var DecodeType_value = map[string]int32{
	"DECODE_TYPE_UNSET":      0,
	"DECODE_TYPE_JSONPB":     1,
	"DECODE_TYPE_PROTOBUF":   2,
	"DECODE_TYPE_AVRO":       3,
	"DECODE_TYPE_THRIFT":     4,
	"DECODE_TYPE_FLATBUFFER": 5,
}

func (x DecodeType) String() string {
	return proto.EnumName(DecodeType_name, int32(x))
}

func (DecodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

type ProtobufSettings struct {
	// @gotags: kong:"help='Input message(s) should be encoded with this message envelope'"
	ProtobufRootMessage string `protobuf:"bytes,3,opt,name=protobuf_root_message,json=protobufRootMessage,proto3" json:"protobuf_root_message,omitempty" kong:"help='Input message(s) should be encoded with this message envelope'"`
	// @gotags: kong:"help='One or more directories which contains protobuf schemas',existingdir"
	ProtobufDirs         []string `protobuf:"bytes,4,rep,name=protobuf_dirs,json=protobufDirs,proto3" json:"protobuf_dirs,omitempty" kong:"help='One or more directories which contains protobuf schemas',existingdir"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtobufSettings) Reset()         { *m = ProtobufSettings{} }
func (m *ProtobufSettings) String() string { return proto.CompactTextString(m) }
func (*ProtobufSettings) ProtoMessage()    {}
func (*ProtobufSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *ProtobufSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtobufSettings.Unmarshal(m, b)
}
func (m *ProtobufSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtobufSettings.Marshal(b, m, deterministic)
}
func (m *ProtobufSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtobufSettings.Merge(m, src)
}
func (m *ProtobufSettings) XXX_Size() int {
	return xxx_messageInfo_ProtobufSettings.Size(m)
}
func (m *ProtobufSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtobufSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ProtobufSettings proto.InternalMessageInfo

func (m *ProtobufSettings) GetProtobufRootMessage() string {
	if m != nil {
		return m.ProtobufRootMessage
	}
	return ""
}

func (m *ProtobufSettings) GetProtobufDirs() []string {
	if m != nil {
		return m.ProtobufDirs
	}
	return nil
}

type EncodeOptions struct {
	// Use an existing schema for encoding (and ignore all other encode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// TODO: Update kong to be able to use proto/smart enums
	// @gotags: kong:"help='Input type (0: Unset, 1: JSONPB)',default=0"
	EncodeType EncodeType `protobuf:"varint,2,opt,name=encode_type,json=encodeType,proto3,enum=protos.encoding.EncodeType" json:"encode_type,omitempty" kong:"help='Input type (0: Unset, 1: JSONPB)',default=0"`
	// @gotags: kong="embed,group=protobuf"
	ProtobufSettings *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty"`
	// @gotags: kong="help='If encode-type is set to avro, must specify avro schema file',existingfile"
	AvroSchemaFile       string   `protobuf:"bytes,4,opt,name=avro_schema_file,json=avroSchemaFile,proto3" json:"avro_schema_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncodeOptions) Reset()         { *m = EncodeOptions{} }
func (m *EncodeOptions) String() string { return proto.CompactTextString(m) }
func (*EncodeOptions) ProtoMessage()    {}
func (*EncodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *EncodeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodeOptions.Unmarshal(m, b)
}
func (m *EncodeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodeOptions.Marshal(b, m, deterministic)
}
func (m *EncodeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodeOptions.Merge(m, src)
}
func (m *EncodeOptions) XXX_Size() int {
	return xxx_messageInfo_EncodeOptions.Size(m)
}
func (m *EncodeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_EncodeOptions proto.InternalMessageInfo

func (m *EncodeOptions) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *EncodeOptions) GetEncodeType() EncodeType {
	if m != nil {
		return m.EncodeType
	}
	return EncodeType_ENCODE_TYPE_UNSET
}

func (m *EncodeOptions) GetProtobufSettings() *ProtobufSettings {
	if m != nil {
		return m.ProtobufSettings
	}
	return nil
}

func (m *EncodeOptions) GetAvroSchemaFile() string {
	if m != nil {
		return m.AvroSchemaFile
	}
	return ""
}

type DecodeOptions struct {
	// Use an existing schema for decoding (and ignore all other decode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// @gotags: kong:"help='Input type (0: Unset, 1: JSONPB, 2: Protobuf, 3: Avro, 4: Thrift, 5: Flatbuffer)',default=0"
	DecodeType DecodeType `protobuf:"varint,2,opt,name=decode_type,json=decodeType,proto3,enum=protos.encoding.DecodeType" json:"decode_type,omitempty" kong:"help='Input type (0: Unset, 1: JSONPB, 2: Protobuf, 3: Avro, 4: Thrift, 5: Flatbuffer)',default=0"`
	// @gotags: kong="embed,group=protobuf"
	ProtobufSettings     *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DecodeOptions) Reset()         { *m = DecodeOptions{} }
func (m *DecodeOptions) String() string { return proto.CompactTextString(m) }
func (*DecodeOptions) ProtoMessage()    {}
func (*DecodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *DecodeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeOptions.Unmarshal(m, b)
}
func (m *DecodeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeOptions.Marshal(b, m, deterministic)
}
func (m *DecodeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeOptions.Merge(m, src)
}
func (m *DecodeOptions) XXX_Size() int {
	return xxx_messageInfo_DecodeOptions.Size(m)
}
func (m *DecodeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeOptions proto.InternalMessageInfo

func (m *DecodeOptions) GetSchemaId() string {
	if m != nil {
		return m.SchemaId
	}
	return ""
}

func (m *DecodeOptions) GetDecodeType() DecodeType {
	if m != nil {
		return m.DecodeType
	}
	return DecodeType_DECODE_TYPE_UNSET
}

func (m *DecodeOptions) GetProtobufSettings() *ProtobufSettings {
	if m != nil {
		return m.ProtobufSettings
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.encoding.EncodeType", EncodeType_name, EncodeType_value)
	proto.RegisterEnum("protos.encoding.DecodeType", DecodeType_name, DecodeType_value)
	proto.RegisterType((*ProtobufSettings)(nil), "protos.encoding.ProtobufSettings")
	proto.RegisterType((*EncodeOptions)(nil), "protos.encoding.EncodeOptions")
	proto.RegisterType((*DecodeOptions)(nil), "protos.encoding.DecodeOptions")
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0x5b, 0xc5, 0x9e, 0xb5, 0xeb, 0xec, 0xb8, 0xbb, 0x04, 0xf7, 0xa6, 0xd6, 0x9b,
	0xb2, 0x60, 0x02, 0xf5, 0x56, 0x91, 0xad, 0x4d, 0x70, 0x45, 0x9b, 0x3a, 0x4d, 0x05, 0xbd, 0x09,
	0x4d, 0x72, 0x36, 0x1d, 0x6c, 0x3b, 0x21, 0x33, 0x15, 0xf6, 0x69, 0x7c, 0x0f, 0x9f, 0xc7, 0x07,
	0x91, 0xcc, 0xf4, 0x4f, 0x88, 0x82, 0x78, 0xb1, 0x57, 0x33, 0xfc, 0x0e, 0xf3, 0x9d, 0xef, 0xfb,
	0x06, 0x3a, 0x22, 0x57, 0x5c, 0xac, 0xa5, 0x93, 0x17, 0x42, 0x09, 0xfa, 0x58, 0x1f, 0xd2, 0xc1,
	0x75, 0x22, 0x52, 0xbe, 0xce, 0x7a, 0xdf, 0x80, 0x4c, 0x4a, 0x14, 0x6f, 0x6e, 0xa6, 0xa8, 0x14,
	0x5f, 0x67, 0x92, 0x0e, 0xe0, 0x2c, 0xdf, 0xb2, 0xa8, 0x10, 0x42, 0x45, 0x2b, 0x94, 0x72, 0x9e,
	0xa1, 0xdd, 0xec, 0x5a, 0xfd, 0x36, 0x7b, 0xb2, 0x1b, 0x32, 0x21, 0xd4, 0x47, 0x33, 0xa2, 0xcf,
	0xa1, 0xb3, 0x7f, 0x93, 0xf2, 0x42, 0xda, 0xad, 0x6e, 0xb3, 0xdf, 0x66, 0x8f, 0x76, 0x70, 0xc4,
	0x0b, 0xd9, 0xfb, 0x65, 0x41, 0xc7, 0x2b, 0x37, 0x63, 0x60, 0x5c, 0xd1, 0x0b, 0x68, 0xcb, 0x64,
	0x81, 0xab, 0x79, 0xc4, 0x53, 0xdb, 0xd2, 0xf2, 0x0f, 0x0d, 0xb8, 0x4e, 0xe9, 0x2b, 0x38, 0xd2,
	0x3e, 0x31, 0x52, 0xb7, 0x39, 0xda, 0x8d, 0xae, 0xd5, 0x3f, 0x1e, 0x5c, 0x38, 0xb5, 0x08, 0x8e,
	0x51, 0x0c, 0x6f, 0x73, 0x64, 0x80, 0xfb, 0x3b, 0x1d, 0xc3, 0xc9, 0xde, 0x91, 0xdc, 0x46, 0xd3,
	0x09, 0x8e, 0x06, 0xcf, 0xfe, 0xd0, 0xa8, 0x77, 0xc0, 0x48, 0x5e, 0x6f, 0xa5, 0x0f, 0x64, 0xfe,
	0xbd, 0x10, 0xd1, 0xd6, 0xef, 0x0d, 0x5f, 0xa2, 0xdd, 0xd2, 0x8e, 0x8f, 0x4b, 0x3e, 0xd5, 0xd8,
	0xe7, 0x4b, 0xec, 0xfd, 0xb4, 0xa0, 0x33, 0xc2, 0xff, 0x89, 0x99, 0xe2, 0xbf, 0x63, 0x1a, 0x45,
	0x13, 0x33, 0xc5, 0xbb, 0x8a, 0x79, 0xf9, 0x09, 0xe0, 0x50, 0x28, 0x3d, 0x83, 0x13, 0x6f, 0xfc,
	0x36, 0x18, 0x79, 0x51, 0xf8, 0x65, 0xe2, 0x45, 0xb3, 0xf1, 0xd4, 0x0b, 0xc9, 0x3d, 0x7a, 0x0e,
	0xb4, 0x8a, 0xdf, 0x4f, 0x83, 0xf1, 0x64, 0x48, 0x2c, 0x7a, 0x0a, 0xa4, 0xca, 0xaf, 0x3e, 0xb3,
	0x80, 0x34, 0x2e, 0x7f, 0x58, 0x00, 0x07, 0xf7, 0xa5, 0xe6, 0xc8, 0xfb, 0xab, 0x66, 0x15, 0xef,
	0x35, 0x6d, 0x38, 0xad, 0xf2, 0x09, 0x0b, 0xc2, 0x60, 0x38, 0xf3, 0x49, 0xa3, 0xdc, 0x56, 0x9d,
	0xe8, 0x6d, 0xcd, 0xba, 0x4e, 0xf8, 0x8e, 0x5d, 0xfb, 0x21, 0x69, 0xd1, 0xa7, 0x70, 0x5e, 0xe5,
	0xfe, 0x87, 0xab, 0x70, 0x38, 0xf3, 0x7d, 0x8f, 0x91, 0xfb, 0xc3, 0x37, 0x5f, 0x5f, 0x67, 0x5c,
	0x2d, 0x36, 0xb1, 0x93, 0x88, 0x95, 0x1b, 0xcf, 0x55, 0xb2, 0x48, 0x44, 0x91, 0xbb, 0xf9, 0x72,
	0xb3, 0x8a, 0xb1, 0x78, 0x61, 0xfe, 0x4a, 0xba, 0xf1, 0x86, 0x2f, 0x53, 0x37, 0x13, 0xae, 0x29,
	0xd6, 0xdd, 0x15, 0x1b, 0x3f, 0xd0, 0xe0, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x5e,
	0x15, 0x5c, 0x6f, 0x03, 0x00, 0x00,
}

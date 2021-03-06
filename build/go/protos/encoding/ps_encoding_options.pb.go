// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_encoding_options.proto

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
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{0}
}

type DecodeType int32

const (
	DecodeType_DECODE_TYPE_UNSET      DecodeType = 0
	DecodeType_DECODE_TYPE_PROTOBUF   DecodeType = 1
	DecodeType_DECODE_TYPE_AVRO       DecodeType = 2
	DecodeType_DECODE_TYPE_THRIFT     DecodeType = 3
	DecodeType_DECODE_TYPE_FLATBUFFER DecodeType = 4
)

var DecodeType_name = map[int32]string{
	0: "DECODE_TYPE_UNSET",
	1: "DECODE_TYPE_PROTOBUF",
	2: "DECODE_TYPE_AVRO",
	3: "DECODE_TYPE_THRIFT",
	4: "DECODE_TYPE_FLATBUFFER",
}

var DecodeType_value = map[string]int32{
	"DECODE_TYPE_UNSET":      0,
	"DECODE_TYPE_PROTOBUF":   1,
	"DECODE_TYPE_AVRO":       2,
	"DECODE_TYPE_THRIFT":     3,
	"DECODE_TYPE_FLATBUFFER": 4,
}

func (x DecodeType) String() string {
	return proto.EnumName(DecodeType_name, int32(x))
}

func (DecodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{1}
}

type EnvelopeType int32

const (
	EnvelopeType_ENVELOPE_TYPE_UNSET   EnvelopeType = 0
	EnvelopeType_ENVELOPE_TYPE_DEEP    EnvelopeType = 1
	EnvelopeType_ENVELOPE_TYPE_SHALLOW EnvelopeType = 2
)

var EnvelopeType_name = map[int32]string{
	0: "ENVELOPE_TYPE_UNSET",
	1: "ENVELOPE_TYPE_DEEP",
	2: "ENVELOPE_TYPE_SHALLOW",
}

var EnvelopeType_value = map[string]int32{
	"ENVELOPE_TYPE_UNSET":   0,
	"ENVELOPE_TYPE_DEEP":    1,
	"ENVELOPE_TYPE_SHALLOW": 2,
}

func (x EnvelopeType) String() string {
	return proto.EnumName(EnvelopeType_name, int32(x))
}

func (EnvelopeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{2}
}

type ProtobufSettings struct {
	// @gotags: kong:"help='Input message(s) should be encoded with this message envelope'"
	ProtobufRootMessage string `protobuf:"bytes,1,opt,name=protobuf_root_message,json=protobufRootMessage,proto3" json:"protobuf_root_message,omitempty" kong:"help='Input message(s) should be encoded with this message envelope'"`
	// Desktop/server should not use this.
	// @gotags: kong:"help='One or more directories which contains protobuf schemas',existingdir"
	ProtobufDirs []string `protobuf:"bytes,2,rep,name=protobuf_dirs,json=protobufDirs,proto3" json:"protobuf_dirs,omitempty" kong:"help='One or more directories which contains protobuf schemas',existingdir"`
	// Directory where protos are stored (used for github import)
	// @gotags: kong:"-"
	XProtobufRootDir string `protobuf:"bytes,3,opt,name=_protobuf_root_dir,json=ProtobufRootDir,proto3" json:"_protobuf_root_dir,omitempty" kong:"-"`
	// Used by server/desktop when creating a read without an existing schema
	// @gotags: kong:"-"
	Archive []byte `protobuf:"bytes,4,opt,name=archive,proto3" json:"archive,omitempty" kong:"-"`
	// Used internally by the server
	// @gotags: kong:"-"
	XMessageDescriptor []byte `protobuf:"bytes,5,opt,name=_message_descriptor,json=MessageDescriptor,proto3" json:"_message_descriptor,omitempty" kong:"-"`
	// @gotags: kong:"help='Envelope type (options: deep, shallow)',type=pbenum,pbenum_strip_prefix=ENVELOPE_TYPE_,pbenum_lowercase,default=deep"
	ProtobufEnvelopeType EnvelopeType `protobuf:"varint,6,opt,name=protobuf_envelope_type,json=protobufEnvelopeType,proto3,enum=protos.encoding.EnvelopeType" json:"protobuf_envelope_type,omitempty" kong:"help='Envelope type (options: deep, shallow)',type=pbenum,pbenum_strip_prefix=ENVELOPE_TYPE_,pbenum_lowercase,default=deep"`
	// @gotags: kong:"help='For shallow envelope messages, the payload field should be encoded with this message name'"
	ShallowEnvelopeMessage string `protobuf:"bytes,7,opt,name=shallow_envelope_message,json=shallowEnvelopeMessage,proto3" json:"shallow_envelope_message,omitempty" kong:"help='For shallow envelope messages, the payload field should be encoded with this message name'"`
	// @gotags: kong:"help='For shallow envelope messages, the field number of the root message that contains the shallow envelope payload'"
	ShallowEnvelopeFieldNumber int32 `protobuf:"varint,8,opt,name=shallow_envelope_field_number,json=shallowEnvelopeFieldNumber,proto3" json:"shallow_envelope_field_number,omitempty" kong:"help='For shallow envelope messages, the field number of the root message that contains the shallow envelope payload'"`
	// Used internally by the server
	// @gotags: kong:"-"
	XShallowEnvelopeMessageDescriptor []byte `protobuf:"bytes,9,opt,name=_shallow_envelope_message_descriptor,json=ShallowEnvelopeMessageDescriptor,proto3" json:"_shallow_envelope_message_descriptor,omitempty" kong:"-"`
	// @gotags: kong:"help='Protobuf descriptor set(.protoset or .fds file)'"
	ProtobufDescriptorSet string   `protobuf:"bytes,10,opt,name=protobuf_descriptor_set,json=protobufDescriptorSet,proto3" json:"protobuf_descriptor_set,omitempty" kong:"help='Protobuf descriptor set(.protoset or .fds file)'"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ProtobufSettings) Reset()         { *m = ProtobufSettings{} }
func (m *ProtobufSettings) String() string { return proto.CompactTextString(m) }
func (*ProtobufSettings) ProtoMessage()    {}
func (*ProtobufSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{0}
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

func (m *ProtobufSettings) GetXProtobufRootDir() string {
	if m != nil {
		return m.XProtobufRootDir
	}
	return ""
}

func (m *ProtobufSettings) GetArchive() []byte {
	if m != nil {
		return m.Archive
	}
	return nil
}

func (m *ProtobufSettings) GetXMessageDescriptor() []byte {
	if m != nil {
		return m.XMessageDescriptor
	}
	return nil
}

func (m *ProtobufSettings) GetProtobufEnvelopeType() EnvelopeType {
	if m != nil {
		return m.ProtobufEnvelopeType
	}
	return EnvelopeType_ENVELOPE_TYPE_UNSET
}

func (m *ProtobufSettings) GetShallowEnvelopeMessage() string {
	if m != nil {
		return m.ShallowEnvelopeMessage
	}
	return ""
}

func (m *ProtobufSettings) GetShallowEnvelopeFieldNumber() int32 {
	if m != nil {
		return m.ShallowEnvelopeFieldNumber
	}
	return 0
}

func (m *ProtobufSettings) GetXShallowEnvelopeMessageDescriptor() []byte {
	if m != nil {
		return m.XShallowEnvelopeMessageDescriptor
	}
	return nil
}

func (m *ProtobufSettings) GetProtobufDescriptorSet() string {
	if m != nil {
		return m.ProtobufDescriptorSet
	}
	return ""
}

type AvroSettings struct {
	// Used by CLI; desktop should not set/use this.
	// @gotags: kong:"help='If encode-type is set to avro, must specify avro schema file',existingfile"
	AvroSchemaFile string `protobuf:"bytes,1,opt,name=avro_schema_file,json=avroSchemaFile,proto3" json:"avro_schema_file,omitempty" kong:"help='If encode-type is set to avro, must specify avro schema file',existingfile"`
	// Used by desktop; CLI should not set/use this.
	// @gotags: kong:"-"
	Schema               []byte   `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty" kong:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvroSettings) Reset()         { *m = AvroSettings{} }
func (m *AvroSettings) String() string { return proto.CompactTextString(m) }
func (*AvroSettings) ProtoMessage()    {}
func (*AvroSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{1}
}

func (m *AvroSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvroSettings.Unmarshal(m, b)
}
func (m *AvroSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvroSettings.Marshal(b, m, deterministic)
}
func (m *AvroSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvroSettings.Merge(m, src)
}
func (m *AvroSettings) XXX_Size() int {
	return xxx_messageInfo_AvroSettings.Size(m)
}
func (m *AvroSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AvroSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AvroSettings proto.InternalMessageInfo

func (m *AvroSettings) GetAvroSchemaFile() string {
	if m != nil {
		return m.AvroSchemaFile
	}
	return ""
}

func (m *AvroSettings) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type JSONSchemaSettings struct {
	// Used by desktop; CLI should not set/use this.
	// @gotags: kong:"-"
	Schema               []byte   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty" kong:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONSchemaSettings) Reset()         { *m = JSONSchemaSettings{} }
func (m *JSONSchemaSettings) String() string { return proto.CompactTextString(m) }
func (*JSONSchemaSettings) ProtoMessage()    {}
func (*JSONSchemaSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{2}
}

func (m *JSONSchemaSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONSchemaSettings.Unmarshal(m, b)
}
func (m *JSONSchemaSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONSchemaSettings.Marshal(b, m, deterministic)
}
func (m *JSONSchemaSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONSchemaSettings.Merge(m, src)
}
func (m *JSONSchemaSettings) XXX_Size() int {
	return xxx_messageInfo_JSONSchemaSettings.Size(m)
}
func (m *JSONSchemaSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONSchemaSettings.DiscardUnknown(m)
}

var xxx_messageInfo_JSONSchemaSettings proto.InternalMessageInfo

func (m *JSONSchemaSettings) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

type EncodeOptions struct {
	// Use an existing schema for encoding (and ignore all other encode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// @gotags: kong:"help='Encode type (options: unset, jsonpb, avro)',default=unset,type=pbenum,pbenum_strip_prefix=ENCODE_TYPE_,pbenum_lowercase"
	EncodeType EncodeType `protobuf:"varint,2,opt,name=encode_type,json=encodeType,proto3,enum=protos.encoding.EncodeType" json:"encode_type,omitempty" kong:"help='Encode type (options: unset, jsonpb, avro)',default=unset,type=pbenum,pbenum_strip_prefix=ENCODE_TYPE_,pbenum_lowercase"`
	// @gotags: kong:"embed,group=protobuf"
	ProtobufSettings *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty" kong:"embed,group=protobuf"`
	// @gotags: kong:"embed,group=avro"
	AvroSettings         *AvroSettings `protobuf:"bytes,4,opt,name=avro_settings,json=avroSettings,proto3" json:"avro_settings,omitempty" kong:"embed,group=avro"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EncodeOptions) Reset()         { *m = EncodeOptions{} }
func (m *EncodeOptions) String() string { return proto.CompactTextString(m) }
func (*EncodeOptions) ProtoMessage()    {}
func (*EncodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{3}
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

func (m *EncodeOptions) GetAvroSettings() *AvroSettings {
	if m != nil {
		return m.AvroSettings
	}
	return nil
}

type DecodeOptions struct {
	// Use an existing schema for decoding (and ignore all other decode settings)
	// @gotags: kong:"-"
	SchemaId string `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty" kong:"-"`
	// @gotags: kong:"help='Decode type (options: unset, protobuf, avro, thrift, flatbuffer)',type=pbenum,pbenum_strip_prefix=DECODE_TYPE_,pbenum_lowercase,default=unset"
	DecodeType DecodeType `protobuf:"varint,2,opt,name=decode_type,json=decodeType,proto3,enum=protos.encoding.DecodeType" json:"decode_type,omitempty" kong:"help='Decode type (options: unset, protobuf, avro, thrift, flatbuffer)',type=pbenum,pbenum_strip_prefix=DECODE_TYPE_,pbenum_lowercase,default=unset"`
	// @gotags: kong:"embed,group=protobuf"
	ProtobufSettings *ProtobufSettings `protobuf:"bytes,3,opt,name=protobuf_settings,json=protobufSettings,proto3" json:"protobuf_settings,omitempty" kong:"embed,group=protobuf"`
	// @gotags: kong:"embed,group=avro"
	AvroSettings         *AvroSettings `protobuf:"bytes,4,opt,name=avro_settings,json=avroSettings,proto3" json:"avro_settings,omitempty" kong:"embed,group=avro"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DecodeOptions) Reset()         { *m = DecodeOptions{} }
func (m *DecodeOptions) String() string { return proto.CompactTextString(m) }
func (*DecodeOptions) ProtoMessage()    {}
func (*DecodeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f6b0c5cf15b9fa, []int{4}
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

func (m *DecodeOptions) GetAvroSettings() *AvroSettings {
	if m != nil {
		return m.AvroSettings
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.encoding.EncodeType", EncodeType_name, EncodeType_value)
	proto.RegisterEnum("protos.encoding.DecodeType", DecodeType_name, DecodeType_value)
	proto.RegisterEnum("protos.encoding.EnvelopeType", EnvelopeType_name, EnvelopeType_value)
	proto.RegisterType((*ProtobufSettings)(nil), "protos.encoding.ProtobufSettings")
	proto.RegisterType((*AvroSettings)(nil), "protos.encoding.AvroSettings")
	proto.RegisterType((*JSONSchemaSettings)(nil), "protos.encoding.JSONSchemaSettings")
	proto.RegisterType((*EncodeOptions)(nil), "protos.encoding.EncodeOptions")
	proto.RegisterType((*DecodeOptions)(nil), "protos.encoding.DecodeOptions")
}

func init() { proto.RegisterFile("ps_encoding_options.proto", fileDescriptor_b8f6b0c5cf15b9fa) }

var fileDescriptor_b8f6b0c5cf15b9fa = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0xda, 0x4a,
	0x14, 0x7d, 0x26, 0xe4, 0x83, 0x1b, 0x92, 0x38, 0x93, 0x84, 0x38, 0x89, 0x22, 0xf1, 0xf2, 0xde,
	0x02, 0xa5, 0x2d, 0x48, 0xa9, 0x54, 0x75, 0xd1, 0xaa, 0x82, 0x62, 0x94, 0x54, 0xd4, 0x76, 0x6d,
	0x27, 0x55, 0xb3, 0x19, 0x19, 0x7b, 0x02, 0x23, 0x19, 0xc6, 0xf2, 0x18, 0xaa, 0xfc, 0x82, 0xfe,
	0xa9, 0xfe, 0xb6, 0xaa, 0xf2, 0xf8, 0x03, 0x03, 0xa9, 0xaa, 0xee, 0xba, 0x02, 0x9f, 0x7b, 0xee,
	0x99, 0x3b, 0xe7, 0x5c, 0x1b, 0x4e, 0x02, 0x8e, 0xc9, 0xc4, 0x65, 0x1e, 0x9d, 0x0c, 0x31, 0x0b,
	0x22, 0xca, 0x26, 0xbc, 0x19, 0x84, 0x2c, 0x62, 0x68, 0x4f, 0xfc, 0xf0, 0x66, 0x56, 0xbe, 0xf8,
	0x5e, 0x06, 0xd9, 0x88, 0xb1, 0xc1, 0xf4, 0xc1, 0x22, 0x51, 0x44, 0x27, 0x43, 0x8e, 0xae, 0xe0,
	0x28, 0x48, 0x31, 0x1c, 0x32, 0x16, 0xe1, 0x31, 0xe1, 0xdc, 0x19, 0x12, 0x45, 0xaa, 0x4b, 0x8d,
	0x8a, 0x79, 0x90, 0x15, 0x4d, 0xc6, 0xa2, 0x8f, 0x49, 0x09, 0xfd, 0x07, 0x3b, 0x79, 0x8f, 0x47,
	0x43, 0xae, 0x94, 0xea, 0x6b, 0x8d, 0x8a, 0x59, 0xcd, 0xc0, 0x2e, 0x0d, 0x39, 0x7a, 0x06, 0x08,
	0x2f, 0x2a, 0x7b, 0x34, 0x54, 0xd6, 0x84, 0xea, 0x9e, 0x51, 0x50, 0xed, 0xd2, 0x10, 0x29, 0xb0,
	0xe9, 0x84, 0xee, 0x88, 0xce, 0x88, 0x52, 0xae, 0x4b, 0x8d, 0xaa, 0x99, 0x3d, 0xa2, 0x26, 0x1c,
	0x64, 0x23, 0x61, 0x8f, 0x70, 0x37, 0xa4, 0x41, 0xc4, 0x42, 0x65, 0x5d, 0xb0, 0xf6, 0xd3, 0x89,
	0xba, 0x79, 0x01, 0x59, 0x50, 0xcb, 0x4f, 0x25, 0x93, 0x19, 0xf1, 0x59, 0x40, 0x70, 0xf4, 0x18,
	0x10, 0x65, 0xa3, 0x2e, 0x35, 0x76, 0xaf, 0xce, 0x9b, 0x4b, 0xb6, 0x34, 0xd5, 0x94, 0x65, 0x3f,
	0x06, 0xc4, 0x3c, 0xcc, 0x9a, 0x8b, 0x28, 0x7a, 0x0d, 0x0a, 0x1f, 0x39, 0xbe, 0xcf, 0xbe, 0xce,
	0x35, 0x33, 0x9f, 0x36, 0xc5, 0x8d, 0x6a, 0x69, 0x3d, 0x6b, 0xcb, 0xac, 0x6a, 0xc3, 0xf9, 0x4a,
	0xe7, 0x03, 0x25, 0xbe, 0x87, 0x27, 0xd3, 0xf1, 0x80, 0x84, 0xca, 0x56, 0x5d, 0x6a, 0xac, 0x9b,
	0xa7, 0x4b, 0xed, 0xbd, 0x98, 0xa2, 0x09, 0x06, 0xd2, 0xe0, 0x7f, 0xfc, 0xab, 0xd3, 0x8b, 0x96,
	0x54, 0x84, 0x25, 0x75, 0xeb, 0xc9, 0x41, 0x0a, 0x0e, 0xbd, 0x82, 0xe3, 0x79, 0x7a, 0x39, 0x8c,
	0x39, 0x89, 0x14, 0x10, 0x77, 0xc9, 0x17, 0x62, 0xde, 0x64, 0x91, 0xe8, 0xc2, 0x80, 0x6a, 0x7b,
	0x16, 0xb2, 0x7c, 0x73, 0x1a, 0x20, 0x3b, 0xb3, 0x90, 0x61, 0xee, 0x8e, 0xc8, 0xd8, 0xc1, 0x0f,
	0xd4, 0xcf, 0x96, 0x66, 0x37, 0xc6, 0x2d, 0x01, 0xf7, 0xa8, 0x4f, 0x50, 0x0d, 0x36, 0x12, 0x92,
	0x52, 0x12, 0x33, 0xa6, 0x4f, 0x17, 0xcf, 0x01, 0x7d, 0xb0, 0x74, 0x2d, 0x61, 0xe6, 0xba, 0x73,
	0xb6, 0xb4, 0xc0, 0xfe, 0x21, 0xc1, 0x8e, 0x1a, 0x87, 0x46, 0xf4, 0x64, 0xcf, 0xd1, 0x19, 0x54,
	0xd2, 0xc3, 0xa9, 0x97, 0x1e, 0xbd, 0x95, 0x00, 0x37, 0x1e, 0x7a, 0x03, 0xdb, 0x22, 0xe2, 0x34,
	0xfd, 0x92, 0x48, 0xff, 0xec, 0x89, 0xf4, 0x63, 0x8e, 0xc8, 0x1e, 0x48, 0xfe, 0x1f, 0x69, 0xb0,
	0x9f, 0x9b, 0xc4, 0xd3, 0xc9, 0xc4, 0xf2, 0x6e, 0x5f, 0xfd, 0xbb, 0xa2, 0xb1, 0xfc, 0x52, 0x99,
	0x72, 0xb0, 0xfc, 0x9a, 0x75, 0x60, 0x27, 0x31, 0x2b, 0xd3, 0x2a, 0x0b, 0xad, 0xd5, 0x6d, 0x2c,
	0x5a, 0x6c, 0x56, 0x9d, 0xc2, 0x93, 0x30, 0xa0, 0x4b, 0xfe, 0xc4, 0x00, 0x8f, 0xfc, 0xde, 0x80,
	0x44, 0x31, 0x31, 0xc0, 0x23, 0x7f, 0xb3, 0x01, 0x97, 0x9f, 0x00, 0xe6, 0x71, 0xa1, 0x23, 0xd8,
	0x57, 0xb5, 0xf7, 0x7a, 0x57, 0xc5, 0xf6, 0x17, 0x43, 0xc5, 0xb7, 0x9a, 0xa5, 0xda, 0xf2, 0x3f,
	0xa8, 0x06, 0xa8, 0x08, 0xc7, 0x0b, 0x66, 0x74, 0x64, 0x09, 0x1d, 0x82, 0x5c, 0xc4, 0xdb, 0x77,
	0xa6, 0x2e, 0x97, 0x2e, 0xbf, 0x49, 0x00, 0x73, 0x07, 0x62, 0xcd, 0xae, 0xba, 0xaa, 0xa9, 0xc0,
	0x61, 0x11, 0x36, 0x4c, 0xdd, 0xd6, 0x3b, 0xb7, 0xbd, 0x44, 0xb5, 0x58, 0x49, 0x54, 0xe3, 0x19,
	0x8a, 0xa8, 0x7d, 0x6d, 0xde, 0xf4, 0x6c, 0x79, 0x0d, 0x9d, 0x42, 0xad, 0x88, 0xf7, 0xfa, 0x6d,
	0xbb, 0x73, 0xdb, 0xeb, 0xa9, 0xa6, 0x5c, 0xbe, 0xbc, 0x87, 0xea, 0xc2, 0x37, 0xe7, 0x18, 0x0e,
	0x54, 0xed, 0x4e, 0xed, 0xeb, 0xc6, 0x53, 0x17, 0x2c, 0x16, 0xba, 0xaa, 0x6a, 0xc8, 0x12, 0x3a,
	0x81, 0xa3, 0x45, 0xdc, 0xba, 0x6e, 0xf7, 0xfb, 0xfa, 0x67, 0xb9, 0xd4, 0x79, 0x77, 0xff, 0x76,
	0x48, 0xa3, 0xd1, 0x74, 0xd0, 0x74, 0xd9, 0xb8, 0x35, 0x70, 0x22, 0x77, 0xe4, 0xb2, 0x30, 0x68,
	0x05, 0xbe, 0xf8, 0xc4, 0xbc, 0x48, 0x76, 0x86, 0xb7, 0x06, 0x53, 0xea, 0x7b, 0xad, 0x21, 0x6b,
	0x25, 0xa1, 0xb4, 0xb2, 0x50, 0x06, 0x1b, 0x02, 0x78, 0xf9, 0x33, 0x00, 0x00, 0xff, 0xff, 0xef,
	0xe8, 0xbd, 0x68, 0x6f, 0x06, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connect.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	conns "github.com/batchcorp/plumber-schemas/build/go/protos/conns"
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

type Connection struct {
	// Friendly name to identify this connection by (used in plumber-server logs)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Any notes associated with this connection (stored plumber-server side)
	Notes string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	// Types that are valid to be assigned to Conn:
	//	*Connection_Kafka
	Conn                 isConnection_Conn `protobuf_oneof:"conn"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{0}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Connection) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

type isConnection_Conn interface {
	isConnection_Conn()
}

type Connection_Kafka struct {
	Kafka *conns.Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

func (*Connection_Kafka) isConnection_Conn() {}

func (m *Connection) GetConn() isConnection_Conn {
	if m != nil {
		return m.Conn
	}
	return nil
}

func (m *Connection) GetKafka() *conns.Kafka {
	if x, ok := m.GetConn().(*Connection_Kafka); ok {
		return x.Kafka
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Connection) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Connection_Kafka)(nil),
	}
}

type GetAllConnectionsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllConnectionsRequest) Reset()         { *m = GetAllConnectionsRequest{} }
func (m *GetAllConnectionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllConnectionsRequest) ProtoMessage()    {}
func (*GetAllConnectionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{1}
}

func (m *GetAllConnectionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllConnectionsRequest.Unmarshal(m, b)
}
func (m *GetAllConnectionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllConnectionsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllConnectionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllConnectionsRequest.Merge(m, src)
}
func (m *GetAllConnectionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllConnectionsRequest.Size(m)
}
func (m *GetAllConnectionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllConnectionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllConnectionsRequest proto.InternalMessageInfo

func (m *GetAllConnectionsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllConnectionsResponse struct {
	Connections          []*Connection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllConnectionsResponse) Reset()         { *m = GetAllConnectionsResponse{} }
func (m *GetAllConnectionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllConnectionsResponse) ProtoMessage()    {}
func (*GetAllConnectionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{2}
}

func (m *GetAllConnectionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllConnectionsResponse.Unmarshal(m, b)
}
func (m *GetAllConnectionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllConnectionsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllConnectionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllConnectionsResponse.Merge(m, src)
}
func (m *GetAllConnectionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllConnectionsResponse.Size(m)
}
func (m *GetAllConnectionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllConnectionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllConnectionsResponse proto.InternalMessageInfo

func (m *GetAllConnectionsResponse) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type GetConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetConnectionRequest) Reset()         { *m = GetConnectionRequest{} }
func (m *GetConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*GetConnectionRequest) ProtoMessage()    {}
func (*GetConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{3}
}

func (m *GetConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionRequest.Unmarshal(m, b)
}
func (m *GetConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionRequest.Marshal(b, m, deterministic)
}
func (m *GetConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionRequest.Merge(m, src)
}
func (m *GetConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_GetConnectionRequest.Size(m)
}
func (m *GetConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionRequest proto.InternalMessageInfo

func (m *GetConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *GetConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type GetConnectionResponse struct {
	Connection           *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetConnectionResponse) Reset()         { *m = GetConnectionResponse{} }
func (m *GetConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*GetConnectionResponse) ProtoMessage()    {}
func (*GetConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{4}
}

func (m *GetConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionResponse.Unmarshal(m, b)
}
func (m *GetConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionResponse.Marshal(b, m, deterministic)
}
func (m *GetConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionResponse.Merge(m, src)
}
func (m *GetConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_GetConnectionResponse.Size(m)
}
func (m *GetConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionResponse proto.InternalMessageInfo

func (m *GetConnectionResponse) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type CreateConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Connection           *Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateConnectionRequest) Reset()         { *m = CreateConnectionRequest{} }
func (m *CreateConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionRequest) ProtoMessage()    {}
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{5}
}

func (m *CreateConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionRequest.Unmarshal(m, b)
}
func (m *CreateConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionRequest.Marshal(b, m, deterministic)
}
func (m *CreateConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionRequest.Merge(m, src)
}
func (m *CreateConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionRequest.Size(m)
}
func (m *CreateConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionRequest proto.InternalMessageInfo

func (m *CreateConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateConnectionRequest) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type CreateConnectionResponse struct {
	// Set with uuid that represents connection if create is successful
	ConnectionId         string   `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConnectionResponse) Reset()         { *m = CreateConnectionResponse{} }
func (m *CreateConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionResponse) ProtoMessage()    {}
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{6}
}

func (m *CreateConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionResponse.Unmarshal(m, b)
}
func (m *CreateConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionResponse.Marshal(b, m, deterministic)
}
func (m *CreateConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionResponse.Merge(m, src)
}
func (m *CreateConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionResponse.Size(m)
}
func (m *CreateConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionResponse proto.InternalMessageInfo

func (m *CreateConnectionResponse) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type TestConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Connection           *Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestConnectionRequest) Reset()         { *m = TestConnectionRequest{} }
func (m *TestConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*TestConnectionRequest) ProtoMessage()    {}
func (*TestConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{7}
}

func (m *TestConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionRequest.Unmarshal(m, b)
}
func (m *TestConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionRequest.Marshal(b, m, deterministic)
}
func (m *TestConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionRequest.Merge(m, src)
}
func (m *TestConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_TestConnectionRequest.Size(m)
}
func (m *TestConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionRequest proto.InternalMessageInfo

func (m *TestConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *TestConnectionRequest) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type TestConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TestConnectionResponse) Reset()         { *m = TestConnectionResponse{} }
func (m *TestConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*TestConnectionResponse) ProtoMessage()    {}
func (*TestConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{8}
}

func (m *TestConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestConnectionResponse.Unmarshal(m, b)
}
func (m *TestConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestConnectionResponse.Marshal(b, m, deterministic)
}
func (m *TestConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestConnectionResponse.Merge(m, src)
}
func (m *TestConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_TestConnectionResponse.Size(m)
}
func (m *TestConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestConnectionResponse proto.InternalMessageInfo

func (m *TestConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type UpdateConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Connection           *Connection  `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateConnectionRequest) Reset()         { *m = UpdateConnectionRequest{} }
func (m *UpdateConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConnectionRequest) ProtoMessage()    {}
func (*UpdateConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{9}
}

func (m *UpdateConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConnectionRequest.Unmarshal(m, b)
}
func (m *UpdateConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConnectionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConnectionRequest.Merge(m, src)
}
func (m *UpdateConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConnectionRequest.Size(m)
}
func (m *UpdateConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConnectionRequest proto.InternalMessageInfo

func (m *UpdateConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *UpdateConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *UpdateConnectionRequest) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type UpdateConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateConnectionResponse) Reset()         { *m = UpdateConnectionResponse{} }
func (m *UpdateConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateConnectionResponse) ProtoMessage()    {}
func (*UpdateConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{10}
}

func (m *UpdateConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConnectionResponse.Unmarshal(m, b)
}
func (m *UpdateConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConnectionResponse.Marshal(b, m, deterministic)
}
func (m *UpdateConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConnectionResponse.Merge(m, src)
}
func (m *UpdateConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateConnectionResponse.Size(m)
}
func (m *UpdateConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConnectionResponse proto.InternalMessageInfo

func (m *UpdateConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteConnectionRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteConnectionRequest) Reset()         { *m = DeleteConnectionRequest{} }
func (m *DeleteConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionRequest) ProtoMessage()    {}
func (*DeleteConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{11}
}

func (m *DeleteConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionRequest.Unmarshal(m, b)
}
func (m *DeleteConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionRequest.Merge(m, src)
}
func (m *DeleteConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionRequest.Size(m)
}
func (m *DeleteConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionRequest proto.InternalMessageInfo

func (m *DeleteConnectionRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteConnectionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type DeleteConnectionResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteConnectionResponse) Reset()         { *m = DeleteConnectionResponse{} }
func (m *DeleteConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionResponse) ProtoMessage()    {}
func (*DeleteConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{12}
}

func (m *DeleteConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionResponse.Unmarshal(m, b)
}
func (m *DeleteConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionResponse.Merge(m, src)
}
func (m *DeleteConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionResponse.Size(m)
}
func (m *DeleteConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionResponse proto.InternalMessageInfo

func (m *DeleteConnectionResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Connection)(nil), "protos.Connection")
	proto.RegisterType((*GetAllConnectionsRequest)(nil), "protos.GetAllConnectionsRequest")
	proto.RegisterType((*GetAllConnectionsResponse)(nil), "protos.GetAllConnectionsResponse")
	proto.RegisterType((*GetConnectionRequest)(nil), "protos.GetConnectionRequest")
	proto.RegisterType((*GetConnectionResponse)(nil), "protos.GetConnectionResponse")
	proto.RegisterType((*CreateConnectionRequest)(nil), "protos.CreateConnectionRequest")
	proto.RegisterType((*CreateConnectionResponse)(nil), "protos.CreateConnectionResponse")
	proto.RegisterType((*TestConnectionRequest)(nil), "protos.TestConnectionRequest")
	proto.RegisterType((*TestConnectionResponse)(nil), "protos.TestConnectionResponse")
	proto.RegisterType((*UpdateConnectionRequest)(nil), "protos.UpdateConnectionRequest")
	proto.RegisterType((*UpdateConnectionResponse)(nil), "protos.UpdateConnectionResponse")
	proto.RegisterType((*DeleteConnectionRequest)(nil), "protos.DeleteConnectionRequest")
	proto.RegisterType((*DeleteConnectionResponse)(nil), "protos.DeleteConnectionResponse")
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor_778b7e3040344da6) }

var fileDescriptor_778b7e3040344da6 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xc9, 0xd2, 0x2d, 0x62, 0x0a, 0x07, 0xbc, 0x5b, 0x6a, 0x38, 0x55, 0xe1, 0x12, 0x09,
	0x91, 0x48, 0x65, 0xc5, 0x15, 0xed, 0x1f, 0xa9, 0x85, 0x4a, 0x48, 0x04, 0xb8, 0x70, 0x41, 0x4e,
	0x32, 0x34, 0x51, 0x13, 0x3b, 0xc4, 0xb6, 0x78, 0x0c, 0x5e, 0x80, 0x07, 0xe4, 0x31, 0x90, 0xed,
	0x40, 0xaa, 0xb4, 0x87, 0xd2, 0x8a, 0x3d, 0x25, 0xf9, 0x66, 0xfc, 0x7d, 0xbf, 0x99, 0x28, 0x81,
	0x87, 0xa9, 0xe0, 0x1c, 0x53, 0x15, 0xd6, 0x8d, 0x50, 0x82, 0x0c, 0xed, 0x45, 0x3e, 0x7d, 0x64,
	0x64, 0x19, 0xad, 0xd9, 0xd7, 0x35, 0x73, 0x25, 0x23, 0x55, 0x95, 0xe0, 0x11, 0xd3, 0x2a, 0x6f,
	0xa5, 0xb3, 0x56, 0x92, 0x8a, 0x29, 0x2d, 0x9d, 0xe8, 0xaf, 0x01, 0xae, 0x9d, 0x67, 0x21, 0x38,
	0x21, 0x30, 0xe0, 0xac, 0x42, 0xea, 0x4d, 0xbd, 0xe0, 0x7e, 0x6c, 0xef, 0xc9, 0x39, 0x9c, 0x72,
	0xa1, 0x50, 0xd2, 0x13, 0x2b, 0xba, 0x07, 0xf2, 0x1c, 0x4e, 0x6d, 0x1c, 0xcd, 0xa6, 0x5e, 0x30,
	0x9a, 0x9d, 0x39, 0x3b, 0x19, 0x5a, 0x92, 0x70, 0x69, 0x4a, 0x8b, 0x3b, 0xb1, 0xeb, 0xb9, 0x1a,
	0xc2, 0xc0, 0xe8, 0xfe, 0x0d, 0xd0, 0x39, 0xaa, 0xcb, 0xb2, 0xec, 0x22, 0x65, 0x8c, 0xdf, 0x34,
	0x4a, 0x45, 0x02, 0x18, 0x18, 0x56, 0xfa, 0xe3, 0x5d, 0xdf, 0xd0, 0x40, 0x87, 0x97, 0x5a, 0xe5,
	0xb1, 0xed, 0xf0, 0xdf, 0xc3, 0x93, 0x1d, 0x2e, 0xb2, 0x16, 0x5c, 0x22, 0xb9, 0x80, 0x51, 0xda,
	0xc9, 0xd4, 0x9b, 0xde, 0x0d, 0x46, 0x33, 0xf2, 0xc7, 0xac, 0x3b, 0x11, 0x6f, 0xb6, 0xf9, 0x08,
	0xe7, 0x73, 0x54, 0x1b, 0xd5, 0x7f, 0x85, 0x22, 0xcf, 0xfe, 0xbe, 0x9b, 0x42, 0xf0, 0x2f, 0x45,
	0xd6, 0xae, 0xf0, 0x41, 0x27, 0xbe, 0xc9, 0xfc, 0x25, 0x8c, 0x7b, 0x31, 0x2d, 0xf5, 0x0c, 0xa0,
	0x6b, 0xb4, 0x47, 0x77, 0x43, 0x6f, 0x74, 0xf9, 0xdf, 0x61, 0x72, 0xdd, 0x20, 0x53, 0x78, 0x0c,
	0xf6, 0x21, 0xc1, 0xaf, 0x81, 0x6e, 0x07, 0xb7, 0x83, 0xec, 0xb5, 0x06, 0x0d, 0xe3, 0x8f, 0x28,
	0xd5, 0x6d, 0x73, 0x2f, 0xe0, 0x71, 0x3f, 0xb6, 0xa5, 0x0e, 0x61, 0xe8, 0x3e, 0x0a, 0xfa, 0xeb,
	0x9e, 0xb5, 0x1a, 0xf7, 0x92, 0x3f, 0xd8, 0x6a, 0xdc, 0x76, 0xf9, 0x3f, 0x3d, 0x98, 0x7c, 0xaa,
	0xb3, 0x23, 0x77, 0xbf, 0xcf, 0xae, 0x7a, 0x83, 0x9e, 0xec, 0x35, 0xe8, 0x5b, 0xa0, 0xdb, 0x74,
	0x07, 0x8e, 0x9a, 0xc3, 0xe4, 0x06, 0x4b, 0xfc, 0xff, 0x93, 0x1a, 0xea, 0xed, 0xa4, 0xc3, 0xa8,
	0xaf, 0x5e, 0x7d, 0xbe, 0x58, 0x15, 0x2a, 0xd7, 0x89, 0xa9, 0x47, 0x09, 0x53, 0x69, 0x9e, 0x8a,
	0xa6, 0x8e, 0xea, 0x52, 0x57, 0x09, 0x36, 0x2f, 0x64, 0x9a, 0x63, 0xc5, 0x64, 0x94, 0xe8, 0xa2,
	0xcc, 0xa2, 0x95, 0x88, 0x9c, 0x5b, 0xe2, 0x7e, 0xa8, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x04, 0xd0, 0x33, 0x2a, 0x68, 0x05, 0x00, 0x00,
}

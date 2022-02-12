// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_base.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("ps_base.proto", fileDescriptor_ef931efdbd582aee) }

var fileDescriptor_ef931efdbd582aee = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xef, 0x44, 0xc7, 0x6e, 0xa9, 0x73, 0xa3, 0xcd, 0xb6, 0x5a, 0x1f, 0xc0, 0x0d, 0xa8,
	0x78, 0x27, 0x58, 0xb7, 0x18, 0x04, 0x41, 0xd9, 0x5a, 0x10, 0x6f, 0x96, 0x24, 0x7b, 0xe8, 0x06,
	0x26, 0x99, 0x31, 0x33, 0x11, 0xfa, 0x0e, 0x3e, 0xb4, 0x34, 0x67, 0x66, 0xe7, 0xcf, 0x49, 0x7a,
	0x55, 0xfa, 0xfd, 0x4e, 0x7e, 0x7c, 0x3b, 0x67, 0x76, 0xc3, 0x16, 0x4a, 0x6f, 0xab, 0x52, 0xc3,
	0x4a, 0xf5, 0xd2, 0x48, 0xfe, 0x68, 0xfc, 0xa3, 0xb3, 0x13, 0xa5, 0xb7, 0xb5, 0xec, 0x3a, 0xa8,
	0x0d, 0x92, 0x31, 0xd9, 0xdd, 0x75, 0x65, 0xdb, 0xd4, 0x36, 0x39, 0x56, 0x7a, 0xdb, 0x83, 0x28,
	0xef, 0xf0, 0xff, 0xb7, 0xff, 0x8e, 0xd8, 0xe2, 0x87, 0x18, 0xda, 0x0a, 0xfa, 0x6b, 0xe8, 0xff,
	0x42, 0xcf, 0x7f, 0xb1, 0x67, 0x05, 0x98, 0x4b, 0x21, 0xd6, 0xa8, 0x6a, 0x64, 0xa7, 0xf9, 0x05,
	0x8e, 0xeb, 0x15, 0x41, 0x1b, 0xf8, 0x33, 0x80, 0x36, 0xd9, 0xeb, 0x07, 0x26, 0xb4, 0x92, 0x9d,
	0x06, 0xfe, 0x8d, 0x2d, 0x0a, 0x30, 0x9e, 0xf0, 0xb3, 0xe0, 0x19, 0x1f, 0x3b, 0xe3, 0xf9, 0x0c,
	0xb5, 0xb6, 0x1b, 0x76, 0xb2, 0xee, 0xa1, 0x34, 0x10, 0x08, 0x5f, 0xb9, 0x47, 0x52, 0xe2, 0x9c,
	0x17, 0xf3, 0x03, 0x56, 0xfb, 0x9d, 0x1d, 0xff, 0x04, 0x1d, 0xb6, 0x3c, 0xf4, 0x88, 0x73, 0xa7,
	0x7c, 0x39, 0x87, 0x7d, 0xcf, 0x1b, 0xb5, 0x9b, 0xe9, 0x99, 0x12, 0xd2, 0x93, 0x0e, 0x78, 0xed,
	0x15, 0x08, 0x98, 0xd6, 0xa6, 0x84, 0x68, 0xe9, 0x80, 0xd5, 0x7e, 0x65, 0x47, 0xb8, 0xc0, 0xcd,
	0xfd, 0x25, 0xd1, 0x7c, 0x19, 0xaf, 0x15, 0x53, 0xa7, 0x3b, 0x9b, 0x86, 0x56, 0xf5, 0x91, 0x3d,
	0x2e, 0xc0, 0x8c, 0x21, 0x7f, 0x1e, 0x4c, 0x8e, 0x89, 0x53, 0xbc, 0xa0, 0xc0, 0x3e, 0xfe, 0x85,
	0x3d, 0xc5, 0x25, 0xa1, 0x21, 0x8b, 0x37, 0x17, 0x49, 0x96, 0x93, 0xcc, 0x7b, 0xf0, 0x10, 0x13,
	0x4f, 0x10, 0x12, 0x4f, 0xc4, 0xbc, 0x67, 0x03, 0x7a, 0x68, 0x53, 0x4f, 0x10, 0x12, 0x4f, 0xc4,
	0xac, 0xe7, 0x13, 0x7b, 0x72, 0x6d, 0xa4, 0x42, 0xcb, 0xe1, 0xe3, 0x1f, 0x22, 0xe7, 0x38, 0x9d,
	0x20, 0xbe, 0x09, 0xee, 0x2f, 0x69, 0x12, 0x84, 0xa4, 0x49, 0xc4, 0xac, 0x67, 0xcd, 0x58, 0x01,
	0xe6, 0x0a, 0x7f, 0x1f, 0xf8, 0x69, 0xb0, 0x09, 0x9b, 0x39, 0x4b, 0x36, 0x85, 0xa2, 0x2f, 0xf5,
	0xa5, 0x10, 0xce, 0x93, 0x5c, 0x8a, 0x44, 0x75, 0x3e, 0x43, 0xbd, 0x0d, 0x77, 0x48, 0x6c, 0x51,
	0x4c, 0x6c, 0x09, 0xf5, 0x07, 0x75, 0x7f, 0x7a, 0xce, 0x95, 0x85, 0x47, 0x9a, 0x98, 0x96, 0x93,
	0xcc, 0xb7, 0xc2, 0x4d, 0x92, 0x56, 0x51, 0x4c, 0x5a, 0x25, 0xd4, 0xdb, 0xf0, 0x7e, 0x11, 0x5b,
	0x14, 0x13, 0x5b, 0x42, 0xbd, 0x0d, 0x77, 0x4b, 0x6c, 0x51, 0x4c, 0x6c, 0x09, 0x45, 0xdb, 0xe7,
	0x0f, 0xbf, 0xdf, 0xdf, 0x36, 0x66, 0x3f, 0x54, 0xab, 0x5a, 0xb6, 0x79, 0x55, 0x9a, 0x7a, 0x5f,
	0xcb, 0x5e, 0xe5, 0x0a, 0x5f, 0x11, 0x6f, 0x74, 0xbd, 0x87, 0xb6, 0xd4, 0x79, 0x35, 0x34, 0x62,
	0x97, 0xdf, 0xca, 0x1c, 0x6d, 0x15, 0xbe, 0x82, 0xde, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3b,
	0x30, 0x0e, 0x37, 0x9a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlumberServerClient is the client API for PlumberServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlumberServerClient interface {
	// List configured/known connections
	GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error)
	GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error)
	CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error)
	UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error)
	ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error)
	StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error)
	DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error)
	// Get a specific dynamic tunnel
	GetDynamic(ctx context.Context, in *GetDynamicRequest, opts ...grpc.CallOption) (*GetDynamicResponse, error)
	GetAllDynamic(ctx context.Context, in *GetAllDynamicRequest, opts ...grpc.CallOption) (*GetAllDynamicResponse, error)
	CreateDynamic(ctx context.Context, in *CreateDynamicRequest, opts ...grpc.CallOption) (*CreateDynamicResponse, error)
	StopDynamic(ctx context.Context, in *StopDynamicRequest, opts ...grpc.CallOption) (*StopDynamicResponse, error)
	ResumeDynamic(ctx context.Context, in *ResumeDynamicRequest, opts ...grpc.CallOption) (*ResumeDynamicResponse, error)
	UpdateDynamic(ctx context.Context, in *UpdateDynamicRequest, opts ...grpc.CallOption) (*UpdateDynamicResponse, error)
	DeleteDynamic(ctx context.Context, in *DeleteDynamicRequest, opts ...grpc.CallOption) (*DeleteDynamicResponse, error)
}

type plumberServerClient struct {
	cc *grpc.ClientConn
}

func NewPlumberServerClient(cc *grpc.ClientConn) PlumberServerClient {
	return &plumberServerClient{cc}
}

func (c *plumberServerClient) GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error) {
	out := new(GetAllConnectionsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error) {
	out := new(UpdateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error) {
	out := new(GetAllRelaysResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllRelays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error) {
	out := new(GetRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error) {
	out := new(CreateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error) {
	out := new(UpdateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error) {
	out := new(ResumeRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error) {
	out := new(StopRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error) {
	out := new(DeleteRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetDynamic(ctx context.Context, in *GetDynamicRequest, opts ...grpc.CallOption) (*GetDynamicResponse, error) {
	out := new(GetDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetAllDynamic(ctx context.Context, in *GetAllDynamicRequest, opts ...grpc.CallOption) (*GetAllDynamicResponse, error) {
	out := new(GetAllDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateDynamic(ctx context.Context, in *CreateDynamicRequest, opts ...grpc.CallOption) (*CreateDynamicResponse, error) {
	out := new(CreateDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopDynamic(ctx context.Context, in *StopDynamicRequest, opts ...grpc.CallOption) (*StopDynamicResponse, error) {
	out := new(StopDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeDynamic(ctx context.Context, in *ResumeDynamicRequest, opts ...grpc.CallOption) (*ResumeDynamicResponse, error) {
	out := new(ResumeDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateDynamic(ctx context.Context, in *UpdateDynamicRequest, opts ...grpc.CallOption) (*UpdateDynamicResponse, error) {
	out := new(UpdateDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteDynamic(ctx context.Context, in *DeleteDynamicRequest, opts ...grpc.CallOption) (*DeleteDynamicResponse, error) {
	out := new(DeleteDynamicResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteDynamic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlumberServerServer is the server API for PlumberServer service.
type PlumberServerServer interface {
	// List configured/known connections
	GetAllConnections(context.Context, *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	GetAllRelays(context.Context, *GetAllRelaysRequest) (*GetAllRelaysResponse, error)
	GetRelay(context.Context, *GetRelayRequest) (*GetRelayResponse, error)
	CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error)
	UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error)
	ResumeRelay(context.Context, *ResumeRelayRequest) (*ResumeRelayResponse, error)
	StopRelay(context.Context, *StopRelayRequest) (*StopRelayResponse, error)
	DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error)
	// Get a specific dynamic tunnel
	GetDynamic(context.Context, *GetDynamicRequest) (*GetDynamicResponse, error)
	GetAllDynamic(context.Context, *GetAllDynamicRequest) (*GetAllDynamicResponse, error)
	CreateDynamic(context.Context, *CreateDynamicRequest) (*CreateDynamicResponse, error)
	StopDynamic(context.Context, *StopDynamicRequest) (*StopDynamicResponse, error)
	ResumeDynamic(context.Context, *ResumeDynamicRequest) (*ResumeDynamicResponse, error)
	UpdateDynamic(context.Context, *UpdateDynamicRequest) (*UpdateDynamicResponse, error)
	DeleteDynamic(context.Context, *DeleteDynamicRequest) (*DeleteDynamicResponse, error)
}

// UnimplementedPlumberServerServer can be embedded to have forward compatible implementations.
type UnimplementedPlumberServerServer struct {
}

func (*UnimplementedPlumberServerServer) GetAllConnections(ctx context.Context, req *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConnections not implemented")
}
func (*UnimplementedPlumberServerServer) GetConnection(ctx context.Context, req *GetConnectionRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (*UnimplementedPlumberServerServer) CreateConnection(ctx context.Context, req *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) TestConnection(ctx context.Context, req *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateConnection(ctx context.Context, req *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteConnection(ctx context.Context, req *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllRelays(ctx context.Context, req *GetAllRelaysRequest) (*GetAllRelaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRelays not implemented")
}
func (*UnimplementedPlumberServerServer) GetRelay(ctx context.Context, req *GetRelayRequest) (*GetRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelay not implemented")
}
func (*UnimplementedPlumberServerServer) CreateRelay(ctx context.Context, req *CreateRelayRequest) (*CreateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateRelay(ctx context.Context, req *UpdateRelayRequest) (*UpdateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeRelay(ctx context.Context, req *ResumeRelayRequest) (*ResumeRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeRelay not implemented")
}
func (*UnimplementedPlumberServerServer) StopRelay(ctx context.Context, req *StopRelayRequest) (*StopRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRelay not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteRelay(ctx context.Context, req *DeleteRelayRequest) (*DeleteRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelay not implemented")
}
func (*UnimplementedPlumberServerServer) GetDynamic(ctx context.Context, req *GetDynamicRequest) (*GetDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllDynamic(ctx context.Context, req *GetAllDynamicRequest) (*GetAllDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) CreateDynamic(ctx context.Context, req *CreateDynamicRequest) (*CreateDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) StopDynamic(ctx context.Context, req *StopDynamicRequest) (*StopDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeDynamic(ctx context.Context, req *ResumeDynamicRequest) (*ResumeDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateDynamic(ctx context.Context, req *UpdateDynamicRequest) (*UpdateDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDynamic not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteDynamic(ctx context.Context, req *DeleteDynamicRequest) (*DeleteDynamicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDynamic not implemented")
}

func RegisterPlumberServerServer(s *grpc.Server, srv PlumberServerServer) {
	s.RegisterService(&_PlumberServer_serviceDesc, srv)
}

func _PlumberServer_GetAllConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllConnections(ctx, req.(*GetAllConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetAllRelays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRelaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllRelays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllRelays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllRelays(ctx, req.(*GetAllRelaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetRelay(ctx, req.(*GetRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateRelay(ctx, req.(*CreateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateRelay(ctx, req.(*UpdateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeRelay(ctx, req.(*ResumeRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRelay(ctx, req.(*StopRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteRelay(ctx, req.(*DeleteRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetDynamic(ctx, req.(*GetDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetAllDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllDynamic(ctx, req.(*GetAllDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateDynamic(ctx, req.(*CreateDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopDynamic(ctx, req.(*StopDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeDynamic(ctx, req.(*ResumeDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateDynamic(ctx, req.(*UpdateDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteDynamic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDynamicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteDynamic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteDynamic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteDynamic(ctx, req.(*DeleteDynamicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlumberServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PlumberServer",
	HandlerType: (*PlumberServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllConnections",
			Handler:    _PlumberServer_GetAllConnections_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _PlumberServer_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _PlumberServer_CreateConnection_Handler,
		},
		{
			MethodName: "TestConnection",
			Handler:    _PlumberServer_TestConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _PlumberServer_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _PlumberServer_DeleteConnection_Handler,
		},
		{
			MethodName: "GetAllRelays",
			Handler:    _PlumberServer_GetAllRelays_Handler,
		},
		{
			MethodName: "GetRelay",
			Handler:    _PlumberServer_GetRelay_Handler,
		},
		{
			MethodName: "CreateRelay",
			Handler:    _PlumberServer_CreateRelay_Handler,
		},
		{
			MethodName: "UpdateRelay",
			Handler:    _PlumberServer_UpdateRelay_Handler,
		},
		{
			MethodName: "ResumeRelay",
			Handler:    _PlumberServer_ResumeRelay_Handler,
		},
		{
			MethodName: "StopRelay",
			Handler:    _PlumberServer_StopRelay_Handler,
		},
		{
			MethodName: "DeleteRelay",
			Handler:    _PlumberServer_DeleteRelay_Handler,
		},
		{
			MethodName: "GetDynamic",
			Handler:    _PlumberServer_GetDynamic_Handler,
		},
		{
			MethodName: "GetAllDynamic",
			Handler:    _PlumberServer_GetAllDynamic_Handler,
		},
		{
			MethodName: "CreateDynamic",
			Handler:    _PlumberServer_CreateDynamic_Handler,
		},
		{
			MethodName: "StopDynamic",
			Handler:    _PlumberServer_StopDynamic_Handler,
		},
		{
			MethodName: "ResumeDynamic",
			Handler:    _PlumberServer_ResumeDynamic_Handler,
		},
		{
			MethodName: "UpdateDynamic",
			Handler:    _PlumberServer_UpdateDynamic_Handler,
		},
		{
			MethodName: "DeleteDynamic",
			Handler:    _PlumberServer_DeleteDynamic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ps_base.proto",
}

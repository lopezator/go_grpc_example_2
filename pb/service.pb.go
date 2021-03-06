// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package echo

/*
Echo Service

Echo Service API consists of a single service which returns a message.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message represents a simple message sent to the Echo service.
type Message struct {
	// Id represents the message identifier.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The message to be sent.
	Msg                  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c3cfa37544d1c3ae, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "echo.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/echo.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Echo(context.Context, *Message) (*Message, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptor_service_c3cfa37544d1c3ae) }

var fileDescriptor_service_c3cfa37544d1c3ae = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4d,
	0xce, 0xc8, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x51, 0xd2, 0xe6, 0x62,
	0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60,
	0x02, 0x0b, 0x80, 0x98, 0x46, 0x11, 0x5c, 0xdc, 0xae, 0xc9, 0x19, 0xf9, 0xc1, 0x10, 0x5b, 0x84,
	0x3c, 0xb9, 0x58, 0x40, 0x5c, 0x21, 0x5e, 0x3d, 0x90, 0x45, 0x7a, 0x50, 0x73, 0xa4, 0x50, 0xb9,
	0x4a, 0xca, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x92, 0x55, 0x92, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0xad,
	0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x07, 0xa9, 0xd0, 0xaf, 0xce, 0x4c, 0xa9, 0xd5, 0xaf, 0xce,
	0x2d, 0x4e, 0xaf, 0x75, 0x62, 0x8b, 0x02, 0x3b, 0x36, 0x89, 0x0d, 0xec, 0x2a, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x6f, 0x0f, 0xfd, 0xcd, 0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event_generic.proto

package event_listener

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The response message
type GenericEventResponse struct {
	// 0 OK
	Response             int32    `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericEventResponse) Reset()         { *m = GenericEventResponse{} }
func (m *GenericEventResponse) String() string { return proto.CompactTextString(m) }
func (*GenericEventResponse) ProtoMessage()    {}
func (*GenericEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae53bbab00c44209, []int{0}
}

func (m *GenericEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericEventResponse.Unmarshal(m, b)
}
func (m *GenericEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericEventResponse.Marshal(b, m, deterministic)
}
func (m *GenericEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericEventResponse.Merge(m, src)
}
func (m *GenericEventResponse) XXX_Size() int {
	return xxx_messageInfo_GenericEventResponse.Size(m)
}
func (m *GenericEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericEventResponse proto.InternalMessageInfo

func (m *GenericEventResponse) GetResponse() int32 {
	if m != nil {
		return m.Response
	}
	return 0
}

// The response message
type GenericEventMessage struct {
	// Generic event response that just contains Json with MessageId as string and other stuff
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericEventMessage) Reset()         { *m = GenericEventMessage{} }
func (m *GenericEventMessage) String() string { return proto.CompactTextString(m) }
func (*GenericEventMessage) ProtoMessage()    {}
func (*GenericEventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae53bbab00c44209, []int{1}
}

func (m *GenericEventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericEventMessage.Unmarshal(m, b)
}
func (m *GenericEventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericEventMessage.Marshal(b, m, deterministic)
}
func (m *GenericEventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericEventMessage.Merge(m, src)
}
func (m *GenericEventMessage) XXX_Size() int {
	return xxx_messageInfo_GenericEventMessage.Size(m)
}
func (m *GenericEventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericEventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GenericEventMessage proto.InternalMessageInfo

func (m *GenericEventMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericEventResponse)(nil), "event_listener.GenericEventResponse")
	proto.RegisterType((*GenericEventMessage)(nil), "event_listener.GenericEventMessage")
}

func init() { proto.RegisterFile("event_generic.proto", fileDescriptor_ae53bbab00c44209) }

var fileDescriptor_ae53bbab00c44209 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0x89, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xca, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x83, 0x08, 0xe6, 0x64, 0x16, 0x97, 0x80, 0xc4, 0x95, 0x8c, 0xb8, 0x44, 0xdc, 0x21, 0x0a,
	0x5c, 0x41, 0x12, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x45,
	0x50, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x9c, 0xaf, 0xa4, 0xcf, 0x25, 0x8c, 0xac,
	0xc7, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x48, 0x82, 0x8b, 0x1d, 0xca, 0x04, 0xeb, 0xe0,
	0x0c, 0x82, 0x71, 0x8d, 0x8a, 0x51, 0x2d, 0xf1, 0x81, 0x5a, 0x2e, 0x14, 0xcd, 0xc5, 0x83, 0x2c,
	0x2e, 0xa4, 0xac, 0x87, 0xea, 0x3a, 0x3d, 0x2c, 0xd6, 0x48, 0xa9, 0xe0, 0x53, 0x04, 0x73, 0xbf,
	0x12, 0x43, 0x12, 0x1b, 0xd8, 0xc3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xca, 0xe0,
	0xce, 0x07, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GenericEventListenerClient is the client API for GenericEventListener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenericEventListenerClient interface {
	// Sends a command
	GenericEvent(ctx context.Context, in *GenericEventMessage, opts ...grpc.CallOption) (*GenericEventResponse, error)
}

type genericEventListenerClient struct {
	cc *grpc.ClientConn
}

func NewGenericEventListenerClient(cc *grpc.ClientConn) GenericEventListenerClient {
	return &genericEventListenerClient{cc}
}

func (c *genericEventListenerClient) GenericEvent(ctx context.Context, in *GenericEventMessage, opts ...grpc.CallOption) (*GenericEventResponse, error) {
	out := new(GenericEventResponse)
	err := c.cc.Invoke(ctx, "/event_listener.GenericEventListener/GenericEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericEventListenerServer is the server API for GenericEventListener service.
type GenericEventListenerServer interface {
	// Sends a command
	GenericEvent(context.Context, *GenericEventMessage) (*GenericEventResponse, error)
}

func RegisterGenericEventListenerServer(s *grpc.Server, srv GenericEventListenerServer) {
	s.RegisterService(&_GenericEventListener_serviceDesc, srv)
}

func _GenericEventListener_GenericEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericEventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericEventListenerServer).GenericEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event_listener.GenericEventListener/GenericEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericEventListenerServer).GenericEvent(ctx, req.(*GenericEventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenericEventListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event_listener.GenericEventListener",
	HandlerType: (*GenericEventListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenericEvent",
			Handler:    _GenericEventListener_GenericEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event_generic.proto",
}

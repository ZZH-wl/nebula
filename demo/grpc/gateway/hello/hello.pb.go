// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package nebula_core_srv_hello

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "nebula.core.srv.hello.Message")
	proto.RegisterType((*Request)(nil), "nebula.core.srv.hello.Request")
	proto.RegisterType((*Response)(nil), "nebula.core.srv.hello.Response")
	proto.RegisterType((*StreamingRequest)(nil), "nebula.core.srv.hello.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "nebula.core.srv.hello.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "nebula.core.srv.hello.Ping")
	proto.RegisterType((*Pong)(nil), "nebula.core.srv.hello.Pong")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbb, 0x34, 0x4d, 0xfb, 0xcd, 0x77, 0xa9, 0x83, 0x8a, 0xa4, 0x5a, 0x65, 0x2f, 0xc6,
	0xcb, 0x52, 0xf4, 0x11, 0xbc, 0xa8, 0x20, 0x48, 0x7c, 0x00, 0xd9, 0x96, 0x61, 0x2d, 0x26, 0xbb,
	0x35, 0x93, 0x08, 0x3e, 0xbd, 0x92, 0xcd, 0x06, 0x44, 0xba, 0x78, 0xdb, 0xe1, 0xf7, 0x9b, 0x61,
	0xfe, 0xc3, 0xc2, 0xff, 0x57, 0x2a, 0x4b, 0xa7, 0x76, 0xb5, 0x6b, 0x1c, 0x1e, 0x59, 0x5a, 0xb7,
	0xa5, 0x56, 0x1b, 0x57, 0x93, 0xe2, 0xfa, 0x43, 0x79, 0x28, 0x17, 0x30, 0x7d, 0x24, 0x66, 0x6d,
	0x08, 0xe7, 0x30, 0x66, 0xfd, 0x79, 0x22, 0x2e, 0x44, 0xfe, 0xaf, 0xe8, 0x9e, 0xf2, 0x0c, 0xa6,
	0x05, 0xbd, 0xb7, 0xc4, 0x0d, 0x22, 0x24, 0x56, 0x57, 0x14, 0xa8, 0x7f, 0xcb, 0x53, 0x98, 0x15,
	0xc4, 0x3b, 0x67, 0xd9, 0x37, 0x57, 0x6c, 0x86, 0xe6, 0x8a, 0x8d, 0xcc, 0x61, 0xfe, 0xdc, 0xd4,
	0xa4, 0xab, 0xad, 0x35, 0xc3, 0x94, 0x43, 0x98, 0x6c, 0x5c, 0x6b, 0x1b, 0xef, 0x8d, 0x8b, 0xbe,
	0x90, 0x57, 0x70, 0xf0, 0xc3, 0x0c, 0x03, 0xf7, 0xab, 0x4b, 0x48, 0x9e, 0xb6, 0xd6, 0xe0, 0x31,
	0xa4, 0xdc, 0xd4, 0xee, 0x8d, 0x02, 0x0e, 0x95, 0xe7, 0x2e, 0xce, 0xaf, 0xbf, 0x04, 0x4c, 0xee,
	0xba, 0xe0, 0x78, 0x0f, 0xc9, 0xad, 0x2e, 0x4b, 0x5c, 0xaa, 0xbd, 0x87, 0x51, 0x61, 0xe5, 0xec,
	0x3c, 0xca, 0xfb, 0x45, 0xe5, 0x08, 0x5f, 0x20, 0xed, 0xf7, 0xc7, 0xcb, 0x88, 0xfc, 0xfb, 0x10,
	0x59, 0xfe, 0xb7, 0x38, 0x8c, 0x5f, 0x09, 0x7c, 0x80, 0x59, 0x97, 0xda, 0x27, 0x5b, 0x44, 0x3a,
	0x3b, 0x21, 0x8b, 0x42, 0x67, 0x8d, 0x1c, 0xe5, 0x62, 0x25, 0xd6, 0xa9, 0xff, 0x0e, 0x37, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x5e, 0x40, 0x09, 0x1d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Hello_StreamClient, error)
	PingPong(ctx context.Context, opts ...grpc.CallOption) (Hello_PingPongClient, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nebula.core.srv.hello.Hello/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Hello_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/nebula.core.srv.hello.Hello/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_StreamClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type helloStreamClient struct {
	grpc.ClientStream
}

func (x *helloStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (Hello_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[1], "/nebula.core.srv.hello.Hello/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloPingPongClient{stream}
	return x, nil
}

type Hello_PingPongClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type helloPingPongClient struct {
	grpc.ClientStream
}

func (x *helloPingPongClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloPingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Call(context.Context, *Request) (*Response, error)
	Stream(*StreamingRequest, Hello_StreamServer) error
	PingPong(Hello_PingPongServer) error
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedHelloServer) Stream(req *StreamingRequest, srv Hello_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedHelloServer) PingPong(srv Hello_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebula.core.srv.hello.Hello/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).Stream(m, &helloStreamServer{stream})
}

type Hello_StreamServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type helloStreamServer struct {
	grpc.ServerStream
}

func (x *helloStreamServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).PingPong(&helloPingPongServer{stream})
}

type Hello_PingPongServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type helloPingPongServer struct {
	grpc.ServerStream
}

func (x *helloPingPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloPingPongServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nebula.core.srv.hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Hello_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Hello_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPong",
			Handler:       _Hello_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
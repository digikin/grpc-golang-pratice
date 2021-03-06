// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calc.proto

package calcpb

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

type Greeting struct {
	FirstValue           int32    `protobuf:"varint,1,opt,name=first_value,json=firstValue,proto3" json:"first_value,omitempty"`
	SecondValue          int32    `protobuf:"varint,2,opt,name=second_value,json=secondValue,proto3" json:"second_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_9193762e063fe25f, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstValue() int32 {
	if m != nil {
		return m.FirstValue
	}
	return 0
}

func (m *Greeting) GetSecondValue() int32 {
	if m != nil {
		return m.SecondValue
	}
	return 0
}

type GreetRequest struct {
	GreetingFirst        int32    `protobuf:"varint,1,opt,name=greeting_first,json=greetingFirst,proto3" json:"greeting_first,omitempty"`
	GreetingSecond       int32    `protobuf:"varint,2,opt,name=greeting_second,json=greetingSecond,proto3" json:"greeting_second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9193762e063fe25f, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreetingFirst() int32 {
	if m != nil {
		return m.GreetingFirst
	}
	return 0
}

func (m *GreetRequest) GetGreetingSecond() int32 {
	if m != nil {
		return m.GreetingSecond
	}
	return 0
}

type GreetResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9193762e063fe25f, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Greeting)(nil), "calc.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "calc.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "calc.GreetResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calc.proto", fileDescriptor_9193762e063fe25f)
}

var fileDescriptor_9193762e063fe25f = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xad, 0xd8, 0x50, 0xa6, 0xad, 0xc2, 0x08, 0x22, 0x5e, 0xd4, 0x05, 0xa9, 0xa7, 0x0a,
	0xf5, 0x1f, 0xe4, 0xa0, 0x37, 0x0f, 0x09, 0x78, 0xf0, 0x60, 0x48, 0xd6, 0x31, 0x04, 0x96, 0xec,
	0xba, 0x1f, 0xf9, 0xfd, 0xb2, 0x93, 0x4d, 0xb4, 0xb7, 0x79, 0x1f, 0xde, 0x79, 0x18, 0x06, 0x84,
	0xac, 0x95, 0x0c, 0xaa, 0xf6, 0xda, 0x3e, 0xfd, 0x8d, 0xa6, 0xe1, 0xb0, 0x37, 0x56, 0x7b, 0x8d,
	0x67, 0x71, 0x16, 0x6f, 0xb0, 0x7a, 0xb5, 0x44, 0xbe, 0xeb, 0x5b, 0xbc, 0x85, 0xf5, 0x77, 0x67,
	0x9d, 0xaf, 0x86, 0x5a, 0x05, 0xba, 0x5e, 0xdc, 0x2d, 0x1e, 0x97, 0x05, 0x30, 0x7a, 0x8f, 0x04,
	0xef, 0x61, 0xe3, 0x48, 0xea, 0xfe, 0x2b, 0x35, 0x4e, 0xb9, 0xb1, 0x1e, 0x19, 0x57, 0xc4, 0x27,
	0x6c, 0xd8, 0x57, 0xd0, 0x4f, 0x20, 0xe7, 0xf1, 0x01, 0xce, 0xdb, 0xe4, 0xaf, 0xd8, 0x94, 0xb4,
	0xdb, 0x89, 0xbe, 0x44, 0x88, 0x3b, 0xb8, 0x98, 0x6b, 0xa3, 0x2e, 0xc9, 0xe7, 0xed, 0x92, 0xa9,
	0xd8, 0xc1, 0x36, 0xf9, 0x9d, 0xd1, 0xbd, 0x23, 0xbc, 0x82, 0xcc, 0x92, 0x0b, 0x6a, 0x12, 0xa7,
	0x74, 0xc8, 0xd3, 0x21, 0x25, 0xd9, 0xa1, 0x93, 0x84, 0x07, 0x58, 0x72, 0x46, 0xdc, 0xf3, 0x13,
	0xfe, 0x5f, 0x79, 0x73, 0x79, 0xc4, 0x46, 0xb3, 0x38, 0xc9, 0x57, 0x1f, 0x59, 0xe4, 0xa6, 0x69,
	0x32, 0xfe, 0xd9, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x75, 0x1a, 0x62, 0x59, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/calc.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calc.proto",
}

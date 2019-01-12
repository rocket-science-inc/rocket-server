// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type GetEventsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_969575fb91067c26, []int{0}
}
func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (dst *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(dst, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

type GetEventsReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsReply) Reset()         { *m = GetEventsReply{} }
func (m *GetEventsReply) String() string { return proto.CompactTextString(m) }
func (*GetEventsReply) ProtoMessage()    {}
func (*GetEventsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_969575fb91067c26, []int{1}
}
func (m *GetEventsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsReply.Unmarshal(m, b)
}
func (m *GetEventsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsReply.Marshal(b, m, deterministic)
}
func (dst *GetEventsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsReply.Merge(dst, src)
}
func (m *GetEventsReply) XXX_Size() int {
	return xxx_messageInfo_GetEventsReply.Size(m)
}
func (m *GetEventsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsReply proto.InternalMessageInfo

type AddEventRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEventRequest) Reset()         { *m = AddEventRequest{} }
func (m *AddEventRequest) String() string { return proto.CompactTextString(m) }
func (*AddEventRequest) ProtoMessage()    {}
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_969575fb91067c26, []int{2}
}
func (m *AddEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventRequest.Unmarshal(m, b)
}
func (m *AddEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventRequest.Marshal(b, m, deterministic)
}
func (dst *AddEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventRequest.Merge(dst, src)
}
func (m *AddEventRequest) XXX_Size() int {
	return xxx_messageInfo_AddEventRequest.Size(m)
}
func (m *AddEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventRequest proto.InternalMessageInfo

type AddEventReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEventReply) Reset()         { *m = AddEventReply{} }
func (m *AddEventReply) String() string { return proto.CompactTextString(m) }
func (*AddEventReply) ProtoMessage()    {}
func (*AddEventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_969575fb91067c26, []int{3}
}
func (m *AddEventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventReply.Unmarshal(m, b)
}
func (m *AddEventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventReply.Marshal(b, m, deterministic)
}
func (dst *AddEventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventReply.Merge(dst, src)
}
func (m *AddEventReply) XXX_Size() int {
	return xxx_messageInfo_AddEventReply.Size(m)
}
func (m *AddEventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetEventsRequest)(nil), "pb.GetEventsRequest")
	proto.RegisterType((*GetEventsReply)(nil), "pb.GetEventsReply")
	proto.RegisterType((*AddEventRequest)(nil), "pb.AddEventRequest")
	proto.RegisterType((*AddEventReply)(nil), "pb.AddEventReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsReply, error)
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventReply, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsReply, error) {
	out := new(GetEventsReply)
	err := c.cc.Invoke(ctx, "/pb.Events/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventReply, error) {
	out := new(AddEventReply)
	err := c.cc.Invoke(ctx, "/pb.Events/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsReply, error)
	AddEvent(context.Context, *AddEventRequest) (*AddEventReply, error)
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Events/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Events/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _Events_GetEvents_Handler,
		},
		{
			MethodName: "AddEvent",
			Handler:    _Events_AddEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_events_969575fb91067c26) }

var fileDescriptor_events_969575fb91067c26 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x12, 0xe2, 0x12,
	0x70, 0x4f, 0x2d, 0x71, 0x05, 0x0b, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x09, 0x70,
	0xf1, 0x21, 0x89, 0x15, 0xe4, 0x54, 0x2a, 0x09, 0x72, 0xf1, 0x3b, 0xa6, 0xa4, 0x80, 0x45, 0x60,
	0x8a, 0xf8, 0xb9, 0x78, 0x11, 0x42, 0x05, 0x39, 0x95, 0x46, 0xc5, 0x5c, 0x6c, 0x10, 0x2d, 0x42,
	0xa6, 0x5c, 0x9c, 0x70, 0xfd, 0x42, 0x22, 0x7a, 0x05, 0x49, 0x7a, 0xe8, 0x56, 0x48, 0x09, 0xa1,
	0x89, 0x16, 0xe4, 0x54, 0x0a, 0x19, 0x71, 0x71, 0xc0, 0x4c, 0x14, 0x12, 0x06, 0xc9, 0xa3, 0x59,
	0x29, 0x25, 0x88, 0x2a, 0x58, 0x90, 0x53, 0x99, 0xc4, 0x06, 0xf6, 0x89, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0xdf, 0xa1, 0x6c, 0xd9, 0x00, 0x00, 0x00,
}

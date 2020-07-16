// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/debug/debug.proto

package debug

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DumpRequest struct {
	// Recursed is true if this request is a recursive call from another request.
	// Callers should leave it unset, it's used to prevent infinite loops of
	// recursive calls.
	Recursed             bool     `protobuf:"varint,1,opt,name=recursed,proto3" json:"recursed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{0}
}
func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(m, src)
}
func (m *DumpRequest) XXX_Size() int {
	return m.Size()
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

func (m *DumpRequest) GetRecursed() bool {
	if m != nil {
		return m.Recursed
	}
	return false
}

type ProfileRequest struct {
	Profile              string          `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Duration             *types.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Worker               *Worker         `protobuf:"bytes,3,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{1}
}
func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *ProfileRequest) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *ProfileRequest) GetWorker() *Worker {
	if m != nil {
		return m.Worker
	}
	return nil
}

type Worker struct {
	Pod                  string   `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Container            string   `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{2}
}
func (m *Worker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return m.Size()
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *Worker) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

type BinaryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinaryRequest) Reset()         { *m = BinaryRequest{} }
func (m *BinaryRequest) String() string { return proto.CompactTextString(m) }
func (*BinaryRequest) ProtoMessage()    {}
func (*BinaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d15a320d0127c22, []int{3}
}
func (m *BinaryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BinaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BinaryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BinaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinaryRequest.Merge(m, src)
}
func (m *BinaryRequest) XXX_Size() int {
	return m.Size()
}
func (m *BinaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BinaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BinaryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DumpRequest)(nil), "debug.DumpRequest")
	proto.RegisterType((*ProfileRequest)(nil), "debug.ProfileRequest")
	proto.RegisterType((*Worker)(nil), "debug.Worker")
	proto.RegisterType((*BinaryRequest)(nil), "debug.BinaryRequest")
}

func init() { proto.RegisterFile("client/debug/debug.proto", fileDescriptor_6d15a320d0127c22) }

var fileDescriptor_6d15a320d0127c22 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x18, 0x74, 0xad, 0x4d, 0xdb, 0xaf, 0x54, 0x65, 0x51, 0x88, 0x55, 0x42, 0x09, 0x08, 0xf5, 0x92,
	0x48, 0x45, 0x10, 0x44, 0xc4, 0xd2, 0x07, 0x90, 0x1c, 0x14, 0xbc, 0xe5, 0xe7, 0x6b, 0x1a, 0x4c,
	0xb3, 0x71, 0x93, 0xa5, 0xf4, 0x0d, 0x7c, 0xb4, 0x1e, 0x7d, 0x04, 0xe9, 0x93, 0x48, 0x77, 0x37,
	0xb5, 0xa5, 0x87, 0x5e, 0xc2, 0x7e, 0xf3, 0xcd, 0xce, 0xcc, 0x0e, 0x01, 0x33, 0x4c, 0x13, 0xcc,
	0x4a, 0x37, 0xc2, 0x40, 0xc4, 0xea, 0xeb, 0xe4, 0x9c, 0x95, 0x8c, 0xd6, 0xe5, 0xd0, 0xb5, 0x62,
	0xc6, 0xe2, 0x14, 0x5d, 0x09, 0x06, 0x62, 0xec, 0xce, 0xb8, 0x9f, 0xe7, 0xc8, 0x0b, 0x45, 0xdb,
	0xdd, 0x47, 0x82, 0xfb, 0x65, 0xc2, 0x32, 0xb5, 0xb7, 0x6f, 0xa0, 0x3d, 0x12, 0xd3, 0xdc, 0xc3,
	0x2f, 0x81, 0x45, 0x49, 0xbb, 0xd0, 0xe4, 0x18, 0x0a, 0x5e, 0x60, 0x64, 0x92, 0x1e, 0xe9, 0x37,
	0xbd, 0xf5, 0x6c, 0x7f, 0x13, 0x38, 0x7e, 0xe5, 0x6c, 0x9c, 0xa4, 0x58, 0xd1, 0x4d, 0x68, 0xe4,
	0x0a, 0x91, 0xec, 0x96, 0x57, 0x8d, 0xf4, 0x1e, 0x9a, 0x95, 0x93, 0x79, 0xd8, 0x23, 0xfd, 0xf6,
	0xe0, 0xc2, 0x51, 0x51, 0x9c, 0x2a, 0x8a, 0x33, 0xd2, 0x04, 0x6f, 0x4d, 0xa5, 0xd7, 0x60, 0xcc,
	0x18, 0xff, 0x44, 0x6e, 0xd6, 0xe4, 0xa5, 0x8e, 0xa3, 0xde, 0xfc, 0x2e, 0x41, 0x4f, 0x2f, 0xed,
	0x07, 0x30, 0x14, 0x42, 0x4f, 0xa1, 0x96, 0xb3, 0x48, 0xbb, 0xaf, 0x8e, 0xf4, 0x0a, 0x5a, 0x21,
	0xcb, 0x4a, 0x3f, 0xc9, 0x90, 0x4b, 0xeb, 0x96, 0xf7, 0x0f, 0xd8, 0x27, 0xd0, 0x19, 0x26, 0x99,
	0xcf, 0xe7, 0xfa, 0x09, 0x83, 0x05, 0x81, 0xfa, 0x68, 0xe5, 0x41, 0x1f, 0xe1, 0x68, 0x55, 0x05,
	0xa5, 0xda, 0x73, 0xa3, 0x97, 0xee, 0xe5, 0x4e, 0xf8, 0xe1, 0xbc, 0xc4, 0xe2, 0xcd, 0x4f, 0x05,
	0xda, 0x07, 0xb7, 0x84, 0xbe, 0x40, 0x43, 0x77, 0x43, 0xcf, 0xf5, 0xfd, 0xed, 0xae, 0xf6, 0x4b,
	0x3c, 0x83, 0xa1, 0xa2, 0xd1, 0x33, 0xad, 0xb0, 0x95, 0x74, 0xaf, 0xc0, 0xf0, 0x69, 0xb1, 0xb4,
	0xc8, 0xcf, 0xd2, 0x22, 0xbf, 0x4b, 0x8b, 0x7c, 0xb8, 0x71, 0x52, 0x4e, 0x44, 0xe0, 0x84, 0x6c,
	0xea, 0xe6, 0x7e, 0x38, 0x99, 0x47, 0xc8, 0x37, 0x4f, 0x05, 0x0f, 0xdd, 0xcd, 0xbf, 0x2b, 0x30,
	0xa4, 0xee, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x2b, 0xac, 0x75, 0x74, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error)
	Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/debug.Debug/Dump", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_DumpClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugDumpClient struct {
	grpc.ClientStream
}

func (x *debugDumpClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[1], "/debug.Debug/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_ProfileClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugProfileClient struct {
	grpc.ClientStream
}

func (x *debugProfileClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[2], "/debug.Debug/Binary", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugBinaryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_BinaryClient interface {
	Recv() (*types.BytesValue, error)
	grpc.ClientStream
}

type debugBinaryClient struct {
	grpc.ClientStream
}

func (x *debugBinaryClient) Recv() (*types.BytesValue, error) {
	m := new(types.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Dump(*DumpRequest, Debug_DumpServer) error
	Profile(*ProfileRequest, Debug_ProfileServer) error
	Binary(*BinaryRequest, Debug_BinaryServer) error
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) Dump(req *DumpRequest, srv Debug_DumpServer) error {
	return status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (*UnimplementedDebugServer) Profile(req *ProfileRequest, srv Debug_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedDebugServer) Binary(req *BinaryRequest, srv Debug_BinaryServer) error {
	return status.Errorf(codes.Unimplemented, "method Binary not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Dump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Dump(m, &debugDumpServer{stream})
}

type Debug_DumpServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugDumpServer struct {
	grpc.ServerStream
}

func (x *debugDumpServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Profile(m, &debugProfileServer{stream})
}

type Debug_ProfileServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugProfileServer struct {
	grpc.ServerStream
}

func (x *debugProfileServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Binary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BinaryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Binary(m, &debugBinaryServer{stream})
}

type Debug_BinaryServer interface {
	Send(*types.BytesValue) error
	grpc.ServerStream
}

type debugBinaryServer struct {
	grpc.ServerStream
}

func (x *debugBinaryServer) Send(m *types.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Dump",
			Handler:       _Debug_Dump_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Profile",
			Handler:       _Debug_Profile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Binary",
			Handler:       _Debug_Binary_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "client/debug/debug.proto",
}

func (m *DumpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DumpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DumpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Recursed {
		i--
		if m.Recursed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Worker != nil {
		{
			size, err := m.Worker.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDebug(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Duration != nil {
		{
			size, err := m.Duration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDebug(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Worker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Worker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Worker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Container) > 0 {
		i -= len(m.Container)
		copy(dAtA[i:], m.Container)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.Container)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pod) > 0 {
		i -= len(m.Pod)
		copy(dAtA[i:], m.Pod)
		i = encodeVarintDebug(dAtA, i, uint64(len(m.Pod)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BinaryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BinaryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BinaryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintDebug(dAtA []byte, offset int, v uint64) int {
	offset -= sovDebug(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DumpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Recursed {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.Duration != nil {
		l = m.Duration.Size()
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.Worker != nil {
		l = m.Worker.Size()
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Worker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pod)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	l = len(m.Container)
	if l > 0 {
		n += 1 + l + sovDebug(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BinaryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDebug(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDebug(x uint64) (n int) {
	return sovDebug(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DumpRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recursed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Recursed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProfileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &types.Duration{}
			}
			if err := m.Duration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Worker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Worker == nil {
				m.Worker = &Worker{}
			}
			if err := m.Worker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Worker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Worker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Worker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Container", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDebug
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Container = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BinaryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BinaryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BinaryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDebug(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDebug
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDebug
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDebug
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDebug
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDebug
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDebug        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDebug          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDebug = fmt.Errorf("proto: unexpected end of group")
)

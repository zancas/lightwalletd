// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package walletrpc

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

// A BlockID message contains identifiers to select a block: a height or a
// hash. If the hash is present it takes precedence.
type BlockID struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{0}
}
func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (dst *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(dst, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// BlockRange technically allows ranging from hash to hash etc but this is not
// currently intended for support, though there is no reason you couldn't do
// it. Further permutations are left as an exercise.
type BlockRange struct {
	Start                *BlockID `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *BlockID `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRange) Reset()         { *m = BlockRange{} }
func (m *BlockRange) String() string { return proto.CompactTextString(m) }
func (*BlockRange) ProtoMessage()    {}
func (*BlockRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{1}
}
func (m *BlockRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRange.Unmarshal(m, b)
}
func (m *BlockRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRange.Marshal(b, m, deterministic)
}
func (dst *BlockRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRange.Merge(dst, src)
}
func (m *BlockRange) XXX_Size() int {
	return xxx_messageInfo_BlockRange.Size(m)
}
func (m *BlockRange) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRange.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRange proto.InternalMessageInfo

func (m *BlockRange) GetStart() *BlockID {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *BlockRange) GetEnd() *BlockID {
	if m != nil {
		return m.End
	}
	return nil
}

// A TxFilter contains the information needed to identify a particular
// transaction: either a block and an index, or a direct transaction hash.
type TxFilter struct {
	Block                *BlockID `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxFilter) Reset()         { *m = TxFilter{} }
func (m *TxFilter) String() string { return proto.CompactTextString(m) }
func (*TxFilter) ProtoMessage()    {}
func (*TxFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{2}
}
func (m *TxFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxFilter.Unmarshal(m, b)
}
func (m *TxFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxFilter.Marshal(b, m, deterministic)
}
func (dst *TxFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxFilter.Merge(dst, src)
}
func (m *TxFilter) XXX_Size() int {
	return xxx_messageInfo_TxFilter.Size(m)
}
func (m *TxFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TxFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TxFilter proto.InternalMessageInfo

func (m *TxFilter) GetBlock() *BlockID {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *TxFilter) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxFilter) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// RawTransaction contains the complete transaction data.
type RawTransaction struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawTransaction) Reset()         { *m = RawTransaction{} }
func (m *RawTransaction) String() string { return proto.CompactTextString(m) }
func (*RawTransaction) ProtoMessage()    {}
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{3}
}
func (m *RawTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawTransaction.Unmarshal(m, b)
}
func (m *RawTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawTransaction.Marshal(b, m, deterministic)
}
func (dst *RawTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTransaction.Merge(dst, src)
}
func (m *RawTransaction) XXX_Size() int {
	return xxx_messageInfo_RawTransaction.Size(m)
}
func (m *RawTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_RawTransaction proto.InternalMessageInfo

func (m *RawTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SendResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{4}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (dst *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(dst, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *SendResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

// Empty placeholder. Someday we may want to specify e.g. a particular chain fork.
type ChainSpec struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainSpec) Reset()         { *m = ChainSpec{} }
func (m *ChainSpec) String() string { return proto.CompactTextString(m) }
func (*ChainSpec) ProtoMessage()    {}
func (*ChainSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_087f17e455cf31eb, []int{5}
}
func (m *ChainSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainSpec.Unmarshal(m, b)
}
func (m *ChainSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainSpec.Marshal(b, m, deterministic)
}
func (dst *ChainSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainSpec.Merge(dst, src)
}
func (m *ChainSpec) XXX_Size() int {
	return xxx_messageInfo_ChainSpec.Size(m)
}
func (m *ChainSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChainSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BlockID)(nil), "cash.z.wallet.sdk.rpc.BlockID")
	proto.RegisterType((*BlockRange)(nil), "cash.z.wallet.sdk.rpc.BlockRange")
	proto.RegisterType((*TxFilter)(nil), "cash.z.wallet.sdk.rpc.TxFilter")
	proto.RegisterType((*RawTransaction)(nil), "cash.z.wallet.sdk.rpc.RawTransaction")
	proto.RegisterType((*SendResponse)(nil), "cash.z.wallet.sdk.rpc.SendResponse")
	proto.RegisterType((*ChainSpec)(nil), "cash.z.wallet.sdk.rpc.ChainSpec")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompactTxStreamerClient is the client API for CompactTxStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompactTxStreamerClient interface {
	GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error)
	GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error)
	GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error)
	GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error)
	SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error)
}

type compactTxStreamerClient struct {
	cc *grpc.ClientConn
}

func NewCompactTxStreamerClient(cc *grpc.ClientConn) CompactTxStreamerClient {
	return &compactTxStreamerClient{cc}
}

func (c *compactTxStreamerClient) GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error) {
	out := new(BlockID)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error) {
	out := new(CompactBlock)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompactTxStreamer_serviceDesc.Streams[0], "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlockRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetBlockRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetBlockRangeClient interface {
	Recv() (*CompactBlock, error)
	grpc.ClientStream
}

type compactTxStreamerGetBlockRangeClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetBlockRangeClient) Recv() (*CompactBlock, error) {
	m := new(CompactBlock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error) {
	out := new(RawTransaction)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompactTxStreamerServer is the server API for CompactTxStreamer service.
type CompactTxStreamerServer interface {
	GetLatestBlock(context.Context, *ChainSpec) (*BlockID, error)
	GetBlock(context.Context, *BlockID) (*CompactBlock, error)
	GetBlockRange(*BlockRange, CompactTxStreamer_GetBlockRangeServer) error
	GetTransaction(context.Context, *TxFilter) (*RawTransaction, error)
	SendTransaction(context.Context, *RawTransaction) (*SendResponse, error)
}

func RegisterCompactTxStreamerServer(s *grpc.Server, srv CompactTxStreamerServer) {
	s.RegisterService(&_CompactTxStreamer_serviceDesc, srv)
}

func _CompactTxStreamer_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, req.(*ChainSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlockRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetBlockRange(m, &compactTxStreamerGetBlockRangeServer{stream})
}

type CompactTxStreamer_GetBlockRangeServer interface {
	Send(*CompactBlock) error
	grpc.ServerStream
}

type compactTxStreamerGetBlockRangeServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetBlockRangeServer) Send(m *CompactBlock) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, req.(*TxFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, req.(*RawTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompactTxStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cash.z.wallet.sdk.rpc.CompactTxStreamer",
	HandlerType: (*CompactTxStreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestBlock",
			Handler:    _CompactTxStreamer_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _CompactTxStreamer_GetBlock_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _CompactTxStreamer_GetTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _CompactTxStreamer_SendTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockRange",
			Handler:       _CompactTxStreamer_GetBlockRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_087f17e455cf31eb) }

var fileDescriptor_service_087f17e455cf31eb = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x1b, 0xda, 0x8e, 0xe5, 0x9a, 0x0d, 0x61, 0x31, 0x34, 0x45, 0x08, 0x8a, 0x01, 0x69,
	0x4f, 0xd1, 0x34, 0xe0, 0x0b, 0xac, 0x88, 0x09, 0x09, 0x24, 0x70, 0xf3, 0x34, 0x1e, 0xa6, 0x9b,
	0x73, 0x34, 0x61, 0xa9, 0x1d, 0xd9, 0x16, 0xab, 0xf8, 0xa0, 0x7c, 0x1e, 0x14, 0xbb, 0x1b, 0x99,
	0x44, 0xd6, 0xbe, 0xf9, 0x9c, 0xdf, 0xfd, 0xff, 0x77, 0xbe, 0x0b, 0xec, 0x59, 0x32, 0xbf, 0x2a,
	0x49, 0x59, 0x63, 0xb4, 0xd3, 0xec, 0x40, 0xa2, 0x2d, 0xb3, 0xdf, 0xd9, 0x35, 0xd6, 0x35, 0xb9,
	0xcc, 0x16, 0x57, 0x99, 0x69, 0x64, 0x7a, 0x20, 0xf5, 0xb2, 0x41, 0xe9, 0x2e, 0x7e, 0x68, 0xb3,
	0x44, 0x67, 0x03, 0xcd, 0xdf, 0xc3, 0xc3, 0xd3, 0x5a, 0xcb, 0xab, 0x4f, 0x1f, 0xd8, 0x53, 0xd8,
	0x29, 0xa9, 0x5a, 0x94, 0xee, 0x30, 0x9a, 0x46, 0x47, 0x23, 0xb1, 0x8e, 0x18, 0x83, 0x51, 0x89,
	0xb6, 0x3c, 0x7c, 0x30, 0x8d, 0x8e, 0x12, 0xe1, 0xcf, 0xdc, 0x01, 0xf8, 0x34, 0x81, 0x6a, 0x41,
	0xec, 0x1d, 0x8c, 0xad, 0x43, 0x13, 0x12, 0x27, 0x27, 0xcf, 0xb3, 0xff, 0x96, 0x90, 0xad, 0x8d,
	0x44, 0x80, 0xd9, 0x31, 0x0c, 0x49, 0x15, 0x5e, 0x76, 0x73, 0x4e, 0x8b, 0xf2, 0x9f, 0xb0, 0x9b,
	0xaf, 0x3e, 0x56, 0xb5, 0x23, 0xd3, 0x7a, 0x5e, 0xb6, 0xdf, 0xb6, 0xf5, 0xf4, 0x30, 0x7b, 0x02,
	0xe3, 0x4a, 0x15, 0xb4, 0xf2, 0xae, 0x23, 0x11, 0x82, 0xdb, 0x0e, 0x87, 0x9d, 0x0e, 0x5f, 0xc3,
	0xbe, 0xc0, 0xeb, 0xdc, 0xa0, 0xb2, 0x28, 0x5d, 0xa5, 0x55, 0x4b, 0x15, 0xe8, 0xd0, 0x1b, 0x26,
	0xc2, 0x9f, 0xf9, 0x57, 0x48, 0xe6, 0xa4, 0x0a, 0x41, 0xb6, 0xd1, 0xca, 0x12, 0x7b, 0x06, 0x31,
	0x19, 0xa3, 0xcd, 0x4c, 0x17, 0xe4, 0xc1, 0xb1, 0xf8, 0x77, 0xc1, 0x38, 0x24, 0x3e, 0xf8, 0x42,
	0xd6, 0xe2, 0x82, 0x7c, 0x11, 0xb1, 0xb8, 0x73, 0xc7, 0x27, 0x10, 0xcf, 0x4a, 0xac, 0xd4, 0xbc,
	0x21, 0x79, 0xf2, 0x67, 0x08, 0x8f, 0x67, 0x61, 0x6e, 0xf9, 0x6a, 0xee, 0x0c, 0xe1, 0x92, 0x0c,
	0xcb, 0x61, 0xff, 0x8c, 0xdc, 0x67, 0x74, 0x64, 0x9d, 0xef, 0x8f, 0x4d, 0x7b, 0xba, 0xbf, 0x55,
	0x4a, 0x37, 0xbc, 0x0f, 0x1f, 0xb0, 0x6f, 0xb0, 0x7b, 0x46, 0x6b, 0xbd, 0x0d, 0x74, 0xfa, 0xaa,
	0xcf, 0x2f, 0xd4, 0xea, 0x31, 0x3e, 0x60, 0xdf, 0x61, 0xef, 0x46, 0x32, 0x2c, 0xca, 0xcb, 0xfb,
	0x74, 0x3d, 0xb2, 0xa5, 0xf4, 0x71, 0xc4, 0xce, 0xfd, 0x2b, 0x74, 0x07, 0xf4, 0xa2, 0x27, 0xf5,
	0x66, 0x67, 0xd2, 0x37, 0x3d, 0xc0, 0xdd, 0x41, 0xf3, 0x01, 0xbb, 0x80, 0x47, 0xed, 0x58, 0xbb,
	0xe2, 0xdb, 0xe5, 0xf6, 0x96, 0xdf, 0xdd, 0x12, 0x3e, 0x38, 0x9d, 0x9c, 0xc7, 0x01, 0x30, 0x8d,
	0xbc, 0xdc, 0xf1, 0xbf, 0xe2, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x6c, 0xab, 0xf9,
	0xc9, 0x03, 0x00, 0x00,
}
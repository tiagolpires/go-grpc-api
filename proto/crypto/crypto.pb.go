// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/crypto.proto

package crypto

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

type Crypto struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Votes                int64    `protobuf:"varint,3,opt,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crypto) Reset()         { *m = Crypto{} }
func (m *Crypto) String() string { return proto.CompactTextString(m) }
func (*Crypto) ProtoMessage()    {}
func (*Crypto) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{0}
}

func (m *Crypto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crypto.Unmarshal(m, b)
}
func (m *Crypto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crypto.Marshal(b, m, deterministic)
}
func (m *Crypto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crypto.Merge(m, src)
}
func (m *Crypto) XXX_Size() int {
	return xxx_messageInfo_Crypto.Size(m)
}
func (m *Crypto) XXX_DiscardUnknown() {
	xxx_messageInfo_Crypto.DiscardUnknown(m)
}

var xxx_messageInfo_Crypto proto.InternalMessageInfo

func (m *Crypto) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Crypto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Crypto) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

type GetCryptoByIdRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCryptoByIdRequest) Reset()         { *m = GetCryptoByIdRequest{} }
func (m *GetCryptoByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetCryptoByIdRequest) ProtoMessage()    {}
func (*GetCryptoByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{1}
}

func (m *GetCryptoByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCryptoByIdRequest.Unmarshal(m, b)
}
func (m *GetCryptoByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCryptoByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetCryptoByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCryptoByIdRequest.Merge(m, src)
}
func (m *GetCryptoByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetCryptoByIdRequest.Size(m)
}
func (m *GetCryptoByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCryptoByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCryptoByIdRequest proto.InternalMessageInfo

func (m *GetCryptoByIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateCryptoRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Votes                int64    `protobuf:"varint,2,opt,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCryptoRequest) Reset()         { *m = CreateCryptoRequest{} }
func (m *CreateCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCryptoRequest) ProtoMessage()    {}
func (*CreateCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{2}
}

func (m *CreateCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCryptoRequest.Unmarshal(m, b)
}
func (m *CreateCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCryptoRequest.Marshal(b, m, deterministic)
}
func (m *CreateCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCryptoRequest.Merge(m, src)
}
func (m *CreateCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCryptoRequest.Size(m)
}
func (m *CreateCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCryptoRequest proto.InternalMessageInfo

func (m *CreateCryptoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCryptoRequest) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

type CryptoResponse struct {
	Crypto               *Crypto  `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoResponse) Reset()         { *m = CryptoResponse{} }
func (m *CryptoResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoResponse) ProtoMessage()    {}
func (*CryptoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{3}
}

func (m *CryptoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoResponse.Unmarshal(m, b)
}
func (m *CryptoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoResponse.Marshal(b, m, deterministic)
}
func (m *CryptoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoResponse.Merge(m, src)
}
func (m *CryptoResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoResponse.Size(m)
}
func (m *CryptoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoResponse proto.InternalMessageInfo

func (m *CryptoResponse) GetCrypto() *Crypto {
	if m != nil {
		return m.Crypto
	}
	return nil
}

func (m *CryptoResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ListCryptosResponse struct {
	Cryptos              []*Crypto `protobuf:"bytes,1,rep,name=cryptos,proto3" json:"cryptos,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListCryptosResponse) Reset()         { *m = ListCryptosResponse{} }
func (m *ListCryptosResponse) String() string { return proto.CompactTextString(m) }
func (*ListCryptosResponse) ProtoMessage()    {}
func (*ListCryptosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{4}
}

func (m *ListCryptosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCryptosResponse.Unmarshal(m, b)
}
func (m *ListCryptosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCryptosResponse.Marshal(b, m, deterministic)
}
func (m *ListCryptosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCryptosResponse.Merge(m, src)
}
func (m *ListCryptosResponse) XXX_Size() int {
	return xxx_messageInfo_ListCryptosResponse.Size(m)
}
func (m *ListCryptosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCryptosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCryptosResponse proto.InternalMessageInfo

func (m *ListCryptosResponse) GetCryptos() []*Crypto {
	if m != nil {
		return m.Cryptos
	}
	return nil
}

func (m *ListCryptosResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SuccessMessageResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessMessageResponse) Reset()         { *m = SuccessMessageResponse{} }
func (m *SuccessMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessMessageResponse) ProtoMessage()    {}
func (*SuccessMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{5}
}

func (m *SuccessMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessMessageResponse.Unmarshal(m, b)
}
func (m *SuccessMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessMessageResponse.Marshal(b, m, deterministic)
}
func (m *SuccessMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessMessageResponse.Merge(m, src)
}
func (m *SuccessMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessMessageResponse.Size(m)
}
func (m *SuccessMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessMessageResponse proto.InternalMessageInfo

func (m *SuccessMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SuccessMessageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SuccessMessageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{6}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Crypto)(nil), "Crypto")
	proto.RegisterType((*GetCryptoByIdRequest)(nil), "GetCryptoByIdRequest")
	proto.RegisterType((*CreateCryptoRequest)(nil), "CreateCryptoRequest")
	proto.RegisterType((*CryptoResponse)(nil), "CryptoResponse")
	proto.RegisterType((*ListCryptosResponse)(nil), "ListCryptosResponse")
	proto.RegisterType((*SuccessMessageResponse)(nil), "SuccessMessageResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
}

func init() { proto.RegisterFile("crypto/crypto.proto", fileDescriptor_27e1fe69dc202ca1) }

var fileDescriptor_27e1fe69dc202ca1 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xcd, 0xb6, 0x08, 0x3a, 0x7c, 0x68, 0x96, 0xaa, 0x0d, 0x17, 0x6b, 0x0f, 0x86, 0xd3, 0x9a,
	0x60, 0x8c, 0x47, 0x13, 0xd0, 0x10, 0x13, 0xf5, 0x50, 0x7e, 0x01, 0x94, 0x89, 0x69, 0x22, 0xb4,
	0xee, 0x2c, 0x24, 0xfc, 0x72, 0xaf, 0x86, 0xee, 0x2e, 0xae, 0x58, 0x3d, 0xb5, 0x6f, 0x3a, 0xf3,
	0xde, 0xcc, 0x7b, 0x85, 0x6e, 0x2a, 0x37, 0x85, 0xca, 0xaf, 0xf5, 0x43, 0x14, 0x32, 0x57, 0x79,
	0x3c, 0x84, 0xfa, 0xa8, 0xc4, 0xbc, 0x03, 0x5e, 0x36, 0x0f, 0x59, 0xc4, 0xfa, 0xb5, 0xc4, 0xcb,
	0xe6, 0x9c, 0x43, 0x6d, 0x39, 0x5d, 0x60, 0xe8, 0x45, 0xac, 0x7f, 0x94, 0x94, 0xef, 0x3c, 0x80,
	0x83, 0x75, 0xae, 0x90, 0x42, 0x3f, 0x62, 0x7d, 0x3f, 0xd1, 0x20, 0xbe, 0x82, 0x60, 0x8c, 0x4a,
	0xd3, 0x0c, 0x37, 0x4f, 0xf3, 0x04, 0x3f, 0x56, 0x48, 0x6a, 0x9f, 0x31, 0xbe, 0x87, 0xee, 0x48,
	0xe2, 0x54, 0xa1, 0x6e, 0xb5, 0x6d, 0x56, 0x88, 0x55, 0x09, 0x79, 0xae, 0xd0, 0x18, 0x3a, 0x76,
	0x94, 0x8a, 0x7c, 0x49, 0xc8, 0x2f, 0xa0, 0xae, 0xcf, 0x29, 0xa7, 0x9b, 0x83, 0x86, 0x30, 0x0d,
	0xa6, 0xbc, 0x25, 0x42, 0x29, 0x73, 0x69, 0xce, 0xd0, 0x20, 0x7e, 0x85, 0xee, 0x73, 0x46, 0x66,
	0x65, 0xda, 0xb1, 0x5d, 0x42, 0x43, 0x8f, 0x51, 0xc8, 0x22, 0xdf, 0xa5, 0xb3, 0xf5, 0x3f, 0xf8,
	0x66, 0x70, 0x36, 0x59, 0xa5, 0x29, 0x12, 0xbd, 0x20, 0xd1, 0xf4, 0x0d, 0x77, 0x94, 0x21, 0x34,
	0x16, 0xba, 0x64, 0xee, 0xb3, 0x70, 0xfb, 0x85, 0xf4, 0x4c, 0xc9, 0x75, 0x98, 0x58, 0xf8, 0xad,
	0xe1, 0xbb, 0x1a, 0x1d, 0x68, 0x3d, 0x2e, 0x0a, 0xb5, 0x31, 0xb6, 0x0d, 0x3e, 0x19, 0xb4, 0xf5,
	0x76, 0x13, 0x94, 0xeb, 0x2c, 0x45, 0x7e, 0x07, 0xed, 0x1f, 0x39, 0xf0, 0x53, 0x51, 0x95, 0x4b,
	0xef, 0x58, 0xec, 0xb9, 0x38, 0x80, 0xa6, 0x63, 0x07, 0x6f, 0x0b, 0x57, 0xa8, 0x17, 0x88, 0x2a,
	0xaf, 0x6e, 0xa1, 0xe5, 0x86, 0xc9, 0x03, 0x51, 0x91, 0xed, 0x6f, 0xa9, 0x21, 0x9c, 0x3c, 0xe0,
	0x3b, 0xda, 0xbe, 0xff, 0xd6, 0x3c, 0x17, 0xd5, 0x9e, 0xce, 0xea, 0xe5, 0xaf, 0x7b, 0xf3, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x41, 0x0d, 0x29, 0x5f, 0xd1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoServiceClient interface {
	GetCryptoById(ctx context.Context, in *GetCryptoByIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	ListCryptos(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListCryptosResponse, error)
	CreateCrypto(ctx context.Context, in *CreateCryptoRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	DeleteCryptoById(ctx context.Context, in *GetCryptoByIdRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error)
}

type cryptoServiceClient struct {
	cc *grpc.ClientConn
}

func NewCryptoServiceClient(cc *grpc.ClientConn) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) GetCryptoById(ctx context.Context, in *GetCryptoByIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
	out := new(CryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/GetCryptoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListCryptos(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListCryptosResponse, error) {
	out := new(ListCryptosResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/ListCryptos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CreateCrypto(ctx context.Context, in *CreateCryptoRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
	out := new(CryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/CreateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteCryptoById(ctx context.Context, in *GetCryptoByIdRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error) {
	out := new(SuccessMessageResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/DeleteCryptoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
type CryptoServiceServer interface {
	GetCryptoById(context.Context, *GetCryptoByIdRequest) (*CryptoResponse, error)
	ListCryptos(context.Context, *EmptyRequest) (*ListCryptosResponse, error)
	CreateCrypto(context.Context, *CreateCryptoRequest) (*CryptoResponse, error)
	DeleteCryptoById(context.Context, *GetCryptoByIdRequest) (*SuccessMessageResponse, error)
}

// UnimplementedCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (*UnimplementedCryptoServiceServer) GetCryptoById(ctx context.Context, req *GetCryptoByIdRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCryptoById not implemented")
}
func (*UnimplementedCryptoServiceServer) ListCryptos(ctx context.Context, req *EmptyRequest) (*ListCryptosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCryptos not implemented")
}
func (*UnimplementedCryptoServiceServer) CreateCrypto(ctx context.Context, req *CreateCryptoRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) DeleteCryptoById(ctx context.Context, req *GetCryptoByIdRequest) (*SuccessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCryptoById not implemented")
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_GetCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCryptoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetCryptoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/GetCryptoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetCryptoById(ctx, req.(*GetCryptoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListCryptos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ListCryptos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/ListCryptos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ListCryptos(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CreateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/CreateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateCrypto(ctx, req.(*CreateCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCryptoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteCryptoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/DeleteCryptoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteCryptoById(ctx, req.(*GetCryptoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCryptoById",
			Handler:    _CryptoService_GetCryptoById_Handler,
		},
		{
			MethodName: "ListCryptos",
			Handler:    _CryptoService_ListCryptos_Handler,
		},
		{
			MethodName: "CreateCrypto",
			Handler:    _CryptoService_CreateCrypto_Handler,
		},
		{
			MethodName: "DeleteCryptoById",
			Handler:    _CryptoService_DeleteCryptoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto/crypto.proto",
}

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
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Votes                int64    `protobuf:"varint,4,opt,name=votes,proto3" json:"votes,omitempty"`
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

func (m *Crypto) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Crypto) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

type CryptoIdRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoIdRequest) Reset()         { *m = CryptoIdRequest{} }
func (m *CryptoIdRequest) String() string { return proto.CompactTextString(m) }
func (*CryptoIdRequest) ProtoMessage()    {}
func (*CryptoIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{1}
}

func (m *CryptoIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoIdRequest.Unmarshal(m, b)
}
func (m *CryptoIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoIdRequest.Marshal(b, m, deterministic)
}
func (m *CryptoIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoIdRequest.Merge(m, src)
}
func (m *CryptoIdRequest) XXX_Size() int {
	return xxx_messageInfo_CryptoIdRequest.Size(m)
}
func (m *CryptoIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoIdRequest proto.InternalMessageInfo

func (m *CryptoIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateCryptoRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Votes                int64    `protobuf:"varint,3,opt,name=votes,proto3" json:"votes,omitempty"`
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

func (m *CreateCryptoRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateCryptoRequest) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

type UpdateCryptoRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Votes                int64    `protobuf:"varint,4,opt,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCryptoRequest) Reset()         { *m = UpdateCryptoRequest{} }
func (m *UpdateCryptoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCryptoRequest) ProtoMessage()    {}
func (*UpdateCryptoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27e1fe69dc202ca1, []int{3}
}

func (m *UpdateCryptoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCryptoRequest.Unmarshal(m, b)
}
func (m *UpdateCryptoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCryptoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCryptoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCryptoRequest.Merge(m, src)
}
func (m *UpdateCryptoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCryptoRequest.Size(m)
}
func (m *UpdateCryptoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCryptoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCryptoRequest proto.InternalMessageInfo

func (m *UpdateCryptoRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateCryptoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCryptoRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UpdateCryptoRequest) GetVotes() int64 {
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
	return fileDescriptor_27e1fe69dc202ca1, []int{4}
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
	return fileDescriptor_27e1fe69dc202ca1, []int{5}
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
	return fileDescriptor_27e1fe69dc202ca1, []int{6}
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
	return fileDescriptor_27e1fe69dc202ca1, []int{7}
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
	proto.RegisterType((*CryptoIdRequest)(nil), "CryptoIdRequest")
	proto.RegisterType((*CreateCryptoRequest)(nil), "CreateCryptoRequest")
	proto.RegisterType((*UpdateCryptoRequest)(nil), "UpdateCryptoRequest")
	proto.RegisterType((*CryptoResponse)(nil), "CryptoResponse")
	proto.RegisterType((*ListCryptosResponse)(nil), "ListCryptosResponse")
	proto.RegisterType((*SuccessMessageResponse)(nil), "SuccessMessageResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
}

func init() { proto.RegisterFile("crypto/crypto.proto", fileDescriptor_27e1fe69dc202ca1) }

var fileDescriptor_27e1fe69dc202ca1 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x4d, 0x5b, 0x04, 0x1d, 0x3e, 0xb3, 0x6d, 0xb4, 0xe1, 0x62, 0xe9, 0x89, 0xd3, 0x9a, 0x60,
	0x88, 0x27, 0x2f, 0x82, 0x21, 0x24, 0xea, 0xa1, 0x04, 0xef, 0xd0, 0x4e, 0x4c, 0x13, 0x61, 0xeb,
	0xee, 0x82, 0xe1, 0x7f, 0xf8, 0x83, 0x4d, 0xbb, 0x2d, 0x6e, 0xa0, 0xc6, 0x90, 0x78, 0xea, 0xce,
	0xeb, 0xcc, 0x7b, 0xdd, 0x37, 0xaf, 0x60, 0x87, 0x7c, 0x97, 0x48, 0x76, 0xa3, 0x1e, 0x34, 0xe1,
	0x4c, 0x32, 0xff, 0x15, 0xaa, 0xa3, 0xac, 0x26, 0x2d, 0x30, 0xe3, 0xc8, 0x35, 0x3c, 0xa3, 0x5f,
	0x09, 0xcc, 0x38, 0x22, 0x04, 0x2a, 0xeb, 0xc5, 0x0a, 0x5d, 0xd3, 0x33, 0xfa, 0x17, 0x41, 0x76,
	0x4e, 0xb1, 0x90, 0x45, 0xe8, 0x5a, 0x0a, 0x4b, 0xcf, 0xc4, 0x81, 0xb3, 0x2d, 0x93, 0x28, 0xdc,
	0x8a, 0x67, 0xf4, 0xad, 0x40, 0x15, 0x7e, 0x0f, 0xda, 0x8a, 0x77, 0x1a, 0x05, 0xf8, 0xb1, 0x41,
	0x21, 0x0f, 0x05, 0xfc, 0x19, 0xd8, 0x23, 0x8e, 0x0b, 0x89, 0xaa, 0xb1, 0x68, 0x2b, 0x74, 0x8d,
	0x12, 0x5d, 0xb3, 0x4c, 0xd7, 0xd2, 0x75, 0x43, 0xb0, 0xe7, 0x49, 0x74, 0x44, 0xfa, 0xbf, 0x97,
	0x9b, 0x40, 0xab, 0xa0, 0x17, 0x09, 0x5b, 0x0b, 0x24, 0xd7, 0x50, 0x55, 0xb6, 0x66, 0x1a, 0xf5,
	0x41, 0x8d, 0xe6, 0x0d, 0x39, 0x9c, 0x12, 0x21, 0xe7, 0x8c, 0xe7, 0x8a, 0xaa, 0xf0, 0x5f, 0xc0,
	0x7e, 0x8a, 0x85, 0x54, 0xbd, 0x62, 0xcf, 0xd6, 0x83, 0x9a, 0x1a, 0x13, 0xae, 0xe1, 0x59, 0x3a,
	0x5d, 0x81, 0xff, 0xc2, 0xb7, 0x84, 0xcb, 0xd9, 0x26, 0x0c, 0x51, 0x88, 0x67, 0x14, 0x62, 0xf1,
	0x86, 0x7b, 0x4a, 0x17, 0x6a, 0x2b, 0x05, 0xe5, 0xc6, 0x16, 0x65, 0xfa, 0x46, 0xa8, 0x99, 0x8c,
	0xeb, 0x3c, 0x28, 0xca, 0x1f, 0x0d, 0x4b, 0xd7, 0x68, 0x41, 0xe3, 0x71, 0x95, 0xc8, 0x5d, 0x6e,
	0xed, 0xe0, 0xcb, 0x82, 0xa6, 0xfa, 0xba, 0x19, 0xf2, 0x6d, 0x1c, 0x22, 0x19, 0x42, 0x67, 0x9e,
	0xa4, 0x4e, 0x29, 0xf8, 0x61, 0x37, 0x8d, 0x48, 0x87, 0x1e, 0xc4, 0xa1, 0xdb, 0xa6, 0x07, 0x1e,
	0xde, 0x01, 0x19, 0xb3, 0xcf, 0xf5, 0xe9, 0x83, 0x03, 0x68, 0x4e, 0x50, 0x9e, 0x3a, 0x53, 0xd7,
	0x9c, 0x27, 0x4d, 0xaa, 0xdf, 0xa9, 0xeb, 0xd0, 0xb2, 0xb5, 0x0c, 0xa1, 0xa1, 0x07, 0x96, 0x38,
	0xb4, 0x24, 0xbf, 0xc7, 0x52, 0x43, 0x68, 0xe8, 0x91, 0x24, 0x0e, 0x2d, 0x49, 0xe8, 0xf1, 0xd8,
	0x3d, 0x74, 0xc6, 0xf8, 0x8e, 0x7f, 0x98, 0x71, 0x45, 0xcb, 0x17, 0xbe, 0xac, 0x66, 0xff, 0xf7,
	0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x89, 0x3b, 0x42, 0xf6, 0x03, 0x00, 0x00,
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
	UpvoteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	DownvoteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	GetCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	ListCryptos(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListCryptosResponse, error)
	CreateCrypto(ctx context.Context, in *CreateCryptoRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	UpdateCrypto(ctx context.Context, in *UpdateCryptoRequest, opts ...grpc.CallOption) (*CryptoResponse, error)
	DeleteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error)
}

type cryptoServiceClient struct {
	cc *grpc.ClientConn
}

func NewCryptoServiceClient(cc *grpc.ClientConn) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) UpvoteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
	out := new(CryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/UpvoteCryptoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DownvoteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
	out := new(CryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/DownvoteCryptoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
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

func (c *cryptoServiceClient) UpdateCrypto(ctx context.Context, in *UpdateCryptoRequest, opts ...grpc.CallOption) (*CryptoResponse, error) {
	out := new(CryptoResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/UpdateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteCryptoById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error) {
	out := new(SuccessMessageResponse)
	err := c.cc.Invoke(ctx, "/CryptoService/DeleteCryptoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
type CryptoServiceServer interface {
	UpvoteCryptoById(context.Context, *CryptoIdRequest) (*CryptoResponse, error)
	DownvoteCryptoById(context.Context, *CryptoIdRequest) (*CryptoResponse, error)
	GetCryptoById(context.Context, *CryptoIdRequest) (*CryptoResponse, error)
	ListCryptos(context.Context, *EmptyRequest) (*ListCryptosResponse, error)
	CreateCrypto(context.Context, *CreateCryptoRequest) (*CryptoResponse, error)
	UpdateCrypto(context.Context, *UpdateCryptoRequest) (*CryptoResponse, error)
	DeleteCryptoById(context.Context, *CryptoIdRequest) (*SuccessMessageResponse, error)
}

// UnimplementedCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (*UnimplementedCryptoServiceServer) UpvoteCryptoById(ctx context.Context, req *CryptoIdRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvoteCryptoById not implemented")
}
func (*UnimplementedCryptoServiceServer) DownvoteCryptoById(ctx context.Context, req *CryptoIdRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownvoteCryptoById not implemented")
}
func (*UnimplementedCryptoServiceServer) GetCryptoById(ctx context.Context, req *CryptoIdRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCryptoById not implemented")
}
func (*UnimplementedCryptoServiceServer) ListCryptos(ctx context.Context, req *EmptyRequest) (*ListCryptosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCryptos not implemented")
}
func (*UnimplementedCryptoServiceServer) CreateCrypto(ctx context.Context, req *CreateCryptoRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) UpdateCrypto(ctx context.Context, req *UpdateCryptoRequest) (*CryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) DeleteCryptoById(ctx context.Context, req *CryptoIdRequest) (*SuccessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCryptoById not implemented")
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_UpvoteCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpvoteCryptoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/UpvoteCryptoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpvoteCryptoById(ctx, req.(*CryptoIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DownvoteCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DownvoteCryptoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/DownvoteCryptoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DownvoteCryptoById(ctx, req.(*CryptoIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoIdRequest)
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
		return srv.(CryptoServiceServer).GetCryptoById(ctx, req.(*CryptoIdRequest))
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

func _CryptoService_UpdateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CryptoService/UpdateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, req.(*UpdateCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteCryptoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoIdRequest)
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
		return srv.(CryptoServiceServer).DeleteCryptoById(ctx, req.(*CryptoIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpvoteCryptoById",
			Handler:    _CryptoService_UpvoteCryptoById_Handler,
		},
		{
			MethodName: "DownvoteCryptoById",
			Handler:    _CryptoService_DownvoteCryptoById_Handler,
		},
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
			MethodName: "UpdateCrypto",
			Handler:    _CryptoService_UpdateCrypto_Handler,
		},
		{
			MethodName: "DeleteCryptoById",
			Handler:    _CryptoService_DeleteCryptoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto/crypto.proto",
}

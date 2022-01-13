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
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x14, 0x44, 0x92, 0x6b, 0xb7, 0xcf, 0x9f, 0xac, 0x44, 0x2b, 0x7c, 0xa9, 0xac, 0x93, 0x4f, 0xdb,
	0xa2, 0x62, 0x0a, 0x85, 0x5e, 0x6a, 0x17, 0x63, 0x48, 0x72, 0x90, 0x70, 0xee, 0xb2, 0xf4, 0x08,
	0x82, 0xc8, 0xab, 0xec, 0xae, 0x1d, 0xfc, 0x2f, 0xf3, 0x93, 0x82, 0xb4, 0x92, 0x23, 0x6c, 0x85,
	0x38, 0x90, 0x93, 0xf7, 0x8d, 0xdf, 0xcc, 0x48, 0xb3, 0x83, 0xc0, 0x8c, 0xf8, 0x21, 0x93, 0xec,
	0x87, 0xfa, 0xa1, 0x19, 0x67, 0x92, 0xb9, 0xb7, 0xd0, 0x9e, 0x17, 0x33, 0x19, 0x80, 0x9e, 0xc4,
	0xb6, 0xe6, 0x68, 0xd3, 0x96, 0xaf, 0x27, 0x31, 0x21, 0xd0, 0xda, 0x86, 0x29, 0xda, 0xba, 0xa3,
	0x4d, 0xbf, 0xf8, 0xc5, 0x39, 0xc7, 0x22, 0x16, 0xa3, 0x6d, 0x28, 0x2c, 0x3f, 0x13, 0x0b, 0x3e,
	0xed, 0x99, 0x44, 0x61, 0xb7, 0x1c, 0x6d, 0x6a, 0xf8, 0x6a, 0x70, 0x27, 0x30, 0x54, 0xba, 0xab,
	0xd8, 0xc7, 0x87, 0x1d, 0x0a, 0x79, 0x6a, 0xe0, 0x06, 0x60, 0xce, 0x39, 0x86, 0x12, 0xd5, 0x62,
	0xb5, 0x56, 0xf9, 0x6a, 0x0d, 0xbe, 0x7a, 0x93, 0xaf, 0x51, 0xf7, 0x8d, 0xc0, 0x5c, 0x67, 0xf1,
	0x99, 0xe8, 0xc7, 0xbe, 0xdc, 0x12, 0x06, 0x95, 0xbc, 0xc8, 0xd8, 0x56, 0x20, 0xf9, 0x0e, 0x6d,
	0x15, 0x6b, 0xe1, 0xd1, 0xf5, 0x3a, 0xb4, 0x5c, 0x28, 0xe1, 0x5c, 0x08, 0x39, 0x67, 0xbc, 0x74,
	0x54, 0x83, 0x7b, 0x03, 0xe6, 0x55, 0x22, 0xa4, 0xda, 0x15, 0x47, 0xb5, 0x09, 0x74, 0x14, 0x4d,
	0xd8, 0x9a, 0x63, 0xd4, 0xe5, 0x2a, 0xfc, 0x15, 0xbd, 0x0d, 0x7c, 0x0d, 0x76, 0x51, 0x84, 0x42,
	0x5c, 0xa3, 0x10, 0xe1, 0x1d, 0x1e, 0x25, 0x6d, 0xe8, 0xa4, 0x0a, 0x2a, 0x83, 0xad, 0xc6, 0xfc,
	0x1f, 0xa1, 0x38, 0x85, 0xd6, 0x67, 0xbf, 0x1a, 0x5f, 0x3c, 0x8c, 0xba, 0xc7, 0x00, 0x7a, 0xff,
	0xd3, 0x4c, 0x1e, 0xca, 0x68, 0xbd, 0x27, 0x03, 0xfa, 0xea, 0xe9, 0x02, 0xe4, 0xfb, 0x24, 0x42,
	0x32, 0x83, 0xd1, 0x3a, 0xcb, 0x93, 0x52, 0xf0, 0xbf, 0xc3, 0x2a, 0x26, 0x23, 0x7a, 0x52, 0x87,
	0xf1, 0x90, 0x9e, 0x64, 0xf8, 0x1b, 0xc8, 0x82, 0x3d, 0x6e, 0xdf, 0x4f, 0xfc, 0x03, 0xe6, 0x12,
	0xcb, 0x10, 0x03, 0xc9, 0x31, 0x4c, 0x2f, 0x64, 0xfe, 0xd4, 0x88, 0x07, 0xfd, 0x23, 0xf7, 0x52,
	0x3f, 0x0f, 0xba, 0xb5, 0x5b, 0x23, 0x7d, 0x5a, 0xcf, 0x63, 0x6c, 0xd1, 0xa6, 0x2b, 0x9d, 0x41,
	0xaf, 0x5e, 0x76, 0x62, 0xd1, 0x86, 0xee, 0x9f, 0x5b, 0xcd, 0xa0, 0x57, 0xaf, 0x33, 0xb1, 0x68,
	0x43, 0xbb, 0xcf, 0x69, 0x7f, 0x61, 0xb4, 0xc0, 0x7b, 0x7c, 0x23, 0xc8, 0x6f, 0xb4, 0xb9, 0x2c,
	0x9b, 0x76, 0xf1, 0x6d, 0xf8, 0xf5, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x20, 0x90, 0x05, 0x7b, 0x32,
	0x04, 0x00, 0x00,
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
	GetCryptoStreamById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (CryptoService_GetCryptoStreamByIdClient, error)
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

func (c *cryptoServiceClient) GetCryptoStreamById(ctx context.Context, in *CryptoIdRequest, opts ...grpc.CallOption) (CryptoService_GetCryptoStreamByIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CryptoService_serviceDesc.Streams[0], "/CryptoService/GetCryptoStreamById", opts...)
	if err != nil {
		return nil, err
	}
	x := &cryptoServiceGetCryptoStreamByIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CryptoService_GetCryptoStreamByIdClient interface {
	Recv() (*CryptoResponse, error)
	grpc.ClientStream
}

type cryptoServiceGetCryptoStreamByIdClient struct {
	grpc.ClientStream
}

func (x *cryptoServiceGetCryptoStreamByIdClient) Recv() (*CryptoResponse, error) {
	m := new(CryptoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	GetCryptoStreamById(*CryptoIdRequest, CryptoService_GetCryptoStreamByIdServer) error
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
func (*UnimplementedCryptoServiceServer) GetCryptoStreamById(req *CryptoIdRequest, srv CryptoService_GetCryptoStreamByIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCryptoStreamById not implemented")
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

func _CryptoService_GetCryptoStreamById_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CryptoIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CryptoServiceServer).GetCryptoStreamById(m, &cryptoServiceGetCryptoStreamByIdServer{stream})
}

type CryptoService_GetCryptoStreamByIdServer interface {
	Send(*CryptoResponse) error
	grpc.ServerStream
}

type cryptoServiceGetCryptoStreamByIdServer struct {
	grpc.ServerStream
}

func (x *cryptoServiceGetCryptoStreamByIdServer) Send(m *CryptoResponse) error {
	return x.ServerStream.SendMsg(m)
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCryptoStreamById",
			Handler:       _CryptoService_GetCryptoStreamById_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crypto/crypto.proto",
}

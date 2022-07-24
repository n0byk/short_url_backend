// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service_proto

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type AddURLHandlerRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddURLHandlerRequest) Reset()         { *m = AddURLHandlerRequest{} }
func (m *AddURLHandlerRequest) String() string { return proto.CompactTextString(m) }
func (*AddURLHandlerRequest) ProtoMessage()    {}
func (*AddURLHandlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *AddURLHandlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddURLHandlerRequest.Unmarshal(m, b)
}
func (m *AddURLHandlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddURLHandlerRequest.Marshal(b, m, deterministic)
}
func (m *AddURLHandlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddURLHandlerRequest.Merge(m, src)
}
func (m *AddURLHandlerRequest) XXX_Size() int {
	return xxx_messageInfo_AddURLHandlerRequest.Size(m)
}
func (m *AddURLHandlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddURLHandlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddURLHandlerRequest proto.InternalMessageInfo

func (m *AddURLHandlerRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddURLHandlerRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AddURLHandlerResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddURLHandlerResponse) Reset()         { *m = AddURLHandlerResponse{} }
func (m *AddURLHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*AddURLHandlerResponse) ProtoMessage()    {}
func (*AddURLHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *AddURLHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddURLHandlerResponse.Unmarshal(m, b)
}
func (m *AddURLHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddURLHandlerResponse.Marshal(b, m, deterministic)
}
func (m *AddURLHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddURLHandlerResponse.Merge(m, src)
}
func (m *AddURLHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_AddURLHandlerResponse.Size(m)
}
func (m *AddURLHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddURLHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddURLHandlerResponse proto.InternalMessageInfo

func (m *AddURLHandlerResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetURLRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetURLRequest) Reset()         { *m = GetURLRequest{} }
func (m *GetURLRequest) String() string { return proto.CompactTextString(m) }
func (*GetURLRequest) ProtoMessage()    {}
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *GetURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetURLRequest.Unmarshal(m, b)
}
func (m *GetURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetURLRequest.Marshal(b, m, deterministic)
}
func (m *GetURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetURLRequest.Merge(m, src)
}
func (m *GetURLRequest) XXX_Size() int {
	return xxx_messageInfo_GetURLRequest.Size(m)
}
func (m *GetURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetURLRequest proto.InternalMessageInfo

func (m *GetURLRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetURLRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetURLResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetURLResponse) Reset()         { *m = GetURLResponse{} }
func (m *GetURLResponse) String() string { return proto.CompactTextString(m) }
func (*GetURLResponse) ProtoMessage()    {}
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *GetURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetURLResponse.Unmarshal(m, b)
}
func (m *GetURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetURLResponse.Marshal(b, m, deterministic)
}
func (m *GetURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetURLResponse.Merge(m, src)
}
func (m *GetURLResponse) XXX_Size() int {
	return xxx_messageInfo_GetURLResponse.Size(m)
}
func (m *GetURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetURLResponse proto.InternalMessageInfo

func (m *GetURLResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BulkDeleteRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDeleteRequest) Reset()         { *m = BulkDeleteRequest{} }
func (m *BulkDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*BulkDeleteRequest) ProtoMessage()    {}
func (*BulkDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *BulkDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDeleteRequest.Unmarshal(m, b)
}
func (m *BulkDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDeleteRequest.Marshal(b, m, deterministic)
}
func (m *BulkDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDeleteRequest.Merge(m, src)
}
func (m *BulkDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_BulkDeleteRequest.Size(m)
}
func (m *BulkDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDeleteRequest proto.InternalMessageInfo

func (m *BulkDeleteRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *BulkDeleteRequest) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type BulkAddURLRequest struct {
	UserID               string                            `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OriginalUrls         []*BulkAddURLRequest_BulkARequest `protobuf:"bytes,2,rep,name=original_urls,json=originalUrls,proto3" json:"original_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BulkAddURLRequest) Reset()         { *m = BulkAddURLRequest{} }
func (m *BulkAddURLRequest) String() string { return proto.CompactTextString(m) }
func (*BulkAddURLRequest) ProtoMessage()    {}
func (*BulkAddURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *BulkAddURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddURLRequest.Unmarshal(m, b)
}
func (m *BulkAddURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddURLRequest.Marshal(b, m, deterministic)
}
func (m *BulkAddURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddURLRequest.Merge(m, src)
}
func (m *BulkAddURLRequest) XXX_Size() int {
	return xxx_messageInfo_BulkAddURLRequest.Size(m)
}
func (m *BulkAddURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddURLRequest proto.InternalMessageInfo

func (m *BulkAddURLRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *BulkAddURLRequest) GetOriginalUrls() []*BulkAddURLRequest_BulkARequest {
	if m != nil {
		return m.OriginalUrls
	}
	return nil
}

type BulkAddURLRequest_BulkARequest struct {
	CorrelationId        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl             string   `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkAddURLRequest_BulkARequest) Reset()         { *m = BulkAddURLRequest_BulkARequest{} }
func (m *BulkAddURLRequest_BulkARequest) String() string { return proto.CompactTextString(m) }
func (*BulkAddURLRequest_BulkARequest) ProtoMessage()    {}
func (*BulkAddURLRequest_BulkARequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6, 0}
}

func (m *BulkAddURLRequest_BulkARequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddURLRequest_BulkARequest.Unmarshal(m, b)
}
func (m *BulkAddURLRequest_BulkARequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddURLRequest_BulkARequest.Marshal(b, m, deterministic)
}
func (m *BulkAddURLRequest_BulkARequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddURLRequest_BulkARequest.Merge(m, src)
}
func (m *BulkAddURLRequest_BulkARequest) XXX_Size() int {
	return xxx_messageInfo_BulkAddURLRequest_BulkARequest.Size(m)
}
func (m *BulkAddURLRequest_BulkARequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddURLRequest_BulkARequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddURLRequest_BulkARequest proto.InternalMessageInfo

func (m *BulkAddURLRequest_BulkARequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *BulkAddURLRequest_BulkARequest) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type BulkAddURLResponse struct {
	ShortUrls            []*BulkAddURLResponse_BulkAResponse `protobuf:"bytes,2,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BulkAddURLResponse) Reset()         { *m = BulkAddURLResponse{} }
func (m *BulkAddURLResponse) String() string { return proto.CompactTextString(m) }
func (*BulkAddURLResponse) ProtoMessage()    {}
func (*BulkAddURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *BulkAddURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddURLResponse.Unmarshal(m, b)
}
func (m *BulkAddURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddURLResponse.Marshal(b, m, deterministic)
}
func (m *BulkAddURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddURLResponse.Merge(m, src)
}
func (m *BulkAddURLResponse) XXX_Size() int {
	return xxx_messageInfo_BulkAddURLResponse.Size(m)
}
func (m *BulkAddURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddURLResponse proto.InternalMessageInfo

func (m *BulkAddURLResponse) GetShortUrls() []*BulkAddURLResponse_BulkAResponse {
	if m != nil {
		return m.ShortUrls
	}
	return nil
}

type BulkAddURLResponse_BulkAResponse struct {
	CorrelationId        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl             string   `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkAddURLResponse_BulkAResponse) Reset()         { *m = BulkAddURLResponse_BulkAResponse{} }
func (m *BulkAddURLResponse_BulkAResponse) String() string { return proto.CompactTextString(m) }
func (*BulkAddURLResponse_BulkAResponse) ProtoMessage()    {}
func (*BulkAddURLResponse_BulkAResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 0}
}

func (m *BulkAddURLResponse_BulkAResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkAddURLResponse_BulkAResponse.Unmarshal(m, b)
}
func (m *BulkAddURLResponse_BulkAResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkAddURLResponse_BulkAResponse.Marshal(b, m, deterministic)
}
func (m *BulkAddURLResponse_BulkAResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkAddURLResponse_BulkAResponse.Merge(m, src)
}
func (m *BulkAddURLResponse_BulkAResponse) XXX_Size() int {
	return xxx_messageInfo_BulkAddURLResponse_BulkAResponse.Size(m)
}
func (m *BulkAddURLResponse_BulkAResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkAddURLResponse_BulkAResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BulkAddURLResponse_BulkAResponse proto.InternalMessageInfo

func (m *BulkAddURLResponse_BulkAResponse) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *BulkAddURLResponse_BulkAResponse) GetShortUrl() string {
	if m != nil {
		return m.ShortUrl
	}
	return ""
}

type StatsResponse struct {
	UrlsCount            int32    `protobuf:"varint,1,opt,name=urls_count,json=urlsCount,proto3" json:"urls_count,omitempty"`
	UsersCount           int32    `protobuf:"varint,2,opt,name=users_count,json=usersCount,proto3" json:"users_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *StatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsResponse.Unmarshal(m, b)
}
func (m *StatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsResponse.Marshal(b, m, deterministic)
}
func (m *StatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsResponse.Merge(m, src)
}
func (m *StatsResponse) XXX_Size() int {
	return xxx_messageInfo_StatsResponse.Size(m)
}
func (m *StatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsResponse proto.InternalMessageInfo

func (m *StatsResponse) GetUrlsCount() int32 {
	if m != nil {
		return m.UrlsCount
	}
	return 0
}

func (m *StatsResponse) GetUsersCount() int32 {
	if m != nil {
		return m.UsersCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "service.grpc.Empty")
	proto.RegisterType((*AddURLHandlerRequest)(nil), "service.grpc.AddURLHandlerRequest")
	proto.RegisterType((*AddURLHandlerResponse)(nil), "service.grpc.AddURLHandlerResponse")
	proto.RegisterType((*GetURLRequest)(nil), "service.grpc.GetURLRequest")
	proto.RegisterType((*GetURLResponse)(nil), "service.grpc.GetURLResponse")
	proto.RegisterType((*BulkDeleteRequest)(nil), "service.grpc.BulkDeleteRequest")
	proto.RegisterType((*BulkAddURLRequest)(nil), "service.grpc.BulkAddURLRequest")
	proto.RegisterType((*BulkAddURLRequest_BulkARequest)(nil), "service.grpc.BulkAddURLRequest.BulkARequest")
	proto.RegisterType((*BulkAddURLResponse)(nil), "service.grpc.BulkAddURLResponse")
	proto.RegisterType((*BulkAddURLResponse_BulkAResponse)(nil), "service.grpc.BulkAddURLResponse.BulkAResponse")
	proto.RegisterType((*StatsResponse)(nil), "service.grpc.StatsResponse")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x53, 0xc5, 0x90, 0x69, 0x5c, 0xc4, 0x52, 0xaa, 0xc8, 0x05, 0x35, 0x5a, 0x04, 0xe2,
	0x00, 0x39, 0x14, 0x6e, 0x88, 0x03, 0xa9, 0x2b, 0xa8, 0x14, 0x04, 0xdd, 0x28, 0x97, 0x5e, 0x22,
	0x63, 0xaf, 0x8c, 0xc5, 0xe2, 0x35, 0xbb, 0xeb, 0x4a, 0xfc, 0x19, 0x17, 0x7e, 0x82, 0x2f, 0x42,
	0xde, 0xb5, 0x5d, 0x6f, 0x71, 0x63, 0x09, 0xf5, 0x64, 0xcd, 0x9b, 0x99, 0x37, 0x6f, 0x66, 0x76,
	0x0c, 0x9e, 0xa4, 0xe2, 0x32, 0x8d, 0xe8, 0x3c, 0x17, 0x5c, 0x71, 0x34, 0xa9, 0xcd, 0x44, 0xe4,
	0x11, 0xbe, 0x03, 0xa3, 0xd3, 0xef, 0xb9, 0xfa, 0x89, 0x03, 0xd8, 0x7f, 0x17, 0xc7, 0x6b, 0xb2,
	0xfc, 0x10, 0x66, 0x31, 0xa3, 0x82, 0xd0, 0x1f, 0x05, 0x95, 0x0a, 0x1d, 0x80, 0x5b, 0x48, 0x2a,
	0xce, 0x82, 0xa9, 0x33, 0x73, 0x9e, 0x8f, 0x49, 0x65, 0xa1, 0x7d, 0x18, 0x5d, 0x86, 0xac, 0xa0,
	0xd3, 0xa1, 0x86, 0x8d, 0x81, 0x5f, 0xc2, 0xc3, 0x6b, 0x2c, 0x32, 0xe7, 0x99, 0xa4, 0x57, 0xe1,
	0x4e, 0x3b, 0xfc, 0x2d, 0x78, 0xef, 0xa9, 0x5a, 0x93, 0xe5, 0xff, 0x55, 0x7b, 0x06, 0x7b, 0x75,
	0xfa, 0xd6, 0x32, 0x27, 0x70, 0x7f, 0x51, 0xb0, 0x6f, 0x01, 0x65, 0x54, 0xd1, 0xbe, 0x52, 0x07,
	0xe0, 0xea, 0x2c, 0x39, 0x1d, 0xce, 0x76, 0x4a, 0xdc, 0x58, 0xf8, 0x8f, 0x63, 0x58, 0x4c, 0x7f,
	0x7d, 0x2c, 0xe7, 0xe0, 0x71, 0x91, 0x26, 0x69, 0x16, 0xb2, 0x4d, 0x21, 0x98, 0x21, 0xdb, 0x3d,
	0x7e, 0x31, 0x6f, 0x4f, 0x7f, 0xfe, 0x0f, 0x9f, 0x41, 0x2a, 0x83, 0x4c, 0x6a, 0x8a, 0xb5, 0x60,
	0xd2, 0x27, 0x30, 0x69, 0x7b, 0xd1, 0x53, 0xd8, 0x8b, 0xb8, 0x10, 0x94, 0x85, 0x2a, 0xe5, 0xd9,
	0x26, 0x8d, 0x2b, 0x09, 0x5e, 0x0b, 0x3d, 0x8b, 0xd1, 0x21, 0x8c, 0xe5, 0x57, 0x2e, 0x54, 0x29,
	0xa3, 0x1a, 0xdf, 0x5d, 0x0d, 0xac, 0x05, 0xc3, 0xbf, 0x1c, 0x40, 0x6d, 0x11, 0xd5, 0x18, 0x3f,
	0x02, 0x34, 0x39, 0xb5, 0xf4, 0xf9, 0xcd, 0xd2, 0x4d, 0x56, 0xad, 0xdd, 0x58, 0x64, 0x5c, 0x17,
	0x91, 0xfe, 0x0a, 0x3c, 0xcb, 0x77, 0x2b, 0xd2, 0x3f, 0x81, 0xb7, 0x52, 0xa1, 0x92, 0x0d, 0xe9,
	0x63, 0x80, 0x52, 0xee, 0x26, 0xe2, 0x45, 0xa6, 0x34, 0xe1, 0x88, 0x8c, 0x4b, 0xe4, 0xa4, 0x04,
	0xd0, 0x11, 0xec, 0x96, 0xbb, 0xa9, 0xfd, 0x43, 0xed, 0x07, 0x0d, 0xe9, 0x80, 0xe3, 0xdf, 0x3b,
	0x30, 0x59, 0x99, 0x16, 0x97, 0x3c, 0x49, 0x23, 0x74, 0x0e, 0x70, 0xd5, 0x25, 0x3a, 0xea, 0x59,
	0x9d, 0x3f, 0xeb, 0x1b, 0x10, 0x1e, 0xa0, 0xc0, 0x50, 0x9a, 0x97, 0xd8, 0x45, 0x69, 0xbd, 0x51,
	0xff, 0x81, 0x1d, 0x60, 0x2e, 0x75, 0x80, 0x5e, 0x83, 0x1b, 0x2c, 0x3e, 0xa7, 0x59, 0x82, 0xba,
	0x02, 0x6e, 0xca, 0xba, 0x00, 0xcf, 0xba, 0x4d, 0x84, 0xed, 0xb8, 0xae, 0xf3, 0xf7, 0x9f, 0x6c,
	0x8d, 0x69, 0xfa, 0x3a, 0x05, 0xd7, 0x5c, 0x22, 0x3a, 0xb4, 0x13, 0xac, 0xf3, 0xf6, 0x1f, 0x75,
	0x3b, 0x1b, 0x9a, 0x37, 0x30, 0xd2, 0x3b, 0xed, 0xee, 0xeb, 0x1a, 0xb5, 0xb5, 0x7d, 0x3c, 0x58,
	0xdc, 0xbb, 0xb0, 0xff, 0x74, 0x5f, 0x5c, 0xfd, 0x79, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x94,
	0xd7, 0x9a, 0x9f, 0x01, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceLogicClient is the client API for ServiceLogic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceLogicClient interface {
	BulkAddURL(ctx context.Context, in *BulkAddURLRequest, opts ...grpc.CallOption) (*BulkAddURLResponse, error)
	BulkDelete(ctx context.Context, in *BulkDeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	DBPing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	AddURLHandler(ctx context.Context, in *AddURLHandlerRequest, opts ...grpc.CallOption) (*AddURLHandlerResponse, error)
	GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error)
	Stats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatsResponse, error)
}

type serviceLogicClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceLogicClient(cc grpc.ClientConnInterface) ServiceLogicClient {
	return &serviceLogicClient{cc}
}

func (c *serviceLogicClient) BulkAddURL(ctx context.Context, in *BulkAddURLRequest, opts ...grpc.CallOption) (*BulkAddURLResponse, error) {
	out := new(BulkAddURLResponse)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/BulkAddURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceLogicClient) BulkDelete(ctx context.Context, in *BulkDeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/BulkDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceLogicClient) DBPing(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/DBPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceLogicClient) AddURLHandler(ctx context.Context, in *AddURLHandlerRequest, opts ...grpc.CallOption) (*AddURLHandlerResponse, error) {
	out := new(AddURLHandlerResponse)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/AddURLHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceLogicClient) GetURL(ctx context.Context, in *GetURLRequest, opts ...grpc.CallOption) (*GetURLResponse, error) {
	out := new(GetURLResponse)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/GetURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceLogicClient) Stats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/service.grpc.ServiceLogic/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceLogicServer is the server API for ServiceLogic service.
type ServiceLogicServer interface {
	BulkAddURL(context.Context, *BulkAddURLRequest) (*BulkAddURLResponse, error)
	BulkDelete(context.Context, *BulkDeleteRequest) (*Empty, error)
	DBPing(context.Context, *Empty) (*Empty, error)
	AddURLHandler(context.Context, *AddURLHandlerRequest) (*AddURLHandlerResponse, error)
	GetURL(context.Context, *GetURLRequest) (*GetURLResponse, error)
	Stats(context.Context, *Empty) (*StatsResponse, error)
}

// UnimplementedServiceLogicServer can be embedded to have forward compatible implementations.
type UnimplementedServiceLogicServer struct {
}

func (*UnimplementedServiceLogicServer) BulkAddURL(ctx context.Context, req *BulkAddURLRequest) (*BulkAddURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkAddURL not implemented")
}
func (*UnimplementedServiceLogicServer) BulkDelete(ctx context.Context, req *BulkDeleteRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDelete not implemented")
}
func (*UnimplementedServiceLogicServer) DBPing(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DBPing not implemented")
}
func (*UnimplementedServiceLogicServer) AddURLHandler(ctx context.Context, req *AddURLHandlerRequest) (*AddURLHandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddURLHandler not implemented")
}
func (*UnimplementedServiceLogicServer) GetURL(ctx context.Context, req *GetURLRequest) (*GetURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURL not implemented")
}
func (*UnimplementedServiceLogicServer) Stats(ctx context.Context, req *Empty) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}

func RegisterServiceLogicServer(s *grpc.Server, srv ServiceLogicServer) {
	s.RegisterService(&_ServiceLogic_serviceDesc, srv)
}

func _ServiceLogic_BulkAddURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkAddURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).BulkAddURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/BulkAddURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).BulkAddURL(ctx, req.(*BulkAddURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceLogic_BulkDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).BulkDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/BulkDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).BulkDelete(ctx, req.(*BulkDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceLogic_DBPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).DBPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/DBPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).DBPing(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceLogic_AddURLHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddURLHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).AddURLHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/AddURLHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).AddURLHandler(ctx, req.(*AddURLHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceLogic_GetURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).GetURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/GetURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).GetURL(ctx, req.(*GetURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceLogic_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceLogicServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.grpc.ServiceLogic/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceLogicServer).Stats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceLogic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.grpc.ServiceLogic",
	HandlerType: (*ServiceLogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BulkAddURL",
			Handler:    _ServiceLogic_BulkAddURL_Handler,
		},
		{
			MethodName: "BulkDelete",
			Handler:    _ServiceLogic_BulkDelete_Handler,
		},
		{
			MethodName: "DBPing",
			Handler:    _ServiceLogic_DBPing_Handler,
		},
		{
			MethodName: "AddURLHandler",
			Handler:    _ServiceLogic_AddURLHandler_Handler,
		},
		{
			MethodName: "GetURL",
			Handler:    _ServiceLogic_GetURL_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _ServiceLogic_Stats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

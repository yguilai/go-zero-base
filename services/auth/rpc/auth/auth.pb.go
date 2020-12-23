// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: auth.proto

package auth

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=AccessExpire,proto3" json:"AccessExpire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=RefreshAfter,proto3" json:"RefreshAfter,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *Token) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Claim) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Claim) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// jwt创建
type JwtCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claims []*Claim `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *JwtCreateReq) Reset() {
	*x = JwtCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtCreateReq) ProtoMessage() {}

func (x *JwtCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtCreateReq.ProtoReflect.Descriptor instead.
func (*JwtCreateReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *JwtCreateReq) GetClaims() []*Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

type JwtCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *JwtCreateResp) Reset() {
	*x = JwtCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtCreateResp) ProtoMessage() {}

func (x *JwtCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtCreateResp.ProtoReflect.Descriptor instead.
func (*JwtCreateResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *JwtCreateResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

// jwt验证
type JwtVerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *JwtVerifyReq) Reset() {
	*x = JwtVerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtVerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtVerifyReq) ProtoMessage() {}

func (x *JwtVerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtVerifyReq.ProtoReflect.Descriptor instead.
func (*JwtVerifyReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *JwtVerifyReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type JwtVerifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *JwtVerifyResp) Reset() {
	*x = JwtVerifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtVerifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtVerifyResp) ProtoMessage() {}

func (x *JwtVerifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtVerifyResp.ProtoReflect.Descriptor instead.
func (*JwtVerifyResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *JwtVerifyResp) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

// jwt刷新
type JwtRefreshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *JwtRefreshReq) Reset() {
	*x = JwtRefreshReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtRefreshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtRefreshReq) ProtoMessage() {}

func (x *JwtRefreshReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtRefreshReq.ProtoReflect.Descriptor instead.
func (*JwtRefreshReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{8}
}

func (x *JwtRefreshReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type JwtRefreshResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *JwtRefreshResp) Reset() {
	*x = JwtRefreshResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtRefreshResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtRefreshResp) ProtoMessage() {}

func (x *JwtRefreshResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtRefreshResp.ProtoReflect.Descriptor instead.
func (*JwtRefreshResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{9}
}

func (x *JwtRefreshResp) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e,
	0x67, 0x22, 0x71, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x33, 0x0a, 0x0c, 0x4a, 0x77, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x0d, 0x4a, 0x77,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24,
	0x0a, 0x0c, 0x4a, 0x77, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x4a, 0x77, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4a,
	0x77, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x33, 0x0a, 0x0e, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc9, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x25, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x77, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x77, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x06, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x77, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4a, 0x77, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a,
	0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4a, 0x77, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_auth_proto_goTypes = []interface{}{
	(*Request)(nil),        // 0: auth.Request
	(*Response)(nil),       // 1: auth.Response
	(*Token)(nil),          // 2: auth.Token
	(*Claim)(nil),          // 3: auth.Claim
	(*JwtCreateReq)(nil),   // 4: auth.JwtCreateReq
	(*JwtCreateResp)(nil),  // 5: auth.JwtCreateResp
	(*JwtVerifyReq)(nil),   // 6: auth.JwtVerifyReq
	(*JwtVerifyResp)(nil),  // 7: auth.JwtVerifyResp
	(*JwtRefreshReq)(nil),  // 8: auth.JwtRefreshReq
	(*JwtRefreshResp)(nil), // 9: auth.JwtRefreshResp
}
var file_auth_proto_depIdxs = []int32{
	3, // 0: auth.JwtCreateReq.claims:type_name -> auth.Claim
	2, // 1: auth.JwtCreateResp.token:type_name -> auth.Token
	2, // 2: auth.JwtRefreshResp.token:type_name -> auth.Token
	0, // 3: auth.Auth.Ping:input_type -> auth.Request
	4, // 4: auth.Auth.Create:input_type -> auth.JwtCreateReq
	6, // 5: auth.Auth.Verify:input_type -> auth.JwtVerifyReq
	8, // 6: auth.Auth.Refresh:input_type -> auth.JwtRefreshReq
	1, // 7: auth.Auth.Ping:output_type -> auth.Response
	5, // 8: auth.Auth.Create:output_type -> auth.JwtCreateResp
	7, // 9: auth.Auth.Verify:output_type -> auth.JwtVerifyResp
	9, // 10: auth.Auth.Refresh:output_type -> auth.JwtRefreshResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claim); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtCreateResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtVerifyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtVerifyResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtRefreshReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtRefreshResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *JwtCreateReq, opts ...grpc.CallOption) (*JwtCreateResp, error)
	Verify(ctx context.Context, in *JwtVerifyReq, opts ...grpc.CallOption) (*JwtVerifyResp, error)
	Refresh(ctx context.Context, in *JwtRefreshReq, opts ...grpc.CallOption) (*JwtRefreshResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/auth.Auth/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Create(ctx context.Context, in *JwtCreateReq, opts ...grpc.CallOption) (*JwtCreateResp, error) {
	out := new(JwtCreateResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Verify(ctx context.Context, in *JwtVerifyReq, opts ...grpc.CallOption) (*JwtVerifyResp, error) {
	out := new(JwtVerifyResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *JwtRefreshReq, opts ...grpc.CallOption) (*JwtRefreshResp, error) {
	out := new(JwtRefreshResp)
	err := c.cc.Invoke(ctx, "/auth.Auth/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Ping(context.Context, *Request) (*Response, error)
	Create(context.Context, *JwtCreateReq) (*JwtCreateResp, error)
	Verify(context.Context, *JwtVerifyReq) (*JwtVerifyResp, error)
	Refresh(context.Context, *JwtRefreshReq) (*JwtRefreshResp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAuthServer) Create(context.Context, *JwtCreateReq) (*JwtCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAuthServer) Verify(context.Context, *JwtVerifyReq) (*JwtVerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (*UnimplementedAuthServer) Refresh(context.Context, *JwtRefreshReq) (*JwtRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Create(ctx, req.(*JwtCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Verify(ctx, req.(*JwtVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwtRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*JwtRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Auth_Ping_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Auth_Create_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Auth_Verify_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

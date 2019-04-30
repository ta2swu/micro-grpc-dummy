// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bidirectionaldummy.proto

package bidirectionaldummy

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Reqmsg               string   `protobuf:"bytes,1,opt,name=Reqmsg,proto3" json:"Reqmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa987b894e3836a, []int{0}
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

func (m *Request) GetReqmsg() string {
	if m != nil {
		return m.Reqmsg
	}
	return ""
}

type Response struct {
	Resmsg               string   `protobuf:"bytes,1,opt,name=Resmsg,proto3" json:"Resmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fa987b894e3836a, []int{1}
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

func (m *Response) GetResmsg() string {
	if m != nil {
		return m.Resmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("bidirectionaldummy.proto", fileDescriptor_4fa987b894e3836a) }

var fileDescriptor_4fa987b894e3836a = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xca, 0x4c, 0xc9,
	0x2c, 0x4a, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x4b, 0xcc, 0x49, 0x29, 0xcd, 0xcd, 0xad, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe3, 0x62, 0x0b, 0x4a, 0x2d, 0xcc, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0xf2, 0x94, 0x94, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x21, 0x6a,
	0x8a, 0x51, 0xd4, 0x80, 0x78, 0x46, 0x46, 0x5c, 0xdc, 0x2e, 0x20, 0x53, 0x83, 0x4b, 0x8a, 0x52,
	0x13, 0x73, 0x85, 0x94, 0xb9, 0xd8, 0xa0, 0x2c, 0x0e, 0x3d, 0xa8, 0xf1, 0x52, 0x9c, 0x7a, 0x30,
	0x53, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x2e, 0x30, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x52, 0xeb, 0x29, 0x9d, 0x00, 0x00, 0x00,
}
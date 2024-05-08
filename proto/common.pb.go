// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package proto

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

type Pager struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalRows            int64    `protobuf:"varint,3,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pager) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetTotalRows() int64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Pager)(nil), "proto.Pager")
	proto.RegisterType((*Error)(nil), "proto.Error")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xbd, 0x0a, 0xc2, 0x30,
	0x10, 0xc7, 0xa9, 0x35, 0x6a, 0x6e, 0xbc, 0x29, 0x20, 0x82, 0x74, 0x72, 0xd2, 0x41, 0x7c, 0x04,
	0x77, 0x89, 0x83, 0x63, 0x89, 0x35, 0x94, 0x82, 0xf1, 0xca, 0x5d, 0xa0, 0xd0, 0xa7, 0x97, 0x9c,
	0x38, 0xfd, 0xbf, 0xe0, 0x77, 0x07, 0x38, 0x32, 0x65, 0x3a, 0x75, 0x94, 0x12, 0x7d, 0x8e, 0x1a,
	0xd0, 0xa8, 0x34, 0x0f, 0x30, 0xb7, 0xd0, 0x47, 0x46, 0x84, 0xe5, 0x18, 0xfa, 0xe8, 0xaa, 0x7d,
	0x75, 0xa8, 0xbd, 0x7a, 0xdc, 0x82, 0x2d, 0xda, 0xca, 0x30, 0x47, 0xb7, 0xd0, 0x61, 0x53, 0x8a,
	0xfb, 0x30, 0x47, 0xdc, 0x01, 0x64, 0xca, 0xe1, 0xdd, 0x32, 0x4d, 0xe2, 0x6a, 0x5d, 0xad, 0x36,
	0x9e, 0x26, 0x69, 0x2e, 0x60, 0xae, 0xcc, 0xa4, 0xe0, 0x8e, 0x5e, 0x3f, 0xb0, 0xf1, 0xea, 0xd1,
	0xc1, 0x3a, 0x45, 0x91, 0x72, 0xaf, 0x60, 0xad, 0xff, 0xc7, 0xe7, 0x4a, 0xdf, 0x3a, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xa7, 0x02, 0x5b, 0x57, 0xb3, 0x00, 0x00, 0x00,
}

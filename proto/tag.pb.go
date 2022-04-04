// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tag.proto

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

type GetTagListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListRequest) Reset()         { *m = GetTagListRequest{} }
func (m *GetTagListRequest) String() string { return proto.CompactTextString(m) }
func (*GetTagListRequest) ProtoMessage()    {}
func (*GetTagListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{0}
}

func (m *GetTagListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListRequest.Unmarshal(m, b)
}
func (m *GetTagListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListRequest.Marshal(b, m, deterministic)
}
func (m *GetTagListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListRequest.Merge(m, src)
}
func (m *GetTagListRequest) XXX_Size() int {
	return xxx_messageInfo_GetTagListRequest.Size(m)
}
func (m *GetTagListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListRequest proto.InternalMessageInfo

func (m *GetTagListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTagListRequest) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type Tag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type GetTagListReply struct {
	List                 []*Tag   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListReply) Reset()         { *m = GetTagListReply{} }
func (m *GetTagListReply) String() string { return proto.CompactTextString(m) }
func (*GetTagListReply) ProtoMessage()    {}
func (*GetTagListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{2}
}

func (m *GetTagListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListReply.Unmarshal(m, b)
}
func (m *GetTagListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListReply.Marshal(b, m, deterministic)
}
func (m *GetTagListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListReply.Merge(m, src)
}
func (m *GetTagListReply) XXX_Size() int {
	return xxx_messageInfo_GetTagListReply.Size(m)
}
func (m *GetTagListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListReply proto.InternalMessageInfo

func (m *GetTagListReply) GetList() []*Tag {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetTagListReply) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTagListRequest)(nil), "proto.GetTagListRequest")
	proto.RegisterType((*Tag)(nil), "proto.Tag")
	proto.RegisterType((*GetTagListReply)(nil), "proto.GetTagListReply")
}

func init() { proto.RegisterFile("proto/tag.proto", fileDescriptor_48dc6f15189f1be6) }

var fileDescriptor_48dc6f15189f1be6 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0x4d, 0xd2, 0x88, 0x4e, 0xd5, 0xe2, 0x20, 0x12, 0x7a, 0x90, 0x90, 0x53, 0x4e, 0x29,
	0xc4, 0xa3, 0x88, 0xe2, 0xc5, 0x8b, 0x88, 0xac, 0xf1, 0xe2, 0x6d, 0x6c, 0x87, 0x65, 0x21, 0x69,
	0x62, 0x32, 0x0a, 0xfd, 0xf7, 0xe2, 0x6c, 0xd1, 0x42, 0x73, 0x9a, 0xd9, 0xef, 0xcd, 0xbe, 0xd9,
	0x85, 0x59, 0xd7, 0xb7, 0xd2, 0x2e, 0x84, 0x6c, 0xa1, 0x1d, 0xc6, 0x5a, 0xe6, 0xe8, 0xf3, 0x65,
	0xdb, 0x34, 0xed, 0xda, 0xa3, 0xec, 0x16, 0xce, 0x1f, 0x59, 0x2a, 0xb2, 0x4f, 0x6e, 0x10, 0xc3,
	0x9f, 0x5f, 0x3c, 0x08, 0x22, 0x4c, 0xd6, 0xd4, 0x70, 0x12, 0xa4, 0x41, 0x7e, 0x6c, 0xb4, 0xc7,
	0x0b, 0x88, 0x07, 0x21, 0xe1, 0x24, 0x4c, 0x83, 0xfc, 0xd4, 0xf8, 0x43, 0x76, 0x07, 0x51, 0x45,
	0x16, 0xcf, 0x20, 0x74, 0x2b, 0x1d, 0x8f, 0x4c, 0xe8, 0x56, 0x7f, 0x82, 0x70, 0x4c, 0x10, 0xed,
	0x0a, 0xde, 0x60, 0xb6, 0xbb, 0xbf, 0xab, 0x37, 0x78, 0x05, 0x93, 0xda, 0x0d, 0x92, 0x04, 0x69,
	0x94, 0x4f, 0x4b, 0xf0, 0x0f, 0x2d, 0x2a, 0xb2, 0x46, 0x73, 0xcc, 0x20, 0xee, 0xc8, 0x72, 0xaf,
	0xf6, 0x69, 0x79, 0xb2, 0x1d, 0x78, 0xf9, 0xcd, 0x8c, 0x47, 0xe5, 0x33, 0x40, 0x45, 0xf6, 0x95,
	0xfb, 0x6f, 0xb7, 0x64, 0xbc, 0x07, 0xf8, 0x5f, 0x82, 0xc9, 0xf6, 0xc2, 0xde, 0xbf, 0xe7, 0x97,
	0x23, 0xa4, 0xab, 0x37, 0xd9, 0xc1, 0x03, 0xbc, 0x1f, 0x15, 0x8b, 0x1b, 0xa5, 0x1f, 0x87, 0x5a,
	0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xc5, 0x53, 0x85, 0x67, 0x01, 0x00, 0x00,
}
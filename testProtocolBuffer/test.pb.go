// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package protocolBuffers

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

type User struct {
	Firstname            *string  `protobuf:"bytes,1,req,name=firstname" json:"firstname,omitempty"`
	Lastname             *string  `protobuf:"bytes,2,req,name=lastname" json:"lastname,omitempty"`
	Id                   *int32   `protobuf:"varint,3,req,name=id" json:"id,omitempty"`
	City                 *string  `protobuf:"bytes,4,opt,name=city" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirstname() string {
	if m != nil && m.Firstname != nil {
		return *m.Firstname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil && m.Lastname != nil {
		return *m.Lastname
	}
	return ""
}

func (m *User) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "protocolBuffers.User")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x07, 0x53, 0xc9, 0xf9, 0x39, 0x4e, 0xa5, 0x69,
	0x69, 0xa9, 0x45, 0xc5, 0x4a, 0x29, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x32, 0x5c, 0x9c,
	0x69, 0x99, 0x45, 0xc5, 0x25, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41,
	0x08, 0x01, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0x44, 0xa8, 0x24, 0x13, 0x58, 0x12, 0xce, 0x17, 0xe2,
	0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x56, 0x60, 0xd2, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x12,
	0xe2, 0x62, 0x49, 0xce, 0x2c, 0xa9, 0x94, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xdd, 0x71, 0xe9, 0x83, 0x00, 0x00, 0x00,
}

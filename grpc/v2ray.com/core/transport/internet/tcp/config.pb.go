// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2ray.com/core/transport/internet/tcp/config.proto

package tcp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	serial "v2ray.com/core/common/serial"
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

type Config struct {
	HeaderSettings       *serial.TypedMessage `protobuf:"bytes,2,opt,name=header_settings,json=headerSettings,proto3" json:"header_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6d289fc61edd40, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetHeaderSettings() *serial.TypedMessage {
	if m != nil {
		return m.HeaderSettings
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tcp.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/tcp/config.proto", fileDescriptor_eb6d289fc61edd40)
}

var fileDescriptor_eb6d289fc61edd40 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0x95, 0x45, 0x2a, 0xa8, 0xec, 0x49, 0x3c, 0xb9, 0x82, 0xe2, 0x69, 0x22, 0xf1,
	0x0d, 0xdc, 0x93, 0x82, 0x28, 0xb5, 0x78, 0xf0, 0x52, 0xe2, 0xec, 0x58, 0x03, 0x26, 0x13, 0x26,
	0x83, 0xd0, 0x57, 0xf2, 0x29, 0x65, 0x37, 0x76, 0x11, 0x2f, 0xde, 0xbf, 0xef, 0xfb, 0xff, 0xc6,
	0x7e, 0x5a, 0x71, 0x23, 0x20, 0x07, 0x83, 0x2c, 0x64, 0x54, 0x5c, 0xcc, 0x89, 0x45, 0x8d, 0x8f,
	0x4a, 0x12, 0x49, 0x8d, 0x62, 0x32, 0xc8, 0xf1, 0xcd, 0x0f, 0x90, 0x84, 0x95, 0xe7, 0x8b, 0xc9,
	0x11, 0x82, 0x2d, 0x0f, 0x13, 0x0f, 0x8a, 0xe9, 0xe4, 0xea, 0x4f, 0x16, 0x39, 0x04, 0x8e, 0x26,
	0x93, 0x78, 0xf7, 0x61, 0x74, 0x4c, 0xb4, 0xea, 0x03, 0xe5, 0xec, 0x06, 0x2a, 0xd1, 0xb3, 0xbe,
	0x99, 0x2d, 0x37, 0x23, 0xf3, 0x87, 0xe6, 0xf0, 0x9d, 0xdc, 0x8a, 0xa4, 0xcf, 0xa4, 0xea, 0xe3,
	0x90, 0x8f, 0xeb, 0xd3, 0xea, 0x72, 0xdf, 0x5e, 0xc0, 0xaf, 0xe1, 0x52, 0x84, 0x52, 0x84, 0x6e,
	0x5d, 0xbc, 0x2f, 0xc1, 0xf6, 0xa0, 0xe8, 0x4f, 0x3f, 0xf6, 0xdd, 0xee, 0x5e, 0x75, 0x54, 0xdf,
	0xb4, 0xcd, 0x39, 0x72, 0x80, 0x7f, 0xbf, 0x3f, 0x56, 0x2f, 0x3b, 0x8a, 0xe9, 0xab, 0x5e, 0x3c,
	0xdb, 0xd6, 0x8d, 0xb0, 0x5c, 0xa3, 0xdd, 0x16, 0xbd, 0x9d, 0xd0, 0x0e, 0xd3, 0xeb, 0x6c, 0xf3,
	0xfd, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf9, 0x1f, 0xa0, 0x46, 0x01, 0x00, 0x00,
}
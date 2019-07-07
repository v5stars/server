// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2ray.com/core/transport/internet/headers/tls/config.proto

package tls

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

type PacketConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketConfig) Reset()         { *m = PacketConfig{} }
func (m *PacketConfig) String() string { return proto.CompactTextString(m) }
func (*PacketConfig) ProtoMessage()    {}
func (*PacketConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55187b75c46dc0d, []int{0}
}

func (m *PacketConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketConfig.Unmarshal(m, b)
}
func (m *PacketConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketConfig.Marshal(b, m, deterministic)
}
func (m *PacketConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketConfig.Merge(m, src)
}
func (m *PacketConfig) XXX_Size() int {
	return xxx_messageInfo_PacketConfig.Size(m)
}
func (m *PacketConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PacketConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PacketConfig)(nil), "v2ray.core.transport.internet.headers.tls.PacketConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/tls/config.proto", fileDescriptor_e55187b75c46dc0d)
}

var fileDescriptor_e55187b75c46dc0d = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2a, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x48,
	0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x2f, 0xc9, 0x29, 0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x84, 0xe9, 0x2d, 0x4a, 0xd5, 0x83, 0xeb, 0xd3,
	0x83, 0xe9, 0xd3, 0x83, 0xea, 0xd3, 0x2b, 0xc9, 0x29, 0x56, 0xe2, 0xe3, 0xe2, 0x09, 0x48, 0x4c,
	0xce, 0x4e, 0x2d, 0x71, 0x06, 0x1b, 0xe0, 0x94, 0xc4, 0xa5, 0x9b, 0x9c, 0x9f, 0xab, 0x47, 0xb4,
	0x01, 0x01, 0x8c, 0x51, 0xcc, 0x25, 0x39, 0xc5, 0xab, 0x98, 0x34, 0xc3, 0x8c, 0x82, 0x12, 0x2b,
	0xf5, 0x9c, 0x41, 0x5a, 0x42, 0xe0, 0x5a, 0x3c, 0x61, 0x5a, 0x3c, 0xa0, 0x5a, 0x42, 0x72, 0x8a,
	0x93, 0xd8, 0xc0, 0xae, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x21, 0xb5, 0x48, 0xe3,
	0x00, 0x00, 0x00,
}

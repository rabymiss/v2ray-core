// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/vmess/inbound/config.proto
// DO NOT EDIT!

/*
Package inbound is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/vmess/inbound/config.proto

It has these top-level messages:
	DetourConfig
	DefaultConfig
	Config
*/
package inbound

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DetourConfig struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *DetourConfig) Reset()                    { *m = DetourConfig{} }
func (m *DetourConfig) String() string            { return proto.CompactTextString(m) }
func (*DetourConfig) ProtoMessage()               {}
func (*DetourConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DefaultConfig struct {
	AlterId uint32 `protobuf:"varint,1,opt,name=alter_id,json=alterId" json:"alter_id,omitempty"`
	Level   uint32 `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
}

func (m *DefaultConfig) Reset()                    { *m = DefaultConfig{} }
func (m *DefaultConfig) String() string            { return proto.CompactTextString(m) }
func (*DefaultConfig) ProtoMessage()               {}
func (*DefaultConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Config struct {
	User    []*v2ray_core_common_protocol.User `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	Default *DefaultConfig                     `protobuf:"bytes,2,opt,name=default" json:"default,omitempty"`
	Detour  *DetourConfig                      `protobuf:"bytes,3,opt,name=detour" json:"detour,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Config) GetUser() []*v2ray_core_common_protocol.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Config) GetDefault() *DefaultConfig {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *Config) GetDetour() *DetourConfig {
	if m != nil {
		return m.Detour
	}
	return nil
}

func init() {
	proto.RegisterType((*DetourConfig)(nil), "v2ray.core.proxy.vmess.inbound.DetourConfig")
	proto.RegisterType((*DefaultConfig)(nil), "v2ray.core.proxy.vmess.inbound.DefaultConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.vmess.inbound.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/vmess/inbound/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x03, 0xd5, 0xa2, 0x8b, 0xf5, 0x40, 0x3c, 0xa0, 0x07, 0x42, 0x38, 0xd5, 0x44, 0x77,
	0x13, 0xf4, 0x01, 0x4c, 0x4b, 0x62, 0x7a, 0x6b, 0x36, 0xb1, 0x07, 0x3d, 0x18, 0xba, 0x6c, 0x0d,
	0x09, 0x30, 0xcd, 0xb2, 0x10, 0x79, 0x25, 0xdf, 0xc5, 0x77, 0x32, 0xcc, 0x42, 0xfc, 0x73, 0xb0,
	0xc7, 0x99, 0x7c, 0xbf, 0x6f, 0xbe, 0x6f, 0x08, 0x6b, 0x63, 0x95, 0x76, 0x54, 0x40, 0xc9, 0x04,
	0x28, 0xc9, 0xf6, 0x0a, 0xde, 0x3b, 0xd6, 0x96, 0xb2, 0xae, 0x59, 0x5e, 0x6d, 0xa1, 0xa9, 0x32,
	0x26, 0xa0, 0xda, 0xe5, 0x6f, 0x74, 0xaf, 0x40, 0x83, 0x17, 0x8c, 0x80, 0x92, 0x14, 0xc5, 0x14,
	0xc5, 0x74, 0x10, 0x5f, 0x5d, 0xff, 0x31, 0x14, 0x50, 0x96, 0x50, 0x31, 0x84, 0x05, 0x14, 0xac,
	0xa9, 0xa5, 0x32, 0x56, 0x51, 0x40, 0xce, 0x12, 0xa9, 0xa1, 0x51, 0x4b, 0x3c, 0xe0, 0x9d, 0x13,
	0x5b, 0x83, 0x6f, 0x85, 0xd6, 0xfc, 0x94, 0xdb, 0x1a, 0xa2, 0x07, 0x32, 0x4b, 0xe4, 0x2e, 0x6d,
	0x0a, 0x3d, 0x08, 0x2e, 0xc9, 0x49, 0x5a, 0x68, 0xa9, 0x5e, 0xf3, 0x0c, 0x65, 0x33, 0xee, 0xe0,
	0xbc, 0xca, 0xbc, 0x0b, 0x72, 0x5c, 0xc8, 0x56, 0x16, 0xbe, 0x8d, 0x7b, 0x33, 0x44, 0x9f, 0x16,
	0x99, 0x0e, 0xec, 0x3d, 0x39, 0xea, 0x4f, 0xfb, 0x56, 0x38, 0x99, 0xbb, 0x71, 0x48, 0x7f, 0xd4,
	0x30, 0x11, 0xe9, 0x18, 0x91, 0x3e, 0xd5, 0x52, 0x71, 0x54, 0x7b, 0x8f, 0xc4, 0xc9, 0x4c, 0x04,
	0x34, 0x76, 0xe3, 0x5b, 0xfa, 0x7f, 0x7f, 0xfa, 0x2b, 0x31, 0x1f, 0x69, 0x2f, 0x21, 0xd3, 0x0c,
	0xbb, 0xfa, 0x13, 0xf4, 0xb9, 0x39, 0xec, 0xf3, 0xfd, 0x19, 0x3e, 0xb0, 0x8b, 0x17, 0x12, 0x09,
	0x28, 0x0f, 0xa0, 0x0b, 0xd7, 0x50, 0xeb, 0xbe, 0xcf, 0xb3, 0x33, 0x6c, 0x3f, 0xec, 0x60, 0x13,
	0xf3, 0xb4, 0xa3, 0xcb, 0x1e, 0x5b, 0x23, 0xb6, 0x41, 0x6c, 0x65, 0x04, 0xdb, 0x29, 0x3e, 0xe0,
	0xee, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x49, 0x28, 0x20, 0xa9, 0x13, 0x02, 0x00, 0x00,
}

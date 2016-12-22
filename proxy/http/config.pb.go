// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/http/config.proto
// DO NOT EDIT!

/*
Package http is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/http/config.proto

It has these top-level messages:
	ServerConfig
	ClientConfig
*/
package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Config for HTTP proxy server.
type ServerConfig struct {
	Timeout uint32 `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ClientConfig for HTTP proxy client.
type ClientConfig struct {
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.http.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.http.ClientConfig")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/http/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x28, 0xca, 0xaf, 0xa8,
	0xd4, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x85, 0xa9, 0x2b, 0x4a, 0xd5, 0x03, 0xab, 0xd1, 0x03, 0xa9, 0x51, 0xd2,
	0xe0, 0xe2, 0x09, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x72, 0x06, 0x2b, 0x16, 0x92, 0xe0, 0x62, 0x2f,
	0xc9, 0xcc, 0x4d, 0xcd, 0x2f, 0x2d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x71, 0x95,
	0xf8, 0xb8, 0x78, 0x9c, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x20, 0x2a, 0x9d, 0xdc, 0xb8, 0x24, 0x93,
	0xf3, 0x73, 0xf5, 0xb0, 0x1a, 0xeb, 0xc4, 0x0d, 0x51, 0x14, 0x00, 0xb2, 0x3a, 0x8a, 0x05, 0x24,
	0xb4, 0x8a, 0x49, 0x34, 0xcc, 0x28, 0x28, 0xb1, 0x52, 0xcf, 0x19, 0xa4, 0x34, 0x00, 0xac, 0xd4,
	0xa3, 0xa4, 0xa4, 0x20, 0x89, 0x0d, 0xec, 0x3e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0x81, 0x13, 0x9c, 0xc9, 0x00, 0x00, 0x00,
}

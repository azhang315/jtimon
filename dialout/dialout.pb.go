// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dialout.proto

/*
Package dialout is a generated protocol buffer package.

It is generated from these files:
	dialout.proto

It has these top-level messages:
	PathsConfig
	DialOutRequest
	DialOutResponse
*/
package dialout

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gnmi "github.com/nileshsimaria/jtimon/gnmi/gnmi"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PathsConfig struct {
	Path      string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Frequency uint64 `protobuf:"varint,2,opt,name=frequency" json:"frequency,omitempty"`
	Mode      string `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
}

func (m *PathsConfig) Reset()                    { *m = PathsConfig{} }
func (m *PathsConfig) String() string            { return proto.CompactTextString(m) }
func (*PathsConfig) ProtoMessage()               {}
func (*PathsConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PathsConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathsConfig) GetFrequency() uint64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *PathsConfig) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type DialOutRequest struct {
	Device         string         `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	DialOutContext []byte         `protobuf:"bytes,2,opt,name=dialOutContext,proto3" json:"dialOutContext,omitempty"`
	RpcType        string         `protobuf:"bytes,3,opt,name=rpcType" json:"rpcType,omitempty"`
	Paths          []*PathsConfig `protobuf:"bytes,4,rep,name=paths" json:"paths,omitempty"`
}

func (m *DialOutRequest) Reset()                    { *m = DialOutRequest{} }
func (m *DialOutRequest) String() string            { return proto.CompactTextString(m) }
func (*DialOutRequest) ProtoMessage()               {}
func (*DialOutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DialOutRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DialOutRequest) GetDialOutContext() []byte {
	if m != nil {
		return m.DialOutContext
	}
	return nil
}

func (m *DialOutRequest) GetRpcType() string {
	if m != nil {
		return m.RpcType
	}
	return ""
}

func (m *DialOutRequest) GetPaths() []*PathsConfig {
	if m != nil {
		return m.Paths
	}
	return nil
}

type DialOutResponse struct {
	Device         string                    `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	DialOutContext []byte                    `protobuf:"bytes,2,opt,name=dialOutContext,proto3" json:"dialOutContext,omitempty"`
	Response       []*gnmi.SubscribeResponse `protobuf:"bytes,3,rep,name=response" json:"response,omitempty"`
}

func (m *DialOutResponse) Reset()                    { *m = DialOutResponse{} }
func (m *DialOutResponse) String() string            { return proto.CompactTextString(m) }
func (*DialOutResponse) ProtoMessage()               {}
func (*DialOutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DialOutResponse) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DialOutResponse) GetDialOutContext() []byte {
	if m != nil {
		return m.DialOutContext
	}
	return nil
}

func (m *DialOutResponse) GetResponse() []*gnmi.SubscribeResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PathsConfig)(nil), "PathsConfig")
	proto.RegisterType((*DialOutRequest)(nil), "DialOutRequest")
	proto.RegisterType((*DialOutResponse)(nil), "DialOutResponse")
}

func init() { proto.RegisterFile("dialout.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x65, 0x12, 0x0a, 0xbd, 0x96, 0x22, 0x79, 0x00, 0x0b, 0x31, 0x44, 0x19, 0x50, 0xa6,
	0x0c, 0xf4, 0x11, 0xca, 0x0e, 0x72, 0x79, 0x81, 0xfc, 0xb9, 0xb6, 0x96, 0x5a, 0xdb, 0xd8, 0x17,
	0x44, 0x5f, 0x80, 0x91, 0x67, 0x46, 0xb1, 0x93, 0x82, 0x58, 0xd9, 0xbe, 0x3b, 0x7f, 0x77, 0xbf,
	0xcf, 0x07, 0x57, 0xad, 0xaa, 0xf6, 0xa6, 0xa3, 0xd2, 0x3a, 0x43, 0xe6, 0x0e, 0xb6, 0xfa, 0xa0,
	0xa2, 0xce, 0xd7, 0x30, 0x7b, 0xa9, 0x68, 0xe7, 0x57, 0x46, 0x6f, 0xd4, 0x96, 0x73, 0x48, 0x6d,
	0x45, 0x3b, 0xc1, 0x32, 0x56, 0x4c, 0x65, 0xd0, 0xfc, 0x1e, 0xa6, 0x1b, 0x87, 0x6f, 0x1d, 0xea,
	0xe6, 0x28, 0xce, 0x32, 0x56, 0xa4, 0xf2, 0xa7, 0xd1, 0x4f, 0x1c, 0x4c, 0x8b, 0x22, 0x89, 0x13,
	0xbd, 0xce, 0xbf, 0x18, 0x2c, 0x9e, 0x54, 0xb5, 0x7f, 0xee, 0x48, 0xf6, 0x3e, 0x4f, 0xfc, 0x06,
	0x26, 0x2d, 0xbe, 0xab, 0x06, 0x87, 0xd5, 0x43, 0xc5, 0x1f, 0x60, 0xd1, 0x46, 0xe7, 0xca, 0x68,
	0xc2, 0x0f, 0x0a, 0x84, 0xb9, 0xfc, 0xd3, 0xe5, 0x02, 0x2e, 0x9c, 0x6d, 0x5e, 0x8f, 0x76, 0x24,
	0x8d, 0x25, 0xcf, 0xe1, 0xbc, 0x8f, 0xe9, 0x45, 0x9a, 0x25, 0xc5, 0xec, 0x71, 0x5e, 0xfe, 0xfa,
	0x8f, 0x8c, 0x4f, 0xf9, 0x27, 0x83, 0xeb, 0x53, 0x20, 0x6f, 0x8d, 0xf6, 0xf8, 0xef, 0x44, 0x4b,
	0xb8, 0x74, 0xc3, 0x2e, 0x91, 0x04, 0xf4, 0x6d, 0x19, 0x0e, 0xbb, 0xee, 0x6a, 0xdf, 0x38, 0x55,
	0xe3, 0x88, 0x92, 0x27, 0x63, 0x3d, 0x09, 0x57, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x25, 0x2d, 0x56, 0x92, 0x01, 0x00, 0x00,
}

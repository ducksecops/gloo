// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/prefix_rewrite.proto

package transformation // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// if set, prefix_rewrite will be used to rewrite the matched HTTP Path prefix on requests to this value.
type PrefixRewrite struct {
	// Set to an empty string to remove the matched HTTP Path prefix
	PrefixRewrite        string   `protobuf:"bytes,1,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrefixRewrite) Reset()         { *m = PrefixRewrite{} }
func (m *PrefixRewrite) String() string { return proto.CompactTextString(m) }
func (*PrefixRewrite) ProtoMessage()    {}
func (*PrefixRewrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_prefix_rewrite_2caab40b039098bf, []int{0}
}
func (m *PrefixRewrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrefixRewrite.Unmarshal(m, b)
}
func (m *PrefixRewrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrefixRewrite.Marshal(b, m, deterministic)
}
func (dst *PrefixRewrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrefixRewrite.Merge(dst, src)
}
func (m *PrefixRewrite) XXX_Size() int {
	return xxx_messageInfo_PrefixRewrite.Size(m)
}
func (m *PrefixRewrite) XXX_DiscardUnknown() {
	xxx_messageInfo_PrefixRewrite.DiscardUnknown(m)
}

var xxx_messageInfo_PrefixRewrite proto.InternalMessageInfo

func (m *PrefixRewrite) GetPrefixRewrite() string {
	if m != nil {
		return m.PrefixRewrite
	}
	return ""
}

func init() {
	proto.RegisterType((*PrefixRewrite)(nil), "transformation.plugins.gloo.solo.io.PrefixRewrite")
}
func (this *PrefixRewrite) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrefixRewrite)
	if !ok {
		that2, ok := that.(PrefixRewrite)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PrefixRewrite != that1.PrefixRewrite {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/prefix_rewrite.proto", fileDescriptor_prefix_rewrite_2caab40b039098bf)
}

var fileDescriptor_prefix_rewrite_2caab40b039098bf = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x8a, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x4b, 0x8a, 0x12,
	0xf3, 0x8a, 0xd3, 0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x0b, 0x8a, 0x52, 0xd3,
	0x32, 0x2b, 0xe2, 0x8b, 0x52, 0xcb, 0x8b, 0x32, 0x4b, 0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x94, 0x51, 0x15, 0xe9, 0x41, 0xf5, 0xea, 0x81, 0xcc, 0xd3, 0x03, 0x59, 0xa5, 0x97, 0x99,
	0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xaf, 0x0f, 0x62, 0x41, 0xb4, 0x2a, 0x99, 0x71,
	0xf1, 0x06, 0x80, 0x8d, 0x0c, 0x82, 0x98, 0x28, 0xa4, 0xca, 0xc5, 0x87, 0x6a, 0x87, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x6f, 0x01, 0xb2, 0x32, 0x27, 0xdf, 0x15, 0x8f, 0xe4, 0x18, 0xa3,
	0xdc, 0x89, 0xf3, 0x52, 0x41, 0x76, 0x3a, 0x7e, 0x6f, 0x25, 0xb1, 0x81, 0x5d, 0x63, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x29, 0xe2, 0x83, 0x57, 0x24, 0x01, 0x00, 0x00,
}

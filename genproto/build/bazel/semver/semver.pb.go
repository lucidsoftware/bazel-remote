// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build/bazel/semver/semver.proto

package semver

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

// The full version of a given tool.
type SemVer struct {
	// The major version, e.g 10 for 10.2.3.
	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	// The minor version, e.g. 2 for 10.2.3.
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	// The patch version, e.g 3 for 10.2.3.
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	// The pre-release version. Either this field or major/minor/patch fields
	// must be filled. They are mutually exclusive. Pre-release versions are
	// assumed to be earlier than any released versions.
	Prerelease           string   `protobuf:"bytes,4,opt,name=prerelease,proto3" json:"prerelease,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SemVer) Reset()         { *m = SemVer{} }
func (m *SemVer) String() string { return proto.CompactTextString(m) }
func (*SemVer) ProtoMessage()    {}
func (*SemVer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d8602885cf21c73, []int{0}
}

func (m *SemVer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SemVer.Unmarshal(m, b)
}
func (m *SemVer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SemVer.Marshal(b, m, deterministic)
}
func (m *SemVer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SemVer.Merge(m, src)
}
func (m *SemVer) XXX_Size() int {
	return xxx_messageInfo_SemVer.Size(m)
}
func (m *SemVer) XXX_DiscardUnknown() {
	xxx_messageInfo_SemVer.DiscardUnknown(m)
}

var xxx_messageInfo_SemVer proto.InternalMessageInfo

func (m *SemVer) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *SemVer) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *SemVer) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *SemVer) GetPrerelease() string {
	if m != nil {
		return m.Prerelease
	}
	return ""
}

func init() {
	proto.RegisterType((*SemVer)(nil), "build.bazel.semver.SemVer")
}

func init() { proto.RegisterFile("build/bazel/semver/semver.proto", fileDescriptor_3d8602885cf21c73) }

var fileDescriptor_3d8602885cf21c73 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2a, 0xcd, 0xcc,
	0x49, 0xd1, 0x4f, 0x4a, 0xac, 0x4a, 0xcd, 0xd1, 0x2f, 0x4e, 0xcd, 0x2d, 0x4b, 0x2d, 0x82, 0x52,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x60, 0x05, 0x7a, 0x60, 0x05, 0x7a, 0x10, 0x19,
	0xa5, 0x2c, 0x2e, 0xb6, 0xe0, 0xd4, 0xdc, 0xb0, 0xd4, 0x22, 0x21, 0x11, 0x2e, 0xd6, 0xdc, 0xc4,
	0xac, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x07, 0x2c, 0x9a, 0x99, 0x97,
	0x5f, 0x24, 0xc1, 0x04, 0x15, 0x05, 0x71, 0x40, 0xa2, 0x05, 0x89, 0x25, 0xc9, 0x19, 0x12, 0xcc,
	0x10, 0x51, 0x30, 0x47, 0x48, 0x8e, 0x8b, 0xab, 0xa0, 0x28, 0xb5, 0x28, 0x35, 0x27, 0x35, 0xb1,
	0x38, 0x55, 0x82, 0x45, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0xc4, 0xc9, 0x8d, 0x0b, 0x8b, 0x0b,
	0x9c, 0xb8, 0x83, 0xc1, 0x74, 0x00, 0xc8, 0x89, 0x01, 0x8c, 0x51, 0x6c, 0x10, 0xe1, 0x45, 0x4c,
	0xcc, 0xc1, 0xbe, 0x61, 0xab, 0x98, 0x84, 0x9c, 0xc0, 0x3a, 0x9c, 0xc0, 0x3a, 0x20, 0x2a, 0x93,
	0xd8, 0xc0, 0xde, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x49, 0x4e, 0xf3, 0x93, 0xf1, 0x00,
	0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: github.com/golang/protobuf/types/empty/empty.proto
// DO NOT EDIT!

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	github.com/golang/protobuf/types/empty/empty.proto

It has these top-level messages:
	Empty
*/
package empty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xcd, 0x2d,
	0x28, 0xa9, 0x84, 0x90, 0x7a, 0x60, 0x29, 0x21, 0xfe, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x3d,
	0x98, 0x42, 0x25, 0x76, 0x2e, 0x56, 0x57, 0x90, 0xbc, 0x53, 0x38, 0x97, 0x30, 0xd0, 0x20, 0x3d,
	0x34, 0x79, 0x27, 0x2e, 0xb0, 0x6c, 0x00, 0x88, 0x1b, 0xc0, 0x18, 0xc5, 0x0a, 0x36, 0x6b, 0x01,
	0x23, 0xe3, 0x0f, 0x46, 0xc6, 0x45, 0x4c, 0xcc, 0xee, 0x01, 0x4e, 0xab, 0x98, 0xe4, 0xdc, 0x21,
	0x5a, 0x02, 0xa0, 0x5a, 0xf4, 0xc2, 0x53, 0x73, 0x72, 0xbc, 0xf3, 0xf2, 0xcb, 0xf3, 0x42, 0x40,
	0x8e, 0x48, 0x62, 0x03, 0x9b, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0x06, 0xae, 0x32,
	0xaf, 0x00, 0x00, 0x00,
}
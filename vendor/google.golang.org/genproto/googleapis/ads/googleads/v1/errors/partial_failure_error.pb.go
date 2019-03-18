// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/partial_failure_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible partial failure errors.
type PartialFailureErrorEnum_PartialFailureError int32

const (
	// Enum unspecified.
	PartialFailureErrorEnum_UNSPECIFIED PartialFailureErrorEnum_PartialFailureError = 0
	// The received error code is not known in this version.
	PartialFailureErrorEnum_UNKNOWN PartialFailureErrorEnum_PartialFailureError = 1
	// The partial failure field was false in the request.
	// This method requires this field be set to true.
	PartialFailureErrorEnum_PARTIAL_FAILURE_MODE_REQUIRED PartialFailureErrorEnum_PartialFailureError = 2
)

var PartialFailureErrorEnum_PartialFailureError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PARTIAL_FAILURE_MODE_REQUIRED",
}
var PartialFailureErrorEnum_PartialFailureError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"PARTIAL_FAILURE_MODE_REQUIRED": 2,
}

func (x PartialFailureErrorEnum_PartialFailureError) String() string {
	return proto.EnumName(PartialFailureErrorEnum_PartialFailureError_name, int32(x))
}
func (PartialFailureErrorEnum_PartialFailureError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_partial_failure_error_fbb749ef1d8095fb, []int{0, 0}
}

// Container for enum describing possible partial failure errors.
type PartialFailureErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartialFailureErrorEnum) Reset()         { *m = PartialFailureErrorEnum{} }
func (m *PartialFailureErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PartialFailureErrorEnum) ProtoMessage()    {}
func (*PartialFailureErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_partial_failure_error_fbb749ef1d8095fb, []int{0}
}
func (m *PartialFailureErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialFailureErrorEnum.Unmarshal(m, b)
}
func (m *PartialFailureErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialFailureErrorEnum.Marshal(b, m, deterministic)
}
func (dst *PartialFailureErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialFailureErrorEnum.Merge(dst, src)
}
func (m *PartialFailureErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PartialFailureErrorEnum.Size(m)
}
func (m *PartialFailureErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialFailureErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PartialFailureErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PartialFailureErrorEnum)(nil), "google.ads.googleads.v1.errors.PartialFailureErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.PartialFailureErrorEnum_PartialFailureError", PartialFailureErrorEnum_PartialFailureError_name, PartialFailureErrorEnum_PartialFailureError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/partial_failure_error.proto", fileDescriptor_partial_failure_error_fbb749ef1d8095fb)
}

var fileDescriptor_partial_failure_error_fbb749ef1d8095fb = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0xff, 0xf2, 0x27, 0x9a, 0x4c, 0x17, 0x36, 0xb8, 0xd0, 0x18, 0x6d, 0x22, 0x0f, 0x30,
	0x84, 0xb8, 0x1b, 0x57, 0x53, 0x99, 0x36, 0xc4, 0x4a, 0x11, 0x05, 0x13, 0x43, 0x42, 0x46, 0xc1,
	0x09, 0x09, 0x9d, 0xc1, 0x19, 0xda, 0x07, 0x72, 0xe9, 0xa3, 0xf8, 0x28, 0x6e, 0x7d, 0x01, 0x03,
	0xd7, 0x76, 0x55, 0x5d, 0x71, 0x72, 0x39, 0xdf, 0xb9, 0x67, 0x2e, 0x22, 0x42, 0x29, 0x51, 0x97,
	0x2e, 0x2f, 0x8c, 0x0b, 0xb2, 0x53, 0x6b, 0xcf, 0x2d, 0xb5, 0x56, 0xda, 0xb8, 0x0d, 0xd7, 0x6d,
	0xc5, 0xeb, 0xfc, 0x85, 0x57, 0xf5, 0x4a, 0x97, 0x79, 0x3f, 0xc6, 0x8d, 0x56, 0xad, 0xb2, 0xc7,
	0x00, 0x60, 0x5e, 0x18, 0xbc, 0x65, 0xf1, 0xda, 0xc3, 0xc0, 0x9e, 0x9c, 0x6e, 0xb2, 0x9b, 0xca,
	0xe5, 0x52, 0xaa, 0x96, 0xb7, 0x95, 0x92, 0x06, 0x68, 0xe7, 0x15, 0x1d, 0x45, 0x10, 0x3e, 0x85,
	0x6c, 0xd6, 0x51, 0x4c, 0xae, 0x96, 0x4e, 0x8a, 0x0e, 0x77, 0xfc, 0xb2, 0x0f, 0xd0, 0x30, 0x09,
	0xef, 0x22, 0x76, 0x15, 0x4c, 0x03, 0xe6, 0x8f, 0xfe, 0xd9, 0x43, 0xb4, 0x9f, 0x84, 0xd7, 0xe1,
	0xe2, 0x21, 0x1c, 0x0d, 0xec, 0x73, 0x74, 0x16, 0xd1, 0xf8, 0x3e, 0xa0, 0xf3, 0x7c, 0x4a, 0x83,
	0x79, 0x12, 0xb3, 0xfc, 0x66, 0xe1, 0xb3, 0x3c, 0x66, 0xb7, 0x49, 0x10, 0x33, 0x7f, 0x64, 0x4d,
	0xbe, 0x06, 0xc8, 0x79, 0x56, 0x4b, 0xfc, 0x77, 0xef, 0xc9, 0xf1, 0x8e, 0xe5, 0x51, 0xd7, 0x39,
	0x1a, 0x3c, 0xfa, 0x3f, 0xac, 0x50, 0x35, 0x97, 0x02, 0x2b, 0x2d, 0x5c, 0x51, 0xca, 0xfe, 0x45,
	0x9b, 0xfb, 0x35, 0x95, 0xf9, 0xed, 0x9c, 0x97, 0xf0, 0x79, 0xb3, 0xfe, 0xcf, 0x28, 0x7d, 0xb7,
	0xc6, 0x33, 0x08, 0xa3, 0x85, 0xc1, 0x20, 0x3b, 0x95, 0x7a, 0xb8, 0x5f, 0x69, 0x3e, 0x36, 0x86,
	0x8c, 0x16, 0x26, 0xdb, 0x1a, 0xb2, 0xd4, 0xcb, 0xc0, 0xf0, 0x69, 0x39, 0x30, 0x25, 0x84, 0x16,
	0x86, 0x90, 0xad, 0x85, 0x90, 0xd4, 0x23, 0x04, 0x4c, 0x4f, 0x7b, 0x7d, 0xbb, 0x8b, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0d, 0xc9, 0x1f, 0x5b, 0xeb, 0x01, 0x00, 0x00,
}

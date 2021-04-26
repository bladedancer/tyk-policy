// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coprocess_response_object.proto

package coprocess

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

type ResponseObject struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	RawBody              []byte            `protobuf:"bytes,2,opt,name=raw_body,json=rawBody,proto3" json:"raw_body,omitempty"`
	Body                 string            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Headers              map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResponseObject) Reset()         { *m = ResponseObject{} }
func (m *ResponseObject) String() string { return proto.CompactTextString(m) }
func (*ResponseObject) ProtoMessage()    {}
func (*ResponseObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_95aa75cec67c3939, []int{0}
}

func (m *ResponseObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseObject.Unmarshal(m, b)
}
func (m *ResponseObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseObject.Marshal(b, m, deterministic)
}
func (m *ResponseObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseObject.Merge(m, src)
}
func (m *ResponseObject) XXX_Size() int {
	return xxx_messageInfo_ResponseObject.Size(m)
}
func (m *ResponseObject) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseObject.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseObject proto.InternalMessageInfo

func (m *ResponseObject) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ResponseObject) GetRawBody() []byte {
	if m != nil {
		return m.RawBody
	}
	return nil
}

func (m *ResponseObject) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *ResponseObject) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseObject)(nil), "coprocess.ResponseObject")
	proto.RegisterMapType((map[string]string)(nil), "coprocess.ResponseObject.HeadersEntry")
}

func init() { proto.RegisterFile("coprocess_response_object.proto", fileDescriptor_95aa75cec67c3939) }

var fileDescriptor_95aa75cec67c3939 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0xa6, 0xb5, 0x66, 0x5a, 0x44, 0x16, 0x0f, 0xb1, 0x97, 0x2e, 0x1e, 0x24, 0xa7,
	0x1c, 0xf4, 0x22, 0x3d, 0x89, 0x22, 0x78, 0x13, 0xf6, 0x0b, 0x2c, 0x9b, 0xec, 0x80, 0xff, 0xc8,
	0x84, 0x99, 0x4d, 0x4b, 0xbe, 0xac, 0x9f, 0x45, 0xdc, 0x68, 0x89, 0xb7, 0xb7, 0x6f, 0x7f, 0xf3,
	0xe6, 0x0d, 0x6c, 0x1b, 0xea, 0x98, 0x1a, 0x14, 0x71, 0x8c, 0xd2, 0x51, 0x2b, 0xe8, 0xa8, 0x7e,
	0xc7, 0x26, 0x56, 0x1d, 0x53, 0x24, 0x9d, 0x1f, 0x81, 0x8d, 0x99, 0xb2, 0xb1, 0xe7, 0xd6, 0xd1,
	0x1e, 0x99, 0xdf, 0x02, 0xca, 0x08, 0x5f, 0x7d, 0x29, 0x38, 0xb3, 0xbf, 0x31, 0x2f, 0x29, 0x45,
	0x6f, 0x61, 0x25, 0xd1, 0xc7, 0x5e, 0x5c, 0x43, 0x01, 0x0b, 0x65, 0x54, 0xb9, 0xb0, 0x30, 0x5a,
	0x8f, 0x14, 0x50, 0x5f, 0xc2, 0x29, 0xfb, 0x83, 0xab, 0x29, 0x0c, 0xc5, 0xcc, 0xa8, 0x72, 0x6d,
	0x97, 0xec, 0x0f, 0x0f, 0x14, 0x06, 0xad, 0x61, 0x9e, 0xec, 0xcc, 0xa8, 0x32, 0xb7, 0x49, 0xeb,
	0x7b, 0x58, 0xbe, 0xa2, 0x0f, 0xc8, 0x52, 0xcc, 0x4d, 0x56, 0xae, 0x6e, 0xae, 0xab, 0x63, 0xad,
	0xea, 0xff, 0xee, 0xea, 0x79, 0x04, 0x9f, 0xda, 0xc8, 0x83, 0xfd, 0x1b, 0xdb, 0xec, 0x60, 0x3d,
	0xfd, 0xd0, 0xe7, 0x90, 0x7d, 0xe0, 0x90, 0x9a, 0xe5, 0xf6, 0x47, 0xea, 0x0b, 0x58, 0xec, 0xfd,
	0x67, 0x8f, 0xa9, 0x4f, 0x6e, 0xc7, 0xc7, 0x6e, 0x76, 0xa7, 0xea, 0x93, 0x74, 0xe7, 0xed, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x37, 0x21, 0x1f, 0x37, 0x01, 0x00, 0x00,
}

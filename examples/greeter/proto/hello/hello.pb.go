// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/kubernetes/examples/greeter/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/kubernetes/examples/greeter/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

func init() {
	proto.RegisterFile("github.com/micro/kubernetes/examples/greeter/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xcd, 0xaa, 0xc2, 0x30,
	0x10, 0x85, 0x6f, 0xe9, 0xf5, 0x2f, 0x2b, 0x09, 0x2e, 0x44, 0xac, 0x48, 0x57, 0xae, 0x12, 0xd0,
	0x57, 0x70, 0xd1, 0x9d, 0x50, 0x9f, 0xa0, 0x2d, 0x87, 0xb4, 0xd8, 0x34, 0x35, 0x93, 0x8a, 0xbe,
	0xbd, 0x34, 0x66, 0xa9, 0x9b, 0xe1, 0x30, 0xdf, 0x0c, 0xdf, 0x61, 0x67, 0xd5, 0xb8, 0x7a, 0x28,
	0x45, 0x65, 0xb4, 0xd4, 0x4d, 0x65, 0x8d, 0xbc, 0x0d, 0x25, 0x6c, 0x07, 0x07, 0x92, 0x78, 0x16,
	0xba, 0x6f, 0x41, 0x52, 0x59, 0xc0, 0xc1, 0xca, 0xde, 0x1a, 0x67, 0x64, 0x8d, 0xb6, 0x0d, 0x53,
	0xf8, 0x0d, 0x5f, 0x29, 0x23, 0xfc, 0xb7, 0x20, 0xfb, 0x10, 0xe1, 0x3a, 0x4d, 0xd8, 0x2c, 0xc7,
	0x7d, 0x00, 0x39, 0xce, 0xd9, 0x7f, 0x57, 0x68, 0xac, 0xa3, 0x7d, 0x74, 0x58, 0xe4, 0x3e, 0xa7,
	0x5b, 0x36, 0xcf, 0x41, 0xbd, 0xe9, 0x08, 0x7c, 0xc9, 0x62, 0x4d, 0x2a, 0xe0, 0x31, 0x1e, 0x2f,
	0x2c, 0xbe, 0x16, 0x2f, 0x9e, 0xb1, 0x49, 0x36, 0x8a, 0x78, 0x22, 0xbe, 0x39, 0x44, 0x10, 0x6c,
	0x76, 0xbf, 0xf0, 0x47, 0x90, 0xfe, 0x95, 0x53, 0x5f, 0xf5, 0xf4, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0xc1, 0x91, 0xcf, 0xf2, 0x00, 0x00, 0x00,
}

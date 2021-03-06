// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authz.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	authz.proto

It has these top-level messages:
	Request
	HTTPRequest
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Response_Status_Code int32

const (
	// https://cloud.google.com/appengine/docs/admin-api/reference/rpc/google.rpc#google.rpc.Code
	Response_Status_OK                  Response_Status_Code = 0
	Response_Status_CANCELLED           Response_Status_Code = 1
	Response_Status_UNKNOWN             Response_Status_Code = 2
	Response_Status_INVALID_ARGUMENT    Response_Status_Code = 3
	Response_Status_DEADLINE_EXCEEDED   Response_Status_Code = 4
	Response_Status_NOT_FOUND           Response_Status_Code = 5
	Response_Status_ALREADY_EXISTS      Response_Status_Code = 6
	Response_Status_PERMISSION_DENIED   Response_Status_Code = 7
	Response_Status_UNAUTHENTICATED     Response_Status_Code = 8
	Response_Status_RESOURCE_EXHAUSTED  Response_Status_Code = 9
	Response_Status_FAILED_PRECONDITION Response_Status_Code = 10
	Response_Status_ABORTED             Response_Status_Code = 11
	Response_Status_OUT_OF_RANGE        Response_Status_Code = 12
	Response_Status_UNIMPLEMENTED       Response_Status_Code = 13
	Response_Status_INTERNAL            Response_Status_Code = 14
	Response_Status_UNAVAILABLE         Response_Status_Code = 15
	Response_Status_DATA_LOSS           Response_Status_Code = 16
)

var Response_Status_Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	8:  "UNAUTHENTICATED",
	9:  "RESOURCE_EXHAUSTED",
	10: "FAILED_PRECONDITION",
	11: "ABORTED",
	12: "OUT_OF_RANGE",
	13: "UNIMPLEMENTED",
	14: "INTERNAL",
	15: "UNAVAILABLE",
	16: "DATA_LOSS",
}
var Response_Status_Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     8,
	"RESOURCE_EXHAUSTED":  9,
	"FAILED_PRECONDITION": 10,
	"ABORTED":             11,
	"OUT_OF_RANGE":        12,
	"UNIMPLEMENTED":       13,
	"INTERNAL":            14,
	"UNAVAILABLE":         15,
	"DATA_LOSS":           16,
}

func (x Response_Status_Code) String() string {
	return proto1.EnumName(Response_Status_Code_name, int32(x))
}
func (Response_Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0, 0} }

type Request struct {
	Subject *Request_Subject `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	Action  *Request_Action  `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetSubject() *Request_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Request) GetAction() *Request_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

type Request_Subject struct {
	ServiceAccount       string            `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	Namespace            string            `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Service              string            `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	Pod                  string            `protobuf:"bytes,4,opt,name=pod" json:"pod,omitempty"`
	IpAddress            string            `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	Port                 string            `protobuf:"bytes,6,opt,name=port" json:"port,omitempty"`
	ServiceAccountLabels map[string]string `protobuf:"bytes,7,rep,name=service_account_labels,json=serviceAccountLabels" json:"service_account_labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request_Subject) Reset()                    { *m = Request_Subject{} }
func (m *Request_Subject) String() string            { return proto1.CompactTextString(m) }
func (*Request_Subject) ProtoMessage()               {}
func (*Request_Subject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Request_Subject) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *Request_Subject) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Request_Subject) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request_Subject) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *Request_Subject) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Request_Subject) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Request_Subject) GetServiceAccountLabels() map[string]string {
	if m != nil {
		return m.ServiceAccountLabels
	}
	return nil
}

type Request_Action struct {
	Namespace string       `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Service   string       `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Port      string       `protobuf:"bytes,3,opt,name=port" json:"port,omitempty"`
	IpAddress string       `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	Http      *HTTPRequest `protobuf:"bytes,5,opt,name=http" json:"http,omitempty"`
}

func (m *Request_Action) Reset()                    { *m = Request_Action{} }
func (m *Request_Action) String() string            { return proto1.CompactTextString(m) }
func (*Request_Action) ProtoMessage()               {}
func (*Request_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Request_Action) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Request_Action) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request_Action) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Request_Action) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Request_Action) GetHttp() *HTTPRequest {
	if m != nil {
		return m.Http
	}
	return nil
}

type HTTPRequest struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Path    string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (m *HTTPRequest) String() string            { return proto1.CompactTextString(m) }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HTTPRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HTTPRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Response struct {
	Status *Response_Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatus() *Response_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// TODO Maybe replace with googleapis
type Response_Status struct {
	Code    Response_Status_Code `protobuf:"varint,1,opt,name=code,enum=authz.v1.Response_Status_Code" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response_Status) Reset()                    { *m = Response_Status{} }
func (m *Response_Status) String() string            { return proto1.CompactTextString(m) }
func (*Response_Status) ProtoMessage()               {}
func (*Response_Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Response_Status) GetCode() Response_Status_Code {
	if m != nil {
		return m.Code
	}
	return Response_Status_OK
}

func (m *Response_Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*Request)(nil), "authz.v1.Request")
	proto1.RegisterType((*Request_Subject)(nil), "authz.v1.Request.Subject")
	proto1.RegisterType((*Request_Action)(nil), "authz.v1.Request.Action")
	proto1.RegisterType((*HTTPRequest)(nil), "authz.v1.HTTPRequest")
	proto1.RegisterType((*Response)(nil), "authz.v1.Response")
	proto1.RegisterType((*Response_Status)(nil), "authz.v1.Response.Status")
	proto1.RegisterEnum("authz.v1.Response_Status_Code", Response_Status_Code_name, Response_Status_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authorization service

type AuthorizationClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/authz.v1.Authorization/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationServer interface {
	Check(context.Context, *Request) (*Response, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authz.v1.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authz.v1.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz.proto",
}

func init() { proto1.RegisterFile("authz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcd, 0x72, 0x9c, 0x46,
	0x10, 0xc7, 0xb3, 0x1f, 0x62, 0xb5, 0xbd, 0xfa, 0x18, 0xb5, 0x65, 0x05, 0x6f, 0x25, 0x29, 0x97,
	0x2e, 0x71, 0x2e, 0x38, 0x5e, 0x5d, 0x52, 0xb9, 0xa4, 0x46, 0x30, 0x92, 0x28, 0xa3, 0x41, 0x35,
	0x80, 0xe3, 0xe4, 0x42, 0x21, 0x76, 0xca, 0x4b, 0x24, 0x2d, 0x84, 0x99, 0x55, 0x95, 0xfd, 0x14,
	0x79, 0x82, 0xbc, 0x46, 0x5e, 0x23, 0xef, 0x90, 0x87, 0xc8, 0x35, 0xc5, 0xc0, 0xda, 0x96, 0x65,
	0xe5, 0xc4, 0xf4, 0xbf, 0xff, 0x3d, 0xd3, 0xfd, 0x63, 0x00, 0x26, 0xd9, 0x4a, 0x2f, 0xde, 0x39,
	0x55, 0x5d, 0xea, 0x12, 0x37, 0xdb, 0xe0, 0xf6, 0xc5, 0xe1, 0xbf, 0x43, 0x18, 0x09, 0xf9, 0xfb,
	0x4a, 0x2a, 0x8d, 0x47, 0x30, 0x52, 0xab, 0xcb, 0xdf, 0x64, 0xae, 0xed, 0xde, 0xd3, 0xde, 0xb3,
	0xc9, 0xec, 0x89, 0xb3, 0xf6, 0x39, 0x9d, 0xc7, 0x89, 0x5a, 0x83, 0x58, 0x3b, 0xf1, 0x7b, 0xb0,
	0xb2, 0x5c, 0x17, 0xe5, 0xd2, 0xee, 0x9b, 0x1a, 0xfb, 0x7e, 0x0d, 0x35, 0x79, 0xd1, 0xf9, 0xa6,
	0xff, 0xf4, 0x61, 0xd4, 0x6d, 0x83, 0xdf, 0xc2, 0xae, 0x92, 0xf5, 0x6d, 0x91, 0xcb, 0x34, 0xcb,
	0xf3, 0x72, 0xb5, 0x6c, 0x8f, 0x1e, 0x8b, 0x9d, 0x4e, 0xa6, 0xad, 0x8a, 0x5f, 0xc1, 0x78, 0x99,
	0xdd, 0x48, 0x55, 0x65, 0xb9, 0x34, 0x27, 0x8d, 0xc5, 0x07, 0x01, 0x6d, 0x18, 0x75, 0x7e, 0x7b,
	0x60, 0x72, 0xeb, 0x10, 0x09, 0x0c, 0xaa, 0x72, 0x6e, 0x0f, 0x8d, 0xda, 0x2c, 0xf1, 0x6b, 0x80,
	0xa2, 0x4a, 0xb3, 0xf9, 0xbc, 0x96, 0x4a, 0xd9, 0x1b, 0xed, 0x56, 0x45, 0x45, 0x5b, 0x01, 0x11,
	0x86, 0x55, 0x59, 0x6b, 0xdb, 0x32, 0x09, 0xb3, 0xc6, 0x02, 0x0e, 0x3e, 0xe9, 0x32, 0xbd, 0xce,
	0x2e, 0xe5, 0xb5, 0xb2, 0x47, 0x4f, 0x07, 0xcf, 0x26, 0xb3, 0xa3, 0x07, 0x39, 0x39, 0xd1, 0x9d,
	0x31, 0x02, 0x53, 0xc5, 0x96, 0xba, 0x7e, 0x2b, 0xf6, 0xd5, 0x67, 0x52, 0xd3, 0x53, 0x78, 0xf2,
	0x60, 0x49, 0x33, 0xcc, 0x95, 0x7c, 0xdb, 0x11, 0x6a, 0x96, 0xb8, 0x0f, 0x1b, 0xb7, 0xd9, 0xf5,
	0x6a, 0x8d, 0xa4, 0x0d, 0x7e, 0xec, 0xff, 0xd0, 0x9b, 0xfe, 0xd9, 0x03, 0xab, 0x05, 0x7f, 0x97,
	0x5d, 0xef, 0x7f, 0xd8, 0xf5, 0xef, 0xb2, 0x5b, 0xa3, 0x18, 0x7c, 0x84, 0xe2, 0x2e, 0xbd, 0xe1,
	0xa7, 0xf4, 0xbe, 0x83, 0xe1, 0x42, 0xeb, 0xca, 0x60, 0x9d, 0xcc, 0x1e, 0x7f, 0xe0, 0x72, 0x16,
	0xc7, 0x17, 0x1d, 0x1b, 0x61, 0x2c, 0x87, 0x11, 0x4c, 0x3e, 0x12, 0x9b, 0x36, 0x6e, 0x65, 0xad,
	0x9a, 0x8b, 0xd4, 0xb6, 0xb8, 0x0e, 0xf1, 0x00, 0xac, 0x1b, 0xa9, 0x17, 0xe5, 0xbc, 0xeb, 0xaf,
	0x8b, 0x4c, 0x7b, 0x99, 0x5e, 0xbc, 0x6f, 0x2f, 0xd3, 0x8b, 0xc3, 0xbf, 0x07, 0xb0, 0x29, 0xa4,
	0xaa, 0xca, 0xa5, 0x92, 0xf8, 0x02, 0x2c, 0xa5, 0x33, 0xbd, 0x52, 0x9f, 0xbb, 0xce, 0xad, 0xc7,
	0x89, 0x8c, 0x41, 0x74, 0xc6, 0xe9, 0x1f, 0x03, 0xb0, 0x5a, 0x09, 0x67, 0x30, 0xcc, 0xcb, 0x79,
	0x0b, 0x6c, 0x67, 0xf6, 0xcd, 0x83, 0xb5, 0x8e, 0x5b, 0xce, 0xa5, 0x30, 0xde, 0x66, 0x88, 0x1b,
	0xa9, 0x54, 0xf6, 0xe6, 0x3d, 0xcb, 0x2e, 0x3c, 0xfc, 0xab, 0x0f, 0xc3, 0xc6, 0x88, 0x16, 0xf4,
	0xc3, 0x97, 0xe4, 0x0b, 0xdc, 0x86, 0xb1, 0x4b, 0xb9, 0xcb, 0x82, 0x80, 0x79, 0xa4, 0x87, 0x13,
	0x18, 0x25, 0xfc, 0x25, 0x0f, 0x7f, 0xe6, 0xa4, 0x8f, 0xfb, 0x40, 0x7c, 0xfe, 0x8a, 0x06, 0xbe,
	0x97, 0x52, 0x71, 0x9a, 0x9c, 0x33, 0x1e, 0x93, 0x01, 0x3e, 0x86, 0x3d, 0x8f, 0x51, 0x2f, 0xf0,
	0x39, 0x4b, 0xd9, 0x6b, 0x97, 0x31, 0x8f, 0x79, 0x64, 0xd8, 0x6c, 0xc4, 0xc3, 0x38, 0x3d, 0x09,
	0x13, 0xee, 0x91, 0x0d, 0x44, 0xd8, 0xa1, 0x81, 0x60, 0xd4, 0xfb, 0x25, 0x65, 0xaf, 0xfd, 0x28,
	0x8e, 0x88, 0xd5, 0x54, 0x5e, 0x30, 0x71, 0xee, 0x47, 0x91, 0x1f, 0xf2, 0xd4, 0x63, 0xdc, 0x67,
	0x1e, 0x19, 0xe1, 0x23, 0xd8, 0x4d, 0x38, 0x4d, 0xe2, 0x33, 0xc6, 0x63, 0xdf, 0xa5, 0x31, 0xf3,
	0xc8, 0x26, 0x1e, 0x00, 0x0a, 0x16, 0x85, 0x89, 0x70, 0x9b, 0x53, 0xce, 0x68, 0x12, 0x35, 0xfa,
	0x18, 0xbf, 0x84, 0x47, 0x27, 0xd4, 0x0f, 0x98, 0x97, 0x5e, 0x08, 0xe6, 0x86, 0xdc, 0xf3, 0x63,
	0x3f, 0xe4, 0x04, 0x9a, 0xce, 0xe9, 0x71, 0x28, 0x1a, 0xd7, 0x04, 0x09, 0x6c, 0x85, 0x49, 0x9c,
	0x86, 0x27, 0xa9, 0xa0, 0xfc, 0x94, 0x91, 0x2d, 0xdc, 0x83, 0xed, 0x84, 0xfb, 0xe7, 0x17, 0x01,
	0x6b, 0xc6, 0x60, 0x1e, 0xd9, 0xc6, 0x2d, 0xd8, 0xf4, 0x79, 0xcc, 0x04, 0xa7, 0x01, 0xd9, 0xc1,
	0x5d, 0x98, 0x24, 0x9c, 0xbe, 0xa2, 0x7e, 0x40, 0x8f, 0x03, 0x46, 0x76, 0x9b, 0x81, 0x3c, 0x1a,
	0xd3, 0x34, 0x08, 0xa3, 0x88, 0x90, 0xd9, 0x4f, 0xb0, 0x4d, 0x57, 0x7a, 0x51, 0xd6, 0xc5, 0xbb,
	0xcc, 0x5c, 0x67, 0x07, 0x36, 0xdc, 0x85, 0xcc, 0xaf, 0x70, 0xef, 0xde, 0x67, 0x37, 0xc5, 0xfb,
	0xaf, 0xe9, 0x78, 0xfa, 0xab, 0xad, 0x8b, 0x37, 0xb2, 0xce, 0x9c, 0xa2, 0x7c, 0x3e, 0x2f, 0xae,
	0x32, 0xa5, 0xa5, 0x7a, 0x6e, 0x7e, 0x84, 0x97, 0x96, 0x79, 0x1c, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x99, 0xe5, 0x09, 0x16, 0x1e, 0x05, 0x00, 0x00,
}

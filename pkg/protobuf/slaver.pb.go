// Code generated by protoc-gen-go. DO NOT EDIT.
// source: slaver.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Check struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Schema               string   `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Check) Reset()         { *m = Check{} }
func (m *Check) String() string { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()    {}
func (*Check) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{0}
}
func (m *Check) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Check.Unmarshal(m, b)
}
func (m *Check) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Check.Marshal(b, m, deterministic)
}
func (dst *Check) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Check.Merge(dst, src)
}
func (m *Check) XXX_Size() int {
	return xxx_messageInfo_Check.Size(m)
}
func (m *Check) XXX_DiscardUnknown() {
	xxx_messageInfo_Check.DiscardUnknown(m)
}

var xxx_messageInfo_Check proto.InternalMessageInfo

func (m *Check) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Check) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Check) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Check) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Check) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Mysql struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mysql) Reset()         { *m = Mysql{} }
func (m *Mysql) String() string { return proto.CompactTextString(m) }
func (*Mysql) ProtoMessage()    {}
func (*Mysql) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{1}
}
func (m *Mysql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mysql.Unmarshal(m, b)
}
func (m *Mysql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mysql.Marshal(b, m, deterministic)
}
func (dst *Mysql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mysql.Merge(dst, src)
}
func (m *Mysql) XXX_Size() int {
	return xxx_messageInfo_Mysql.Size(m)
}
func (m *Mysql) XXX_DiscardUnknown() {
	xxx_messageInfo_Mysql.DiscardUnknown(m)
}

var xxx_messageInfo_Mysql proto.InternalMessageInfo

func (m *Mysql) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Mysql) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Mysql) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Mysql) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type RouteSql struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Scope                string   `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Sql                  string   `protobuf:"bytes,3,opt,name=sql,proto3" json:"sql,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Mysql                *Mysql   `protobuf:"bytes,5,opt,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSql) Reset()         { *m = RouteSql{} }
func (m *RouteSql) String() string { return proto.CompactTextString(m) }
func (*RouteSql) ProtoMessage()    {}
func (*RouteSql) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{2}
}
func (m *RouteSql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSql.Unmarshal(m, b)
}
func (m *RouteSql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSql.Marshal(b, m, deterministic)
}
func (dst *RouteSql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSql.Merge(dst, src)
}
func (m *RouteSql) XXX_Size() int {
	return xxx_messageInfo_RouteSql.Size(m)
}
func (m *RouteSql) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSql.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSql proto.InternalMessageInfo

func (m *RouteSql) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RouteSql) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *RouteSql) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *RouteSql) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RouteSql) GetMysql() *Mysql {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type RouteCommand struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode           int32    `protobuf:"varint,5,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteCommand) Reset()         { *m = RouteCommand{} }
func (m *RouteCommand) String() string { return proto.CompactTextString(m) }
func (*RouteCommand) ProtoMessage()    {}
func (*RouteCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{3}
}
func (m *RouteCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteCommand.Unmarshal(m, b)
}
func (m *RouteCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteCommand.Marshal(b, m, deterministic)
}
func (dst *RouteCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteCommand.Merge(dst, src)
}
func (m *RouteCommand) XXX_Size() int {
	return xxx_messageInfo_RouteCommand.Size(m)
}
func (m *RouteCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteCommand.DiscardUnknown(m)
}

var xxx_messageInfo_RouteCommand proto.InternalMessageInfo

func (m *RouteCommand) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RouteCommand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteCommand) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RouteCommand) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RouteCommand) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

type RouteRequest struct {
	Method               string                  `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Schema               string                  `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Host                 string                  `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32                   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Path                 string                  `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Body                 string                  `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Header               map[string]*HeaderValue `protobuf:"bytes,7,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RouteRequest) Reset()         { *m = RouteRequest{} }
func (m *RouteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()    {}
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{4}
}
func (m *RouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRequest.Unmarshal(m, b)
}
func (m *RouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRequest.Marshal(b, m, deterministic)
}
func (dst *RouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRequest.Merge(dst, src)
}
func (m *RouteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteRequest.Size(m)
}
func (m *RouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRequest proto.InternalMessageInfo

func (m *RouteRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *RouteRequest) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *RouteRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *RouteRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RouteRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RouteRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *RouteRequest) GetHeader() map[string]*HeaderValue {
	if m != nil {
		return m.Header
	}
	return nil
}

type HeaderValue struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValue) Reset()         { *m = HeaderValue{} }
func (m *HeaderValue) String() string { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()    {}
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{5}
}
func (m *HeaderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValue.Unmarshal(m, b)
}
func (m *HeaderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValue.Marshal(b, m, deterministic)
}
func (dst *HeaderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValue.Merge(dst, src)
}
func (m *HeaderValue) XXX_Size() int {
	return xxx_messageInfo_HeaderValue.Size(m)
}
func (m *HeaderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValue proto.InternalMessageInfo

func (m *HeaderValue) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type Result struct {
	// The name of the feature.
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	StatusCode           int32    `protobuf:"varint,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_slaver_cf2ee650c0caa11f, []int{6}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Result) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func init() {
	proto.RegisterType((*Check)(nil), "proto.Check")
	proto.RegisterType((*Mysql)(nil), "proto.Mysql")
	proto.RegisterType((*RouteSql)(nil), "proto.RouteSql")
	proto.RegisterType((*RouteCommand)(nil), "proto.RouteCommand")
	proto.RegisterType((*RouteRequest)(nil), "proto.RouteRequest")
	proto.RegisterMapType((map[string]*HeaderValue)(nil), "proto.RouteRequest.HeaderEntry")
	proto.RegisterType((*HeaderValue)(nil), "proto.HeaderValue")
	proto.RegisterType((*Result)(nil), "proto.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteCallClient is the client API for RouteCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteCallClient interface {
	CheckHealth(ctx context.Context, in *Check, opts ...grpc.CallOption) (*Result, error)
	ExecuteSql(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteSqlClient, error)
	ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteCommandClient, error)
	ExecuteRequest(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteRequestClient, error)
}

type routeCallClient struct {
	cc *grpc.ClientConn
}

func NewRouteCallClient(cc *grpc.ClientConn) RouteCallClient {
	return &routeCallClient{cc}
}

func (c *routeCallClient) CheckHealth(ctx context.Context, in *Check, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/proto.RouteCall/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeCallClient) ExecuteSql(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteSqlClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteCall_serviceDesc.Streams[0], "/proto.RouteCall/ExecuteSql", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCallExecuteSqlClient{stream}
	return x, nil
}

type RouteCall_ExecuteSqlClient interface {
	Send(*RouteSql) error
	Recv() (*RouteSql, error)
	grpc.ClientStream
}

type routeCallExecuteSqlClient struct {
	grpc.ClientStream
}

func (x *routeCallExecuteSqlClient) Send(m *RouteSql) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeCallExecuteSqlClient) Recv() (*RouteSql, error) {
	m := new(RouteSql)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeCallClient) ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteCall_serviceDesc.Streams[1], "/proto.RouteCall/ExecuteCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCallExecuteCommandClient{stream}
	return x, nil
}

type RouteCall_ExecuteCommandClient interface {
	Send(*RouteCommand) error
	Recv() (*RouteCommand, error)
	grpc.ClientStream
}

type routeCallExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *routeCallExecuteCommandClient) Send(m *RouteCommand) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeCallExecuteCommandClient) Recv() (*RouteCommand, error) {
	m := new(RouteCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeCallClient) ExecuteRequest(ctx context.Context, opts ...grpc.CallOption) (RouteCall_ExecuteRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteCall_serviceDesc.Streams[2], "/proto.RouteCall/ExecuteRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeCallExecuteRequestClient{stream}
	return x, nil
}

type RouteCall_ExecuteRequestClient interface {
	Send(*RouteRequest) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type routeCallExecuteRequestClient struct {
	grpc.ClientStream
}

func (x *routeCallExecuteRequestClient) Send(m *RouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeCallExecuteRequestClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteCallServer is the server API for RouteCall service.
type RouteCallServer interface {
	CheckHealth(context.Context, *Check) (*Result, error)
	ExecuteSql(RouteCall_ExecuteSqlServer) error
	ExecuteCommand(RouteCall_ExecuteCommandServer) error
	ExecuteRequest(RouteCall_ExecuteRequestServer) error
}

func RegisterRouteCallServer(s *grpc.Server, srv RouteCallServer) {
	s.RegisterService(&_RouteCall_serviceDesc, srv)
}

func _RouteCall_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteCallServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RouteCall/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteCallServer).CheckHealth(ctx, req.(*Check))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteCall_ExecuteSql_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteCallServer).ExecuteSql(&routeCallExecuteSqlServer{stream})
}

type RouteCall_ExecuteSqlServer interface {
	Send(*RouteSql) error
	Recv() (*RouteSql, error)
	grpc.ServerStream
}

type routeCallExecuteSqlServer struct {
	grpc.ServerStream
}

func (x *routeCallExecuteSqlServer) Send(m *RouteSql) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeCallExecuteSqlServer) Recv() (*RouteSql, error) {
	m := new(RouteSql)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteCall_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteCallServer).ExecuteCommand(&routeCallExecuteCommandServer{stream})
}

type RouteCall_ExecuteCommandServer interface {
	Send(*RouteCommand) error
	Recv() (*RouteCommand, error)
	grpc.ServerStream
}

type routeCallExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *routeCallExecuteCommandServer) Send(m *RouteCommand) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeCallExecuteCommandServer) Recv() (*RouteCommand, error) {
	m := new(RouteCommand)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteCall_ExecuteRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteCallServer).ExecuteRequest(&routeCallExecuteRequestServer{stream})
}

type RouteCall_ExecuteRequestServer interface {
	Send(*Result) error
	Recv() (*RouteRequest, error)
	grpc.ServerStream
}

type routeCallExecuteRequestServer struct {
	grpc.ServerStream
}

func (x *routeCallExecuteRequestServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeCallExecuteRequestServer) Recv() (*RouteRequest, error) {
	m := new(RouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteCall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RouteCall",
	HandlerType: (*RouteCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _RouteCall_CheckHealth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteSql",
			Handler:       _RouteCall_ExecuteSql_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ExecuteCommand",
			Handler:       _RouteCall_ExecuteCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ExecuteRequest",
			Handler:       _RouteCall_ExecuteRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "slaver.proto",
}

func init() { proto.RegisterFile("slaver.proto", fileDescriptor_slaver_cf2ee650c0caa11f) }

var fileDescriptor_slaver_cf2ee650c0caa11f = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0x5d, 0xc7, 0x71, 0x76, 0x33, 0x4e, 0x3f, 0xd0, 0x96, 0x62, 0x72, 0xd9, 0xa0, 0x5e, 0x72,
	0x28, 0xa1, 0xa4, 0x85, 0x96, 0x3d, 0x15, 0xc2, 0xc2, 0x5e, 0xf6, 0xa2, 0x85, 0xde, 0x4a, 0xd1,
	0xda, 0x43, 0x0c, 0x2b, 0x47, 0x89, 0x25, 0x6f, 0xeb, 0xbf, 0xd0, 0xde, 0xfb, 0x2f, 0xfb, 0x1f,
	0x8a, 0x46, 0x72, 0x30, 0x69, 0x08, 0x3d, 0x79, 0xe6, 0x49, 0x33, 0xf3, 0xf4, 0xde, 0x18, 0x26,
	0x46, 0xc9, 0x27, 0xac, 0x17, 0xdb, 0x5a, 0x5b, 0xcd, 0x12, 0xfa, 0x70, 0x0d, 0xc9, 0xaa, 0xc4,
	0xfc, 0x91, 0x31, 0x18, 0xda, 0x76, 0x8b, 0x59, 0x34, 0x8b, 0xe6, 0x63, 0x41, 0xb1, 0xc3, 0x4a,
	0x6d, 0x6c, 0x36, 0xf0, 0x98, 0x8b, 0xd9, 0x6b, 0x18, 0x99, 0xbc, 0xc4, 0x4a, 0x66, 0x31, 0xa1,
	0x21, 0x73, 0x77, 0xb7, 0xba, 0xb6, 0xd9, 0x70, 0x16, 0xcd, 0x13, 0x41, 0x31, 0x61, 0xd2, 0x96,
	0x59, 0xe2, 0xeb, 0x5d, 0xcc, 0xd7, 0x90, 0xdc, 0xb5, 0x66, 0xa7, 0xd8, 0x14, 0x2e, 0x1a, 0x83,
	0xf5, 0x46, 0x56, 0xdd, 0xd0, 0x7d, 0xee, 0xce, 0xb6, 0xd2, 0x98, 0xef, 0xba, 0x2e, 0xc2, 0xf0,
	0x7d, 0xbe, 0x27, 0x15, 0xf7, 0x48, 0x1d, 0x19, 0xce, 0x7f, 0x46, 0x70, 0x21, 0x74, 0x63, 0xf1,
	0x7e, 0xa7, 0x58, 0x06, 0xe7, 0xa6, 0xc9, 0x73, 0x34, 0x86, 0x66, 0x5d, 0x88, 0x2e, 0x65, 0xaf,
	0x20, 0x31, 0xb9, 0xde, 0x62, 0x98, 0xe3, 0x13, 0xf6, 0x12, 0x62, 0xb3, 0x53, 0x61, 0x86, 0x0b,
	0x5d, 0x87, 0x0a, 0x8d, 0x91, 0x6b, 0xa4, 0x29, 0x63, 0xd1, 0xa5, 0x8c, 0x43, 0x52, 0xb9, 0x17,
	0xd1, 0x33, 0xd3, 0xe5, 0xc4, 0x0b, 0xbc, 0xa0, 0x57, 0x0a, 0x7f, 0xc4, 0x7f, 0x45, 0x30, 0x21,
	0x32, 0x2b, 0x5d, 0x55, 0x72, 0x53, 0x9c, 0x20, 0xc4, 0x60, 0x48, 0x9a, 0x04, 0xd1, 0x49, 0x0f,
	0x06, 0x43, 0x59, 0xaf, 0x4d, 0x16, 0xcf, 0x62, 0x87, 0xb9, 0xf8, 0x04, 0xa1, 0x2b, 0x48, 0x8d,
	0x95, 0xb6, 0x31, 0xdf, 0x72, 0x5d, 0x20, 0xd1, 0x4a, 0x04, 0x78, 0x68, 0xa5, 0x0b, 0xe4, 0xbf,
	0x07, 0x81, 0x8d, 0xc0, 0x5d, 0x83, 0xde, 0xd4, 0x0a, 0x6d, 0xa9, 0x8b, 0xe0, 0x44, 0xc8, 0x7a,
	0x66, 0x0f, 0x0e, 0xcd, 0xfe, 0x1f, 0x0f, 0x8e, 0x2d, 0x80, 0xc3, 0x1e, 0x74, 0xd1, 0x66, 0x23,
	0x8f, 0xb9, 0x98, 0x7d, 0x84, 0x51, 0x89, 0xb2, 0xc0, 0x3a, 0x3b, 0x9f, 0xc5, 0xf3, 0x74, 0x79,
	0x15, 0x34, 0xec, 0x93, 0x5c, 0xdc, 0xd2, 0x8d, 0x9b, 0x8d, 0xad, 0x5b, 0x11, 0xae, 0x4f, 0xef,
	0x20, 0xed, 0xc1, 0xce, 0xb6, 0x47, 0x6c, 0xc3, 0x23, 0x5c, 0xc8, 0xe6, 0x90, 0x3c, 0x49, 0xd5,
	0x78, 0x39, 0xd3, 0x25, 0x0b, 0x8d, 0x7d, 0xd1, 0x17, 0x77, 0x22, 0xfc, 0x85, 0xeb, 0xc1, 0xa7,
	0x88, 0xbf, 0xe9, 0xda, 0xd1, 0x89, 0xdb, 0x0d, 0x5f, 0x1c, 0x91, 0xee, 0x3e, 0xe1, 0x5f, 0x61,
	0x24, 0xd0, 0x34, 0xca, 0x9e, 0x30, 0xb1, 0x67, 0xce, 0xe0, 0xa4, 0x39, 0xf1, 0xa1, 0x39, 0xcb,
	0x3f, 0x11, 0x8c, 0xfd, 0xaa, 0x48, 0xa5, 0xd8, 0x5b, 0x48, 0xe9, 0xff, 0xbc, 0x45, 0xa9, 0x6c,
	0xc9, 0xba, 0xe5, 0x22, 0x6c, 0xfa, 0xac, 0x93, 0x89, 0xe8, 0xf0, 0x33, 0xf6, 0x01, 0xe0, 0xe6,
	0x07, 0xe6, 0x61, 0xe9, 0x5f, 0xf4, 0x55, 0xbc, 0xdf, 0xa9, 0xe9, 0x21, 0xc0, 0xcf, 0xe6, 0xd1,
	0xbb, 0x88, 0x7d, 0x86, 0xe7, 0xa1, 0xaa, 0xdb, 0xce, 0xcb, 0xfe, 0xc5, 0x00, 0x4e, 0x8f, 0x81,
	0xa1, 0xc3, 0xf5, 0xbe, 0x43, 0xb7, 0x51, 0x97, 0x47, 0x1c, 0xfc, 0x87, 0xaf, 0xab, 0x7d, 0x18,
	0x11, 0xf6, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xa4, 0x2c, 0x76, 0x9f, 0x04, 0x00,
	0x00,
}
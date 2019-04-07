// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Employee
	GetAllRequest
	GetByBadgeNumberRequest
	EmployeeRequest
	EmployeeResponse
	AddPhotoRequest
	AddPhotoResponse
*/
package pb

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

type Employee struct {
	Id                  int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BadgeNumber         int32   `protobuf:"varint,2,opt,name=badgeNumber" json:"badgeNumber,omitempty"`
	FirstName           string  `protobuf:"bytes,3,opt,name=firstName" json:"firstName,omitempty"`
	LastName            string  `protobuf:"bytes,4,opt,name=lastName" json:"lastName,omitempty"`
	VacationAccrualRate float32 `protobuf:"fixed32,5,opt,name=vacationAccrualRate" json:"vacationAccrualRate,omitempty"`
	VacationAccrued     float32 `protobuf:"fixed32,6,opt,name=vacationAccrued" json:"vacationAccrued,omitempty"`
}

func (m *Employee) Reset()                    { *m = Employee{} }
func (m *Employee) String() string            { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()               {}
func (*Employee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Employee) GetVacationAccrualRate() float32 {
	if m != nil {
		return m.VacationAccrualRate
	}
	return 0
}

func (m *Employee) GetVacationAccrued() float32 {
	if m != nil {
		return m.VacationAccrued
	}
	return 0
}

type GetAllRequest struct {
}

func (m *GetAllRequest) Reset()                    { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()               {}
func (*GetAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetByBadgeNumberRequest struct {
	BadgeNumber int32 `protobuf:"varint,1,opt,name=badgeNumber" json:"badgeNumber,omitempty"`
}

func (m *GetByBadgeNumberRequest) Reset()                    { *m = GetByBadgeNumberRequest{} }
func (m *GetByBadgeNumberRequest) String() string            { return proto.CompactTextString(m) }
func (*GetByBadgeNumberRequest) ProtoMessage()               {}
func (*GetByBadgeNumberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetByBadgeNumberRequest) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

type EmployeeRequest struct {
	Employee *Employee `protobuf:"bytes,1,opt,name=employee" json:"employee,omitempty"`
}

func (m *EmployeeRequest) Reset()                    { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string            { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()               {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EmployeeRequest) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type EmployeeResponse struct {
	Employee *Employee `protobuf:"bytes,1,opt,name=employee" json:"employee,omitempty"`
}

func (m *EmployeeResponse) Reset()                    { *m = EmployeeResponse{} }
func (m *EmployeeResponse) String() string            { return proto.CompactTextString(m) }
func (*EmployeeResponse) ProtoMessage()               {}
func (*EmployeeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EmployeeResponse) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type AddPhotoRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AddPhotoRequest) Reset()                    { *m = AddPhotoRequest{} }
func (m *AddPhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPhotoRequest) ProtoMessage()               {}
func (*AddPhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddPhotoRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AddPhotoResponse struct {
	IsOk bool `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
}

func (m *AddPhotoResponse) Reset()                    { *m = AddPhotoResponse{} }
func (m *AddPhotoResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPhotoResponse) ProtoMessage()               {}
func (*AddPhotoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddPhotoResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*GetAllRequest)(nil), "GetAllRequest")
	proto.RegisterType((*GetByBadgeNumberRequest)(nil), "GetByBadgeNumberRequest")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRequest")
	proto.RegisterType((*EmployeeResponse)(nil), "EmployeeResponse")
	proto.RegisterType((*AddPhotoRequest)(nil), "AddPhotoRequest")
	proto.RegisterType((*AddPhotoResponse)(nil), "AddPhotoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EmployeeService service

type EmployeeServiceClient interface {
	GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error)
	Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error)
	AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := grpc.Invoke(ctx, "/EmployeeService/GetByBadgeNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[0], c.cc, "/EmployeeService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetAllClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := grpc.Invoke(ctx, "/EmployeeService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[1], c.cc, "/EmployeeService/SaveAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceSaveAllClient{stream}
	return x, nil
}

type EmployeeService_SaveAllClient interface {
	Send(*EmployeeRequest) error
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceSaveAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceSaveAllClient) Send(m *EmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceSaveAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[2], c.cc, "/EmployeeService/AddPhoto", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceAddPhotoClient{stream}
	return x, nil
}

type EmployeeService_AddPhotoClient interface {
	Send(*AddPhotoRequest) error
	CloseAndRecv() (*AddPhotoResponse, error)
	grpc.ClientStream
}

type employeeServiceAddPhotoClient struct {
	grpc.ClientStream
}

func (x *employeeServiceAddPhotoClient) Send(m *AddPhotoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoClient) CloseAndRecv() (*AddPhotoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddPhotoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EmployeeService service

type EmployeeServiceServer interface {
	GetByBadgeNumber(context.Context, *GetByBadgeNumberRequest) (*EmployeeResponse, error)
	GetAll(*GetAllRequest, EmployeeService_GetAllServer) error
	Save(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	SaveAll(EmployeeService_SaveAllServer) error
	AddPhoto(EmployeeService_AddPhotoServer) error
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetByBadgeNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByBadgeNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/GetByBadgeNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, req.(*GetByBadgeNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetAll(m, &employeeServiceGetAllServer{stream})
}

type EmployeeService_GetAllServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Save(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_SaveAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).SaveAll(&employeeServiceSaveAllServer{stream})
}

type EmployeeService_SaveAllServer interface {
	Send(*EmployeeResponse) error
	Recv() (*EmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceSaveAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceSaveAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceSaveAllServer) Recv() (*EmployeeRequest, error) {
	m := new(EmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmployeeService_AddPhoto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).AddPhoto(&employeeServiceAddPhotoServer{stream})
}

type EmployeeService_AddPhotoServer interface {
	SendAndClose(*AddPhotoResponse) error
	Recv() (*AddPhotoRequest, error)
	grpc.ServerStream
}

type employeeServiceAddPhotoServer struct {
	grpc.ServerStream
}

func (x *employeeServiceAddPhotoServer) SendAndClose(m *AddPhotoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoServer) Recv() (*AddPhotoRequest, error) {
	m := new(AddPhotoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByBadgeNumber",
			Handler:    _EmployeeService_GetByBadgeNumber_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _EmployeeService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _EmployeeService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveAll",
			Handler:       _EmployeeService_SaveAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddPhoto",
			Handler:       _EmployeeService_AddPhoto_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6f, 0xa2, 0x40,
	0x14, 0xc7, 0x33, 0x2c, 0xba, 0xf8, 0xdc, 0x15, 0x9c, 0x3d, 0x2c, 0x31, 0x7b, 0x20, 0x24, 0x6e,
	0x48, 0x36, 0x3b, 0x71, 0xdd, 0x4b, 0x9b, 0x1e, 0x1a, 0x4d, 0x1a, 0x6f, 0xb6, 0xc1, 0x5b, 0x6f,
	0x03, 0xbc, 0x5a, 0x52, 0x10, 0xca, 0x8c, 0x26, 0xfe, 0x21, 0xfd, 0xc7, 0xfa, 0x17, 0x35, 0x4e,
	0x41, 0x2a, 0xda, 0xc4, 0xdb, 0xf0, 0xfd, 0x31, 0x8f, 0xf7, 0x49, 0x06, 0x7a, 0x29, 0x0a, 0xc1,
	0x97, 0x28, 0x58, 0x5e, 0x64, 0x32, 0x73, 0x5f, 0x09, 0x18, 0x37, 0x69, 0x9e, 0x64, 0x5b, 0x44,
	0xda, 0x03, 0x2d, 0x8e, 0x6c, 0xe2, 0x10, 0xaf, 0xe5, 0x6b, 0x71, 0x44, 0x1d, 0xe8, 0x06, 0x3c,
	0x5a, 0xe2, 0x7c, 0x9d, 0x06, 0x58, 0xd8, 0x9a, 0x32, 0x3e, 0x4a, 0xf4, 0x17, 0x74, 0x1e, 0xe2,
	0x42, 0xc8, 0x39, 0x4f, 0xd1, 0xfe, 0xe2, 0x10, 0xaf, 0xe3, 0xd7, 0x02, 0x1d, 0x80, 0x91, 0xf0,
	0xd2, 0xd4, 0x95, 0xb9, 0xff, 0xa6, 0x23, 0xf8, 0xb1, 0xe1, 0x21, 0x97, 0x71, 0xb6, 0x9a, 0x84,
	0x61, 0xb1, 0xe6, 0x89, 0xcf, 0x25, 0xda, 0x2d, 0x87, 0x78, 0x9a, 0x7f, 0xca, 0xa2, 0x1e, 0x98,
	0x07, 0x32, 0x46, 0x76, 0x5b, 0xa5, 0x9b, 0xb2, 0x6b, 0xc2, 0xf7, 0x19, 0xca, 0x49, 0x92, 0xf8,
	0xf8, 0xbc, 0x46, 0x21, 0xdd, 0x2b, 0xf8, 0x39, 0x43, 0x39, 0xdd, 0x4e, 0xeb, 0x5f, 0x2f, 0xad,
	0xe6, 0x8e, 0xe4, 0x68, 0x47, 0xf7, 0x02, 0xcc, 0x8a, 0x50, 0x55, 0x1a, 0x82, 0x81, 0xa5, 0xa4,
	0x1a, 0xdd, 0x71, 0x87, 0xed, 0x33, 0x7b, 0xcb, 0xbd, 0x04, 0xab, 0x6e, 0x8a, 0x3c, 0x5b, 0x09,
	0x3c, 0xb7, 0x3a, 0x04, 0x73, 0x12, 0x45, 0x77, 0x8f, 0x99, 0xcc, 0xaa, 0xa1, 0x14, 0xf4, 0x88,
	0x4b, 0xae, 0x5a, 0xdf, 0x7c, 0x75, 0x76, 0x7f, 0x83, 0x55, 0xc7, 0xca, 0x09, 0x14, 0xf4, 0x58,
	0xdc, 0x3e, 0xa9, 0x9c, 0xe1, 0xab, 0xf3, 0xf8, 0x45, 0xab, 0x97, 0x58, 0x60, 0xb1, 0x89, 0x43,
	0xa4, 0xd7, 0x60, 0x35, 0xa1, 0x50, 0x9b, 0x7d, 0xc2, 0x69, 0xd0, 0x67, 0x47, 0xab, 0xfc, 0x85,
	0xf6, 0x3b, 0x66, 0xda, 0x63, 0x07, 0xbc, 0x4f, 0x84, 0x47, 0x84, 0xfe, 0x01, 0x7d, 0xc1, 0x37,
	0x48, 0x2d, 0xd6, 0xc0, 0x79, 0xea, 0xee, 0x31, 0x7c, 0xdd, 0x85, 0x77, 0x97, 0x9f, 0x93, 0xf7,
	0xc8, 0x88, 0xd0, 0x7f, 0x60, 0x54, 0x30, 0xa8, 0xc5, 0x1a, 0xf8, 0x06, 0x7d, 0xd6, 0x24, 0xe5,
	0x91, 0xa9, 0x7e, 0xaf, 0xe5, 0x41, 0xd0, 0x56, 0x6f, 0xe1, 0xff, 0x5b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x94, 0xce, 0x62, 0xe1, 0x1d, 0x03, 0x00, 0x00,
}

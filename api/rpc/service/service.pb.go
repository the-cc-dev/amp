// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/appcelerator/amp/api/rpc/service/service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/service/service.proto

It has these top-level messages:
	TasksRequest
	Task
	TasksReply
	ServiceListRequest
	ServiceListReply
	ServiceListEntry
	ServiceEntity
	ServiceInspectRequest
	ServiceInspectReply
	ServiceScaleRequest
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type TasksRequest struct {
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
}

func (m *TasksRequest) Reset()                    { *m = TasksRequest{} }
func (m *TasksRequest) String() string            { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()               {}
func (*TasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TasksRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type Task struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Image        string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	CurrentState string `protobuf:"bytes,3,opt,name=current_state,json=currentState" json:"current_state,omitempty"`
	DesiredState string `protobuf:"bytes,4,opt,name=desired_state,json=desiredState" json:"desired_state,omitempty"`
	NodeId       string `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Error        string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetCurrentState() string {
	if m != nil {
		return m.CurrentState
	}
	return ""
}

func (m *Task) GetDesiredState() string {
	if m != nil {
		return m.DesiredState
	}
	return ""
}

func (m *Task) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Task) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TasksReply struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TasksReply) Reset()                    { *m = TasksReply{} }
func (m *TasksReply) String() string            { return proto.CompactTextString(m) }
func (*TasksReply) ProtoMessage()               {}
func (*TasksReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TasksReply) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ServiceListRequest struct {
	StackName string `protobuf:"bytes,1,opt,name=stack_name,json=stackName" json:"stack_name,omitempty"`
}

func (m *ServiceListRequest) Reset()                    { *m = ServiceListRequest{} }
func (m *ServiceListRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceListRequest) ProtoMessage()               {}
func (*ServiceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServiceListRequest) GetStackName() string {
	if m != nil {
		return m.StackName
	}
	return ""
}

type ServiceListReply struct {
	Entries []*ServiceListEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ServiceListReply) Reset()                    { *m = ServiceListReply{} }
func (m *ServiceListReply) String() string            { return proto.CompactTextString(m) }
func (*ServiceListReply) ProtoMessage()               {}
func (*ServiceListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ServiceListReply) GetEntries() []*ServiceListEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type ServiceListEntry struct {
	Service     *ServiceEntity `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	FailedTasks int32          `protobuf:"varint,2,opt,name=failed_tasks,json=failedTasks" json:"failed_tasks,omitempty"`
	ReadyTasks  int32          `protobuf:"varint,3,opt,name=ready_tasks,json=readyTasks" json:"ready_tasks,omitempty"`
	TotalTasks  int32          `protobuf:"varint,4,opt,name=total_tasks,json=totalTasks" json:"total_tasks,omitempty"`
	Status      string         `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *ServiceListEntry) Reset()                    { *m = ServiceListEntry{} }
func (m *ServiceListEntry) String() string            { return proto.CompactTextString(m) }
func (*ServiceListEntry) ProtoMessage()               {}
func (*ServiceListEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServiceListEntry) GetService() *ServiceEntity {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceListEntry) GetFailedTasks() int32 {
	if m != nil {
		return m.FailedTasks
	}
	return 0
}

func (m *ServiceListEntry) GetReadyTasks() int32 {
	if m != nil {
		return m.ReadyTasks
	}
	return 0
}

func (m *ServiceListEntry) GetTotalTasks() int32 {
	if m != nil {
		return m.TotalTasks
	}
	return 0
}

func (m *ServiceListEntry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ServiceEntity struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mode  string `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Image string `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Tag   string `protobuf:"bytes,5,opt,name=tag" json:"tag,omitempty"`
}

func (m *ServiceEntity) Reset()                    { *m = ServiceEntity{} }
func (m *ServiceEntity) String() string            { return proto.CompactTextString(m) }
func (*ServiceEntity) ProtoMessage()               {}
func (*ServiceEntity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceEntity) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ServiceEntity) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ServiceEntity) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ServiceInspectRequest struct {
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
}

func (m *ServiceInspectRequest) Reset()                    { *m = ServiceInspectRequest{} }
func (m *ServiceInspectRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceInspectRequest) ProtoMessage()               {}
func (*ServiceInspectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ServiceInspectRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type ServiceInspectReply struct {
	ServiceEntity string `protobuf:"bytes,1,opt,name=service_entity,json=serviceEntity" json:"service_entity,omitempty"`
}

func (m *ServiceInspectReply) Reset()                    { *m = ServiceInspectReply{} }
func (m *ServiceInspectReply) String() string            { return proto.CompactTextString(m) }
func (*ServiceInspectReply) ProtoMessage()               {}
func (*ServiceInspectReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ServiceInspectReply) GetServiceEntity() string {
	if m != nil {
		return m.ServiceEntity
	}
	return ""
}

type ServiceScaleRequest struct {
	ServiceId      string `protobuf:"bytes,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	ReplicasNumber uint64 `protobuf:"varint,2,opt,name=replicas_number,json=replicasNumber" json:"replicas_number,omitempty"`
}

func (m *ServiceScaleRequest) Reset()                    { *m = ServiceScaleRequest{} }
func (m *ServiceScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceScaleRequest) ProtoMessage()               {}
func (*ServiceScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ServiceScaleRequest) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceScaleRequest) GetReplicasNumber() uint64 {
	if m != nil {
		return m.ReplicasNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*TasksRequest)(nil), "service.TasksRequest")
	proto.RegisterType((*Task)(nil), "service.Task")
	proto.RegisterType((*TasksReply)(nil), "service.TasksReply")
	proto.RegisterType((*ServiceListRequest)(nil), "service.ServiceListRequest")
	proto.RegisterType((*ServiceListReply)(nil), "service.ServiceListReply")
	proto.RegisterType((*ServiceListEntry)(nil), "service.ServiceListEntry")
	proto.RegisterType((*ServiceEntity)(nil), "service.ServiceEntity")
	proto.RegisterType((*ServiceInspectRequest)(nil), "service.ServiceInspectRequest")
	proto.RegisterType((*ServiceInspectReply)(nil), "service.ServiceInspectReply")
	proto.RegisterType((*ServiceScaleRequest)(nil), "service.ServiceScaleRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error)
	ListService(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListReply, error)
	InspectService(ctx context.Context, in *ServiceInspectRequest, opts ...grpc.CallOption) (*ServiceInspectReply, error)
	ScaleService(ctx context.Context, in *ServiceScaleRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksReply, error) {
	out := new(TasksReply)
	err := grpc.Invoke(ctx, "/service.Service/Tasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListService(ctx context.Context, in *ServiceListRequest, opts ...grpc.CallOption) (*ServiceListReply, error) {
	out := new(ServiceListReply)
	err := grpc.Invoke(ctx, "/service.Service/ListService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) InspectService(ctx context.Context, in *ServiceInspectRequest, opts ...grpc.CallOption) (*ServiceInspectReply, error) {
	out := new(ServiceInspectReply)
	err := grpc.Invoke(ctx, "/service.Service/InspectService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ScaleService(ctx context.Context, in *ServiceScaleRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/service.Service/ScaleService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Tasks(context.Context, *TasksRequest) (*TasksReply, error)
	ListService(context.Context, *ServiceListRequest) (*ServiceListReply, error)
	InspectService(context.Context, *ServiceInspectRequest) (*ServiceInspectReply, error)
	ScaleService(context.Context, *ServiceScaleRequest) (*google_protobuf1.Empty, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Tasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/ListService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListService(ctx, req.(*ServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_InspectService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).InspectService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/InspectService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).InspectService(ctx, req.(*ServiceInspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ScaleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ScaleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/ScaleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ScaleService(ctx, req.(*ServiceScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tasks",
			Handler:    _Service_Tasks_Handler,
		},
		{
			MethodName: "ListService",
			Handler:    _Service_ListService_Handler,
		},
		{
			MethodName: "InspectService",
			Handler:    _Service_InspectService_Handler,
		},
		{
			MethodName: "ScaleService",
			Handler:    _Service_ScaleService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/service/service.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x6a, 0xdb, 0x4a,
	0x14, 0xc7, 0xf1, 0x37, 0x39, 0xfe, 0xb8, 0x61, 0x72, 0xe3, 0xeb, 0x28, 0xb9, 0x69, 0xa2, 0x50,
	0x1a, 0x0a, 0xb5, 0x92, 0x18, 0x0a, 0x85, 0xbe, 0x86, 0x12, 0x28, 0x79, 0x90, 0xfb, 0x5a, 0xcc,
	0x58, 0x9a, 0x28, 0x43, 0x24, 0xcd, 0x74, 0x34, 0x0e, 0x88, 0x90, 0x97, 0x6e, 0xa1, 0x6b, 0xe8,
	0x06, 0xba, 0x83, 0xae, 0xa1, 0x5b, 0xe8, 0x42, 0xca, 0x7c, 0x39, 0x96, 0x93, 0x42, 0x5e, 0x6c,
	0xcd, 0xff, 0xfc, 0xcf, 0xfc, 0x34, 0xe7, 0x1c, 0x0d, 0xbc, 0x4b, 0xa8, 0xbc, 0x5e, 0xcc, 0xc7,
	0x11, 0xcb, 0x02, 0xcc, 0x79, 0x44, 0x52, 0x22, 0xb0, 0x64, 0x22, 0xc0, 0x19, 0x0f, 0x30, 0xa7,
	0x81, 0xe0, 0x51, 0x50, 0x10, 0x71, 0x4b, 0x23, 0xe2, 0xfe, 0xc7, 0x5c, 0x30, 0xc9, 0x50, 0xc7,
	0x2e, 0xbd, 0xbd, 0x84, 0xb1, 0x24, 0x25, 0xda, 0x8e, 0xf3, 0x9c, 0x49, 0x2c, 0x29, 0xcb, 0x0b,
	0x63, 0xf3, 0x26, 0x2b, 0x84, 0x84, 0xa5, 0x38, 0x4f, 0x02, 0x1d, 0x98, 0x2f, 0xae, 0x02, 0x2e,
	0x4b, 0x4e, 0x8a, 0x80, 0x64, 0x5c, 0x96, 0xe6, 0xd7, 0x24, 0xf9, 0x6f, 0xa0, 0xf7, 0x09, 0x17,
	0x37, 0x45, 0x48, 0xbe, 0x2c, 0x48, 0x21, 0xd1, 0xff, 0x00, 0x96, 0x36, 0xa3, 0xf1, 0xa8, 0x76,
	0x50, 0x3b, 0xde, 0x08, 0x37, 0xac, 0x72, 0x11, 0xfb, 0xdf, 0x6b, 0xd0, 0x54, 0x7e, 0x34, 0x80,
	0xfa, 0x32, 0x5e, 0xa7, 0x31, 0xfa, 0x17, 0x5a, 0x34, 0xc3, 0x09, 0x19, 0xd5, 0xb5, 0x64, 0x16,
	0xe8, 0x08, 0xfa, 0xd1, 0x42, 0x08, 0x92, 0xcb, 0x59, 0x21, 0xb1, 0x24, 0xa3, 0x86, 0x8e, 0xf6,
	0xac, 0x38, 0x55, 0x9a, 0x32, 0xc5, 0xa4, 0xa0, 0x82, 0xc4, 0xd6, 0xd4, 0x34, 0x26, 0x2b, 0x1a,
	0xd3, 0x7f, 0xd0, 0xc9, 0x59, 0xac, 0x5f, 0xaa, 0xa5, 0xc3, 0x6d, 0xb5, 0xbc, 0xd0, 0x60, 0x22,
	0x04, 0x13, 0xa3, 0xb6, 0x01, 0xeb, 0x85, 0x7f, 0x0a, 0x60, 0x8f, 0xc5, 0xd3, 0x12, 0x1d, 0x41,
	0x4b, 0xaa, 0xd5, 0xa8, 0x76, 0xd0, 0x38, 0xee, 0x9e, 0xf5, 0xc7, 0xae, 0xbe, 0xca, 0x13, 0x9a,
	0x98, 0x3f, 0x01, 0x34, 0x35, 0xf2, 0x47, 0x5a, 0xc8, 0xd5, 0x7a, 0x48, 0x1c, 0xdd, 0xcc, 0x72,
	0x9c, 0x91, 0x65, 0x3d, 0x94, 0x72, 0x89, 0x33, 0xe2, 0x7f, 0x80, 0xcd, 0x4a, 0x92, 0xa2, 0x4d,
	0xa0, 0x43, 0x72, 0x29, 0x28, 0x71, 0xbc, 0x9d, 0x25, 0x6f, 0xc5, 0x7b, 0x9e, 0x4b, 0x51, 0x86,
	0xce, 0xe9, 0xff, 0xac, 0x55, 0x76, 0xd2, 0x51, 0x74, 0x02, 0xae, 0xf5, 0x9a, 0xdc, 0x3d, 0x1b,
	0xae, 0xef, 0x74, 0x9e, 0x4b, 0x2a, 0xcb, 0xd0, 0xd9, 0xd0, 0x21, 0xf4, 0xae, 0x30, 0x4d, 0x49,
	0x3c, 0x33, 0x07, 0x56, 0xdd, 0x68, 0x85, 0x5d, 0xa3, 0xe9, 0x8a, 0xa0, 0x17, 0xd0, 0x15, 0x04,
	0xc7, 0xa5, 0x75, 0x34, 0xb4, 0x03, 0xb4, 0xb4, 0x34, 0x48, 0x26, 0x71, 0x6a, 0x0d, 0x4d, 0x63,
	0xd0, 0x92, 0x31, 0x0c, 0xa1, 0xad, 0x1a, 0xb5, 0x28, 0x5c, 0x2b, 0xcc, 0xca, 0x67, 0xd0, 0xaf,
	0xbc, 0xd6, 0xa3, 0x21, 0x41, 0xd0, 0xd4, 0x65, 0x34, 0x33, 0xa2, 0x9f, 0x95, 0x96, 0xb1, 0xd8,
	0x4d, 0x86, 0x7e, 0x7e, 0x18, 0xa6, 0xe6, 0xea, 0x30, 0x6d, 0x42, 0x43, 0xe2, 0xc4, 0x32, 0xd5,
	0xa3, 0xff, 0x16, 0xb6, 0x2d, 0xf0, 0x22, 0x2f, 0x38, 0x89, 0xe4, 0x33, 0xa7, 0xf8, 0x3d, 0x6c,
	0xad, 0xe7, 0xa9, 0xc6, 0xbd, 0x84, 0x81, 0xcb, 0x22, 0xfa, 0x00, 0x36, 0xb3, 0x5f, 0xac, 0x9e,
	0xca, 0xff, 0xbc, 0xcc, 0x9e, 0x46, 0x38, 0x25, 0xcf, 0x63, 0xa2, 0x57, 0xf0, 0x8f, 0x20, 0x3c,
	0xa5, 0x11, 0x2e, 0x66, 0xf9, 0x22, 0x9b, 0x13, 0xa1, 0xcb, 0xd0, 0x0c, 0x07, 0x4e, 0xbe, 0xd4,
	0xea, 0xd9, 0x8f, 0x06, 0x74, 0xec, 0xfe, 0x68, 0x0a, 0x2d, 0x53, 0xf2, 0xed, 0xca, 0xc8, 0xba,
	0xaf, 0xd5, 0xdb, 0x5a, 0x97, 0x79, 0x5a, 0xfa, 0xfb, 0x5f, 0x7f, 0xfd, 0xfe, 0x56, 0x1f, 0xa1,
	0x61, 0x70, 0x7b, 0x1a, 0xe8, 0x1e, 0x06, 0x77, 0x0f, 0x6f, 0x76, 0x8f, 0xae, 0xa1, 0xab, 0x46,
	0xcc, 0x31, 0x76, 0x9f, 0x9a, 0x4e, 0x07, 0xd8, 0x79, 0x3a, 0xa8, 0x30, 0x87, 0x1a, 0xb3, 0x8b,
	0x76, 0x14, 0xc6, 0xba, 0x14, 0x69, 0xf9, 0xb5, 0xdc, 0x23, 0x0e, 0x03, 0x5b, 0x60, 0x07, 0xdb,
	0x5f, 0xdf, 0xaf, 0xda, 0x38, 0x6f, 0xef, 0xaf, 0x71, 0x85, 0x3c, 0xd0, 0x48, 0x0f, 0x8d, 0x14,
	0x92, 0x9a, 0x48, 0xf5, 0x6c, 0xb7, 0xd0, 0xd3, 0x4d, 0x71, 0xbc, 0x47, 0xfb, 0xad, 0xb6, 0xcc,
	0x1b, 0x8e, 0xcd, 0x85, 0x3a, 0x76, 0xf7, 0xe4, 0xf8, 0x5c, 0x5d, 0x8d, 0xfe, 0x89, 0xe6, 0xbc,
	0xf6, 0x8e, 0xf5, 0xd1, 0x54, 0x46, 0x85, 0x12, 0xdc, 0xad, 0x75, 0xf2, 0x7e, 0xde, 0xd6, 0x3b,
	0x4c, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x20, 0x35, 0xbf, 0xe6, 0x05, 0x00, 0x00,
}

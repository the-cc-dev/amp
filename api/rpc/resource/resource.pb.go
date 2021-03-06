// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/appcelerator/amp/api/rpc/resource/resource.proto

/*
Package resource is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/resource/resource.proto

It has these top-level messages:
	ResourceEntry
	ListRequest
	ListReply
	AddToTeamRequest
	RemoveFromTeamRequest
	ChangePermissionLevelRequest
*/
package resource

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import accounts "github.com/appcelerator/amp/data/accounts"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ResourceType int32

const (
	ResourceType_RESOURCE_STACK        ResourceType = 0
	ResourceType_RESOURCE_DASHBOARD    ResourceType = 1
	ResourceType_RESOURCE_USER         ResourceType = 2
	ResourceType_RESOURCE_ORGANIZATION ResourceType = 3
	ResourceType_RESOURCE_TEAM         ResourceType = 4
)

var ResourceType_name = map[int32]string{
	0: "RESOURCE_STACK",
	1: "RESOURCE_DASHBOARD",
	2: "RESOURCE_USER",
	3: "RESOURCE_ORGANIZATION",
	4: "RESOURCE_TEAM",
}
var ResourceType_value = map[string]int32{
	"RESOURCE_STACK":        0,
	"RESOURCE_DASHBOARD":    1,
	"RESOURCE_USER":         2,
	"RESOURCE_ORGANIZATION": 3,
	"RESOURCE_TEAM":         4,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Action int32

const (
	Action_ACTION_CREATE Action = 0
	Action_ACTION_READ   Action = 1
	Action_ACTION_UPDATE Action = 2
	Action_ACTION_DELETE Action = 3
)

var Action_name = map[int32]string{
	0: "ACTION_CREATE",
	1: "ACTION_READ",
	2: "ACTION_UPDATE",
	3: "ACTION_DELETE",
}
var Action_value = map[string]int32{
	"ACTION_CREATE": 0,
	"ACTION_READ":   1,
	"ACTION_UPDATE": 2,
	"ACTION_DELETE": 3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ResourceEntry struct {
	Id    string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type  ResourceType      `protobuf:"varint,2,opt,name=type,enum=resource.ResourceType" json:"type,omitempty"`
	Name  string            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Owner *accounts.Account `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
}

func (m *ResourceEntry) Reset()                    { *m = ResourceEntry{} }
func (m *ResourceEntry) String() string            { return proto.CompactTextString(m) }
func (*ResourceEntry) ProtoMessage()               {}
func (*ResourceEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceEntry) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_RESOURCE_STACK
}

func (m *ResourceEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceEntry) GetOwner() *accounts.Account {
	if m != nil {
		return m.Owner
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListReply struct {
	Resources []*ResourceEntry `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListReply) GetResources() []*ResourceEntry {
	if m != nil {
		return m.Resources
	}
	return nil
}

type AddToTeamRequest struct {
	ResourceId       string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	OrganizationName string `protobuf:"bytes,2,opt,name=organization_name,json=organizationName" json:"organization_name,omitempty"`
	TeamName         string `protobuf:"bytes,3,opt,name=team_name,json=teamName" json:"team_name,omitempty"`
}

func (m *AddToTeamRequest) Reset()                    { *m = AddToTeamRequest{} }
func (m *AddToTeamRequest) String() string            { return proto.CompactTextString(m) }
func (*AddToTeamRequest) ProtoMessage()               {}
func (*AddToTeamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddToTeamRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AddToTeamRequest) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *AddToTeamRequest) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

type RemoveFromTeamRequest struct {
	ResourceId       string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	OrganizationName string `protobuf:"bytes,2,opt,name=organization_name,json=organizationName" json:"organization_name,omitempty"`
	TeamName         string `protobuf:"bytes,3,opt,name=team_name,json=teamName" json:"team_name,omitempty"`
}

func (m *RemoveFromTeamRequest) Reset()                    { *m = RemoveFromTeamRequest{} }
func (m *RemoveFromTeamRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveFromTeamRequest) ProtoMessage()               {}
func (*RemoveFromTeamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveFromTeamRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *RemoveFromTeamRequest) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *RemoveFromTeamRequest) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

type ChangePermissionLevelRequest struct {
	OrganizationName string                       `protobuf:"bytes,1,opt,name=organization_name,json=organizationName" json:"organization_name,omitempty"`
	TeamName         string                       `protobuf:"bytes,2,opt,name=team_name,json=teamName" json:"team_name,omitempty"`
	ResourceId       string                       `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	PermissionLevel  accounts.TeamPermissionLevel `protobuf:"varint,4,opt,name=permission_level,json=permissionLevel,enum=accounts.TeamPermissionLevel" json:"permission_level,omitempty"`
}

func (m *ChangePermissionLevelRequest) Reset()                    { *m = ChangePermissionLevelRequest{} }
func (m *ChangePermissionLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangePermissionLevelRequest) ProtoMessage()               {}
func (*ChangePermissionLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChangePermissionLevelRequest) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *ChangePermissionLevelRequest) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *ChangePermissionLevelRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ChangePermissionLevelRequest) GetPermissionLevel() accounts.TeamPermissionLevel {
	if m != nil {
		return m.PermissionLevel
	}
	return accounts.TeamPermissionLevel_TEAM_READ
}

func init() {
	proto.RegisterType((*ResourceEntry)(nil), "resource.ResourceEntry")
	proto.RegisterType((*ListRequest)(nil), "resource.ListRequest")
	proto.RegisterType((*ListReply)(nil), "resource.ListReply")
	proto.RegisterType((*AddToTeamRequest)(nil), "resource.AddToTeamRequest")
	proto.RegisterType((*RemoveFromTeamRequest)(nil), "resource.RemoveFromTeamRequest")
	proto.RegisterType((*ChangePermissionLevelRequest)(nil), "resource.ChangePermissionLevelRequest")
	proto.RegisterEnum("resource.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("resource.Action", Action_name, Action_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resource service

type ResourceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	AddToTeam(ctx context.Context, in *AddToTeamRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ChangePermissionLevel(ctx context.Context, in *ChangePermissionLevelRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	RemoveFromTeam(ctx context.Context, in *RemoveFromTeamRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type resourceClient struct {
	cc *grpc.ClientConn
}

func NewResourceClient(cc *grpc.ClientConn) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/resource.Resource/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) AddToTeam(ctx context.Context, in *AddToTeamRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/resource.Resource/AddToTeam", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) ChangePermissionLevel(ctx context.Context, in *ChangePermissionLevelRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/resource.Resource/ChangePermissionLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) RemoveFromTeam(ctx context.Context, in *RemoveFromTeamRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/resource.Resource/RemoveFromTeam", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Resource service

type ResourceServer interface {
	List(context.Context, *ListRequest) (*ListReply, error)
	AddToTeam(context.Context, *AddToTeamRequest) (*google_protobuf.Empty, error)
	ChangePermissionLevel(context.Context, *ChangePermissionLevelRequest) (*google_protobuf.Empty, error)
	RemoveFromTeam(context.Context, *RemoveFromTeamRequest) (*google_protobuf.Empty, error)
}

func RegisterResourceServer(s *grpc.Server, srv ResourceServer) {
	s.RegisterService(&_Resource_serviceDesc, srv)
}

func _Resource_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_AddToTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).AddToTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/AddToTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).AddToTeam(ctx, req.(*AddToTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_ChangePermissionLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePermissionLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).ChangePermissionLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/ChangePermissionLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).ChangePermissionLevel(ctx, req.(*ChangePermissionLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_RemoveFromTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).RemoveFromTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/RemoveFromTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).RemoveFromTeam(ctx, req.(*RemoveFromTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Resource_List_Handler,
		},
		{
			MethodName: "AddToTeam",
			Handler:    _Resource_AddToTeam_Handler,
		},
		{
			MethodName: "ChangePermissionLevel",
			Handler:    _Resource_ChangePermissionLevel_Handler,
		},
		{
			MethodName: "RemoveFromTeam",
			Handler:    _Resource_RemoveFromTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/resource/resource.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/resource/resource.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xce, 0x18, 0x6e, 0x14, 0x0e, 0x17, 0xe2, 0xcc, 0x15, 0xb9, 0x5c, 0x6e, 0xaa, 0x20, 0x2f,
	0x5a, 0x44, 0x25, 0x46, 0x25, 0xaa, 0x54, 0xb5, 0x2b, 0x07, 0xdc, 0x06, 0x35, 0x0d, 0xe9, 0xe0,
	0xa8, 0x52, 0x36, 0x68, 0x02, 0x53, 0x62, 0x09, 0x7b, 0xa6, 0xb6, 0x49, 0x45, 0x11, 0x9b, 0xac,
	0xba, 0xef, 0xb2, 0xbb, 0xbe, 0x42, 0xdf, 0xa3, 0x9b, 0xbc, 0x42, 0x1f, 0xa4, 0xb2, 0xc1, 0xc6,
	0x90, 0x34, 0xed, 0xa2, 0x55, 0x37, 0x68, 0xfc, 0xcd, 0xf9, 0xf9, 0xbe, 0xe1, 0x3b, 0x07, 0x9e,
	0x0c, 0x2c, 0xff, 0x7c, 0x74, 0x56, 0xeb, 0x09, 0x9b, 0x30, 0x29, 0x7b, 0x7c, 0xc8, 0x5d, 0xe6,
	0x0b, 0x97, 0x30, 0x5b, 0x12, 0x26, 0x2d, 0xe2, 0xca, 0x1e, 0x71, 0xb9, 0x27, 0x46, 0x6e, 0x8f,
	0xc7, 0x87, 0x9a, 0x74, 0x85, 0x2f, 0xf0, 0x46, 0xf4, 0x5d, 0x7a, 0x74, 0x5b, 0x99, 0x3e, 0xf3,
	0x19, 0x61, 0xbd, 0x9e, 0x18, 0x39, 0xbe, 0x17, 0x1f, 0x66, 0x35, 0x4a, 0x7b, 0x89, 0xcc, 0x81,
	0x18, 0x32, 0x67, 0x40, 0xc2, 0x8b, 0xb3, 0xd1, 0x6b, 0x22, 0xfd, 0xb1, 0xe4, 0x1e, 0xe1, 0xb6,
	0xf4, 0xc7, 0xb3, 0xdf, 0x79, 0xd2, 0xce, 0x40, 0x88, 0xc1, 0x90, 0x87, 0x04, 0x99, 0xe3, 0x08,
	0x9f, 0xf9, 0x96, 0x70, 0xe6, 0x25, 0xb5, 0xf7, 0x08, 0x72, 0x74, 0xce, 0xcc, 0x70, 0x7c, 0x77,
	0x8c, 0xf3, 0xa0, 0x58, 0xfd, 0x22, 0x2a, 0xa3, 0x4a, 0x86, 0x2a, 0x56, 0x1f, 0x57, 0x21, 0x1d,
	0x94, 0x2e, 0x2a, 0x65, 0x54, 0xc9, 0xd7, 0xb7, 0x6b, 0xb1, 0xae, 0x28, 0xcd, 0x1c, 0x4b, 0x4e,
	0xc3, 0x18, 0x8c, 0x21, 0xed, 0x30, 0x9b, 0x17, 0x53, 0x61, 0x76, 0x78, 0xc6, 0xf7, 0xe0, 0x2f,
	0xf1, 0xd6, 0xe1, 0x6e, 0x31, 0x5d, 0x46, 0x95, 0x6c, 0x7d, 0xab, 0x16, 0x8b, 0xd2, 0x67, 0x07,
	0x3a, 0xbb, 0xd7, 0x72, 0x90, 0x3d, 0xb4, 0x3c, 0x9f, 0xf2, 0x37, 0x23, 0xee, 0xf9, 0xda, 0x3e,
	0x64, 0x66, 0x9f, 0x72, 0x38, 0xc6, 0x0f, 0x21, 0x13, 0xf5, 0xf5, 0x8a, 0xa8, 0x9c, 0xaa, 0x64,
	0xeb, 0xff, 0x5e, 0x67, 0x12, 0x0a, 0xa0, 0x8b, 0x48, 0x6d, 0x0a, 0xaa, 0xde, 0xef, 0x9b, 0xc2,
	0xe4, 0xcc, 0x9e, 0xd7, 0xc5, 0xbb, 0x90, 0x8d, 0x02, 0xba, 0xb1, 0x50, 0x88, 0xa0, 0x56, 0x1f,
	0xdf, 0x87, 0x2d, 0xe1, 0x0e, 0x98, 0x63, 0xbd, 0x0b, 0x5f, 0xaa, 0x1b, 0x2a, 0x52, 0xc2, 0x30,
	0x35, 0x79, 0x71, 0x14, 0xa8, 0xfb, 0x1f, 0x32, 0x3e, 0x67, 0x76, 0x37, 0x21, 0x7b, 0x23, 0x00,
	0x82, 0x4b, 0xed, 0x12, 0x41, 0x81, 0x72, 0x5b, 0x5c, 0xf0, 0xa7, 0xae, 0xb0, 0xff, 0x10, 0x89,
	0x2b, 0x04, 0x3b, 0x8d, 0x73, 0xe6, 0x0c, 0xf8, 0x31, 0x77, 0x6d, 0xcb, 0xf3, 0x2c, 0xe1, 0x1c,
	0xf2, 0x0b, 0x3e, 0x8c, 0xb8, 0xdc, 0xd8, 0x0a, 0xfd, 0x4c, 0x2b, 0x65, 0xb9, 0xd5, 0xaa, 0xaa,
	0xd4, 0x35, 0x55, 0x07, 0xa0, 0xca, 0x98, 0x44, 0x77, 0x18, 0xb0, 0x08, 0x6d, 0x91, 0xaf, 0xdf,
	0x59, 0xd8, 0x22, 0x78, 0xa7, 0x55, 0xaa, 0x9b, 0x72, 0x19, 0xa8, 0x4e, 0xe0, 0xef, 0xa4, 0xff,
	0x30, 0x86, 0x3c, 0x35, 0x3a, 0xed, 0x13, 0xda, 0x30, 0xba, 0x1d, 0x53, 0x6f, 0x3c, 0x57, 0xd7,
	0xf0, 0x36, 0xe0, 0x18, 0x6b, 0xea, 0x9d, 0x83, 0xfd, 0xb6, 0x4e, 0x9b, 0x2a, 0xc2, 0x5b, 0x90,
	0x8b, 0xf1, 0x93, 0x8e, 0x41, 0x55, 0x05, 0xff, 0x07, 0x85, 0x18, 0x6a, 0xd3, 0x67, 0xfa, 0x51,
	0xeb, 0x54, 0x37, 0x5b, 0xed, 0x23, 0x35, 0xb5, 0x14, 0x6d, 0x1a, 0xfa, 0x0b, 0x35, 0x5d, 0xa5,
	0xb0, 0xae, 0xf7, 0x82, 0x27, 0x09, 0x2e, 0xf5, 0x46, 0x10, 0xd8, 0x6d, 0x50, 0x43, 0x37, 0x0d,
	0x75, 0x0d, 0x6f, 0x42, 0x76, 0x0e, 0x51, 0x43, 0x9f, 0xb7, 0x9b, 0x03, 0x27, 0xc7, 0xcd, 0x20,
	0x46, 0x49, 0x40, 0x4d, 0xe3, 0xd0, 0x30, 0x0d, 0x35, 0x55, 0xff, 0x92, 0x86, 0x8d, 0x48, 0x11,
	0x6e, 0x41, 0x3a, 0xf0, 0x3e, 0x2e, 0x2c, 0x3c, 0x9e, 0x18, 0x8d, 0xd2, 0x3f, 0xab, 0xb0, 0x1c,
	0x8e, 0xb5, 0xc2, 0xe5, 0xd5, 0xd7, 0x0f, 0xca, 0x26, 0xce, 0x91, 0x8b, 0x07, 0xf1, 0xf2, 0xf1,
	0xf0, 0x47, 0x04, 0x99, 0x78, 0x06, 0x70, 0x69, 0x91, 0xb9, 0x3a, 0x18, 0xa5, 0xed, 0xda, 0x6c,
	0x53, 0xd4, 0xa2, 0x9d, 0x52, 0x33, 0x82, 0x35, 0xa2, 0x9d, 0x86, 0x85, 0x4d, 0xad, 0xbd, 0x54,
	0x98, 0x4c, 0x12, 0x7f, 0xf5, 0x94, 0x24, 0x8d, 0xe2, 0x91, 0xc9, 0x35, 0x43, 0x4d, 0x49, 0xe0,
	0x12, 0x8f, 0x4c, 0x62, 0xf7, 0x4c, 0x1f, 0xa3, 0x2a, 0xfe, 0x8c, 0xa0, 0x70, 0xa3, 0x39, 0xf1,
	0xdd, 0x05, 0xd3, 0xdb, 0xdc, 0xfb, 0x23, 0xd6, 0xa5, 0xdf, 0xc1, 0xfa, 0x13, 0x82, 0xfc, 0xf2,
	0x5c, 0xe3, 0xdd, 0xe4, 0x36, 0xba, 0x61, 0xe2, 0xbf, 0xcb, 0xf3, 0x55, 0xc8, 0xf3, 0x65, 0xf5,
	0x57, 0xf3, 0x3c, 0x5b, 0x0f, 0x1b, 0xed, 0x7d, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x8d, 0xe5,
	0x48, 0xb6, 0x06, 0x00, 0x00,
}

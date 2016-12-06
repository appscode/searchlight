// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	build.proto
	job.proto
	master.proto

It has these top-level messages:
	AgentListRequest
	AgentListResponse
	Agent
	AgentCreateRequest
	PortInfo
	AgentDescribeRequest
	AgentDescribeResponse
	AgentDeleteRequest
	AgentRestartRequest
	AgentRestartResponse
	BuildDescribeRequest
	BuildDescribeResponse
	BuildListRequest
	BuildListResponse
	Build
	JobListRequest
	JobListResponse
	JobBuildRequest
	JobDescribeRequest
	JobDescribeResponse
	JobHealth
	JobDeleteRequest
	JobCreateRequest
	JobCopyRequest
	Job
	MasterCreateRequest
	MasterDeleteRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type AgentListRequest struct {
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   ONLINE
	//   OFFLINE
	//   DELETED
	Status []string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *AgentListRequest) Reset()                    { *m = AgentListRequest{} }
func (m *AgentListRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentListRequest) ProtoMessage()               {}
func (*AgentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentListRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type AgentListResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Agents []*Agent                `protobuf:"bytes,2,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Agent struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	AgentStatus string `protobuf:"bytes,3,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed,json=isRefreshed" json:"is_refreshed,omitempty"`
	CreatedAt   int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Agent) GetAgentStatus() string {
	if m != nil {
		return m.AgentStatus
	}
	return ""
}

func (m *Agent) GetIsRefreshed() int32 {
	if m != nil {
		return m.IsRefreshed
	}
	return 0
}

func (m *Agent) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Agent) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type AgentCreateRequest struct {
	Sku               string      `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Executors         int32       `protobuf:"varint,2,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,4,opt,name=user_startup_script,json=userStartupScript" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,5,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,6,opt,name=ci_starter_version,json=ciStarterVersion" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,7,rep,name=ports" json:"ports,omitempty"`
	Role              string      `protobuf:"bytes,8,opt,name=role" json:"role,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentCreateRequest) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *AgentCreateRequest) GetExecutors() int32 {
	if m != nil {
		return m.Executors
	}
	return 0
}

func (m *AgentCreateRequest) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *AgentCreateRequest) GetUserStartupScript() string {
	if m != nil {
		return m.UserStartupScript
	}
	return ""
}

func (m *AgentCreateRequest) GetSaltbaseVersion() string {
	if m != nil {
		return m.SaltbaseVersion
	}
	return ""
}

func (m *AgentCreateRequest) GetCiStarterVersion() string {
	if m != nil {
		return m.CiStarterVersion
	}
	return ""
}

func (m *AgentCreateRequest) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *AgentCreateRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type PortInfo struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	PortRange string `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
}

func (m *PortInfo) Reset()                    { *m = PortInfo{} }
func (m *PortInfo) String() string            { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()               {}
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PortInfo) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PortInfo) GetPortRange() string {
	if m != nil {
		return m.PortRange
	}
	return ""
}

type AgentDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDescribeRequest) Reset()                    { *m = AgentDescribeRequest{} }
func (m *AgentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeRequest) ProtoMessage()               {}
func (*AgentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AgentDescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AgentDescribeResponse struct {
	Status           *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name             string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Executors        int64                   `protobuf:"varint,3,opt,name=executors" json:"executors,omitempty"`
	AgentStatus      string                  `protobuf:"bytes,4,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	AgentStatusCause string                  `protobuf:"bytes,5,opt,name=agent_status_cause,json=agentStatusCause" json:"agent_status_cause,omitempty"`
	Label            string                  `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	Provider         string                  `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Sku              string                  `protobuf:"bytes,8,opt,name=sku" json:"sku,omitempty"`
	StartupScript    string                  `protobuf:"bytes,9,opt,name=startup_script,json=startupScript" json:"startup_script,omitempty"`
	CreatedAt        int64                   `protobuf:"varint,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt        int64                   `protobuf:"varint,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *AgentDescribeResponse) Reset()                    { *m = AgentDescribeResponse{} }
func (m *AgentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeResponse) ProtoMessage()               {}
func (*AgentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AgentDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentDescribeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentDescribeResponse) GetExecutors() int64 {
	if m != nil {
		return m.Executors
	}
	return 0
}

func (m *AgentDescribeResponse) GetAgentStatus() string {
	if m != nil {
		return m.AgentStatus
	}
	return ""
}

func (m *AgentDescribeResponse) GetAgentStatusCause() string {
	if m != nil {
		return m.AgentStatusCause
	}
	return ""
}

func (m *AgentDescribeResponse) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AgentDescribeResponse) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *AgentDescribeResponse) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *AgentDescribeResponse) GetStartupScript() string {
	if m != nil {
		return m.StartupScript
	}
	return ""
}

func (m *AgentDescribeResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AgentDescribeResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type AgentDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AgentDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AgentRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentRestartRequest) Reset()                    { *m = AgentRestartRequest{} }
func (m *AgentRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartRequest) ProtoMessage()               {}
func (*AgentRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AgentRestartRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AgentRestartResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AgentRestartResponse) Reset()                    { *m = AgentRestartResponse{} }
func (m *AgentRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartResponse) ProtoMessage()               {}
func (*AgentRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AgentRestartResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentListRequest)(nil), "appscode.ci.v1beta1.AgentListRequest")
	proto.RegisterType((*AgentListResponse)(nil), "appscode.ci.v1beta1.AgentListResponse")
	proto.RegisterType((*Agent)(nil), "appscode.ci.v1beta1.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "appscode.ci.v1beta1.AgentCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "appscode.ci.v1beta1.PortInfo")
	proto.RegisterType((*AgentDescribeRequest)(nil), "appscode.ci.v1beta1.AgentDescribeRequest")
	proto.RegisterType((*AgentDescribeResponse)(nil), "appscode.ci.v1beta1.AgentDescribeResponse")
	proto.RegisterType((*AgentDeleteRequest)(nil), "appscode.ci.v1beta1.AgentDeleteRequest")
	proto.RegisterType((*AgentRestartRequest)(nil), "appscode.ci.v1beta1.AgentRestartRequest")
	proto.RegisterType((*AgentRestartResponse)(nil), "appscode.ci.v1beta1.AgentRestartResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
	Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error) {
	out := new(AgentDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error) {
	out := new(AgentRestartResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Agents/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *AgentListRequest) (*AgentListResponse, error)
	Describe(context.Context, *AgentDescribeRequest) (*AgentDescribeResponse, error)
	Create(context.Context, *AgentCreateRequest) (*appscode_dtypes.LongRunningResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*appscode_dtypes.LongRunningResponse, error)
	Restart(context.Context, *AgentRestartRequest) (*AgentRestartResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*AgentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Describe(ctx, req.(*AgentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Agents/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Restart(ctx, req.(*AgentRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Agents_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Agents_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 841 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0x97, 0xf3, 0xc7, 0x89, 0x5f, 0xf8, 0x93, 0xce, 0x16, 0xb0, 0x4c, 0x4b, 0x53, 0xab, 0x2d,
	0x4e, 0x54, 0x6c, 0x9a, 0x16, 0x21, 0x71, 0xdb, 0x6e, 0x11, 0x42, 0xea, 0x61, 0xe5, 0x48, 0x1c,
	0xb8, 0x44, 0x13, 0x67, 0x1a, 0x2c, 0x52, 0x8f, 0x99, 0x19, 0xaf, 0x16, 0x21, 0x2e, 0x2b, 0x2e,
	0x1c, 0x38, 0x71, 0xe0, 0x03, 0x70, 0xe7, 0x86, 0xf8, 0x20, 0x7c, 0x05, 0x6e, 0x5c, 0xf9, 0x00,
	0x68, 0xfe, 0xd8, 0xb1, 0xb3, 0x9b, 0x6c, 0xb4, 0x97, 0x68, 0xe6, 0xbd, 0xdf, 0xcc, 0xfb, 0xbd,
	0xf7, 0x7e, 0xf3, 0x62, 0x18, 0xe0, 0x15, 0xc9, 0x44, 0x98, 0x33, 0x2a, 0x28, 0x3a, 0xc2, 0x79,
	0xce, 0x13, 0xba, 0x24, 0x61, 0x92, 0x86, 0x67, 0x4f, 0x16, 0x44, 0xe0, 0x27, 0xde, 0x9d, 0x15,
	0xa5, 0xab, 0x35, 0x89, 0x70, 0x9e, 0x46, 0x38, 0xcb, 0xa8, 0xc0, 0x22, 0xa5, 0x19, 0xd7, 0x47,
	0xbc, 0x0f, 0xca, 0x23, 0x3b, 0xfc, 0xf7, 0x1a, 0xfe, 0xa5, 0xf8, 0x3e, 0x27, 0x3c, 0x52, 0xbf,
	0x1a, 0xe0, 0x4f, 0x60, 0x78, 0x2c, 0x29, 0xbc, 0x4c, 0xb9, 0x88, 0xc9, 0x77, 0x05, 0xe1, 0x02,
	0xbd, 0x0b, 0x36, 0x17, 0x58, 0x14, 0xdc, 0xb5, 0x46, 0xed, 0xc0, 0x89, 0xcd, 0xce, 0x3f, 0x87,
	0x5b, 0x35, 0x2c, 0xcf, 0x69, 0xc6, 0x09, 0x8a, 0x6a, 0x60, 0x2b, 0x18, 0x4c, 0xdf, 0x0b, 0xab,
	0x2c, 0x74, 0xb8, 0x70, 0xa6, 0xdc, 0xe5, 0x2d, 0x68, 0x0a, 0xb6, 0x4a, 0x9a, 0xbb, 0xad, 0x51,
	0x3b, 0x18, 0x4c, 0xbd, 0xf0, 0x8a, 0xb4, 0x43, 0x15, 0x28, 0x36, 0x48, 0xff, 0x2f, 0x0b, 0xba,
	0xca, 0x82, 0x10, 0x74, 0x32, 0xfc, 0x9a, 0xa8, 0x60, 0x4e, 0xac, 0xd6, 0x35, 0xbe, 0x2d, 0x65,
	0x2d, 0x23, 0xdd, 0x87, 0x37, 0xd4, 0xf9, 0xb9, 0xf1, 0xb6, 0x95, 0x57, 0x97, 0x7c, 0x56, 0x41,
	0x52, 0x3e, 0x67, 0xe4, 0x15, 0x23, 0xfc, 0x1b, 0xb2, 0x74, 0x3b, 0x23, 0x2b, 0xe8, 0xc6, 0x83,
	0x94, 0xc7, 0xa5, 0x09, 0xdd, 0x05, 0x48, 0x18, 0xc1, 0x82, 0x2c, 0xe7, 0x58, 0xb8, 0xdd, 0x91,
	0x15, 0xb4, 0x63, 0xc7, 0x58, 0x8e, 0x85, 0x74, 0x17, 0xf9, 0xb2, 0x74, 0xdb, 0xda, 0x6d, 0x2c,
	0xc7, 0xc2, 0xff, 0xa3, 0x05, 0x48, 0x31, 0x3f, 0x51, 0x27, 0xca, 0x12, 0x0f, 0xa1, 0xcd, 0xbf,
	0x2d, 0x4c, 0x16, 0x72, 0x89, 0xee, 0x80, 0x43, 0xce, 0x49, 0x52, 0x08, 0xca, 0x74, 0x1e, 0xdd,
	0x78, 0x63, 0x90, 0x29, 0xae, 0xf1, 0x82, 0xac, 0xcb, 0x24, 0xcc, 0x0e, 0x85, 0x70, 0x54, 0x70,
	0xc2, 0x64, 0x86, 0x4c, 0x14, 0xf9, 0x9c, 0x27, 0x2c, 0xcd, 0x85, 0x4a, 0xc3, 0x89, 0x6f, 0x49,
	0xd7, 0x4c, 0x7b, 0x66, 0xca, 0x81, 0xc6, 0x30, 0xe4, 0x78, 0x2d, 0x16, 0x98, 0x93, 0xf9, 0x19,
	0x61, 0x3c, 0xa5, 0x99, 0x4a, 0xc9, 0x89, 0xdf, 0x2e, 0xed, 0x5f, 0x69, 0x33, 0x7a, 0x0c, 0x28,
	0x49, 0xf5, 0xc5, 0x84, 0x55, 0x60, 0x5b, 0x81, 0x87, 0x49, 0x3a, 0xd3, 0x8e, 0x12, 0xfd, 0x14,
	0xba, 0x39, 0x65, 0x82, 0xbb, 0x3d, 0xd5, 0xd4, 0xbb, 0x57, 0x36, 0xf5, 0x94, 0x32, 0xf1, 0x65,
	0xf6, 0x8a, 0xc6, 0x1a, 0x2b, 0x9b, 0xc9, 0xe8, 0x9a, 0xb8, 0x7d, 0xdd, 0x4c, 0xb9, 0xf6, 0x3f,
	0x87, 0x7e, 0x09, 0x43, 0x1e, 0xf4, 0x95, 0x4a, 0x13, 0xba, 0x36, 0xa5, 0xaa, 0xf6, 0xb2, 0xee,
	0xf2, 0x92, 0x39, 0xc3, 0xd9, 0x8a, 0x98, 0xc6, 0x3b, 0xd2, 0x12, 0x4b, 0x83, 0x3f, 0x81, 0xdb,
	0xaa, 0xec, 0x2f, 0x88, 0x2c, 0xc9, 0xa2, 0x2a, 0xfc, 0x15, 0xfa, 0xf1, 0xff, 0x6b, 0xc1, 0x3b,
	0x5b, 0xe0, 0x9b, 0x8a, 0xbb, 0xbc, 0xbe, 0x55, 0x93, 0x67, 0xa3, 0xb3, 0x6d, 0x2d, 0x90, 0x4d,
	0x67, 0xb7, 0x45, 0xda, 0xb9, 0x2c, 0xd2, 0xc7, 0x80, 0xea, 0x90, 0x79, 0x82, 0x0b, 0x4e, 0x4c,
	0xdb, 0x86, 0x35, 0xe0, 0x89, 0xb4, 0xa3, 0xdb, 0xd0, 0x55, 0xe2, 0x30, 0xad, 0xd2, 0x1b, 0x53,
	0xca, 0xb3, 0x74, 0x49, 0x98, 0xdb, 0xab, 0x4a, 0xa9, 0xf6, 0xa5, 0x18, 0xfb, 0x1b, 0x31, 0x3e,
	0x84, 0xb7, 0xb6, 0x14, 0xe5, 0x28, 0xe7, 0x9b, 0xbc, 0xa1, 0xa6, 0xe6, 0xd3, 0x80, 0xfd, 0x4f,
	0x63, 0xb0, 0xfd, 0x34, 0x02, 0xf3, 0x32, 0x5e, 0x90, 0x35, 0x11, 0x7b, 0x1b, 0x34, 0x86, 0x23,
	0x3d, 0x0f, 0x88, 0x8a, 0xbf, 0x0f, 0xfa, 0x85, 0xe9, 0x7b, 0x05, 0xbd, 0x61, 0x27, 0xa7, 0xff,
	0x76, 0xc1, 0x56, 0x37, 0x71, 0xf4, 0x93, 0x05, 0x1d, 0x39, 0xf3, 0xd0, 0xc3, 0xdd, 0xa3, 0xaa,
	0x36, 0x3f, 0xbd, 0x47, 0xd7, 0xc1, 0x34, 0x27, 0xff, 0xa3, 0x8b, 0x3f, 0xdd, 0x56, 0xdf, 0xba,
	0xf8, 0xfb, 0x9f, 0x5f, 0x5b, 0xf7, 0xd1, 0xbd, 0xa8, 0x31, 0xab, 0x93, 0x34, 0x32, 0x27, 0x23,
	0x3d, 0x04, 0xd1, 0x6f, 0x16, 0xf4, 0x4b, 0x85, 0xa2, 0xf1, 0xee, 0x18, 0x5b, 0x92, 0xf7, 0x26,
	0x87, 0x40, 0x0d, 0xa5, 0x67, 0x35, 0x4a, 0x01, 0x7a, 0x74, 0x0d, 0xa5, 0xe8, 0x07, 0x59, 0xf3,
	0x1f, 0xd1, 0xcf, 0x16, 0xd8, 0x7a, 0xbe, 0xa1, 0x0f, 0x77, 0x07, 0x6b, 0x4c, 0x40, 0xef, 0xc1,
	0xa5, 0x06, 0xbc, 0xa4, 0xd9, 0x2a, 0x2e, 0xb2, 0x2c, 0xcd, 0x56, 0x15, 0x9f, 0x8f, 0x6b, 0x7c,
	0x1e, 0xf8, 0xd7, 0x95, 0xe8, 0x33, 0x6b, 0x82, 0x7e, 0xb1, 0xc0, 0xd6, 0x8a, 0xda, 0xc7, 0xa5,
	0xa1, 0xb9, 0x03, 0xb9, 0x34, 0x6a, 0x33, 0x39, 0xb4, 0x36, 0xbf, 0x5b, 0xd0, 0x33, 0x62, 0x44,
	0xc1, 0x9e, 0xbf, 0xba, 0x86, 0xb4, 0xbd, 0xf1, 0x01, 0x48, 0x43, 0xeb, 0xa4, 0x46, 0xeb, 0x53,
	0xff, 0x93, 0xc3, 0x68, 0x45, 0x38, 0x51, 0x9f, 0x09, 0x11, 0x23, 0x0b, 0x4a, 0xc5, 0xf3, 0x67,
	0xf0, 0x7e, 0x42, 0x5f, 0x6f, 0x82, 0xe2, 0x3c, 0xad, 0x05, 0x7e, 0x0e, 0x2a, 0xf2, 0xa9, 0x9c,
	0xbd, 0xa7, 0xd6, 0xd7, 0x3d, 0x63, 0x5e, 0xd8, 0x6a, 0x1a, 0x3f, 0xfd, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x70, 0x74, 0x2b, 0xc6, 0x08, 0x00, 0x00,
}

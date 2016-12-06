// Code generated by protoc-gen-go.
// source: artifact.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	artifact.proto
	version.proto

It has these top-level messages:
	ArtifactSearchRequest
	ArtifactSearchResponse
	ArtifactListRequest
	ArtifactListResponse
	Artifact
	VersionListRequest
	VersionListResponse
	VersionDescribeRequest
	VersionDescribeResponse
	ArtifactVersion
	JavaSpec
	DockerSpec
	PhpSpec
	NpmSpec
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

type ArtifactSearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactSearchRequest) Reset()                    { *m = ArtifactSearchRequest{} }
func (m *ArtifactSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchRequest) ProtoMessage()               {}
func (*ArtifactSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ArtifactSearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ArtifactSearchRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ArtifactSearchResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact             `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactSearchResponse) Reset()                    { *m = ArtifactSearchResponse{} }
func (m *ArtifactSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchResponse) ProtoMessage()               {}
func (*ArtifactSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArtifactSearchResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactSearchResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type ArtifactListRequest struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactListRequest) Reset()                    { *m = ArtifactListRequest{} }
func (m *ArtifactListRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListRequest) ProtoMessage()               {}
func (*ArtifactListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ArtifactListRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ArtifactListResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact             `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactListResponse) Reset()                    { *m = ArtifactListResponse{} }
func (m *ArtifactListResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListResponse) ProtoMessage()               {}
func (*ArtifactListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArtifactListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactListResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type Artifact struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	ModifiedAt int64  `protobuf:"varint,3,opt,name=modified_at,json=modifiedAt" json:"modified_at,omitempty"`
}

func (m *Artifact) Reset()                    { *m = Artifact{} }
func (m *Artifact) String() string            { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()               {}
func (*Artifact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Artifact) GetModifiedAt() int64 {
	if m != nil {
		return m.ModifiedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*ArtifactSearchRequest)(nil), "appscode.artifactory.v1beta1.ArtifactSearchRequest")
	proto.RegisterType((*ArtifactSearchResponse)(nil), "appscode.artifactory.v1beta1.ArtifactSearchResponse")
	proto.RegisterType((*ArtifactListRequest)(nil), "appscode.artifactory.v1beta1.ArtifactListRequest")
	proto.RegisterType((*ArtifactListResponse)(nil), "appscode.artifactory.v1beta1.ArtifactListResponse")
	proto.RegisterType((*Artifact)(nil), "appscode.artifactory.v1beta1.Artifact")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Artifacts service

type ArtifactsClient interface {
	Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error)
	List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error)
}

type artifactsClient struct {
	cc *grpc.ClientConn
}

func NewArtifactsClient(cc *grpc.ClientConn) ArtifactsClient {
	return &artifactsClient{cc}
}

func (c *artifactsClient) Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error) {
	out := new(ArtifactSearchResponse)
	err := grpc.Invoke(ctx, "/appscode.artifactory.v1beta1.Artifacts/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error) {
	out := new(ArtifactListResponse)
	err := grpc.Invoke(ctx, "/appscode.artifactory.v1beta1.Artifacts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Artifacts service

type ArtifactsServer interface {
	Search(context.Context, *ArtifactSearchRequest) (*ArtifactSearchResponse, error)
	List(context.Context, *ArtifactListRequest) (*ArtifactListResponse, error)
}

func RegisterArtifactsServer(s *grpc.Server, srv ArtifactsServer) {
	s.RegisterService(&_Artifacts_serviceDesc, srv)
}

func _Artifacts_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.artifactory.v1beta1.Artifacts/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).Search(ctx, req.(*ArtifactSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.artifactory.v1beta1.Artifacts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).List(ctx, req.(*ArtifactListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Artifacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.artifactory.v1beta1.Artifacts",
	HandlerType: (*ArtifactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Artifacts_Search_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Artifacts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "artifact.proto",
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x93, 0xbf, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0xb5, 0x86, 0x10, 0x18, 0x94, 0x14, 0x1b, 0x92, 0x58, 0x16, 0x0a, 0xc8, 0x8a, 0x22,
	0x27, 0x85, 0x57, 0x18, 0x94, 0x1a, 0x50, 0xca, 0x14, 0xc8, 0x74, 0x69, 0xa2, 0xc5, 0x5e, 0x88,
	0xa5, 0xe0, 0x35, 0xde, 0x25, 0x12, 0x8a, 0xd2, 0xd0, 0x47, 0x57, 0xdc, 0x63, 0x5c, 0x71, 0xd5,
	0x3d, 0xc9, 0xbd, 0xc2, 0x75, 0xf7, 0x12, 0x27, 0xff, 0x59, 0xb0, 0x4f, 0xdc, 0x1d, 0x54, 0xd7,
	0x58, 0xeb, 0x99, 0x9d, 0xf9, 0x7e, 0x3b, 0xdf, 0x2e, 0xbc, 0xa6, 0xb1, 0x0c, 0xe6, 0xd4, 0x93,
	0x76, 0x14, 0x73, 0xc9, 0x71, 0x9b, 0x46, 0x91, 0xf0, 0xb8, 0xcf, 0x6c, 0x95, 0xe0, 0xf1, 0xc6,
	0xfe, 0xd3, 0x9b, 0x31, 0x49, 0x7b, 0x46, 0x7b, 0xc1, 0xf9, 0xe2, 0x37, 0x23, 0x34, 0x0a, 0x08,
	0x0d, 0x43, 0x2e, 0xa9, 0x0c, 0x78, 0x28, 0xb2, 0x5a, 0xe3, 0x83, 0xaa, 0x7d, 0x20, 0xdf, 0x29,
	0xe5, 0x7d, 0xb9, 0x89, 0x98, 0x20, 0xe9, 0x37, 0xdb, 0x60, 0x8e, 0xe0, 0xed, 0x28, 0x57, 0x9d,
	0x32, 0x1a, 0x7b, 0xbf, 0x5c, 0xb6, 0x5a, 0x33, 0x21, 0x71, 0x0b, 0x5e, 0xac, 0xd6, 0x2c, 0xde,
	0xe8, 0xa8, 0x8b, 0xac, 0x86, 0x9b, 0xfd, 0x60, 0x0c, 0xd5, 0xa4, 0x5a, 0xd7, 0xd2, 0x60, 0xba,
	0x36, 0xcf, 0x10, 0xbc, 0xbb, 0xdf, 0x43, 0x44, 0x3c, 0x14, 0x0c, 0x13, 0xa8, 0x09, 0x49, 0xe5,
	0x5a, 0xa4, 0x5d, 0x9a, 0xce, 0x7b, 0x7b, 0x77, 0xd6, 0x8c, 0xc5, 0x9e, 0xa6, 0x69, 0x37, 0xdf,
	0x86, 0xbf, 0x41, 0x43, 0x0d, 0x41, 0xe8, 0x5a, 0xb7, 0x62, 0x35, 0x9d, 0x4f, 0xf6, 0x63, 0xf3,
	0xb1, 0x95, 0xb2, 0xbb, 0x2f, 0x34, 0x3f, 0xc3, 0x1b, 0x15, 0xfe, 0x1e, 0x08, 0xa9, 0x8e, 0xa4,
	0xe0, 0x51, 0x01, 0xfe, 0x3f, 0x82, 0x56, 0x79, 0xef, 0xf3, 0xa2, 0x4f, 0xa1, 0xae, 0xc2, 0x09,
	0x6f, 0x48, 0x97, 0x3b, 0xde, 0x64, 0x7d, 0xc8, 0x00, 0xdc, 0x81, 0xe6, 0x92, 0xfb, 0xc1, 0x3c,
	0x60, 0xfe, 0x4f, 0x2a, 0xf5, 0x4a, 0x17, 0x59, 0x15, 0x17, 0x54, 0x68, 0x24, 0x9d, 0x5b, 0x0d,
	0x1a, 0xaa, 0xab, 0xc0, 0x17, 0x08, 0x6a, 0x99, 0x4f, 0xb8, 0x7f, 0x1c, 0x60, 0xe9, 0x66, 0x18,
	0x83, 0xd3, 0x8a, 0xb2, 0x79, 0x9a, 0x5f, 0xb7, 0x57, 0xba, 0x56, 0x47, 0xdb, 0xeb, 0x9b, 0x73,
	0xed, 0x0b, 0xb6, 0x48, 0xf9, 0xe2, 0xee, 0xfb, 0x90, 0xbc, 0x0f, 0x11, 0x19, 0xe2, 0x25, 0x82,
	0x6a, 0x62, 0x0c, 0xee, 0x1d, 0x27, 0x5b, 0x30, 0xdc, 0x70, 0x4e, 0x29, 0xc9, 0x39, 0x87, 0x05,
	0xce, 0x01, 0x76, 0x9e, 0xe6, 0xdc, 0x79, 0x47, 0xfe, 0x26, 0x6e, 0xfc, 0x1b, 0x0f, 0xe1, 0xa3,
	0xc7, 0x97, 0x05, 0xe9, 0x28, 0x38, 0x24, 0x3f, 0x7e, 0xa5, 0xf4, 0x27, 0xc9, 0x4b, 0x9c, 0xa0,
	0x1f, 0x2f, 0xf3, 0xcc, 0xac, 0x96, 0xbe, 0xcd, 0xfe, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x37, 0xdf, 0x46, 0x2a, 0x04, 0x00, 0x00,
}

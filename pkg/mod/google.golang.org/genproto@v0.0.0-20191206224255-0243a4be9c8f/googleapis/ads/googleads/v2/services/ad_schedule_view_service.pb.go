// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_schedule_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [AdScheduleViewService.GetAdScheduleView][google.ads.googleads.v2.services.AdScheduleViewService.GetAdScheduleView].
type GetAdScheduleViewRequest struct {
	// The resource name of the ad schedule view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdScheduleViewRequest) Reset()         { *m = GetAdScheduleViewRequest{} }
func (m *GetAdScheduleViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdScheduleViewRequest) ProtoMessage()    {}
func (*GetAdScheduleViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0b48d75ffab38c, []int{0}
}

func (m *GetAdScheduleViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdScheduleViewRequest.Unmarshal(m, b)
}
func (m *GetAdScheduleViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdScheduleViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAdScheduleViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdScheduleViewRequest.Merge(m, src)
}
func (m *GetAdScheduleViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdScheduleViewRequest.Size(m)
}
func (m *GetAdScheduleViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdScheduleViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdScheduleViewRequest proto.InternalMessageInfo

func (m *GetAdScheduleViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdScheduleViewRequest)(nil), "google.ads.googleads.v2.services.GetAdScheduleViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_schedule_view_service.proto", fileDescriptor_fc0b48d75ffab38c)
}

var fileDescriptor_fc0b48d75ffab38c = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x26, 0x11, 0x04, 0x83, 0x2e, 0x0c, 0x88, 0x25, 0xba, 0x28, 0xb5, 0x0b, 0xe9, 0x62, 0x86,
	0xa6, 0x1b, 0x1d, 0x91, 0x92, 0x6e, 0xea, 0x4a, 0x4a, 0x0b, 0x59, 0x48, 0x20, 0x8c, 0x99, 0x43,
	0x0c, 0x24, 0x99, 0x9a, 0x93, 0xa4, 0x0b, 0x71, 0xa1, 0xaf, 0xe0, 0x1b, 0xb8, 0xf4, 0x1d, 0x7c,
	0x81, 0x6e, 0x7d, 0x05, 0x57, 0xae, 0x7d, 0x80, 0x4b, 0x3a, 0x99, 0xf4, 0x86, 0xdb, 0xd0, 0xdd,
	0xc7, 0x9c, 0xef, 0xe7, 0x9c, 0x2f, 0xb1, 0x96, 0xb1, 0x94, 0x71, 0x0a, 0x94, 0x0b, 0xa4, 0x0a,
	0x36, 0xa8, 0x76, 0x29, 0x42, 0x51, 0x27, 0x11, 0x20, 0xe5, 0x22, 0xc4, 0xe8, 0x13, 0x88, 0x2a,
	0x85, 0xb0, 0x4e, 0xe0, 0x10, 0xb6, 0x13, 0xb2, 0x2f, 0x64, 0x29, 0xed, 0xb1, 0x52, 0x11, 0x2e,
	0x90, 0x74, 0x06, 0xa4, 0x76, 0x89, 0x36, 0x70, 0x5e, 0x0d, 0x45, 0x14, 0x80, 0xb2, 0x2a, 0x2e,
	0x65, 0x28, 0x6f, 0xe7, 0xb9, 0x56, 0xee, 0x13, 0xca, 0xf3, 0x5c, 0x96, 0xbc, 0x4c, 0x64, 0x8e,
	0xed, 0xf4, 0xe9, 0xad, 0x69, 0x94, 0x26, 0x90, 0x97, 0x6a, 0x30, 0x59, 0x5a, 0xa3, 0x35, 0x94,
	0x9e, 0xd8, 0xb5, 0x96, 0x7e, 0x02, 0x87, 0x2d, 0x7c, 0xae, 0x00, 0x4b, 0xfb, 0x85, 0xf5, 0x48,
	0xc7, 0x86, 0x39, 0xcf, 0x60, 0x64, 0x8c, 0x8d, 0x97, 0x0f, 0xb6, 0x0f, 0xf5, 0xe3, 0x7b, 0x9e,
	0x81, 0xfb, 0xdf, 0xb0, 0x9e, 0xf4, 0xe5, 0x3b, 0x75, 0x8c, 0xfd, 0xdb, 0xb0, 0x1e, 0xdf, 0xf1,
	0xb6, 0x19, 0xb9, 0x56, 0x02, 0x19, 0x5a, 0xc8, 0x99, 0x0f, 0x6a, 0xbb, 0x7a, 0x48, 0x5f, 0x39,
	0x79, 0xfd, 0xfd, 0xcf, 0xdf, 0x1f, 0xe6, 0xc2, 0x9e, 0x37, 0x25, 0x7e, 0xe9, 0x9d, 0xf3, 0x36,
	0xaa, 0xb0, 0x94, 0x19, 0x14, 0x48, 0x67, 0x94, 0xf7, 0x64, 0x48, 0x67, 0x5f, 0x9d, 0x67, 0x47,
	0x6f, 0x74, 0x0e, 0x69, 0xd1, 0x3e, 0x41, 0x12, 0xc9, 0x6c, 0xf5, 0xcd, 0xb4, 0xa6, 0x91, 0xcc,
	0xae, 0x1e, 0xb3, 0x72, 0x2e, 0x96, 0xb3, 0x69, 0xca, 0xdf, 0x18, 0x1f, 0xde, 0xb5, 0xfa, 0x58,
	0xa6, 0x3c, 0x8f, 0x89, 0x2c, 0x62, 0x1a, 0x43, 0x7e, 0xfa, 0x34, 0xf4, 0x9c, 0x38, 0xfc, 0xc7,
	0xbd, 0xd1, 0xe0, 0xa7, 0x79, 0x6f, 0xed, 0x79, 0xbf, 0xcc, 0xf1, 0x5a, 0x19, 0x7a, 0x02, 0x89,
	0x82, 0x0d, 0xf2, 0x5d, 0xd2, 0x06, 0xe3, 0x51, 0x53, 0x02, 0x4f, 0x60, 0xd0, 0x51, 0x02, 0xdf,
	0x0d, 0x34, 0xe5, 0x9f, 0x39, 0x55, 0xef, 0x8c, 0x79, 0x02, 0x19, 0xeb, 0x48, 0x8c, 0xf9, 0x2e,
	0x63, 0x9a, 0xf6, 0xf1, 0xfe, 0x69, 0xcf, 0xc5, 0x4d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0xf6,
	0x94, 0x3b, 0x18, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdScheduleViewServiceClient is the client API for AdScheduleViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdScheduleViewServiceClient interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error)
}

type adScheduleViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdScheduleViewServiceClient(cc *grpc.ClientConn) AdScheduleViewServiceClient {
	return &adScheduleViewServiceClient{cc}
}

func (c *adScheduleViewServiceClient) GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error) {
	out := new(resources.AdScheduleView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdScheduleViewService/GetAdScheduleView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdScheduleViewServiceServer is the server API for AdScheduleViewService service.
type AdScheduleViewServiceServer interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(context.Context, *GetAdScheduleViewRequest) (*resources.AdScheduleView, error)
}

// UnimplementedAdScheduleViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdScheduleViewServiceServer struct {
}

func (*UnimplementedAdScheduleViewServiceServer) GetAdScheduleView(ctx context.Context, req *GetAdScheduleViewRequest) (*resources.AdScheduleView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdScheduleView not implemented")
}

func RegisterAdScheduleViewServiceServer(s *grpc.Server, srv AdScheduleViewServiceServer) {
	s.RegisterService(&_AdScheduleViewService_serviceDesc, srv)
}

func _AdScheduleViewService_GetAdScheduleView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdScheduleViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdScheduleViewService/GetAdScheduleView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, req.(*GetAdScheduleViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdScheduleViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdScheduleViewService",
	HandlerType: (*AdScheduleViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdScheduleView",
			Handler:    _AdScheduleViewService_GetAdScheduleView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_schedule_view_service.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package feed

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	CreatFeed(ctx context.Context, in *CreatFeedRequest, opts ...grpc.CallOption) (*CreatFeedResponse, error)
	GetProfileFeed(ctx context.Context, in *GetProfileFeedRequest, opts ...grpc.CallOption) (*GetProfileFeedResponse, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) CreatFeed(ctx context.Context, in *CreatFeedRequest, opts ...grpc.CallOption) (*CreatFeedResponse, error) {
	out := new(CreatFeedResponse)
	err := c.cc.Invoke(ctx, "/feed.FeedService/CreatFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) GetProfileFeed(ctx context.Context, in *GetProfileFeedRequest, opts ...grpc.CallOption) (*GetProfileFeedResponse, error) {
	out := new(GetProfileFeedResponse)
	err := c.cc.Invoke(ctx, "/feed.FeedService/GetProfileFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations must embed UnimplementedFeedServiceServer
// for forward compatibility
type FeedServiceServer interface {
	CreatFeed(context.Context, *CreatFeedRequest) (*CreatFeedResponse, error)
	GetProfileFeed(context.Context, *GetProfileFeedRequest) (*GetProfileFeedResponse, error)
	mustEmbedUnimplementedFeedServiceServer()
}

// UnimplementedFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (UnimplementedFeedServiceServer) CreatFeed(context.Context, *CreatFeedRequest) (*CreatFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatFeed not implemented")
}
func (UnimplementedFeedServiceServer) GetProfileFeed(context.Context, *GetProfileFeedRequest) (*GetProfileFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileFeed not implemented")
}
func (UnimplementedFeedServiceServer) mustEmbedUnimplementedFeedServiceServer() {}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_CreatFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).CreatFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.FeedService/CreatFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).CreatFeed(ctx, req.(*CreatFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_GetProfileFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetProfileFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.FeedService/GetProfileFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetProfileFeed(ctx, req.(*GetProfileFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatFeed",
			Handler:    _FeedService_CreatFeed_Handler,
		},
		{
			MethodName: "GetProfileFeed",
			Handler:    _FeedService_GetProfileFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed.proto",
}

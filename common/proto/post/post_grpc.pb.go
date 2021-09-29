// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package post

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

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatPost(ctx context.Context, in *CreatPostRequest, opts ...grpc.CallOption) (*CreatPostResponse, error)
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error)
	GetPostsByProfileId(ctx context.Context, in *GetPostsByProfileIdRequest, opts ...grpc.CallOption) (*GetPostsByProfileIdResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatPost(ctx context.Context, in *CreatPostRequest, opts ...grpc.CallOption) (*CreatPostResponse, error) {
	out := new(CreatPostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreatPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error) {
	out := new(GetPostByIdResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetPostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostsByProfileId(ctx context.Context, in *GetPostsByProfileIdRequest, opts ...grpc.CallOption) (*GetPostsByProfileIdResponse, error) {
	out := new(GetPostsByProfileIdResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetPostsByProfileId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	CreatPost(context.Context, *CreatPostRequest) (*CreatPostResponse, error)
	GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error)
	GetPostsByProfileId(context.Context, *GetPostsByProfileIdRequest) (*GetPostsByProfileIdResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) CreatPost(context.Context, *CreatPostRequest) (*CreatPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatPost not implemented")
}
func (UnimplementedPostServiceServer) GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedPostServiceServer) GetPostsByProfileId(context.Context, *GetPostsByProfileIdRequest) (*GetPostsByProfileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByProfileId not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreatPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatPost(ctx, req.(*CreatPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetPostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostsByProfileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsByProfileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostsByProfileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetPostsByProfileId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostsByProfileId(ctx, req.(*GetPostsByProfileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatPost",
			Handler:    _PostService_CreatPost_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _PostService_GetPostById_Handler,
		},
		{
			MethodName: "GetPostsByProfileId",
			Handler:    _PostService_GetPostsByProfileId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}

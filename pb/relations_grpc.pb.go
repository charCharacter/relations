// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: relations.proto

package relations

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RelationsServiceClient is the client API for RelationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationsServiceClient interface {
	RelationsCreateKeywords(ctx context.Context, in *RelationsCreateKeywordsRequest, opts ...grpc.CallOption) (*RelationsCreateKeywordsResponse, error)
	ArticleCreate(ctx context.Context, in *ArticleCreateRequest, opts ...grpc.CallOption) (*ArticleCreateResponse, error)
	ArticleFind(ctx context.Context, in *ArticleFindRequest, opts ...grpc.CallOption) (*ArticleFindResponse, error)
	ArticleFindAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ArticleFindAllResponse, error)
	RelationsList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RelationsListResponse, error)
}

type relationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationsServiceClient(cc grpc.ClientConnInterface) RelationsServiceClient {
	return &relationsServiceClient{cc}
}

func (c *relationsServiceClient) RelationsCreateKeywords(ctx context.Context, in *RelationsCreateKeywordsRequest, opts ...grpc.CallOption) (*RelationsCreateKeywordsResponse, error) {
	out := new(RelationsCreateKeywordsResponse)
	err := c.cc.Invoke(ctx, "/relations.RelationsService/RelationsCreateKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationsServiceClient) ArticleCreate(ctx context.Context, in *ArticleCreateRequest, opts ...grpc.CallOption) (*ArticleCreateResponse, error) {
	out := new(ArticleCreateResponse)
	err := c.cc.Invoke(ctx, "/relations.RelationsService/ArticleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationsServiceClient) ArticleFind(ctx context.Context, in *ArticleFindRequest, opts ...grpc.CallOption) (*ArticleFindResponse, error) {
	out := new(ArticleFindResponse)
	err := c.cc.Invoke(ctx, "/relations.RelationsService/ArticleFind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationsServiceClient) ArticleFindAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ArticleFindAllResponse, error) {
	out := new(ArticleFindAllResponse)
	err := c.cc.Invoke(ctx, "/relations.RelationsService/ArticleFindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationsServiceClient) RelationsList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RelationsListResponse, error) {
	out := new(RelationsListResponse)
	err := c.cc.Invoke(ctx, "/relations.RelationsService/RelationsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationsServiceServer is the server API for RelationsService service.
// All implementations must embed UnimplementedRelationsServiceServer
// for forward compatibility
type RelationsServiceServer interface {
	RelationsCreateKeywords(context.Context, *RelationsCreateKeywordsRequest) (*RelationsCreateKeywordsResponse, error)
	ArticleCreate(context.Context, *ArticleCreateRequest) (*ArticleCreateResponse, error)
	ArticleFind(context.Context, *ArticleFindRequest) (*ArticleFindResponse, error)
	ArticleFindAll(context.Context, *emptypb.Empty) (*ArticleFindAllResponse, error)
	RelationsList(context.Context, *emptypb.Empty) (*RelationsListResponse, error)
	mustEmbedUnimplementedRelationsServiceServer()
}

// UnimplementedRelationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationsServiceServer struct {
}

func (UnimplementedRelationsServiceServer) RelationsCreateKeywords(context.Context, *RelationsCreateKeywordsRequest) (*RelationsCreateKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationsCreateKeywords not implemented")
}
func (UnimplementedRelationsServiceServer) ArticleCreate(context.Context, *ArticleCreateRequest) (*ArticleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleCreate not implemented")
}
func (UnimplementedRelationsServiceServer) ArticleFind(context.Context, *ArticleFindRequest) (*ArticleFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleFind not implemented")
}
func (UnimplementedRelationsServiceServer) ArticleFindAll(context.Context, *emptypb.Empty) (*ArticleFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleFindAll not implemented")
}
func (UnimplementedRelationsServiceServer) RelationsList(context.Context, *emptypb.Empty) (*RelationsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationsList not implemented")
}
func (UnimplementedRelationsServiceServer) mustEmbedUnimplementedRelationsServiceServer() {}

// UnsafeRelationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationsServiceServer will
// result in compilation errors.
type UnsafeRelationsServiceServer interface {
	mustEmbedUnimplementedRelationsServiceServer()
}

func RegisterRelationsServiceServer(s grpc.ServiceRegistrar, srv RelationsServiceServer) {
	s.RegisterService(&RelationsService_ServiceDesc, srv)
}

func _RelationsService_RelationsCreateKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationsCreateKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).RelationsCreateKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relations.RelationsService/RelationsCreateKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).RelationsCreateKeywords(ctx, req.(*RelationsCreateKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationsService_ArticleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).ArticleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relations.RelationsService/ArticleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).ArticleCreate(ctx, req.(*ArticleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationsService_ArticleFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).ArticleFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relations.RelationsService/ArticleFind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).ArticleFind(ctx, req.(*ArticleFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationsService_ArticleFindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).ArticleFindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relations.RelationsService/ArticleFindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).ArticleFindAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationsService_RelationsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationsServiceServer).RelationsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relations.RelationsService/RelationsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationsServiceServer).RelationsList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationsService_ServiceDesc is the grpc.ServiceDesc for RelationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relations.RelationsService",
	HandlerType: (*RelationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationsCreateKeywords",
			Handler:    _RelationsService_RelationsCreateKeywords_Handler,
		},
		{
			MethodName: "ArticleCreate",
			Handler:    _RelationsService_ArticleCreate_Handler,
		},
		{
			MethodName: "ArticleFind",
			Handler:    _RelationsService_ArticleFind_Handler,
		},
		{
			MethodName: "ArticleFindAll",
			Handler:    _RelationsService_ArticleFindAll_Handler,
		},
		{
			MethodName: "RelationsList",
			Handler:    _RelationsService_RelationsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relations.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: product_image.proto

package product_image

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

// ProductImageServiceClient is the client API for ProductImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductImageServiceClient interface {
	GetImagesByProductIds(ctx context.Context, in *GetImagesRequest, opts ...grpc.CallOption) (*GetImagesResponse, error)
}

type productImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductImageServiceClient(cc grpc.ClientConnInterface) ProductImageServiceClient {
	return &productImageServiceClient{cc}
}

func (c *productImageServiceClient) GetImagesByProductIds(ctx context.Context, in *GetImagesRequest, opts ...grpc.CallOption) (*GetImagesResponse, error) {
	out := new(GetImagesResponse)
	err := c.cc.Invoke(ctx, "/product_image.ProductImageService/GetImagesByProductIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductImageServiceServer is the server API for ProductImageService service.
// All implementations must embed UnimplementedProductImageServiceServer
// for forward compatibility
type ProductImageServiceServer interface {
	GetImagesByProductIds(context.Context, *GetImagesRequest) (*GetImagesResponse, error)
	mustEmbedUnimplementedProductImageServiceServer()
}

// UnimplementedProductImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductImageServiceServer struct {
}

func (UnimplementedProductImageServiceServer) GetImagesByProductIds(context.Context, *GetImagesRequest) (*GetImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImagesByProductIds not implemented")
}
func (UnimplementedProductImageServiceServer) mustEmbedUnimplementedProductImageServiceServer() {}

// UnsafeProductImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductImageServiceServer will
// result in compilation errors.
type UnsafeProductImageServiceServer interface {
	mustEmbedUnimplementedProductImageServiceServer()
}

func RegisterProductImageServiceServer(s grpc.ServiceRegistrar, srv ProductImageServiceServer) {
	s.RegisterService(&ProductImageService_ServiceDesc, srv)
}

func _ProductImageService_GetImagesByProductIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductImageServiceServer).GetImagesByProductIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_image.ProductImageService/GetImagesByProductIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductImageServiceServer).GetImagesByProductIds(ctx, req.(*GetImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductImageService_ServiceDesc is the grpc.ServiceDesc for ProductImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_image.ProductImageService",
	HandlerType: (*ProductImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImagesByProductIds",
			Handler:    _ProductImageService_GetImagesByProductIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_image.proto",
}

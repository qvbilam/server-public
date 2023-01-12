// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: image.proto

package publicV1

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

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	Create(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *SearchImageRequest, opts ...grpc.CallOption) (*ImagesResponse, error)
	GetDetail(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	Exists(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ExistsImageResponse, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) Create(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) Delete(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) Get(ctx context.Context, in *SearchImageRequest, opts ...grpc.CallOption) (*ImagesResponse, error) {
	out := new(ImagesResponse)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetDetail(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) Exists(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*ExistsImageResponse, error) {
	out := new(ExistsImageResponse)
	err := c.cc.Invoke(ctx, "/publicPb.v1.Image/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	Create(context.Context, *UpdateImageRequest) (*ImageResponse, error)
	Update(context.Context, *UpdateImageRequest) (*emptypb.Empty, error)
	Delete(context.Context, *UpdateImageRequest) (*emptypb.Empty, error)
	Get(context.Context, *SearchImageRequest) (*ImagesResponse, error)
	GetDetail(context.Context, *GetImageRequest) (*ImageResponse, error)
	Exists(context.Context, *GetImageRequest) (*ExistsImageResponse, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) Create(context.Context, *UpdateImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedImageServer) Update(context.Context, *UpdateImageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedImageServer) Delete(context.Context, *UpdateImageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedImageServer) Get(context.Context, *SearchImageRequest) (*ImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedImageServer) GetDetail(context.Context, *GetImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedImageServer) Exists(context.Context, *GetImageRequest) (*ExistsImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Create(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Update(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Delete(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Get(ctx, req.(*SearchImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetDetail(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicPb.v1.Image/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).Exists(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publicPb.v1.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Image_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Image_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Image_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Image_Get_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _Image_GetDetail_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _Image_Exists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}

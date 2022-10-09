package api

import (
	"context"
	proto "file/api/qvbilam/file/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ImageServer struct {
	proto.UnimplementedImageServer
}

func (s *ImageServer) Create(ctx context.Context, request *proto.UpdateImageRequest) (*proto.ImagesResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *ImageServer) Update(ctx context.Context, request *proto.UpdateImageRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *ImageServer) Delete(ctx context.Context, request *proto.UpdateImageRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *ImageServer) Get(ctx context.Context, request *proto.SearchImageRequest) (*proto.ImagesResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *ImageServer) GetDetail(ctx context.Context, request *proto.GetImageRequest) (*proto.ImageResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

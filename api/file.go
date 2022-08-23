package api

import (
	"context"
	proto "file/api/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileServer struct {
	proto.UnimplementedFileServer
}

func (s *FileServer) Exists(ctx context.Context, request *proto.GetRequest) (*proto.FileResponse, error) {
	return nil, nil
}

func (s *FileServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.FileResponse, error) {
	return nil, nil
}

func (s *FileServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.FileResponse, error) {
	return nil, nil
}

func (s *FileServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.FileResponse, error) {
	return nil, nil
}

func (s *FileServer) Delete(ctx context.Context, request *proto.UpdateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

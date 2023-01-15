package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "public/api/qvbilam/public/v1"
	"public/business"
)

type FileServer struct {
	proto.UnimplementedFileServer
}

func (s *FileServer) Create(ctx context.Context, request *proto.UpdateFileRequest) (*proto.FileResponse, error) {
	b := business.FileBusiness{
		Sha1:        request.Sha1,
		Url:         request.Url,
		ContentType: request.ContentType,
		Extra:       request.Extra,
		Channel:     request.Channel,
	}

	file, err := b.Create()
	if err != nil {
		return nil, err
	}
	return &proto.FileResponse{
		Id:          file.ID,
		Channel:     file.Channel,
		Sha1:        file.Sha1,
		Url:         file.Url,
		ContentType: file.ContentType,
		Size:        file.Size,
		Extra:       file.Extra,
		Callback:    file.Callback,
	}, nil
}

func (s *FileServer) Update(ctx context.Context, request *proto.UpdateFileRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *FileServer) Delete(ctx context.Context, request *proto.UpdateFileRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *FileServer) Get(ctx context.Context, request *proto.SearchFileRequest) (*proto.FilesResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *FileServer) GetDetail(ctx context.Context, request *proto.FileDetailRequest) (*proto.FileResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *FileServer) Exists(ctx context.Context, request *proto.FileDetailRequest) (*proto.ExistsFileResponse, error) {
	b := business.FileBusiness{
		Sha1: request.Sha1,
	}
	f := b.GetBySha1()
	if f == nil {
		return &proto.ExistsFileResponse{IsExists: false, File: nil}, nil
	}

	return &proto.ExistsFileResponse{
		IsExists: true,
		File: &proto.FileResponse{
			Id:          f.ID,
			Channel:     f.Channel,
			Sha1:        f.Sha1,
			Url:         f.Url,
			ContentType: f.ContentType,
			Size:        f.Size,
			Extra:       f.Extra,
		},
	}, nil
}

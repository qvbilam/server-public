package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "public/api/qvbilam/public/v1"
	"public/business"
)

type VideoServer struct {
	proto.UnimplementedVideoServer
}

func (s *VideoServer) Exists(ctx context.Context, request *proto.GetVideoRequest) (*proto.ExistsVideoResponse, error) {
	if request.Id == 0 && request.BusinessId == "" && request.FileSha1 == "" {
		return nil, status.Errorf(codes.InvalidArgument, "参数错误")
	}

	b := business.VideoBusiness{
		ID:         request.Id,
		BusinessID: request.BusinessId,
		Sha1:       request.FileSha1,
	}

	entity := b.GetDetail()

	video := &proto.VideoResponse{}
	isExists := false
	if entity != nil {
		isExists = true
		video = &proto.VideoResponse{
			Id:         entity.ID,
			BusinessId: entity.BusinessId,
			Url:        entity.Url,
			Status:     entity.Status,
			Channel:    entity.Channel,
		}
	}

	return &proto.ExistsVideoResponse{
		IsExists: isExists,
		Video:    video,
	}, nil
}

func (s *VideoServer) Create(ctx context.Context, request *proto.UpdateVideoRequest) (*proto.VideoResponse, error) {
	b := business.VideoBusiness{
		UserID:      request.UserId,
		BusinessID:  request.BusinessId,
		Sha1:        request.Sha1,
		Url:         request.Url,
		Size:        request.Size,
		Duration:    request.Duration,
		Status:      request.Status,
		ContentType: request.ContentType,
		Extra:       request.Extra,
		Channel:     request.Channel,
	}

	entity, err := b.Create()
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		return nil, status.Errorf(st.Code(), st.Message())
	}

	return &proto.VideoResponse{
		Id:         entity.ID,
		BusinessId: entity.BusinessId,
		Url:        entity.Url,
		Channel:    entity.Channel,
	}, nil
}

func (s *VideoServer) Update(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *VideoServer) Delete(ctx context.Context, request *proto.UpdateVideoRequest) (*emptypb.Empty, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *VideoServer) Get(ctx context.Context, request *proto.SearchVideoRequest) (*proto.VideosResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unavailable, "服务不可用")
}

func (s *VideoServer) GetDetail(ctx context.Context, request *proto.GetVideoRequest) (*proto.VideoResponse, error) {
	//TODO implement me
	b := business.VideoBusiness{
		ID:         request.Id,
		BusinessID: request.BusinessId,
		Sha1:       request.FileSha1,
	}
	entity := b.GetDetail()
	if entity == nil {
		return nil, status.Errorf(codes.NotFound, "文件不存在")
	}

	return &proto.VideoResponse{
		Id:         entity.ID,
		BusinessId: entity.BusinessId,
		Url:        entity.Url,
		Channel:    entity.Channel,
	}, nil
}

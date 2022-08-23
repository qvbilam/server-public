package api

import (
	"context"
	proto "file/api/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConfigServer struct {
	proto.UnimplementedConfigServer
}

func (*ConfigServer) GetConfig(ctx context.Context, request *proto.UpdateConfigRequest) (*proto.ConfigResponse, error) {
	return nil, nil
}

func (*ConfigServer) CreateConfig(ctx context.Context, request *proto.UpdateConfigRequest) (*proto.ConfigResponse, error) {
	return nil, nil
}

func (*ConfigServer) UpdateConfig(ctx context.Context, request *proto.UpdateConfigRequest) (*proto.ConfigResponse, error) {
	return nil, nil
}

func (*ConfigServer) DeleteConfig(ctx context.Context, request *proto.UpdateConfigRequest) (*emptypb.Empty, error) {
	return nil, nil
}

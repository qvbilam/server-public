package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "public/api/qvbilam/public/v1"
)

type SmsServer struct {
	proto.UnimplementedSmsServer
}

func (s *SmsServer) Send(ctx context.Context, request *proto.SendSmsRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *SmsServer) Check(ctx context.Context, request *proto.CheckSmsRequest) (*emptypb.Empty, error) {
	return nil, nil
}

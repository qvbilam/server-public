package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "public/api/qvbilam/public/v1"
	"public/business"
)

type SmsServer struct {
	proto.UnimplementedSmsServer
}

func (s *SmsServer) Send(ctx context.Context, request *proto.SendSmsRequest) (*emptypb.Empty, error) {
	b := business.SmsBusiness{
		Mobile:   request.Mobile,
		Type:     request.Type,
		ClientIP: request.ClientIP,
	}

	if err := b.Send(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *SmsServer) Check(ctx context.Context, request *proto.CheckSmsRequest) (*emptypb.Empty, error) {
	b := business.SmsBusiness{
		Mobile:    request.Mobile,
		Type:      request.Type,
		CheckCode: request.Code,
	}
	if err := b.Check(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

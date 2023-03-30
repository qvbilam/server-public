package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "public/api/qvbilam/public/v1"
	"public/business"
	"public/enum"
)

type SmsServer struct {
	proto.UnimplementedSmsServer
}

func (s *SmsServer) SendLogin(ctx context.Context, request *proto.SendSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypeLogin
	return s.Send(ctx, request)
}

func (s *SmsServer) SendLogout(ctx context.Context, request *proto.SendSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypeLogout
	return s.Send(ctx, request)
}

func (s *SmsServer) SendPassword(ctx context.Context, request *proto.SendSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypePassword
	return s.Send(ctx, request)
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

func (s *SmsServer) CheckLogin(ctx context.Context, request *proto.CheckSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypeLogin
	return s.Check(ctx, request)
}

func (s *SmsServer) CheckLogout(ctx context.Context, request *proto.CheckSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypeLogout
	return s.Check(ctx, request)
}

func (s *SmsServer) CheckPassword(ctx context.Context, request *proto.CheckSmsRequest) (*emptypb.Empty, error) {
	request.Type = enum.SmsTypePassword
	return s.Check(ctx, request)
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

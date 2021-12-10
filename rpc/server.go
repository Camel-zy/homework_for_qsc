package rpc

import (
	"context"

	// SMSService "git.zjuqsc.com/rop/rop-sms/gRPC"
	SMSService "gRPC"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
	connection *grpc.ClientConn
	client     SMSService.SMSClient
}

func (s *server) Init() {
	var err error
	s.connection, err = grpc.DialContext(context.Background(), viper.GetString("rpc.endpoint"), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeout)) // TODO(TO/GA): remove insecure
	if err != nil {
		panic(err)
	}
	s.client = SMSService.NewSMSClient(s.connection)
}

func (s *server) Ping(in *SMSService.PingReq) (*SMSService.PingReply, error) {
	if !enable {
		return nil, ErrSMSUninitialized
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	r, err := s.client.Ping(ctx, in)
	return r, err
}

func (s *server) SendMsgByText(in *SMSService.MsgReq) (*SMSService.MsgReply, error) {
	if !enable {
		return nil, ErrSMSUninitialized
	}
	in.UserID = uint32(app_id)
	in.UserKey = app_key
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	r, err := s.client.SendMsgByText(ctx, in)
	return r, err
}

func (s *server) UserBalance(in *SMSService.UsrReq) (*SMSService.UsrReply, error) {
	if !enable {
		return nil, ErrSMSUninitialized
	}
	in.UserID = uint32(app_id)
	in.UserKey = app_key
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	r, err := s.client.UserBalance(ctx, in)
	return r, err
}

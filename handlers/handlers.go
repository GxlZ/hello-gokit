package handlers

import (
	"context"

	pb "hello-gokit/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.HelloServer {
	return helloService{}
}

type helloService struct{}

// HelloV1 implements Service.
func (s helloService) HelloV1(ctx context.Context, in *pb.Req) (*pb.Resp, error) {
	var resp pb.Resp
	resp = pb.Resp{
		// Code:
		// Msg:
	}
	return &resp, nil
}

package main

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	greetv1 "github.com/somasekimoto/connect-go-test/server/protocolbuffers/greet/v1"
	"github.com/somasekimoto/connect-go-test/server/protocolbuffers/greet/v1/greetv1connect"
)

type GreetServer struct {
	greetv1connect.UnimplementedGreetServiceHandler
}

func (s *GreetServer) GreetUnary(
	ctx context.Context,
	req *connect.Request[greetv1.GreetUnaryRequest],
) (*connect.Response[greetv1.GreetUnaryResponse], error) {
	res := connect.NewResponse(&greetv1.GreetUnaryResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}

package controller

import (
	greetv1 "atom/http/proto/gen/greet/v1"
	"context"

	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom-addons/providers/grpcs"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
	"google.golang.org/grpc"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(NewGreetController, atom.GroupGrpcServer)
}

// c *GreetController greetv1.GreeterServiceServer
var _ greetv1.GreeterServiceServer = (*GreetController)(nil)

type GreetController struct {
	greetv1.UnimplementedGreeterServiceServer
}

func NewGreetController() grpcs.ServerService {
	return &GreetController{}
}

func (c *GreetController) Register(s *grpc.Server) {
	greetv1.RegisterGreeterServiceServer(s, c)
}

func (c *GreetController) Name() string {
	return greetv1.GreeterService_ServiceDesc.ServiceName
}

func (c *GreetController) SayHello(ctx context.Context, req *greetv1.SayHelloRequest) (*greetv1.SayHelloResponse, error) {
	return &greetv1.SayHelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

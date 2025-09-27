package server

import (
	"github.com/Dizzrt/ellie-layout/api/gen/example"
	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie-layout/internal/iface"
	"github.com/Dizzrt/ellie/log"
	"github.com/Dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.Bootstrap, logger log.LogWriter, exampleHandler *iface.ExampleHandler) *grpc.Server {
	opts := []grpc.ServerOption{}

	grpcServerConf := c.Server.GRPC
	if grpcServerConf.Addr != "" {
		opts = append(opts, grpc.Address(grpcServerConf.Addr))
	}

	srv := grpc.NewServer(opts...)
	example.RegisterExampleServer(srv, exampleHandler)

	return srv
}

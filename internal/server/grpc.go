package server

import (
	"github.com/Dizzrt/ellie-layout/api/gen/example"
	"github.com/Dizzrt/ellie-layout/internal/iface"
	"github.com/Dizzrt/ellie/log"
	"github.com/Dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(logger log.LogWriter, exampleHandler *iface.ExampleHandler) *grpc.Server {
	opts := []grpc.ServerOption{}

	// if c.Addr != "" {
	// 	opts = append(opts, grpc.Address(c.Addr))
	// }

	srv := grpc.NewServer(opts...)
	example.RegisterExampleServer(srv, exampleHandler)

	return srv
}

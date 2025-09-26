package server

import (
	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie/log"
	"github.com/Dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.GRPCServer, logger log.LogWriter) *grpc.Server {
	opts := []grpc.ServerOption{}

	if c.Addr != "" {
		opts = append(opts, grpc.Address(c.Addr))
	}

	srv := grpc.NewServer(opts...)
	// srv.RegisterService()

	return srv
}

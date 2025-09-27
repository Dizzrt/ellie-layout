package server

import (
	"github.com/Dizzrt/ellie-layout/api/gen/example"
	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie-layout/internal/iface"
	"github.com/Dizzrt/ellie/log"
	"github.com/Dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.Bootstrap, logger log.LogWriter, exampleHandler *iface.ExampleHandler) *http.Server {
	opts := []http.ServerOption{}

	httpServerConf := c.Server.HTTP

	if httpServerConf.Addr != "" {
		opts = append(opts, http.Address(httpServerConf.Addr))
	}

	srv := http.NewServer(opts...)
	example.RegisterExampleHTTPServer(srv, exampleHandler)

	return srv
}

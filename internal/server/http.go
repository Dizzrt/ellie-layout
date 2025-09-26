package server

import (
	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie/log"
	"github.com/Dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.HTTPServer, logger log.LogWriter) *http.Server {
	opts := []http.ServerOption{}

	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}

	srv := http.NewServer(opts...)
	// TODO register

	return srv
}

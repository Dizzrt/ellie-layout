//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie-layout/internal/conf"
	"github.com/dizzrt/ellie-layout/internal/iface"
	"github.com/dizzrt/ellie-layout/internal/server"
	"github.com/dizzrt/ellie/log"
	"github.com/google/wire"
)

func wireApp(bootstrap *conf.Bootstrap, logger log.LogWriter) (*ellie.App, func(), error) {
	panic(wire.Build(newApp, server.ProviderSet, iface.ProviderSet))
}

//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/Dizzrt/ellie"
	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie-layout/internal/server"
	"github.com/Dizzrt/ellie/log"
	"github.com/google/wire"
)

func wireApp(bootstrap *conf.Bootstrap, logger log.LogWriter) (*ellie.App, func(), error) {
	panic(wire.Build(newApp, server.ProviderSet))
}

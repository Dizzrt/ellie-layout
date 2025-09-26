package cmd

import (
	"os"

	"github.com/Dizzrt/ellie-layout/internal/conf"
	"github.com/Dizzrt/ellie/transport/grpc"
	"github.com/Dizzrt/ellie/transport/http"

	"github.com/Dizzrt/ellie"
	"github.com/Dizzrt/ellie/config"
	"github.com/Dizzrt/ellie/log"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     string
	Version  string
	flagConf string
	id, _    = os.Hostname()
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start this service",
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.NewStdLogger(os.Stdout)
		conf := config.NewStdViperConfig()
		if err := conf.Load(); err != nil {
			panic(err)
		}

		bootstrap, err := buildBootstrap(conf)
		if err != nil {
			panic(err)
		}

		app, cleanup, err := wireApp(bootstrap, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func buildBootstrap(c config.Config) (*conf.Bootstrap, error) {
	var bootstrap conf.Bootstrap
	if err := c.UnmarshalKey("server", &bootstrap.Server); err != nil {
		return nil, err
	}

	return &bootstrap, nil
}

func newApp(logger log.LogWriter, gs *grpc.Server, hs *http.Server) *ellie.App {
	return ellie.New(
		ellie.ID(id),
		ellie.Name(Name),
		ellie.Version(Version),
		ellie.Metadata(map[string]string{}),
		ellie.Logger(logger),
		ellie.Server(gs, hs),
	)
}

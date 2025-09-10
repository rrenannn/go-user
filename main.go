package main

import (
	"github.com/rrenannn/go-user/cmd"
	"github.com/rrenannn/go-user/config"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			config.NewContainerDI,
			config.NewDB,
			config.NewQueries,
			cmd.NewEcho,
			cmd.NewRoutes,
		),
		config.AllModules,
		fx.Invoke(cmd.RegisterRoutes),
	)
	app.Run()
}

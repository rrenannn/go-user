package cmd

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rrenannn/go-user/internal/user"
	"go.uber.org/fx"
)

type Routes struct {
	lc          fx.Lifecycle
	e           *echo.Echo
	userHandler *user.Handler
}

func NewRoutes(lc fx.Lifecycle, e *echo.Echo, uHandler *user.Handler) Routes {
	return Routes{
		lc:          lc,
		e:           e,
		userHandler: uHandler,
	}
}

func NewEcho() *echo.Echo {
	return echo.New()
}

func RegisterRoutes(routes Routes) {
	routes.lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			users := routes.e.Group("user")
			users.POST("/create", routes.userHandler.CreateUser)
			users.GET("/:id", routes.userHandler.GetUserById)
			users.GET("/email", routes.userHandler.GetUserByEmail)

			routes.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: []string{"*"}, // ou "*"
				AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS, echo.DELETE},
			}))

			go func() {
				routes.e.Logger.Fatal(routes.e.Start(":8080"))
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			routes.e.Logger.Info("shutting down...")
			return routes.e.Shutdown(ctx)
		},
	})
}

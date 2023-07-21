package main

import (
	"context"
	"fmt"

	"github.com/LehcimW4R/inventory-manager/database"
	"github.com/LehcimW4R/inventory-manager/internal/api"
	"github.com/LehcimW4R/inventory-manager/internal/repository"
	"github.com/LehcimW4R/inventory-manager/internal/service"
	"github.com/LehcimW4R/inventory-manager/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx" //para inyección de dependencias
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

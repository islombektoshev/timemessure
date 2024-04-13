package app

import (
	"github.com/islombektoshev/timemessure/app/app_props"
	"github.com/islombektoshev/timemessure/app/gin_config"
	"github.com/islombektoshev/timemessure/app/gin_handlers"
	. "github.com/islombektoshev/timemessure/app/utils"
	"go.uber.org/fx"
)

func CreateFxApp() *fx.App {
	app := fx.New(
		fx.Provide(
			Spreed(
				app_props.FxDeps(),
				gin_config.FxDeps(),
				gin_handlers.FxDeps(),
			)...,
		),
		fx.Invoke(gin_config.GinStart),
	)
	return app
}

package gin_config

import (
	"github.com/gin-gonic/gin"
	"github.com/islombektoshev/timemessure/app/app_props"
	"go.uber.org/fx"
)

type StartParams struct {
	fx.In
	Configurers []GinConfigurer `group:"gin_configurer"`
	Engine      *gin.Engine
	Props       app_props.AppProperties
}
type GinStartRes = int

func init() {
	provide(func() *gin.Engine {
		return gin.Default()
	})
}

func GinStart(params StartParams) (GinStartRes, error) {
	for _, configurer := range params.Configurers {
		configurer(params.Engine)
	}
	return 0, params.Engine.Run(params.Props["server.addr"])
}

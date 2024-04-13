package gin_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/islombektoshev/timemessure/app/gin_config"
	"go.uber.org/fx"
)

type FxGinConfigurerResult struct {
	fx.Out
	GinConfigurer gin_config.GinConfigurer `group:"gin_configurer"`
}

func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	provide(func() FxGinConfigurerResult {
		return FxGinConfigurerResult{
			GinConfigurer: func(g *gin.Engine) {
				g.Handle(httpMethod, relativePath, handlers...)
			},
		}
	})
}

func GET(relativePath string, handlers ...gin.HandlerFunc) {
	provide(func() FxGinConfigurerResult {
		return FxGinConfigurerResult{
			GinConfigurer: func(g *gin.Engine) {
				g.GET(relativePath, handlers...)
			},
		}
	})
}

func POST(relativePath string, handlers ...gin.HandlerFunc) {
	provide(func() FxGinConfigurerResult {
		return FxGinConfigurerResult{
			GinConfigurer: func(g *gin.Engine) {
				g.POST(relativePath, handlers...)
			},
		}
	})
}

func PUT(relativePath string, handlers ...gin.HandlerFunc) {
	provide(func() FxGinConfigurerResult {
		return FxGinConfigurerResult{
			GinConfigurer: func(g *gin.Engine) {
				g.PUT(relativePath, handlers...)
			},
		}
	})
}

func DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	provide(func() FxGinConfigurerResult {
		return FxGinConfigurerResult{
			GinConfigurer: func(g *gin.Engine) {
				g.DELETE(relativePath, handlers...)
			},
		}
	})
}

package gin_handlers

import (
	"github.com/gin-gonic/gin"
	. "github.com/islombektoshev/timemessure/app/utils"
	"github.com/islombektoshev/timemessure/app/view"
)

func init() {
	GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Render(200, ToRender(view.Index()))
	})

}

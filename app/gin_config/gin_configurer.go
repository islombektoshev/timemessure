package gin_config

import "github.com/gin-gonic/gin"

// GinConfigurer arbitrary configuration will apply before run
type GinConfigurer func(engine *gin.Engine)

package server

import (
	"github/kunhou/gl_exercise/internal/deliver/http/router"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(debug bool, taskRouter *router.TaskRouter) *gin.Engine {
	ginMode := gin.ReleaseMode
	if debug {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	r := gin.New()

	r.GET("/_health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	return r
}

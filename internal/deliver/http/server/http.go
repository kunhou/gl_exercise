package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github/kunhou/gl_exercise/document/swagger"
	"github/kunhou/gl_exercise/internal/deliver/http/router"
)

// @title Task API
// @version 1.0
// @description This is a Task API server.
// @termsOfService

// @contact.name Kun Hou
// @contact.url
// @contact.email

// @host localhost:8080
// @BasePath /
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

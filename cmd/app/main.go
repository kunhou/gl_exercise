package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github/kunhou/gl_exercise/internal/common/config"
	"github/kunhou/gl_exercise/internal/pkg/srvmgmt"
	"github/kunhou/gl_exercise/internal/pkg/srvmgmt/httpsrv"
)

func main() {
	runApp()
}

func runApp() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApplication(
		cfg.Debug, &cfg.Server)
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

func newApplication(serverConf *config.Server, engine *gin.Engine) *srvmgmt.Application {
	return srvmgmt.NewApp(
		srvmgmt.WithServer(httpsrv.NewServer(engine, &serverConf.HTTP)),
	)
}

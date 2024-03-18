//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github/kunhou/gl_exercise/internal/common/config"
	"github/kunhou/gl_exercise/internal/deliver/http/router"
	"github/kunhou/gl_exercise/internal/deliver/http/server"
	"github/kunhou/gl_exercise/internal/pkg/srvmgmt"
	"github/kunhou/gl_exercise/internal/repository"
	"github/kunhou/gl_exercise/internal/service"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *config.Server,
) (*srvmgmt.Application, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		service.ProviderSetService,
		repository.ProviderSetRepository,
		newApplication,
	))
}

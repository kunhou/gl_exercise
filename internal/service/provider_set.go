package service

import (
	"github.com/google/wire"

	"github/kunhou/gl_exercise/internal/service/task"
)

var ProviderSetUsecase = wire.NewSet(
	task.New,
)

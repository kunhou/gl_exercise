package repository

import (
	"github.com/google/wire"

	task "github/kunhou/gl_exercise/internal/repository/task"
)

var ProviderSetRepository = wire.NewSet(
	task.NewTaskRepo,
)

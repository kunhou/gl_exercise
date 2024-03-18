package task

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github/kunhou/gl_exercise/internal/common/reason"
	"github/kunhou/gl_exercise/internal/entity"
	"github/kunhou/gl_exercise/internal/pkg/errors"
)

type taskTestSuite struct {
	suite.Suite

	repo *TaskRepo
}

func (suite *taskTestSuite) SetupTest() {
	suite.repo = &TaskRepo{
		taskStorage: make(map[int]entity.Task),
	}
}

func (suite *taskTestSuite) TearDownTest() {
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(taskTestSuite))
}

func (suite *taskTestSuite) TestList() {
	suite.repo.taskStorage[1] = entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}

	suite.repo.taskStorage[2] = entity.Task{
		Id:     2,
		Name:   "task 2",
		Status: 1,
	}

	suite.repo.taskStorage[3] = entity.Task{
		Id:     3,
		Name:   "task 3",
		Status: 2,
	}

	tasks, err := suite.repo.List(context.Background())

	suite.NoError(err)

	suite.Equal([]entity.Task{
		{
			Id:     1,
			Name:   "task 1",
			Status: 0,
		},
		{
			Id:     2,
			Name:   "task 2",
			Status: 1,
		},
		{
			Id:     3,
			Name:   "task 3",
			Status: 2,
		},
	}, tasks)
}

func (suite *taskTestSuite) TestListEmpty() {
	tasks, err := suite.repo.List(context.Background())

	suite.NoError(err)

	suite.Equal([]entity.Task{}, tasks)
}

func (suite *taskTestSuite) TestCreate() {
	suite.repo.taskStorage[1] = entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}

	suite.repo.lastID = 1

	task, err := suite.repo.Create(context.Background(), entity.Task{
		Name:   "task 2",
		Status: 0,
	})

	suite.NoError(err)

	suite.Equal(entity.Task{
		Id:     2,
		Name:   "task 2",
		Status: 0,
	}, task)
}

func (suite *taskTestSuite) TestCreateFirst() {
	task, err := suite.repo.Create(context.Background(), entity.Task{
		Name:   "task 1",
		Status: 0,
	})

	suite.NoError(err)

	suite.Equal(entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}, task)
}

func (suite *taskTestSuite) TestUpdate() {
	suite.repo.taskStorage[1] = entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}

	suite.repo.taskStorage[2] = entity.Task{
		Id:     2,
		Name:   "task 2",
		Status: 1,
	}

	task, err := suite.repo.Update(context.Background(), 1, entity.Task{
		Name:   "task 1 updated",
		Status: 1,
	})

	suite.NoError(err)

	suite.Equal(entity.Task{
		Id:     1,
		Name:   "task 1 updated",
		Status: 1,
	}, task)
}

func (suite *taskTestSuite) TestUpdateNotExist() {
	suite.repo.taskStorage[1] = entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}

	task, err := suite.repo.Update(context.Background(), 2, entity.Task{
		Name:   "task 2 updated",
		Status: 0,
	})

	var myErr *errors.Error
	suite.ErrorAs(err, &myErr)
	suite.Equal(404, myErr.Code)
	suite.Equal(reason.TaskNotFound, myErr.Reason)

	suite.Equal(entity.Task{}, task)
}

func (suite *taskTestSuite) TestDelete() {
	suite.repo.taskStorage[1] = entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}

	suite.repo.taskStorage[2] = entity.Task{
		Id:     2,
		Name:   "task 2",
		Status: 1,
	}

	err := suite.repo.Delete(context.Background(), 1)

	suite.NoError(err)

	suite.Equal(1, len(suite.repo.taskStorage))
}

func (suite *taskTestSuite) TestDeleteNotExist() {
	err := suite.repo.Delete(context.Background(), 1)

	var myErr *errors.Error
	suite.ErrorAs(err, &myErr)
	suite.Equal(404, myErr.Code)
	suite.Equal(reason.TaskNotFound, myErr.Reason)
}

package task

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github/kunhou/gl_exercise/internal/entity"
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

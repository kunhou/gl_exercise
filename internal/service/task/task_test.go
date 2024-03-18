package task

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github/kunhou/gl_exercise/internal/entity"
	"github/kunhou/gl_exercise/internal/service/mocks"
)

type taskTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	repo *mocks.MockITaskRepository

	taskSrv *Service
}

func (suite *taskTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockITaskRepository(suite.ctrl)

	suite.taskSrv = New(suite.repo)
}

func (suite *taskTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(taskTestSuite))
}

func (suite *taskTestSuite) TestList() {
	suite.repo.EXPECT().List(gomock.Any()).Return([]entity.Task{
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
	}, nil)

	data, err := suite.taskSrv.List(context.Background())
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
	}, data)
}

func (suite *taskTestSuite) TestCreate() {
	suite.repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}, nil)

	data, err := suite.taskSrv.Create(context.Background(), entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	})

	suite.NoError(err)

	suite.Equal(entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}, data)
}

func (suite *taskTestSuite) TestUpdate() {
	suite.repo.EXPECT().Update(gomock.Any(), 1, entity.Task{
		Name:   "task 1 updated",
		Status: 1,
	}).Return(entity.Task{
		Id:     1,
		Name:   "task 1 updated",
		Status: 1,
	}, nil)

	data, err := suite.taskSrv.Update(context.Background(), 1, entity.Task{
		Name:   "task 1 updated",
		Status: 1,
	})

	suite.NoError(err)

	suite.Equal(entity.Task{
		Id:     1,
		Name:   "task 1 updated",
		Status: 1,
	}, data)
}

func (suite *taskTestSuite) TestDelete() {
	suite.repo.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	err := suite.taskSrv.Delete(context.Background(), 1)

	suite.NoError(err)
}

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github/kunhou/gl_exercise/internal/common/reason"
	"github/kunhou/gl_exercise/internal/deliver/http/mocks"
	"github/kunhou/gl_exercise/internal/entity"
	"github/kunhou/gl_exercise/internal/pkg/errors"
)

type taskTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	srv  *mocks.MockITaskService

	router *gin.Engine
}

func (suite *taskTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.srv = mocks.NewMockITaskService(suite.ctrl)

	tRouter := NewTaskRouter(suite.srv)

	gin.SetMode(gin.TestMode)
	suite.router = gin.New()
	tRouter.RegisterRouter(suite.router)
}

func (suite *taskTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(taskTestSuite))
}

func (suite *taskTestSuite) TestList() {
	suite.srv.EXPECT().List(gomock.Any()).Return([]entity.Task{
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

	request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(`{"result":[{"id":1,"name":"task 1","status":0},{"id":2,"name":"task 2","status":1}]}`, response.Body.String())
}

func (suite *taskTestSuite) TestListWithEmptyList() {
	suite.srv.EXPECT().List(gomock.Any()).Return([]entity.Task{}, nil)

	request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(`{"result":[]}`, response.Body.String())
}

func (suite *taskTestSuite) TestListWithError() {
	suite.srv.EXPECT().List(gomock.Any()).Return(nil, errors.InternalServer(reason.UnknownError))

	request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusInternalServerError, response.Code)
	suite.Equal(`{"result":"unknown_error"}`, response.Body.String())
}

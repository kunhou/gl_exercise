package router

import (
	"bytes"
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

func (suite *taskTestSuite) TestCreate() {
	suite.srv.EXPECT().Create(gomock.Any(), entity.Task{
		Name: "task 1",
	}).Return(entity.Task{
		Id:     1,
		Name:   "task 1",
		Status: 0,
	}, nil)

	request, _ := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer([]byte(`{"name":"task 1"}`)))
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusCreated, response.Code)
	suite.Equal(`{"result":{"id":1,"name":"task 1","status":0}}`, response.Body.String())
}

func (suite *taskTestSuite) TestCreateWithoutName() {
	request, _ := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer([]byte(`{}`)))
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusBadRequest, response.Code)
	suite.Equal(`{"result":"request_format_error"}`, response.Body.String())
}

// test update task
func (suite *taskTestSuite) TestUpdateSuccess() {
	tests := []struct {
		name                string
		body                string
		mockUpdateInputId   int
		mockUpdateInputTask entity.Task
		mockUpdateResult    entity.Task
		responseCode        int
		responseBody        string
	}{
		{
			name:              "update success",
			body:              `{"name":"task 1","status":1}`,
			mockUpdateInputId: 1,
			mockUpdateInputTask: entity.Task{
				Name:   "task 1",
				Status: 1,
			},
			mockUpdateResult: entity.Task{
				Id:     1,
				Name:   "task 1",
				Status: 1,
			},
			responseCode: http.StatusOK,
			responseBody: `{"result":{"id":1,"name":"task 1","status":1}}`,
		},
		{
			name:              "update success with status 0",
			body:              `{"name":"task 1","status":0}`,
			mockUpdateInputId: 1,
			mockUpdateInputTask: entity.Task{
				Name:   "task 1",
				Status: 0,
			},
			mockUpdateResult: entity.Task{
				Id:     1,
				Name:   "task 1",
				Status: 0,
			},
			responseCode: http.StatusOK,
			responseBody: `{"result":{"id":1,"name":"task 1","status":0}}`,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.srv.EXPECT().Update(gomock.Any(), tt.mockUpdateInputId, tt.mockUpdateInputTask).Return(tt.mockUpdateResult, nil)

			request, _ := http.NewRequest(http.MethodPut, "/task/1", bytes.NewBuffer([]byte(tt.body)))
			response := httptest.NewRecorder()
			suite.router.ServeHTTP(response, request)

			suite.Equal(tt.responseCode, response.Code)
			suite.Equal(tt.responseBody, response.Body.String())
		})
	}
}

func (suite *taskTestSuite) TestUpdateWithInvalidInput() {
	tests := []struct {
		name         string
		id           string
		body         string
		responseCode int
		responseBody string
	}{
		{
			name:         "update without name",
			id:           "1",
			body:         `{"status": 0}`,
			responseCode: http.StatusBadRequest,
			responseBody: `{"result":"request_format_error"}`,
		},
		{
			name:         "update without status",
			id:           "1",
			body:         `{"name": "task 1"}`,
			responseCode: http.StatusBadRequest,
			responseBody: `{"result":"request_format_error"}`,
		},
		{
			name:         "update with invalid status",
			id:           "1",
			body:         `{"name":"task 1","status":2}`,
			responseCode: http.StatusBadRequest,
			responseBody: `{"result":"request_format_error"}`,
		},
		{
			name:         "update with invalid id",
			id:           "a",
			body:         `{"name":"task 1","status":1}`,
			responseCode: http.StatusBadRequest,
			responseBody: `{"result":"request_format_error"}`,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			request, _ := http.NewRequest(http.MethodPut, "/task/"+tt.id, bytes.NewBuffer([]byte(tt.body)))
			response := httptest.NewRecorder()
			suite.router.ServeHTTP(response, request)

			suite.Equal(tt.responseCode, response.Code)
			suite.Equal(tt.responseBody, response.Body.String())
		})
	}
}

func (suite *taskTestSuite) TestUpdateNotFound() {
	suite.srv.EXPECT().Update(gomock.Any(), 1, entity.Task{
		Name:   "task 1",
		Status: 0,
	}).Return(entity.Task{}, errors.NotFound(reason.TaskNotFound))

	request, _ := http.NewRequest(http.MethodPut, "/task/1", bytes.NewBuffer([]byte(`{"name":"task 1","status":0}`)))
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusNotFound, response.Code)
	suite.Equal(`{"result":"error.task.not_found"}`, response.Body.String())
}

func (suite *taskTestSuite) TestDelete() {
	suite.srv.EXPECT().Delete(gomock.Any(), 1).Return(nil)

	request, _ := http.NewRequest(http.MethodDelete, "/task/1", nil)
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusOK, response.Code)
	suite.Equal(`{"result":"success"}`, response.Body.String())
}

func (suite *taskTestSuite) TestDeleteNotFound() {
	suite.srv.EXPECT().Delete(gomock.Any(), 1).Return(errors.NotFound(reason.TaskNotFound))

	request, _ := http.NewRequest(http.MethodDelete, "/task/1", nil)
	response := httptest.NewRecorder()
	suite.router.ServeHTTP(response, request)

	suite.Equal(http.StatusNotFound, response.Code)
	suite.Equal(`{"result":"error.task.not_found"}`, response.Body.String())
}

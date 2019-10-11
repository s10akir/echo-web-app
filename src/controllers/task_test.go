package controller

import (
	"errors"
	"fmt"
	"github.com/s10akir/echo-web-app/src/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

type MockRepo struct {
}

func (MockRepo) CreateNewTask(title string, content string) (*models.Task, error) {
	if title != "" && content != "" {
		return &models.Task{}, nil
	} else {
		return nil, errors.New("")
	}
}

func (MockRepo) FindTaskByID(id int) (*models.Task, error) {
	return &models.Task{}, nil
}

func (MockRepo) UpdateTask(id int, title string, content string) error {
	return nil
}

func (MockRepo) DeleteTask(id int) error {
	return nil
}

func (MockRepo) Close() error {
	return nil
}

var taskJson = `{"title":"This is Test Title", "content":"This is Test Content"}`

func TestTaskController_New(test *testing.T) {
	app := echo.New()

	// 正しいパラメータでリクエストを投げられた際に正常にレスポンスが変えることを保証する
	{
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(taskJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := app.NewContext(req, rec)

		taskController := TaskController{MockRepo{}}

		if assert.NoError(test, taskController.New(context)) {
			assert.Equal(test, http.StatusOK, rec.Code)
		}
	}

	// 正しくないパラメータでリクエストを投げられた際に正常にエラーを返すことを保証する
	{
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		context := app.NewContext(req, rec)

		taskController := TaskController{MockRepo{}}

		if assert.NoError(test, taskController.New(context)) {
			fmt.Println(rec.Body)
			assert.Equal(test, http.StatusBadRequest, rec.Code)
		}
	}
}

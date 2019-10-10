package controller

import (
	"encoding/json"
	"github.com/s10akir/echo-web-app/src/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"

	"github.com/s10akir/echo-web-app/src/repository"
)

type TaskController struct {
	Repo repository.Repository
}

var warn = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.LUTC)

func (taskController TaskController) Index(context echo.Context) error {
	return context.String(http.StatusOK, "task#index")
}

func (taskController TaskController) New(context echo.Context) error {
	type param struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	value := new(param)
	if err := context.Bind(value); err != nil {
		handleError(err)

		return context.String(http.StatusBadRequest, "error")
	}

	task, err := taskController.Repo.CreateNewTask(value.Title, value.Content)
	if err != nil {
		handleError(err)

		return context.String(http.StatusBadRequest, "error")

	}

	var jsonByte []byte
	{
		var err error

		if jsonByte, err = json.Marshal(task); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	return context.String(http.StatusOK, string(jsonByte))
}

func (taskController TaskController) Show(context echo.Context) error {
	var id int
	{
		var err error

		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	var task *models.Task
	{
		var err error

		if task, err = taskController.Repo.FindTaskByID(id); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	var jsonByte []byte
	{
		var err error

		if jsonByte, err = json.Marshal(task); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	return context.String(http.StatusOK, string(jsonByte))
}

func (taskController TaskController) Update(context echo.Context) error {
	type param struct {
		Id      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var id int
	{
		var err error

		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	value := new(param)
	if err := context.Bind(value); err != nil {
		handleError(err)

		return context.String(http.StatusBadRequest, "error")
	}

	if err := taskController.Repo.UpdateTask(id, value.Title, value.Content); err != nil {
		handleError(err)

		return context.String(http.StatusBadRequest, "error")
	}

	return context.String(http.StatusOK, "update success.")
}

func (taskController TaskController) Delete(context echo.Context) error {
	var id int
	{
		var err error
		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			handleError(err)

			return context.String(http.StatusBadRequest, "error")
		}
	}

	if err := taskController.Repo.DeleteTask(id); err != nil {
		handleError(err)

		return context.String(http.StatusBadRequest, "error")
	}

	return context.String(http.StatusOK, "delete success.")
}

func handleError(err error) {
	warn.Print(err)
}

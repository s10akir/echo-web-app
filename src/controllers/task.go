package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/s10akir/echo-web-app/src/repository"
)

type TaskController struct {
	Repo repository.Repository
}

func (taskController TaskController) Index(context echo.Context) error {
	return context.String(http.StatusOK, "task#index")
}

func (taskController TaskController) New(context echo.Context) error {
	type param struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	value := new(param)
	context.Bind(value)

	task, err := taskController.Repo.CreateNewTask(value.Title, value.Content)
	if err != nil {
		fmt.Println(err)
		return context.String(http.StatusOK, "error")
	}

	jsonByte, _ := json.Marshal(task)

	return context.String(http.StatusOK, string(jsonByte))
}

func (taskController TaskController) Show(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))
	task, err := taskController.Repo.FindTaskByID(id)

	if err != nil {
		fmt.Println(err)
		return context.String(http.StatusOK, "error")
	}

	jsonByte, _ := json.Marshal(task)
	return context.String(http.StatusOK, string(jsonByte))
}

func (taskController TaskController) Update(context echo.Context) error {
	type param struct {
		Id      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	id, _ := strconv.Atoi(context.Param("id"))
	value := new(param)
	context.Bind(value)

	err := taskController.Repo.UpdateTask(id, value.Title, value.Content)

	if err != nil {
		fmt.Println(err)
		return context.String(http.StatusOK, "error")
	}

	return context.String(http.StatusOK, "update success.")
}

func (taskController TaskController) Delete(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))
	err := taskController.Repo.DeleteTask(id)

	if err != nil {
		fmt.Println(err)
		return context.String(http.StatusOK, "error")
	}

	return context.String(http.StatusOK, "delete success.")
}

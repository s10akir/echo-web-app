package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type TaskController struct{}

func (taskController TaskController) Index(context echo.Context) error {
	return context.String(http.StatusOK, "task#index")
}

func (taskController TaskController) New(context echo.Context) error {
	return context.String(http.StatusOK, "task#new")
}

func (taskController TaskController) Show(context echo.Context) error {
	return context.String(http.StatusOK, "task#show: "+context.Param("id"))
}

func (taskController TaskController) Update(context echo.Context) error {
	return context.String(http.StatusOK, "task#update: "+context.Param("id"))
}

func (taskController TaskController) Delete(context echo.Context) error {
	return context.String(http.StatusOK, "task#delete: "+context.Param("id"))
}

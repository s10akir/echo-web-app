package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/s10akir/echo-web-app/src/controllers"
	"github.com/s10akir/echo-web-app/src/repository"
)

const PORT = ":8080"

func main() {
	// debug
	repo, _ := repository.New()

	app := echo.New()

	app.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})

	task := app.Group("/task")
	{
		taskController := controller.TaskController{Repo: repo}
		task.GET("", taskController.Index)
		task.POST("", taskController.New)
		task.GET("/:id", taskController.Show)
		task.POST("/:id", taskController.Update)
		task.DELETE("/:id", taskController.Delete)
	}

	app.Start(PORT)

	defer repo.Close()
}

package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/tylerb/graceful"

	"github.com/s10akir/echo-web-app/src/controllers"
	"github.com/s10akir/echo-web-app/src/repository"
)

const PORT = ":8080"
const TIMEOUT = 30 * time.Second

func main() {
	// debug
	repo, _ := repository.New()
	defer repo.Close()

	app := echo.New()
	app.Server.Addr = PORT

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

	graceful.ListenAndServe(app.Server, TIMEOUT)
}

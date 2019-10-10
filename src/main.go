package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/tylerb/graceful"

	"github.com/s10akir/echo-web-app/src/controllers"
	"github.com/s10akir/echo-web-app/src/repository"
)

const PORT = ":8080"
const TIMEOUT = 30 * time.Second

var warn = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.LUTC)

func main() {
	var repo repository.Repository
	{
		var err error

		repo, err = repository.New()
		if err != nil {
			handleError(err)
			// dbに接続できない以上継続不可なので強制終了
			panic(err)
		}
	}

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

	if err := graceful.ListenAndServe(app.Server, TIMEOUT); err != nil {
		handleError(err)
	}

	defer func() {
		if err := repo.Close(); err != nil {
			handleError(err)
		}
	}()
}

func handleError(err error) {
	warn.Print(err)
}

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

const TIMEOUT = 30 * time.Second

var env = os.Getenv("ECHO_ENV")
var warn = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.LUTC)

func main() {
	var port string
	var driver string
	var dataSource string

	switch env {
	// TODO: このあたりのパラメータを外で定義するようにする
	case "development":
		port = ":8080"
		driver = "mysql"
		dataSource = "user:password@tcp(db)/echo?parseTime=True"
		break

	case "production":
		port = ":80"
		driver = ""
		dataSource = ""
		break

	default:
		warn.Print("ECHO_ENV is required.")
	}

	var repo repository.Repository
	{
		var err error

		if repo, err = repository.New(driver, dataSource); err != nil {
			handleError(err)
			// dbに接続できない以上継続不可なので強制終了
			panic(err)
		}
	}

	app := echo.New()
	app.Server.Addr = port

	app.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})

	task := app.Group("/task")
	{
		taskController := controller.TaskController{Repo: repo}
		task.GET("", taskController.Index)
		task.POST("", taskController.New)
		task.GET("/:id", taskController.Show)
		task.PUT("/:id", taskController.Update)
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

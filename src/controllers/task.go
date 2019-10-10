package controller

import (
	"encoding/json"
	"github.com/s10akir/echo-web-app/src/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/s10akir/echo-web-app/src/repository"
)

type TaskController struct {
	Repo repository.Repository
}

// TODO: 各アクションで行っているエラー処理、こういうときこそpanic()をつかうべきなのではという感じがする

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
		return err
	}

	task, err := taskController.Repo.CreateNewTask(value.Title, value.Content)
	if err != nil {
		// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
		// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
		if err := context.String(http.StatusOK, "error"); err != nil {
			return err
		}

		return err
	}

	var jsonByte []byte
	{
		var err error

		if jsonByte, err = json.Marshal(task); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	return context.String(http.StatusOK, string(jsonByte))
}

func (taskController TaskController) Show(context echo.Context) error {
	var id int
	{
		var err error
		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	var task *models.Task
	{
		var err error

		if task, err = taskController.Repo.FindTaskByID(id); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	var jsonByte []byte
	{
		var err error

		if jsonByte, err = json.Marshal(task); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
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
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	value := new(param)
	if err := context.Bind(value); err != nil {
		// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
		// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
		if err := context.String(http.StatusOK, "error"); err != nil {
			return err
		}
	}

	if err := taskController.Repo.UpdateTask(id, value.Title, value.Content); err != nil {
		// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
		// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
		if err := context.String(http.StatusOK, "error"); err != nil {
			return err
		}
	}

	return context.String(http.StatusOK, "update success.")
}

func (taskController TaskController) Delete(context echo.Context) error {
	var id int
	{
		var err error
		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	if err := taskController.Repo.DeleteTask(id); err != nil {
		if id, err = strconv.Atoi(context.Param("id")); err != nil {
			// TODO: ここはcontext.String()のエラーを優先してよいのだろうか
			// どちらかを捨てるよりは両方エラー出力したいが、エラー処理はmain()に集約したい気もする
			if err := context.String(http.StatusOK, "error"); err != nil {
				return err
			}
		}
	}

	return context.String(http.StatusOK, "delete success.")
}

package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/s10akir/echo-web-app/src/models"
)

func (repo *repository) CreateNewTask(title string, content string) (*models.Task, error) {
	var result sql.Result
	{
		var err error

		if result, err = repo.db.Exec(
			`
		INSERT INTO tasks(title, content)
			VALUES(?, ?)
		`,
			title, content,
		); err != nil {
			return nil, err
		}
	}

	var id int
	{
		var err error
		var lastInsertedId int64

		if lastInsertedId, err = result.LastInsertId(); err != nil {
			return nil, err
		}

		id = int(lastInsertedId)
	}

	var task *models.Task
	{
		var err error

		if task, err = repo.FindTaskByID(int(id)); err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (repo *repository) FindTaskByID(id int) (*models.Task, error) {
	var task models.Task

	if err := repo.db.Get(
		&task,
		// idはPrimary Keyなので常に一件しか帰ってこないことが保証できる
		`
		SELECT * FROM tasks
			WHERE id = ?
		`,
		id,
	); err != nil {
		return nil, err
	}

	return &task, nil
}

func (repo *repository) UpdateTask(id int, title string, content string) error {
	now := time.Now()

	if _, err := repo.db.Exec(
		`
		UPDATE tasks SET title = ?, content = ?, updated_at = ?
			WHERE id = ?
		`,
		title, content, now, id,
	); err != nil {
		return err
	}

	return nil
}

func (repo *repository) DeleteTask(id int) error {
	if _, err := repo.db.Exec(
		`
		DELETE FROM tasks
		WHERE id = ?
		`,
		id,
	); err != nil {
		return err
	}

	return nil
}

package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/s10akir/echo-web-app/src/models"
)

func (r *repository) CreateNewTask(title string, content string) (*models.Task, error) {
	result, err := r.db.Exec(
		`
		INSERT INTO tasks(title, content)
			VALUES(?, ?)
		`,
		title, content,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	task, err := r.FindTaskByID(uint64(id))
	if err != nil {
		return nil, err
	}

	return task, err
}

func (r *repository) FindTaskByID(id uint64) (*models.Task, error) {
	var task models.Task
	err := r.db.Get(
		&task,
		// idはPrimary Keyなので常に一件しか帰ってこないことが保証できる
		`
		SELECT * FROM tasks
			WHERE id = ?
		`,
		id,
	)

	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *repository) UpdateTask(id uint64, title string, content string) error {
	now := time.Now()
	_, err := r.db.Exec(
		`
		UPDATE tasks SET title = ?, content = ?, updated_at = ?
			WHERE id = ?
		`,
		title, content, now, id,
	)

	return err
}

func (r *repository) DeleteTask(id uint64) error {
	_, err := r.db.Exec(
		`
		DELETE FROM tasks
		WHERE id = ?
		`,
		id,
	)

	return err
}

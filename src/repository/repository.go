package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/s10akir/echo-web-app/src/models"
)

type repository struct {
	db *sqlx.DB
}

type Repository interface {
	CreateNewTask(title string, content string) (*models.Task, error)
	FindTaskByID(id int) (*models.Task, error)
	UpdateTask(id int, title string, content string) error
	DeleteTask(id int) error
	Close() error
}

func New() (Repository, error) {
	var db *sqlx.DB
	{
		var err error

		if db, err = sqlx.Connect(
			"mysql",
			"user:password@tcp(db)/echo?parseTime=True",
		); err != nil {
			return nil, err
		}
	}

	return &repository{db: db}, nil
}

func (repo *repository) Close() error {
	return repo.db.Close()
}

func (repo *repository) generateID() (int, error) {
	var id int

	err := repo.db.Get(
		&id,
		`
		SELECT auto_increment FROM information_schema.tables
			WHERE table_name = 'tasks'
		`,
	)

	return id, err
}

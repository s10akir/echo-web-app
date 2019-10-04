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
	FindTaskByID(id uint64) (*models.Task, error)
	UpdateTask(id uint64, title string, content string) error
	DeleteTask(id uint64) error
	Close() error
}

func New() (Repository, error) {
	db, err := sqlx.Connect("mysql", "user:password@tcp(0.0.0.0)/echo?parseTime=True")

	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}

func (repo *repository) Close() error {
	return repo.db.Close()
}

func (repo *repository) generateID() (uint64, error) {
	var id uint64
	err := repo.db.Get(
		&id,
		`
		SELECT auto_increment FROM information_schema.tables
			WHERE table_name = 'tasks'
		`,
	)

	return id, err
}

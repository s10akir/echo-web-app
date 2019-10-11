package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_CreateNewTask(test *testing.T) {
	var db *sql.DB
	var mock sqlmock.Sqlmock
	{
		var err error
		if db, mock, err = sqlmock.New(); err != nil {
			panic(err)
		}
	}

	mockDB := sqlx.NewDb(db, "sqlmock")

	title := "title"
	content := "content"
	rows := sqlmock.NewRows([]string{"id", "title", "content", "created_at", "updated_at"}).AddRow(1, title, content, time.Now().UTC(), time.Now().UTC())

	mock.ExpectExec("INSERT INTO tasks").WithArgs(title, content).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(rows)

	repo := &repository{db: mockDB}

	// 正常にタスクが生成できることを保証する
	{
		_, err := repo.CreateNewTask(title, content)
		assert.NoError(test, err)
	}
}

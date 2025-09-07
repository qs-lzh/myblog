package data

import (
	"database/sql"
	"time"
)

type Todo struct {
	Title    string
	Content  string
	CratedAt time.Time
	DueDate  time.Time
}

type TodoModel struct {
	DB *sql.DB
}

func (model *TodoModel) Insert(title string, content string, dueDate time.Time) error {
	stmt := `
    insert into todos (title, content, due_date)
	  values (?, ?, ?)
  `
	_, err := model.DB.Exec(stmt, title, content, dueDate)
	return err
}

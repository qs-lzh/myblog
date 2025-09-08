package data

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID       int
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

func (model *TodoModel) GetAll() ([]*Todo, error) {
	stmt := `
	  select id, title, content, created_at, due_date from todos
	`
	rows, err := model.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		todo := &Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.CratedAt, &todo.DueDate)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (model *TodoModel) Delete(id int) error {
	stmt := `
	  delete from todos where id = ?
	`
	_, err := model.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

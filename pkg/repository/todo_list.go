package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo"
)

type TodoList struct {
	db *sqlx.DB
}
type TodoService struct {
	repo Todo
}

func NewTodoPostgres(db *sqlx.DB) *TodoList {
	return &TodoList{db: db}
}

func (r *TodoList) CreateTodo(todoL todo.Todo) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", TODO_TABLE)
	row := r.db.QueryRow(query, todoL.Title, todoL.Description)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *TodoList) AllContent() ([]*todo.Todo, error) {

	rows, err := r.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*todo.Todo, 0)
	for rows.Next() {
		bk := new(todo.Todo)
		err := rows.Scan(&bk.Id, &bk.Title, &bk.Description)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s\n", bk.Id, bk.Title, bk.Description)
	}
	return bks, err
}

func (r *TodoList) DeleteTodo(title string) (string, error) {
	query := `DELETE FROM todo WHERE title=$1;`
	_, err := r.db.Exec(query, title)
	if err != nil {
		return "", err
	}
	return "successfully", err
}

func (r *TodoList) UpdateTodo(todo todo.Todo) (string, error) {

	query := `UPDATE todo SET title=$2, description=$3 where id=$1;`
	row := r.db.QueryRow(query, todo.Id, todo.Title, todo.Description)
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return "ok", nil
	case nil:
		fmt.Println(todo.Id)
	default:
		return "ok", err
	}
	return "complete", nil
}

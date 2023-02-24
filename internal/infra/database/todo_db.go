package database

import "database/sql"

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {

}

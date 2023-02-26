// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int32          `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Done        sql.NullBool   `json:"done"`
	CreatedAt   time.Time      `json:"created_at"`
}
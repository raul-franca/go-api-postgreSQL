package entity

import "time"

type Todo struct {
	ID          int32     `json:"ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

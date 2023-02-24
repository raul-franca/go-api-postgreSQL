package entity

type Todo struct {
	ID          int64  `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

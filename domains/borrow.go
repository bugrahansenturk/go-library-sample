package domain

import "time"

type Borrow struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	BookID   int       `json:"book_id"`
	DueDate  time.Time `json:"due_date"`
	Returned bool      `json:"returned"`
}

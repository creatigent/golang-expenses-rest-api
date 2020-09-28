package models

import "time"

// Expense represents the expense model
type Expense struct {
	ID         string    `json:"id"`
	Price      float64   `json:"price"`
	Title      string    `json:"title"`
	Currency   string    `json:"currency"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

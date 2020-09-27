package models

import "time"

// Expense represents the expense model
type Expense struct {
	ID         string
	Price      float64
	Title      string
	Currency   string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

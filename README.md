# Expenses REST API in Go

This application is a fully featured REST API written in Go

#### Expense Model

```go
type Expense struct {
	ID         string
	Price      float64
	Title      string
    Currency   string
	CreatedAt  time.Time
	ModifiedAt  time.Time
}
```

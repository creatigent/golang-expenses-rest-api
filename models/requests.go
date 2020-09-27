package models

// GetAllExpensesRequest represents http request for fetching all expenses with pagination
type GetAllExpensesRequest struct {
	Page     int
	PageSize int
}

// GetExpensesByIDsRequest represents http request for fetching a list of expenses by ids
type GetExpensesByIDsRequest struct {
	IDs []string
}

// CreateExpenseRequest represents http request for creating an expense
type CreateExpenseRequest struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

// UpdateExpenseRequest represents http request for updating an expense
type UpdateExpenseRequest struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}

// DeleteExpensesRequest represents http request for deleting multiple expenses
type DeleteExpensesRequest struct {
	IDs []string
}

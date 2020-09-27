package repositories

import (
	"github.com/steevehook/expenses-rest-api/models"
)

type Expenses interface {
	GetAllExpenses(page, size int) ([]models.Expense, error)
	GetExpensesByIDs(ids []string) ([]models.Expense, error)
	CreateExpense(title, currency string, price float64) error
	UpdateExpense(title, currency string, price float64) error
	DeleteExpenses(ids []string) error
	Stop() error
}

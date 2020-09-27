package services

import (
	"github.com/steevehook/expenses-rest-api/models"
	"github.com/steevehook/expenses-rest-api/repositories"
)

type Expenses struct {
	ExpensesRepo repositories.Expenses
}

func (s Expenses) GetAllExpenses(req models.GetAllExpensesRequest) []models.Expense {
	return nil
}

func (s Expenses) GetExpensesByIDs(req models.GetAllExpensesRequest) []models.Expense {
	return nil
}

func (s Expenses) CreateExpense(req models.CreateExpenseRequest) error {
	return nil
}

func (s Expenses) UpdateExpense(req models.UpdateExpenseRequest) error {
	return nil
}

func (s Expenses) DeleteExpenses(req models.DeleteExpensesRequest) error {
	return nil
}

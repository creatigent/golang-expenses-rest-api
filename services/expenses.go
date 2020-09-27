package services

import (
	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/logging"
	"github.com/steevehook/expenses-rest-api/models"
	"github.com/steevehook/expenses-rest-api/repositories"
)

type Expenses struct {
	ExpensesRepo repositories.Expenses
}

func (s Expenses) GetAllExpenses(req models.GetAllExpensesRequest) ([]models.Expense, error) {
	expenses, err := s.ExpensesRepo.GetAllExpenses(req.Page, req.PageSize)
	if err != nil {
		logging.Logger.Error("could not fetch all expenses from db", zap.Error(err))
		return []models.Expense{}, err
	}
	return expenses, nil
}

func (s Expenses) GetExpensesByIDs(req models.GetAllExpensesRequest) ([]models.Expense, error) {
	return []models.Expense{}, nil
}

func (s Expenses) CreateExpense(req models.CreateExpenseRequest) error {
	err := s.ExpensesRepo.CreateExpense(req.Title, req.Currency, req.Price)
	if err != nil {
		logging.Logger.Error("could not create expense in db", zap.Error(err))
	}
	return nil
}

func (s Expenses) UpdateExpense(req models.UpdateExpenseRequest) error {
	return nil
}

func (s Expenses) DeleteExpenses(req models.DeleteExpensesRequest) error {
	return nil
}

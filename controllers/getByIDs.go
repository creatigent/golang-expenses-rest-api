package controllers

import (
	"net/http"

	"github.com/steevehook/expenses-rest-api/models"
)

type expensesByIDsGetter interface {
	GetExpensesByIDs(models.GetAllExpensesRequest) []models.Expense
}

func getExpensesByIDs(service expensesByIDsGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

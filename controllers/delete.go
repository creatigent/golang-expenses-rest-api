package controllers

import (
	"net/http"

	"github.com/steevehook/expenses-rest-api/models"
)

type expensesDeleter interface {
	DeleteExpenses(models.DeleteExpensesRequest) error
}

func deleteExpense(service expensesDeleter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

package controllers

import (
	"net/http"

	"github.com/steevehook/expenses-rest-api/models"
)

type expenseDeleter interface {
	DeleteExpense(models.DeleteExpenseRequest) error
}

func deleteExpense(service expenseDeleter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

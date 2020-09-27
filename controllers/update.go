package controllers

import (
	"net/http"

	"github.com/steevehook/expenses-rest-api/models"
)

type expenseUpdater interface {
	UpdateExpense(models.UpdateExpenseRequest) error
}

func updateExpense(service expenseUpdater) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("update one"))
	})
}

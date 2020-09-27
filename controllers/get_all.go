package controllers

import (
	"net/http"

	"github.com/steevehook/expenses-rest-api/models"
)

type allExpensesGetter interface {
	GetAllExpenses(models.GetAllExpensesRequest) []models.Expense
}

func getAllExpenses(service allExpensesGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get all"))
	})
}

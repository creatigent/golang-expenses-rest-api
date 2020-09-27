package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/steevehook/expenses-rest-api/middleware"
)

type ExpensesService interface {
	allExpensesGetter
	expensesByIDsGetter
	expenseCreator
	expenseUpdater
	expensesDeleter
}

type RouterConfig struct {
	ExpensesSvc ExpensesService
}

// NewRouter creates a new application HTTP router
func NewRouter(cfg RouterConfig) http.Handler {
	chain := alice.New(
		middleware.HTTPLogger,
	)
	jsonBodyChain := chain.Append(
		middleware.JSONBody,
	)
	route := func(h http.Handler) http.Handler {
		return chain.Then(h)
	}
	routeWithBody := func(h http.Handler) http.Handler {
		return jsonBodyChain.Then(h)
	}

	router := httprouter.New()
	router.Handler(http.MethodGet, "/expenses", route(getAllExpenses(cfg.ExpensesSvc)))
	router.Handler(http.MethodGet, "/expenses/:ids", route(getExpensesByIDs(cfg.ExpensesSvc)))
	router.Handler(http.MethodPost, "/expenses", routeWithBody(createExpense(cfg.ExpensesSvc)))
	router.Handler(http.MethodPatch, "/expenses/:id", routeWithBody(updateExpense(cfg.ExpensesSvc)))
	router.Handler(http.MethodDelete, "/expenses/:ids", route(deleteExpense(cfg.ExpensesSvc)))
	router.NotFound = route(NotFound())

	return router
}

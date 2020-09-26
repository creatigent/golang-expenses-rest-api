package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"github.com/steevehook/expenses-rest-api/middleware"
)

// NewRouter creates a new application HTTP router
func NewRouter() http.Handler {
	middlewareChain := alice.New(
		middleware.HTTPLogger,
	)
	route := func(h http.Handler) http.Handler {
		return middlewareChain.Then(h)
	}

	router := httprouter.New()
	router.Handler(http.MethodGet, "/expenses", route(getAllExpenses()))
	router.Handler(http.MethodGet, "/expenses/:ids", route(getExpensesByIDs()))
	router.Handler(http.MethodPost, "/expenses", route(createExpense()))
	router.Handler(http.MethodPatch, "/expenses/:id", route(updateExpense()))
	router.Handler(http.MethodDelete, "/expenses/:ids", route(deleteExpense()))

	return router
}

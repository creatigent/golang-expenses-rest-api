package controllers

import "net/http"

func getExpensesByIDs() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

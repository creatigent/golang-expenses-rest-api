package controllers

import "net/http"

func createExpense() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("create one"))
	})
}

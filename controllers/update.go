package controllers

import "net/http"

func updateExpense() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("update one"))
	})
}

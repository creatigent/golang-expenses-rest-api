package controllers

import "net/http"

func getAllExpenses() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get all"))
	})
}

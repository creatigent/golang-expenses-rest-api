package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/steevehook/expenses-rest-api/models"
	"github.com/steevehook/expenses-rest-api/transport"
)

const (
	pageQueryParam     = "page"
	pageSizeQueryParam = "page_size"

	defaultPage     = 1
	defaultPageSize = 10
)

type allExpensesGetter interface {
	GetAllExpenses(models.GetAllExpensesRequest) ([]models.Expense, error)
}

func getAllExpenses(service allExpensesGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := parseQueryParam(r, pageQueryParam, defaultPage)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}
		pageSize, err := parseQueryParam(r, pageSizeQueryParam, defaultPageSize)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}
		req := models.GetAllExpensesRequest{
			Page:     page,
			PageSize: pageSize,
		}

		expenses, err := service.GetAllExpenses(req)
		if err != nil {
			transport.SendHTTPError(w, err)
			return
		}
		transport.SendJSON(w, http.StatusOK, expenses)
	})
}

func parseQueryParam(r *http.Request, paramName string, defaultValue int) (int, error) {
	param := r.URL.Query().Get(paramName)
	if strings.TrimSpace(param) == "" {
		return defaultValue, nil
	}
	intParam, err := strconv.Atoi(param)
	if err != nil || intParam < 1 {
		e := models.DataValidationError{
			Message: fmt.Sprintf("invalid value: %s for param: %s", param, paramName),
		}
		return 0, e
	}
	return intParam, nil
}

package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/steevehook/expenses-rest-api/logging"
)

// params fetches params from context and converts it into julienschmidt/httprouter.Params struct
func params(r *http.Request) httprouter.Params {
	ctx := r.Context()
	psCtx := ctx.Value(httprouter.ParamsKey)
	ps, ok := psCtx.(httprouter.Params)

	if !ok {
		logging.Logger.Error("could not extract params from context")
		return httprouter.Params{}
	}
	return ps
}

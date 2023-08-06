package handler

import (
	"net/http"

	"github.com/PestovOleg/mini-bank/pkg/util"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	log := util.Getlogger("API")
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/v1").Subrouter()
	SetHandler(subRouter, BaseRoutes(log))

	return r
}

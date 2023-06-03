package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

func (a *api) Router() http.Handler {
	return a.router
}

type Server interface {
	Router() http.Handler
}

func NewServer() Server {
	a := &api{}
	r := mux.NewRouter()
	a.router = r
	return a
}

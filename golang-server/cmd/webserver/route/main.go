package route

import (
	"net/http"
	"rz-server/internal/common/interfaces"

	"github.com/gorilla/mux"
)

var _ interfaces.Route = (*Route)(nil)

type Route struct {
	router *mux.Router
	util   *interfaces.Util
}

func New(router *mux.Router, util *interfaces.Util) *Route {
	subRouter := router.NewRoute().Subrouter()

	return &Route{
		util:   util,
		router: subRouter,
	}
}

func (r *Route) Get(path string, handler http.HandlerFunc) {
	r.router.Handle(path, handler).Methods("GET")
}

func (r *Route) Post(path string, handler http.HandlerFunc) {
	r.router.Handle(path, handler).Methods("POST")
}

func (r *Route) Put(path string, handler http.HandlerFunc) {
	r.router.Handle(path, handler).Methods("PUT")
}

func (r *Route) Delete(path string, handler http.HandlerFunc) {
	r.router.Handle(path, handler).Methods("DELETE")
}

func (r *Route) SetPathPrefix(path string) {
	r.router = r.router.PathPrefix(path).Subrouter()
}

func (r *Route) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		r.router.Use(middleware)
	}
}

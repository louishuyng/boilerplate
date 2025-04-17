package interfaces

import "net/http"

type Route interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	SetPathPrefix(path string)
	Get(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Put(route string, handler http.HandlerFunc)
	Delete(route string, handler http.HandlerFunc)
}

type Server interface {
	NewRoute() Route
}

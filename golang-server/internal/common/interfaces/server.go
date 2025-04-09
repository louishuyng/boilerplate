package interfaces

import "net/http"

type Server interface {
	RegisterMiddlewares(handlers []func(http.Handler) http.Handler)

	GET(route string, handler http.HandlerFunc)
	POST(route string, handler http.HandlerFunc)
	PUT(route string, handler http.HandlerFunc)
	DELETE(route string, handler http.HandlerFunc)
}

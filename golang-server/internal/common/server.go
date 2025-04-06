package common

import "net/http"

type BasicServer interface {
	http.Handler
	Start()
	GET(route string, handler http.HandlerFunc)
	POST(route string, handler http.HandlerFunc)
}

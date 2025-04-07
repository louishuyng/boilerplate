package common

import "net/http"

type BasicServer interface {
	GET(route string, handler http.HandlerFunc)
	POST(route string, handler http.HandlerFunc)
}

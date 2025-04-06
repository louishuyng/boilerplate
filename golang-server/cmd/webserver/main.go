package main

import (
	"go-server/internal/user"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	user.RegisterUserServer(router)

	log.Fatal(http.ListenAndServe(":5001", router))
}

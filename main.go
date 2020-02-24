package main

import (
	"go-API/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":8080", mux)
}

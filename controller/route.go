package controller

import "net/http"

func Register() *http.ServeMux {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", home())
	return mux
}
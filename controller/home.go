package controller

import (
	"encoding/json"
	"go-API/views"
	"net/http"
)

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "Hello, everybody!!!",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}

package middlewares

import (
	"encoding/json"
	"net/http"
	"topcoin/ratingservice/response"
)

// Recovery - Middleware to handle Unhandled Error
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				jsonBody, _ := json.Marshal(response.APIResponse{Error: "Internal Server Error", Status: false, Response: ""})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}

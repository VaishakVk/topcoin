package middlewares

import (
	"net/http"
)

// SetHeaders MW
func SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(res, req)
	})
}

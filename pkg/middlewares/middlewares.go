package middleware

import (
	"fmt"
	"net/http"
	"time"
)

const timeFormat = time.RFC3339

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("[%s] %s %s \n", time.Now().UTC().Format(timeFormat), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

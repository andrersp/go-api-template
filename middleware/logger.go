package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		request := r.URL

		fmt.Println(request)

		next.ServeHTTP(w, r)

	})
}

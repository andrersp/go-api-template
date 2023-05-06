package middlewares

import (
	"context"
	"net/http"

	"github.com/andrersp/go-api-template/internal/adapters/api/helpers"
	"github.com/andrersp/go-api-template/internal/core/dto"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

}

func LoggerMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {
			// defer func() {
			// 	if err := recover(); err != nil {
			// 		w.WriteHeader(http.StatusInternalServerError)
			// 	}
			// }()

			tokenData, err := helpers.ExtractTokenData(r)
			if err != nil {
				erroResponse := dto.ErrorResponse{
					Success: false,
					Msg:     err.Error(),
				}

				helpers.Responder(http.StatusForbidden, w, erroResponse)
				return
			}

			ctx := context.WithValue(r.Context(), "tokenData", tokenData)

			next.ServeHTTP(w, r.WithContext(ctx))

		}

		response := http.HandlerFunc(fn)

		return response
	}
}

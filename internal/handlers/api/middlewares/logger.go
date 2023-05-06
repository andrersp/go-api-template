package middlewares

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/andrersp/go-api-template/pkg/logger"
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

		log := logger.GetLogger()

		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Error().Bytes("trace", debug.Stack()).Send()
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			log.Info().Int("status", wrapped.status).
				Str("method", r.Method).
				Str("path", r.URL.EscapedPath()).
				Dur("duration", time.Since(start)).
				Send()

		}

		response := http.HandlerFunc(fn)

		return response
	}
}

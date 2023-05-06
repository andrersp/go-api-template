package middlewares

import (
	"context"
	"net/http"

	"github.com/andrersp/go-api-template/internal/handlers/api/helpers"
	apperrors "github.com/andrersp/go-api-template/pkg/app-errors"
	secutiry "github.com/andrersp/go-api-template/pkg/security"
)

func JwtMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			tokenData, err := secutiry.ExtractTokenData(r)
			if err != nil {
				err := apperrors.NewAppError(err.Error())

				helpers.Responder(http.StatusForbidden, w, err)
				return
			}

			ctx := context.WithValue(r.Context(), secutiry.TokenData{}, tokenData)

			next.ServeHTTP(w, r.WithContext(ctx))

		}

		response := http.HandlerFunc(fn)

		return response
	}
}

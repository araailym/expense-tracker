package middleware

import (
	"github.com/araailym/expense-tracker/internal/auth"
	"net/http"
)

func (m *Middleware) Authorizer(
	role string,
	next http.Handler,
) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := m.log.With("middleware", "Authorizer")

		userData, ok := ctx.Value("user").(*auth.UserData)
		if !ok {
			log.ErrorContext(
				ctx,
				"fail authentication",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		if role != userData.Role {
			log.ErrorContext(
				ctx,
				"fail authentication role mismatch",
				"middleware role", role,
				"user role", userData.Role,
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		next.ServeHTTP(w, r)
	}
	return m.Authenticator(http.HandlerFunc(h))
}

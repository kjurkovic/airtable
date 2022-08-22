package middleware

import (
	"context"
	"net/http"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/models"
)

/**
 * Role middleware should be used after auth middleware since it's using auth Claims model
 * which is extracted from Authorization header value (JWT token)
 *
 */

type KeyAuthor struct{}

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims := r.Context().Value(KeyClaims{}).(*models.Claims)
		email := claims.Username

		user, err := datastore.UserDao.GetByEmail(email)

		if err != nil {
			models.UserNotFoundError.SendErrorResponse(w, http.StatusBadRequest)
			return
		}

		if user.Role != models.ADMIN && user.Role != models.MANAGER {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), KeyAuthor{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

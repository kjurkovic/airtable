package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kjurkovic/airtable/service/workspace/models"
)

type KeyWorkspace struct{}

func MiddlewareValidateWorkspace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		workspace := models.Workspace{}

		err := workspace.Deserialize(r.Body)
		if err != nil {
			http.Error(rw, "Unable to parse request", http.StatusBadRequest)
			return
		}

		err = workspace.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating workspace: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyWorkspace{}, workspace)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

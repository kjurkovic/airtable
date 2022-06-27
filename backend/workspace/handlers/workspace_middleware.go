package handlers

import (
	"context"
	"fmt"
	"net/http"
	"workspace/errors"
	"workspace/models"
)

// middleware
func (handler *WorkspaceHandler) MiddlewareAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")

		// TODO: workspace routes not working
		handler.logger.Print("Hello middleware")

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodGet, handler.config.Address, nil)
		req.Header.Set("Authorization", authorizationHeader)

		handler.logger.Print(req)
		res, err := client.Do(req)

		handler.logger.Print(res)

		if err != nil {
			errors.ServerError.SendErrorResponse(rw, http.StatusInternalServerError)
			return
		}

		if res.StatusCode == http.StatusOK {
			user := &models.User{}
			err = user.Deserialize(res.Body)

			if err != nil {
				errors.SerializationError.SendErrorResponse(rw, http.StatusBadGateway)
				return
			}

			ctx := context.WithValue(r.Context(), models.UserKey{}, user)
			next.ServeHTTP(rw, r.WithContext(ctx))
		} else {
			errors.WrongCredentials.SendErrorResponse(rw, http.StatusUnauthorized)
		}
	})
}

func (handler *WorkspaceHandler) MiddlewareValidateWorkspace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		workspace := models.Workspace{}

		err := workspace.Deserialize(r.Body)
		if err != nil {
			handler.logger.Println("[ERROR] deserializing workspace", err)
			http.Error(rw, "Unable to parse request", http.StatusBadRequest)
			return
		}

		err = workspace.Validate()
		if err != nil {
			handler.logger.Println("[ERROR] validating workspace", err)
			http.Error(rw, fmt.Sprintf("Error validating workspace: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyWorkspace{}, workspace)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

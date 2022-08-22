package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/models"
)

type KeyClaims struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("authorization")
		authorizationHeader = strings.TrimPrefix(authorizationHeader, "Bearer ")

		if len(authorizationHeader) == 0 {
			models.WrongCredentials.SendErrorResponse(w, http.StatusUnauthorized)
			return
		}

		claims, err := validateToken(authorizationHeader)

		if err != nil {
			models.WrongCredentials.SendErrorResponse(w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), KeyClaims{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(input string) (*models.Claims, error) {
	config, _ := config.Load()
	signingKey := []byte(config.Server.Secret)

	token, _ := jwt.ParseWithClaims(input, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("wrong JWT Claim model")
	}
}

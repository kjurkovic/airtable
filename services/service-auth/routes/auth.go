package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/middleware"
	"github.com/kjurkovic/airtable/service/auth/service"
	"github.com/kjurkovic/airtable/service/auth/util"
)

type AuthorizationRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (authRoutes *AuthorizationRoutes) Prefix() string {
	return "/auth"
}

func (authRoutes *AuthorizationRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.JsonResponseMiddleware}
}

func (authRoutes *AuthorizationRoutes) Setup(router *mux.Router) {
	authRoutes.Log.Info("Initializing auth service routes")
	authService := &service.AuthService{
		Log:    authRoutes.Log,
		Config: authRoutes.Config,
	}

	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/login", authService.Login)
	post.HandleFunc("/register", authService.Register)
	post.HandleFunc("/refresh-token", authService.RefreshToken)
}

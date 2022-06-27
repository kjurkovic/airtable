package router

import (
	"auth/config"
	"auth/database"
	"auth/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	log    *log.Logger
	config *config.Config
}

func (wr *AuthRouter) pathPrefix() string {
	return "/auth"
}

func (wr *AuthRouter) setup(router *mux.Router) {
	wh := handlers.Auth(wr.log, &database.UserDao{Conn: *database.Instance}, &wr.config.Authentication)

	registerRouter := router.Methods(http.MethodPost).Subrouter()
	registerRouter.HandleFunc("/register", wh.Register)
	registerRouter.Use(wh.MiddlewareUser, wh.MiddlewareRegistration)

	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", wh.Login)
	loginRouter.Use(wh.MiddlewareUser, wh.MiddlewareLogin)

	authRouter := router.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/refresh-token", wh.RefreshToken)
	authRouter.HandleFunc("/forgot-password", wh.ForgetPassword)
	authRouter.HandleFunc("/reset-password", wh.ResetPassword)
}

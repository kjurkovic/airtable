package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/middleware"
	"github.com/kjurkovic/airtable/service/auth/service"
	"github.com/kjurkovic/airtable/service/auth/util"
)

type UserRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (userRoutes *UserRoutes) Prefix() string {
	return "/users"
}

func (userRoutes *UserRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.AuthMiddleware, middleware.RoleMiddleware, middleware.JsonResponseMiddleware}
}

var userIdRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

func (userRoutes *UserRoutes) Setup(router *mux.Router) {
	userRoutes.Log.Info("Initializing user service routes")
	userService := &service.UserService{
		Log:    userRoutes.Log,
		Config: userRoutes.Config,
	}

	// TODO only for product admins - add middleware and roles -> available only for other services -> add middleware and secret for that as well

	// GET router
	get := router.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("", userService.GetUsers)
	get.HandleFunc(fmt.Sprintf("/{id:%s}", userIdRegex), userService.GetUser)

	// PATCH router
	patch := router.Methods(http.MethodPatch).Subrouter()
	patch.HandleFunc(fmt.Sprintf("/{id:%s}", userIdRegex), userService.EditUser)
	patch.HandleFunc(fmt.Sprintf("/{id:%s}/password", userIdRegex), userService.EditUserPassword)

	// POST router
	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", userService.CreateUser)

	// DELETE router
	delete := router.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc(fmt.Sprintf("/{id:%s}", userIdRegex), userService.DeleteUser)
}

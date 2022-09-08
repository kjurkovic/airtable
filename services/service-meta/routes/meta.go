package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/meta/config"
	"github.com/kjurkovic/airtable/service/meta/middleware"
	"github.com/kjurkovic/airtable/service/meta/service"
	"github.com/kjurkovic/airtable/service/meta/util"
)

var uuidRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type MetaRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (routes *MetaRoutes) Prefix() string {
	return "/meta"
}

func (routes *MetaRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.JsonResponseMiddleware}
}

func (routes *MetaRoutes) Setup(router *mux.Router) {
	routes.Log.Info("Initializing meta service routes")
	service := &service.MetaService{
		Log:    routes.Log,
		Config: routes.Config,
	}

	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", service.Create)
	post.Use(middleware.AuthMiddleware)

	put := router.Methods(http.MethodPut).Subrouter()
	put.HandleFunc(fmt.Sprintf("/{id:%s}", uuidRegex), service.Update)
	put.Use(middleware.AuthMiddleware)

	protectedGet := router.Methods(http.MethodGet).Subrouter()
	protectedGet.HandleFunc(fmt.Sprintf("/user/{userId:%s}", uuidRegex), service.GetAll)
	protectedGet.Use(middleware.AuthMiddleware)

	get := router.Methods(http.MethodGet).Subrouter()
	get.HandleFunc(fmt.Sprintf("/{metaId:%s}", uuidRegex), service.GetOne)

	delete := router.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc(fmt.Sprintf("/{id:%s}", uuidRegex), service.Delete)
	delete.Use(middleware.AuthMiddleware)
}

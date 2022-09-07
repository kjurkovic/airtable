package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/data/config"
	"github.com/kjurkovic/airtable/service/data/middleware"
	"github.com/kjurkovic/airtable/service/data/service"
	"github.com/kjurkovic/airtable/service/data/util"
)

var uuidRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type DataRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (routes *DataRoutes) Prefix() string {
	return "/data"
}

func (routes *DataRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{middleware.JsonResponseMiddleware}
}

func (routes *DataRoutes) Setup(router *mux.Router) {
	routes.Log.Info("Initializing data service routes")
	service := &service.DataService{
		Log:    routes.Log,
		Config: routes.Config,
	}

	get := router.Methods(http.MethodGet).Subrouter()
	get.HandleFunc(fmt.Sprintf("/{metaId:%s}", uuidRegex), service.Get)
	get.Use(middleware.AuthMiddleware)

	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc(fmt.Sprintf("/{metaId:%s}", uuidRegex), service.Create)
}

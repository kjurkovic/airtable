package routes

import (
	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/notification/config"
	"github.com/kjurkovic/airtable/service/notification/util"
)

type IRoutes interface {
	Setup(router *mux.Router)
	Prefix() string
	RouteMiddleware() []mux.MiddlewareFunc
}

type Routes struct {
	Router *mux.Router
	Logger *util.Logger
	Config *config.Config
}

func (r *Routes) Initialize() {
	r.Logger.Info("Initializing service routes")

	routes := []IRoutes{
		&HealthCheck{Log: r.Logger, Config: r.Config},
		&MailerRoutes{Log: r.Logger, Config: r.Config},
	}

	for _, route := range routes {
		subrouter := r.Router.PathPrefix(route.Prefix()).Subrouter()
		for _, mware := range route.RouteMiddleware() {
			subrouter.Use(mware)
		}
		route.Setup(subrouter)
	}
}

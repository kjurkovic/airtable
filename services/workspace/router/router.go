package router

import (
	"log"
	"workspace/config"

	"github.com/gorilla/mux"
)

type Router struct{}

type RouteHandler interface {
	pathPrefix() string
	setup(r *mux.Router)
}

func (r *Router) Setup(log *log.Logger, router *mux.Router, config *config.Config) {
	routeHandlers := []RouteHandler{
		&SwaggerRouter{log},
		&WorkspaceRouter{log, &config.Authorization},
	}

	for _, handler := range routeHandlers {
		if len(handler.pathPrefix()) == 0 {
			handler.setup(router)
		} else {
			handler.setup(router.PathPrefix(handler.pathPrefix()).Subrouter())
		}
	}
}

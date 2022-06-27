package router

import (
	"auth/config"
	"log"

	"github.com/gorilla/mux"
)

type Router struct{}

type RouteHandler interface {
	pathPrefix() string
	setup(r *mux.Router)
}

func (r *Router) Setup(log *log.Logger, config *config.Config, router *mux.Router) {
	routeHandlers := []RouteHandler{
		&SwaggerRouter{log},
		&AuthRouter{log, config},
		&UserRouter{log, config},
	}

	for _, handler := range routeHandlers {
		if len(handler.pathPrefix()) == 0 {
			handler.setup(router)
		} else {
			handler.setup(router.PathPrefix(handler.pathPrefix()).Subrouter())
		}
	}
}

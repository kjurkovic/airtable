package router

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

type SwaggerRouter struct {
	log *log.Logger
}

func (sw *SwaggerRouter) pathPrefix() string {
	return ""
}

func (sw *SwaggerRouter) setup(router *mux.Router) {
	sw.log.Print("Swagger routes setup")
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := middleware.Redoc(ops, nil)
	swagger := router.Methods(http.MethodGet).Subrouter()
	swagger.Handle("/docs", docsHandler)
	swagger.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
}

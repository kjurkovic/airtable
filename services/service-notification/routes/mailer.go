package routes

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/notification/config"
	"github.com/kjurkovic/airtable/service/notification/mailer"
	"github.com/kjurkovic/airtable/service/notification/service"
	"github.com/kjurkovic/airtable/service/notification/util"

	"github.com/gorilla/mux"
)

type MailerRoutes struct {
	Log    *util.Logger
	Config *config.Config
}

func (routes *MailerRoutes) Prefix() string {
	return "/notification"
}

func (routes *MailerRoutes) RouteMiddleware() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{}
}

func (routes *MailerRoutes) Setup(router *mux.Router) {
	service := service.MailerService{
		Log:    routes.Log,
		Config: routes.Config,
		Sender: &mailer.Sender{
			Config: &routes.Config.Mailer,
		},
	}

	// POST router
	post := router.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("", service.Post)
}

package service

import (
	"net/http"

	"github.com/kjurkovic/airtable/service/notification/config"
	"github.com/kjurkovic/airtable/service/notification/mailer"
	"github.com/kjurkovic/airtable/service/notification/models"
	"github.com/kjurkovic/airtable/service/notification/util"
)

type MailerService struct {
	Log    *util.Logger
	Config *config.Config
	Sender *mailer.Sender
}

func (service *MailerService) Post(rw http.ResponseWriter, r *http.Request) {
	service.Log.Info("New notification received")

	message := &models.Message{}
	message.Deserialize(r.Body)

	service.Sender.Mail(message)
}

package wrappers

import (
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/util"
)

type NotificationServiceInitializer struct{}

type notificationService struct {
	client *NotificationClient
}

var Notification *notificationService

func (*NotificationServiceInitializer) Initialize() {

	logger := util.New()
	config, _ := config.Load()

	Notification = &notificationService{
		client: &NotificationClient{
			BaseUrl: config.Services.Notification,
			Log:     logger.Error,
		},
	}
}

func (service *notificationService) SendEmail(to string, email string, subject string, text string) {
	service.client.SendEmail(to, email, subject, text)
}

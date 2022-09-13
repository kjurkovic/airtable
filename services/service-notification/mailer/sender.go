package mailer

import (
	"log"

	"github.com/kjurkovic/airtable/service/notification/config"
	"github.com/kjurkovic/airtable/service/notification/models"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Sender struct {
	Config *config.Mailer
}

func (sender *Sender) Mail(message *models.Message) {

	from := mail.NewEmail(sender.Config.From.Name, sender.Config.From.Email)

	recipient := mail.NewEmail(message.To, message.Email)
	plainTextContent := message.Text
	htmlContent := message.Text
	body := mail.NewSingleEmail(from, message.Subject, recipient, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sender.Config.Key)
	resp, err := client.Send(body)

	log.Println(resp, err)
}

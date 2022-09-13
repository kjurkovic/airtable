package wrappers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Log func(error)

type AuditClient struct {
	BaseUrl string
	Log     Log
}

type NotificationClient struct {
	BaseUrl string
	Log     Log
}

var (
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
	}
)

func (service *NotificationClient) SendEmail(to string, email string, subject string, text string) {
	message := &Message{
		To:      to,
		Email:   email,
		Subject: subject,
		Text:    text,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(message)

	if err != nil {
		service.Log(err)
		return
	}

	resp, err := client.Post(fmt.Sprintf("%s/%s", service.BaseUrl, "notification"), "application/json", &buf)

	if err != nil {
		service.Log(err)
		return
	}

	resp.Body.Close()
}

func (service *AuditClient) WriteLog(userId uuid.UUID, obj string, auditType AuditType) {
	event := &Event{
		UserId:      userId,
		AuditObject: obj,
		Type:        auditType,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(event)

	if err != nil {
		service.Log(err)
		return
	}

	resp, err := client.Post(fmt.Sprintf("%s/%s", service.BaseUrl, "logs"), "application/json", &buf)

	fmt.Print("Sent request to audit:", resp, err)

	if err != nil {
		service.Log(err)
		return
	}

	resp.Body.Close()
}

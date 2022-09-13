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

type MetaClient struct {
	BaseUrl string
	Log     Log
}

type AuditClient struct {
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

func (service *MetaClient) Get(metaId uuid.UUID) (*Meta, error) {
	resp, err := client.Get(fmt.Sprintf("%s/%s/%s", service.BaseUrl, "meta", metaId.String()))

	if err != nil {
		service.Log(err)
		return nil, err
	}

	meta := &Meta{}
	meta.Deserialize(resp.Body)

	if err != nil {
		service.Log(err)
		return nil, err
	}

	resp.Body.Close()

	return meta, nil
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

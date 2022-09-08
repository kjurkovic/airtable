package wrappers

import (
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

package wrappers

import (
	"github.com/google/uuid"

	"github.com/kjurkovic/airtable/service/data/config"
	"github.com/kjurkovic/airtable/service/data/util"
)

type MetaServiceInitializer struct{}

type metaService struct {
	client *MetaClient
}

var MetaApi *metaService

func (*MetaServiceInitializer) Initialize() {

	logger := util.New()
	config, _ := config.Load()

	MetaApi = &metaService{
		client: &MetaClient{
			BaseUrl: config.Services.Meta,
			Log:     logger.Error,
		},
	}
}

func (service *metaService) Get(metaId uuid.UUID) (*Meta, error) {
	return service.client.Get(metaId)
}

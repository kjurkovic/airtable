package wrappers

import (
	"github.com/google/uuid"

	"github.com/kjurkovic/airtable/service/meta/config"
	"github.com/kjurkovic/airtable/service/meta/util"
)

type AuditServiceInitializer struct{}

type auditService struct {
	client *AuditClient
}

var Audit *auditService

func (*AuditServiceInitializer) Initialize() {

	logger := util.New()
	config, _ := config.Load()

	Audit = &auditService{
		client: &AuditClient{
			BaseUrl: config.Services.Audit,
			Log:     logger.Error,
		},
	}
}

func (service *auditService) SendEvent(userId uuid.UUID, obj string, auditType AuditType) {
	service.client.WriteLog(userId, obj, auditType)
}

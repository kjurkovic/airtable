package wrappers

import (
	"github.com/google/uuid"
	audit "github.com/kjurkovic/airtable/clients/service-audit"
	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/util"
)

type AuditServiceInitializer struct{}

type auditService struct {
	client *audit.AuditClient
}

var Audit *auditService

func (*AuditServiceInitializer) Initialize() {

	logger := util.New()
	config, _ := config.Load()

	Audit = &auditService{
		client: &audit.AuditClient{
			BaseUrl: config.Services.Audit,
			Log:     logger.Error,
		},
	}
}

func (service *auditService) SendEvent(userId uuid.UUID, obj string, auditType audit.AuditType) {
	service.client.WriteLog(userId, obj, auditType)
}

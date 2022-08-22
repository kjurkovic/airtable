package service

import (
	"github.com/kjurkovic/airtable/services/service-audit/config"
	"github.com/kjurkovic/airtable/services/service-audit/util"
)

type AuditService struct {
	Log    *util.Logger
	Config *config.Config
}

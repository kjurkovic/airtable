package service

import (
	"github.com/kjurkovic/airtable/service/workspace/config"
	"github.com/kjurkovic/airtable/service/workspace/util"
)

type WorkspaceService struct {
	Log    *util.Logger
	Config *config.Config
}

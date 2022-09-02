package service

import (
	"github.com/kjurkovic/airtable/service/meta/config"
	"github.com/kjurkovic/airtable/service/meta/util"
)

type MetaService struct {
	Log    *util.Logger
	Config *config.Config
}

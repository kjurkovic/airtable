package service

import (
	"github.com/kjurkovic/airtable/service/data/config"
	"github.com/kjurkovic/airtable/service/data/util"
)

type DataService struct {
	Log    *util.Logger
	Config *config.Config
}

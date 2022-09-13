package main

import (
	"github.com/kjurkovic/airtable/service/meta/loaders"
	"github.com/kjurkovic/airtable/service/meta/wrappers"
)

func main() {

	initializers := []loaders.Loader{
		&wrappers.AuditServiceInitializer{},
		&loaders.Database{},
		&loaders.App{},
	}

	for _, initializer := range initializers {
		initializer.Initialize()
	}
}

package main

import (
	"github.com/kjurkovic/airtable/service/workspace/loaders"
	"github.com/kjurkovic/airtable/service/workspace/wrappers"
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

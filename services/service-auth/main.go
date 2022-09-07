package main

import (
	"github.com/kjurkovic/airtable/service/auth/loaders"
	"github.com/kjurkovic/airtable/service/auth/wrappers"
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

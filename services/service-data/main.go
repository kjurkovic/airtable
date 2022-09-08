package main

import (
	"github.com/kjurkovic/airtable/service/data/loaders"
	"github.com/kjurkovic/airtable/service/data/wrappers"
)

func main() {

	initializers := []loaders.Loader{
		&wrappers.MetaServiceInitializer{},
		&loaders.Database{},
		&loaders.App{},
	}

	for _, initializer := range initializers {
		initializer.Initialize()
	}
}

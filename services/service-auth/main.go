package main

import (
	"github.com/kjurkovic/airtable/service/auth/loaders"
)

func main() {

	initializers := []loaders.Loader{
		&loaders.Database{},
		&loaders.App{},
	}

	for _, initializer := range initializers {
		initializer.Initialize()
	}
}

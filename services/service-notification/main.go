package main

import (
	"github.com/kjurkovic/airtable/service/notification/loaders"
)

func main() {

	initializers := []loaders.Loader{
		&loaders.App{},
	}

	for _, initializer := range initializers {
		initializer.Initialize()
	}
}

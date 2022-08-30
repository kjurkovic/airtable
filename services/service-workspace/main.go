package main

import (
	"path/filepath"
	"runtime"

	"github.com/kjurkovic/airtable/service/workspace/loaders"
)

var (
	_, b, _, _  = runtime.Caller(0)
	projectRoot = filepath.Dir(b)
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

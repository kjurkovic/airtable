package loaders

import (
	"fmt"

	"github.com/kjurkovic/airtable/service/auth/config"
	"github.com/kjurkovic/airtable/service/auth/datastore"
)

type Database struct{}

func (database *Database) Initialize() {
	config, err := config.Load()

	if err != nil {
		panic(err)
	}

	db := datastore.Database{Config: config}
	db.Connect()
	fmt.Println("Database loaded successfully")
}

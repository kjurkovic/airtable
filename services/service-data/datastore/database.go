package datastore

import (
	"fmt"

	"github.com/kjurkovic/airtable/service/data/config"
	"github.com/kjurkovic/airtable/service/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Config *config.Config
}

var (
	MetaDao  *MetaRepository
	FieldDao *FieldRepository
)

func (database *Database) Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=Europe/Zagreb",
		database.Config.Database.Ip,
		database.Config.Database.Username,
		database.Config.Database.Password,
		database.Config.Database.Name,
		database.Config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Field{},
		&models.Meta{},
	)

	MetaDao = &MetaRepository{database: db}
	FieldDao = &FieldRepository{database: db}
}

package datastore

import (
	"fmt"

	"github.com/kjurkovic/airtable/services/service-audit/config"
	"github.com/kjurkovic/airtable/services/service-audit/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Config *config.Config
}

var (
	AuditDao *AuditRepository
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
		&models.Event{},
	)

	AuditDao = &AuditRepository{database: db}
}

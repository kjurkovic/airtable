package datastore

import (
	"fmt"

	"github.com/kjurkovic/airtable/service/auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"honnef.co/go/tools/config"
)

type Database struct {
	Config *config.Config
}

var (
	AuthDao         *AuthRepository
	UserDao         *UserRepository
	RefreshTokenDao *RefreshTokenRepository
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
		&models.User{},
		&models.RefreshToken{},
	)

	AuthDao = &AuthRepository{database: db}
	UserDao = &UserRepository{database: db}
	RefreshTokenDao = &RefreshTokenRepository{database: db}
}

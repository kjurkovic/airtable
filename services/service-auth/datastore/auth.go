package datastore

import "gorm.io/gorm"

type AuthRepository struct {
	database *gorm.DB
}

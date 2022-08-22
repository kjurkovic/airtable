package datastore

import (
	"time"

	"github.com/kjurkovic/airtable/service/auth/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	database *gorm.DB
}

func (repo *RefreshTokenRepository) GetByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	tx := repo.database.First(&refreshToken, "token = ?", token)
	return &refreshToken, tx.Error
}

func (repo *RefreshTokenRepository) Save(refreshToken *models.RefreshToken) (int64, error) {
	tx := repo.database.Create(refreshToken)
	return tx.RowsAffected, tx.Error
}

func (repo *RefreshTokenRepository) DeleteOutdated() (int64, error) {
	now := time.Now().UnixMilli()
	tx := repo.database.Unscoped().Where("expires_at < ?", now).Delete(&models.RefreshToken{})
	return tx.RowsAffected, tx.Error
}

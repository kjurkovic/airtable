package datastore

import (
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/meta/models"
	"gorm.io/gorm"
)

type FieldRepository struct {
	database *gorm.DB
}

func (repo *FieldRepository) Create(field *models.Field) (*models.Field, error) {
	tx := repo.database.Create(field)
	return field, tx.Error
}

func (repo *FieldRepository) GetAll(metaId uuid.UUID, page int, size int) ([]models.Field, error) {
	var data []models.Field
	tx := repo.database.Where("meta_id = ?", metaId).Find(&data)
	return data, tx.Error
}

func (repo *FieldRepository) GetOne(id uuid.UUID) (*models.Meta, error) {
	var model *models.Meta
	tx := repo.database.First(model, "id = ?", id)
	return model, tx.Error
}

func (repo *FieldRepository) Delete(metaId uuid.UUID) error {
	tx := repo.database.Where("metaId = ? ", metaId).Delete(&models.Field{})
	return tx.Error
}

package datastore

import (
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/data/models"
	"gorm.io/gorm"
)

type DataRepository struct {
	database *gorm.DB
}

func (repo *DataRepository) Create(meta *models.Data) (*models.Data, error) {
	tx := repo.database.Create(meta)
	return meta, tx.Error
}

func (repo *DataRepository) GetAll(id uuid.UUID, page int, size int) (*models.Pageable[models.Data], error) {
	var data []models.Data
	var count int64
	var tx *gorm.DB

	tx = repo.database.Model(&models.Data{}).Where("meta_id = ?", id).Count(&count)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = repo.database.Limit(size).Offset((page-1)*size).Where("meta_id = ?", id).Find(&data)

	pageable := models.Paginate(&data, count, page, size)
	return pageable, tx.Error
}

func (repo *DataRepository) GetOne(id uuid.UUID) (*models.Data, error) {
	var model *models.Data
	tx := repo.database.First(model, "id = ?", id)
	return model, tx.Error
}

func (repo *DataRepository) Update(meta models.Data) error {
	tx := repo.database.Model(&models.Data{Id: meta.Id}).Updates(meta)
	return tx.Error
}

func (repo *DataRepository) Delete(id uuid.UUID, userId uuid.UUID) (int64, error) {
	tx := repo.database.Where("id = ? AND userId = ?", id, userId).Delete(&models.Data{})
	return tx.RowsAffected, tx.Error
}

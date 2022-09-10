package datastore

import (
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/meta/models"
	"gorm.io/gorm"
)

type MetaRepository struct {
	database *gorm.DB
}

func (repo *MetaRepository) Create(meta *models.Meta) (*models.Meta, error) {
	tx := repo.database.Create(meta)
	return meta, tx.Error
}

func (repo *MetaRepository) GetAllWorkspace(id uuid.UUID, workspaceId uuid.UUID, page int, size int) (*models.Pageable[models.Meta], error) {
	var data []models.Meta
	var count int64
	var tx *gorm.DB

	tx = repo.database.Model(&models.Meta{}).Where("user_id = ? AND workspace_id = ?", id, workspaceId).Count(&count)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = repo.database.Preload("Fields").Limit(size).Offset((page-1)*size).Where("user_id = ? AND workspace_id = ?", id, workspaceId).Find(&data)

	pageable := models.Paginate(&data, count, page, size)
	return pageable, tx.Error
}

func (repo *MetaRepository) GetAll(id uuid.UUID, page int, size int) (*models.Pageable[models.Meta], error) {
	var data []models.Meta
	var count int64
	var tx *gorm.DB

	tx = repo.database.Model(&models.Meta{}).Where("user_id = ?", id).Count(&count)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = repo.database.Preload("Fields").Limit(size).Offset((page-1)*size).Where("user_id = ?", id).Find(&data)

	pageable := models.Paginate(&data, count, page, size)
	return pageable, tx.Error
}

func (repo *MetaRepository) GetOne(id uuid.UUID) (*models.Meta, error) {
	var model *models.Meta
	tx := repo.database.Preload("Fields").First(&model, "id = ?", id)
	return model, tx.Error
}

func (repo *MetaRepository) Update(meta models.Meta) error {
	tx := repo.database.Model(&models.Meta{Id: meta.Id}).Updates(meta)
	return tx.Error
}

func (repo *MetaRepository) Delete(id uuid.UUID, userId uuid.UUID) error {
	return repo.database.Transaction(func(tx *gorm.DB) error {
		if txMeta := repo.database.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Meta{}); txMeta.Error != nil {
			return txMeta.Error
		}
		if txFields := repo.database.Where("meta_id = ? ", id).Delete(&models.Field{}); txFields.Error != nil {
			return txFields.Error
		}
		return nil
	})
}

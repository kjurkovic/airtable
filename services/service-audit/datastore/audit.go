package datastore

import (
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/services/service-audit/models"
	"gorm.io/gorm"
)

type AuditRepository struct {
	database *gorm.DB
}

func (repo *AuditRepository) AddLog(event *models.Event) error {
	tx := repo.database.Create(event)
	return tx.Error
}

func (repo *AuditRepository) GetUserLogs(id uuid.UUID, auditType string, page int, size int) (*models.Pageable[models.Event], error) {
	var events []models.Event
	var count int64
	var tx *gorm.DB

	if auditType == "" {
		tx = repo.database.Model(&models.Event{}).Where("user_id = ?", id).Count(&count)
	} else {
		tx = repo.database.Model(&models.Event{}).Where("user_id = ? AND type = ?", id, auditType).Count(&count)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	if auditType == "" {
		tx = repo.database.Limit(size).Offset((page-1)*size).Where("user_id = ?", id).Find(&events)
	} else {
		tx = repo.database.Limit(size).Offset((page-1)*size).Where("user_id = ? AND type = ?", id, auditType).Find(&events)
	}

	pageable := models.Paginate(&events, count, page, size)
	return pageable, tx.Error
}

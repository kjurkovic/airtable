package datastore

import (
	"github.com/kjurkovic/airtable/service/workspace/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	database *gorm.DB
}

func (repo *WorkspaceRepository) Insert(item *models.Workspace) (int64, error) {
	tx := repo.database.Create(item)
	return tx.RowsAffected, tx.Error
}

func (repo *WorkspaceRepository) Get(id uuid.UUID, userId uuid.UUID) (*models.Workspace, error) {
	var workspace models.Workspace
	tx := repo.database.First(&workspace, "id = ? AND user_id = ?", id, userId)
	return &workspace, tx.Error
}

func (repo *WorkspaceRepository) GetAll(userId uuid.UUID) (models.Workspaces, error) {
	var workspaces models.Workspaces
	tx := repo.database.Find(&workspaces, "user_id = ?", userId)
	return workspaces, tx.Error
}

func (repo *WorkspaceRepository) Update(id uuid.UUID, userId uuid.UUID, item *models.Workspace) (*models.Workspace, error) {
	var workspace models.Workspace
	tx := repo.database.First(&workspace, "id = ? AND user_id = ?", id, userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// on workspace only name is allowed to be updated
	tx = repo.database.Model(&workspace).Where("id = ?", id).Update("name", item.Name)
	return &workspace, tx.Error
}

func (repo *WorkspaceRepository) Delete(id uuid.UUID, userId uuid.UUID) (int64, error) {
	tx := repo.database.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Workspace{})
	return tx.RowsAffected, tx.Error
}

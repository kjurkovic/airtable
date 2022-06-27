package database

import (
	"workspace/models"

	"github.com/google/uuid"
)

type WorkspaceDao struct {
	Conn Database
}

func (dao *WorkspaceDao) Insert(item *models.Workspace) (int64, error) {
	tx := dao.Conn.Db.Create(item)
	return tx.RowsAffected, tx.Error
}

func (dao *WorkspaceDao) Get(id uuid.UUID, userId uuid.UUID) (*models.Workspace, error) {
	var workspace models.Workspace
	tx := dao.Conn.Db.First(&workspace, "id = ? AND user_id = ?", id, userId)
	return &workspace, tx.Error
}

func (dao *WorkspaceDao) GetAll(userId uuid.UUID) (models.Workspaces, error) {
	var workspaces models.Workspaces
	tx := dao.Conn.Db.Find(&workspaces, "user_id = ?", userId)
	return workspaces, tx.Error
}

func (dao *WorkspaceDao) Update(id uuid.UUID, userId uuid.UUID, item *models.Workspace) (*models.Workspace, error) {
	var workspace models.Workspace
	tx := dao.Conn.Db.First(&workspace, "id = ? AND user_id = ?", id, userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// on workspace only name is allowed to be updated
	tx = dao.Conn.Db.Model(&workspace).Where("id = ?", id).Update("name", item.Name)
	return &workspace, tx.Error
}

func (dao *WorkspaceDao) Delete(id uuid.UUID, userId uuid.UUID) (int64, error) {
	tx := dao.Conn.Db.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Workspace{})
	return tx.RowsAffected, tx.Error
}

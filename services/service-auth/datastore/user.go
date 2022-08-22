package datastore

import (
	"github.com/google/uuid"
	"github.com/kjurkovic/airtable/service/auth/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func (repo *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	tx := repo.database.First(&user, "email = ?", email)
	return &user, tx.Error
}

func (repo *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User
	tx := repo.database.First(&user, "id = ?", id)
	return &user, tx.Error
}

func (repo *UserRepository) GetAll(page int, size int) (*models.Pageable[models.User], error) {
	var users []models.User
	var count int64
	tx := repo.database.Model(&models.User{}).Count(&count)

	if tx.Error != nil {
		return nil, tx.Error
	}

	tx = repo.database.Limit(size).Offset((page - 1) * size).Find(&users)

	pageable := models.Paginate(&users, count, page, size)
	return pageable, tx.Error
}

func (repo *UserRepository) Save(user *models.User) (int64, error) {
	tx := repo.database.Create(user)
	return tx.RowsAffected, tx.Error
}

func (repo *UserRepository) Update(user *models.User) error {
	tx := repo.database.Model(&models.User{Id: user.Id}).Updates(user)
	return tx.Error
}

func (repo *UserRepository) Delete(id uuid.UUID) (int64, error) {
	tx := repo.database.Delete(&models.User{}, id)
	return tx.RowsAffected, tx.Error
}

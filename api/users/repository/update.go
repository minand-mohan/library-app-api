package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/database/models"
)

// Update/Partial update a user by id
func (repo *UserRepositoryImpl) UpdateByUserId(id uuid.UUID, user *models.User) (*models.User, error) {
	result := repo.db.Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/database/models"
)

func (repo *UserRepositoryImpl) DeleteByUserId(id uuid.UUID) error {
	var user models.User
	result := repo.db.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

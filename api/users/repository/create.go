package repository

import "github.com/minand-mohan/library-app-api/database/models"

// CreateUser creates a new user
func (repo *UserRepositoryImpl) CreateUser(userObj *models.User) error {
	result := repo.db.Create(&userObj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"
)

// List all users
func (repo *UserRepositoryImpl) FindAllUsers(queryParams *dto.UserQueryParams) ([]models.User, error) {
	var users []models.User
	dbQuery := GenerateDbQueries(queryParams)
	result := repo.db.
		Where(dbQuery.Email).
		Where(dbQuery.Username).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Retrieve a user by their ID
func (repo *UserRepositoryImpl) FindByUserId(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Used by create to check for any duplicate values
func (repo *UserRepositoryImpl) FindByEmailOrUsernameOrPhone(email string, username string, phone string) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "email = ? OR username = ? OR phone = ?", email, username, phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

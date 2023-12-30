package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"
)

func (repo *UserRepositoryImpl) FindByEmailOrUsernameOrPhoneNotUuid(email string, username string, phone string, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "(email = ? OR username = ? OR phone = ?) AND id != ?", email, username, phone, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

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

func (repo *UserRepositoryImpl) FindByUserId(id uuid.UUID) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

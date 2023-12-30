package repository

import "github.com/minand-mohan/library-app-api/database/models"

func (repo *UserRepositoryImpl) CreateUser(userObj *models.User) error {
	result := repo.db.Create(&userObj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepositoryImpl) FindByEmailOrUsernameOrPhone(email string, username string, phone string) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "email = ? OR username = ? OR phone = ?", email, username, phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

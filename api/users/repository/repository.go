package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/database/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(userObj *models.User) error
	FindByEmailOrUsernameOrPhone(email string, username string, phone string) (*models.User, error)
	FindByEmailOrUsernameOrPhoneNotUuid(email string, username string, phone string, id uuid.UUID) (*models.User, error)
	FindAllUsers() ([]models.User, error)
	FindByUserId(id uuid.UUID) (*models.User, error)
	UpdateByUserId(id uuid.UUID, user *models.User) (*models.User, error)
	// DeleteById(id uint) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

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

func (repo *UserRepositoryImpl) FindByEmailOrUsernameOrPhoneNotUuid(email string, username string, phone string, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := repo.db.First(&user, "(email = ? OR username = ? OR phone = ?) AND id != ?", email, username, phone, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
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

func (repo *UserRepositoryImpl) UpdateByUserId(id uuid.UUID, user *models.User) (*models.User, error) {
	result := repo.db.Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

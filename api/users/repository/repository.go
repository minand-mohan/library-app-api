package repository

import (
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(userObj *models.User) error
	FindByEmailOrUsernameOrPhone(email string, username string, phone string) (*models.User, error)
	FindByEmailOrUsernameOrPhoneNotUuid(email string, username string, phone string, id uuid.UUID) (*models.User, error)
	FindAllUsers(queryParams *dto.UserQueryParams) ([]models.User, error)
	FindByUserId(id uuid.UUID) (*models.User, error)
	UpdateByUserId(id uuid.UUID, user *models.User) (*models.User, error)
	DeleteByUserId(id uuid.UUID) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

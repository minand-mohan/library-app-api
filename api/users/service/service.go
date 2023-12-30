package service

import (
	"errors"
	"fmt"

	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/api/users/repository"
	"github.com/minand-mohan/library-app-api/database/models"
	"github.com/minand-mohan/library-app-api/utils"
)

type UserService interface {
	CreateUser(userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error)
	// FindAll() ([]model.User, error)
	// FindById(id uint) (*model.User, error)
	// UpdateById(id uint, user *model.User) (*model.User, error)
	// DeleteById(id uint) error
}

type UserServiceImpl struct {
	repo   repository.UserRepository
	logger *utils.AppLogger
}

func NewUserService(repo repository.UserRepository, logger utils.AppLogger) UserService {
	return &UserServiceImpl{
		repo:   repo,
		logger: &logger,
	}
}

func (service *UserServiceImpl) CreateUser(userReq *dto.UserRequestBody) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Create user")
	userObj := &models.User{
		Username: &userReq.Username,
		Email:    &userReq.Email,
		Phone:    &userReq.Phone,
	}

	existingUser, err := service.repo.FindByEmailOrUsernameOrPhone(*userObj.Email, *userObj.Username, *userObj.Phone)
	if err == nil {
		service.logger.Error(fmt.Sprintf("UserService: User with email %s, username %s or phone %s already exists", *userObj.Email, *userObj.Username, *userObj.Phone))
		responseContent := map[string]interface{}{
			"id":       existingUser.ID,
			"username": existingUser.Username,
			"email":    existingUser.Email,
			"phone":    existingUser.Phone,
		}
		responseBody := response.HTTPResponse{
			Code:    400,
			Message: "Bad request, user already exists",
			Content: responseContent,
		}
		return &responseBody, errors.New("User already exists")
	}

	err = service.repo.CreateUser(userObj)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while creating user: %s", err))
		responseBody := response.HTTPResponse{
			Code:    500,
			Message: "Internal Server Error",
			Content: map[string]interface{}{},
		}
		return &responseBody, err
	}

	responseContent := map[string]interface{}{
		"id":       userObj.ID,
		"username": userObj.Username,
		"email":    userObj.Email,
		"phone":    userObj.Phone,
	}

	responseBody := response.HTTPResponse{
		Code:    200,
		Message: "User created successfully",
		Content: responseContent,
	}

	return &responseBody, nil
}

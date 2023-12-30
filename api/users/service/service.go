package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/api/users/repository"
	"github.com/minand-mohan/library-app-api/database/models"
	"github.com/minand-mohan/library-app-api/utils"
)

type UserService interface {
	CreateUser(userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error)
	FindAllUsers(queryParams *dto.UserQueryParams) (*response.HTTPResponse, error)
	FindByUserId(id uuid.UUID) (*response.HTTPResponse, error)
	UpdateByUserId(id uuid.UUID, userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error)
	DeleteByUserId(id uuid.UUID) (*response.HTTPResponse, error)
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

func (service *UserServiceImpl) FindAllUsers(queryParams *dto.UserQueryParams) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Find all users")
	users, err := service.repo.FindAllUsers(queryParams)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while finding all users: %s", err))
		return nil, err
	}
	var usersMap []map[string]interface{}
	for _, user := range users {
		userMap := map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"phone":    user.Phone,
		}
		usersMap = append(usersMap, userMap)
	}
	responseContent := response.HTTPResponseContent{
		Count:    len(users),
		Previous: nil,
		Next:     nil,
		Results:  usersMap,
	}
	responseBody := response.HTTPResponse{
		Code:    200,
		Message: "Users found",
		Content: responseContent,
	}
	return &responseBody, nil
}

func (service *UserServiceImpl) FindByUserId(id uuid.UUID) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Find user by id")
	user, err := service.repo.FindByUserId(id)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while finding user by id: %s", err))
		responseBody := response.HTTPResponse{
			Code:    404,
			Message: "User not found.",
			Content: map[string]interface{}{},
		}
		return &responseBody, nil
	}
	responseContent := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.Phone,
	}
	responseBody := response.HTTPResponse{
		Code:    200,
		Message: "User found",
		Content: responseContent,
	}
	return &responseBody, nil
}

func (service *UserServiceImpl) UpdateByUserId(id uuid.UUID, userReqBody *dto.UserRequestBody) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Update user by id")
	userObj := &models.User{
		Username: &userReqBody.Username,
		Email:    &userReqBody.Email,
		Phone:    &userReqBody.Phone,
	}
	_, err := service.repo.FindByUserId(id)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while finding user by id: %s", err))
		responseBody := response.HTTPResponse{
			Code:    404,
			Message: "User not found.",
			Content: map[string]interface{}{},
		}
		return &responseBody, nil
	}

	existingUser, err := service.repo.FindByEmailOrUsernameOrPhoneNotUuid(*userObj.Email, *userObj.Username, *userObj.Phone, id)
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
			Message: "Bad request, values already exists",
			Content: responseContent,
		}
		return &responseBody, errors.New("User already exists")
	}
	updatedUserObj, err := service.repo.UpdateByUserId(id, userObj)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while updating user: %s", err))
		responseBody := response.HTTPResponse{
			Code:    500,
			Message: "Internal Server Error",
			Content: map[string]interface{}{},
		}
		return &responseBody, err
	}
	responseContent := map[string]interface{}{
		"id":       id,
		"username": updatedUserObj.Username,
		"email":    updatedUserObj.Email,
		"phone":    updatedUserObj.Phone,
	}
	responseBody := response.HTTPResponse{
		Code:    200,
		Message: "User updated successfully",
		Content: responseContent,
	}
	return &responseBody, nil
}

func (service *UserServiceImpl) DeleteByUserId(id uuid.UUID) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Delete user by id")
	_, err := service.repo.FindByUserId(id)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while finding user by id: %s", err))
		responseBody := response.HTTPResponse{
			Code:    404,
			Message: "User not found.",
			Content: map[string]interface{}{},
		}
		return &responseBody, nil
	}
	err = service.repo.DeleteByUserId(id)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while deleting user: %s", err))
		responseBody := response.HTTPResponse{
			Code:    500,
			Message: "Internal Server Error",
			Content: map[string]interface{}{},
		}
		return &responseBody, err
	}
	responseBody := response.HTTPResponse{
		Code:    200,
		Message: "User deleted successfully",
		Content: map[string]interface{}{},
	}
	return &responseBody, nil
}

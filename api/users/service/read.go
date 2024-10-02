package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
)

func (service *UserServiceImpl) FindAllUsers(queryParams *dto.UserQueryParams) (*response.HTTPResponse, error) {
	service.logger.Info("User Service: Find all users")
	users, err := service.repo.FindAllUsers(queryParams)
	if err != nil {
		service.logger.Error(fmt.Sprintf("UserService: Error while finding all users: %s", err))
		responseBody := response.HTTPResponse{
			Code:    500,
			Message: "Internal Server Error",
			Content: map[string]interface{}{},
		}
		return &responseBody, err
	}
	if len(users) == 0 {
		service.logger.Error(fmt.Sprintf("UserService: No users found"))
		responseBody := response.HTTPResponse{
			Code:    404,
			Message: "No users found",
			Content: map[string]interface{}{},
		}
		return &responseBody, nil
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
		Message: "Users found successfully",
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

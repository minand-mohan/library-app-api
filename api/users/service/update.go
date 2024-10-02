package service

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"
)

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
		return &responseBody, err
	}

	updatedUserObj, err := service.repo.UpdateByUserId(id, userObj)
	if err != nil {
		// if duplicate key value error return 400
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			service.logger.Error(fmt.Sprintf("UserService: Error while updating user: %s", err))
			responseBody := response.HTTPResponse{
				Code:    400,
				Message: "Bad request, non-unique values",
				Content: map[string]interface{}{},
			}
			return &responseBody, err
		}

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

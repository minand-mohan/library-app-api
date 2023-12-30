package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
)

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

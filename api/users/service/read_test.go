package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	repomocks "github.com/minand-mohan/library-app-api/api/users/repository/mocks"
	"github.com/minand-mohan/library-app-api/database/models"
	"github.com/minand-mohan/library-app-api/utils"
)

func TestListUser(t *testing.T) {

	internalServerError := errors.New("Internal Server Error")

	tc := []struct {
		name                   string
		queryParams            *dto.UserQueryParams
		expectedResponse       *response.HTTPResponse
		expectedError          error
		mockFindAllUsersReturn []models.User
		mockFindAllUserError   error
	}{
		{
			name:        "Find all users successfully",
			queryParams: nil,
			expectedResponse: &response.HTTPResponse{
				Code:    200,
				Message: "Users found successfully",
			},
			expectedError: nil,
			mockFindAllUsersReturn: []models.User{
				generateRandomUser01(),
				generateRandomUser02(),
			},
			mockFindAllUserError: nil,
		},
		{
			name: "Find no users",
			queryParams: &dto.UserQueryParams{
				Username: "test",
			},
			expectedResponse: &response.HTTPResponse{
				Code:    404,
				Message: "No users found",
			},
			expectedError:          nil,
			mockFindAllUsersReturn: []models.User{},
			mockFindAllUserError:   nil,
		},
		{
			name: "Find all users with error",
			queryParams: &dto.UserQueryParams{
				Username: "test",
			},
			expectedResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal Server Error",
			},
			expectedError:          internalServerError,
			mockFindAllUsersReturn: nil,
			mockFindAllUserError:   internalServerError,
		},
	}

	for _, tt := range tc {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := repomocks.NewMockUserRepository(mockCtrl)
			mockUserRepo.EXPECT().FindAllUsers(tt.queryParams).Return(tt.mockFindAllUsersReturn, tt.mockFindAllUserError)

			userService := NewUserService(mockUserRepo, *utils.NewLogger())
			response, err := userService.FindAllUsers(tt.queryParams)
			if err != tt.expectedError {
				t.Errorf("Expected error to be %v, but got %v", tt.expectedError, err)
			}
			if response.Code != tt.expectedResponse.Code {
				t.Errorf("Expected code to be %d, but got %d", tt.expectedResponse.Code, response.Code)
			}
			if response.Message != tt.expectedResponse.Message {
				t.Errorf("Expected message to be %s, but got %s", tt.expectedResponse.Message, response.Message)
			}
		})
	}
}

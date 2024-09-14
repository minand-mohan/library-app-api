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

func TestCreateUser(t *testing.T) {

	test_email := "test@example.com"
	test_name := "test"
	test_phone := "1234567890"

	test_cases_that_require_create_user := map[string]bool{
		"Create User sucessfully":        true,
		"Create User with service error": true,
	}
	test_cases_that_require_find_user := map[string]bool{
		"Create User sucessfully":        true,
		"Create User with existing user": true,
		"Create User with service error": true,
	}
	tc := []struct {
		name                string
		requestbody         *dto.UserRequestBody
		expectedResponse    *response.HTTPResponse
		expectedError       error
		mockFindUserReturn  *models.User
		mockFindUserError   error
		mockCreateUserError error
	}{
		{
			name: "Create User sucessfully",
			requestbody: &dto.UserRequestBody{
				Email:    test_email,
				Username: test_name,
				Phone:    test_phone,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User created successfully",
				Content: map[string]interface{}{},
			},
			expectedError:       nil,
			mockFindUserReturn:  nil,
			mockFindUserError:   errors.New("cannot find user"),
			mockCreateUserError: nil,
		},
		{
			name: "Create User with existing user",
			requestbody: &dto.UserRequestBody{
				Email:    test_email,
				Username: test_name,
				Phone:    test_phone,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    400,
				Message: "Bad request, user already exists",
				Content: map[string]interface{}{},
			},
			expectedError: errors.New("user already exists"),
			mockFindUserReturn: &models.User{
				Email:    &test_email,
				Phone:    &test_phone,
				Username: &test_name,
			},
			mockFindUserError:   nil,
			mockCreateUserError: nil,
		},
		{
			name: "Create User with service error",
			requestbody: &dto.UserRequestBody{
				Email:    test_email,
				Username: test_name,
				Phone:    test_phone,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal Server Error",
				Content: map[string]interface{}{},
			},
			expectedError:       errors.New("Internal Server Error"),
			mockFindUserReturn:  nil,
			mockFindUserError:   errors.New("cannot find user"),
			mockCreateUserError: errors.New("Internal Server Error"),
		},
	}

	for _, tt := range tc {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		t.Run(tt.name, func(t *testing.T) {
			logger := utils.NewLogger()

			mockRepo := repomocks.NewMockUserRepository(mockCtrl)
			if test_cases_that_require_find_user[tt.name] {
				mockRepo.EXPECT().FindByEmailOrUsernameOrPhone(tt.requestbody.Email, tt.requestbody.Username, tt.requestbody.Phone).Return(tt.mockFindUserReturn, tt.mockFindUserError)
			}
			if test_cases_that_require_create_user[tt.name] {
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(tt.mockCreateUserError)
			}
			service := NewUserService(mockRepo, *logger)

			// invoke the method
			response, err := service.CreateUser(tt.requestbody)

			// Assert
			if err != nil && tt.expectedError != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
				}
			}
			if response.Code != tt.expectedResponse.Code {
				t.Errorf("Expected code %d, got %d", tt.expectedResponse.Code, response.Code)
			}
			if response.Message != tt.expectedResponse.Message {
				t.Errorf("Expected message %s, got %s", tt.expectedResponse.Message, response.Message)
			}
		})
	}

}

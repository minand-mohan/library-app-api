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

func TestUpdateUser(t *testing.T) {

	test_email := "test@example.com"

	internalServerError := errors.New("Internal Server Error")
	userNotFoundError := errors.New("User not found")
	duplicateKeyError := errors.New("pq: duplicate key value violates unique constraint \"users_email_key\"")

	test_user := generateRandomUser01()
	test_id := *test_user.ID

	test_cases_that_require_find_user := map[string]bool{
		"Update User sucessfully":       true,
		"Update User with error":        true,
		"Cannot find user":              true,
		"Update User non-unique values": true,
	}

	test_cases_that_require_update_user := map[string]bool{
		"Update User sucessfully":       true,
		"Update User with error":        true,
		"Update User non-unique values": true,
	}

	tc := []struct {
		name                 string
		requestbody          *dto.UserRequestBody
		expectedResponse     *response.HTTPResponse
		expectedError        error
		mockFindUserReturn   *models.User
		mockFindUserError    error
		mockUpdateUserReturn *models.User
		mockUpdateUserError  error
	}{
		{
			name: "Update User sucessfully",
			requestbody: &dto.UserRequestBody{
				Email: test_email,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User updated successfully",
				Content: map[string]interface{}{},
			},
			expectedError:      nil,
			mockFindUserReturn: &test_user,
			mockFindUserError:  nil,
			mockUpdateUserReturn: &models.User{
				ID:       test_user.ID,
				Email:    &test_email,
				Username: test_user.Username,
				Phone:    test_user.Phone,
			},
			mockUpdateUserError: nil,
		},
		{
			name: "Update User with error",
			requestbody: &dto.UserRequestBody{
				Email: test_email,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal Server Error",
				Content: map[string]interface{}{},
			},
			expectedError:        internalServerError,
			mockFindUserReturn:   &test_user,
			mockFindUserError:    nil,
			mockUpdateUserReturn: nil,
			mockUpdateUserError:  internalServerError,
		},
		{
			name: "Cannot find user",
			requestbody: &dto.UserRequestBody{
				Email: test_email,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    404,
				Message: "User not found.",
				Content: map[string]interface{}{},
			},
			expectedError:        userNotFoundError,
			mockFindUserReturn:   nil,
			mockFindUserError:    userNotFoundError,
			mockUpdateUserReturn: nil,
			mockUpdateUserError:  nil,
		},
		{
			name: "Update User non-unique values",
			requestbody: &dto.UserRequestBody{
				Email: test_email,
			},
			expectedResponse: &response.HTTPResponse{
				Code:    400,
				Message: "Bad request, non-unique values",
				Content: map[string]interface{}{},
			},
			expectedError:        duplicateKeyError,
			mockFindUserReturn:   &test_user,
			mockFindUserError:    nil,
			mockUpdateUserReturn: nil,
			mockUpdateUserError:  duplicateKeyError,
		},
	}

	for _, tt := range tc {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		t.Run(tt.name, func(t *testing.T) {
			logger := utils.NewLogger()

			test_input := &models.User{
				Username: &tt.requestbody.Username,
				Email:    &tt.requestbody.Email,
				Phone:    &tt.requestbody.Phone,
			}

			// Arrange
			mockUserRepo := repomocks.NewMockUserRepository(mockCtrl)
			service := NewUserService(mockUserRepo, *utils.NewLogger())
			if test_cases_that_require_find_user[tt.name] {
				mockUserRepo.EXPECT().FindByUserId(test_id).Return(tt.mockFindUserReturn, tt.mockFindUserError)
			}
			if test_cases_that_require_update_user[tt.name] {
				mockUserRepo.EXPECT().UpdateByUserId(test_id, test_input).Return(tt.mockUpdateUserReturn, tt.mockUpdateUserError)
			}
			// Act
			response, err := service.UpdateByUserId(test_id, tt.requestbody)
			logger.Info("Response: " + response.Message)
			// Assert
			// if !reflect.DeepEqual(response, tt.expectedResponse) {
			// 	t.Errorf("expected response %v, got %v", tt.expectedResponse, response)
			// }
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
			}
			if response.Code != tt.expectedResponse.Code {
				t.Errorf("expected code %d, got %d", tt.expectedResponse.Code, response.Code)
			}
			if response.Message != tt.expectedResponse.Message {
				t.Errorf("expected message %s, got %s", tt.expectedResponse.Message, response.Message)
			}

		})
	}
}

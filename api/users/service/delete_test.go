package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	repomocks "github.com/minand-mohan/library-app-api/api/users/repository/mocks"
	"github.com/minand-mohan/library-app-api/utils"
)

func TestDeleteUser(t *testing.T) {

	test_cases_that_require_find_user := map[string]bool{
		"Delete User by id successfully": true,
		"Cannot find user":               true,
		"Delete User by id with error":   true,
	}
	test_cases_that_require_delete_user := map[string]bool{
		"Delete User by id successfully": true,
		"Delete User by id with error":   true,
	}
	tc := []struct {
		name                string
		id                  string
		expectedResponse    *response.HTTPResponse
		expectedError       error
		mockDeleteUserError error
		mockFindUserError   error
	}{
		{
			name: "Delete User by id successfully",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			expectedResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User deleted successfully",
				Content: map[string]interface{}{},
			},
			expectedError:       nil,
			mockDeleteUserError: nil,
			mockFindUserError:   nil,
		},
		{
			name: "Cannot find user",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			expectedResponse: &response.HTTPResponse{
				Code:    404,
				Message: "User not found.",
				Content: map[string]interface{}{},
			},
			expectedError:       nil,
			mockDeleteUserError: nil,
			mockFindUserError:   errors.New("User not found"),
		},
		{
			name: "Delete User by id with error",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			expectedResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal Server Error",
				Content: map[string]interface{}{},
			},
			expectedError:       errors.New("Internal Server Error"),
			mockDeleteUserError: errors.New("Internal Server Error"),
			mockFindUserError:   nil,
		},
	}

	for _, tc := range tc {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		t.Run(tc.name, func(t *testing.T) {
			id, _ := uuid.Parse(tc.id)
			mockRepo := repomocks.NewMockUserRepository(mockCtrl)

			if test_cases_that_require_find_user[tc.name] {
				mockRepo.EXPECT().FindByUserId(id).Return(nil, tc.mockFindUserError)
			}
			if test_cases_that_require_delete_user[tc.name] {
				mockRepo.EXPECT().DeleteByUserId(id).Return(tc.mockDeleteUserError)
			}
			service := UserServiceImpl{
				repo:   mockRepo,
				logger: utils.NewLogger(),
			}

			response, err := service.DeleteByUserId(id)
			if err != nil && tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
				}
			}

			if !reflect.DeepEqual(response, tc.expectedResponse) {
				t.Errorf("Expected response: %v, got: %v", tc.expectedResponse, response)
			}
		})
	}
}

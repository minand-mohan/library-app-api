package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/service/mocks"
	"github.com/minand-mohan/library-app-api/api/users/validator"
	"github.com/minand-mohan/library-app-api/utils"
)

func TestDeleteByUserId(t *testing.T) {
	testCases := []struct {
		name                      string
		id                        string
		mockServiceExpectResponse *response.HTTPResponse
		mockServiceExpectError    error
		expectedStatus            int
		expectedMessage           string
	}{
		{
			name: "Delete user by id with valid id",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User deleted successfully",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError: nil,
			expectedStatus:         200,
			expectedMessage:        "User deleted successfully",
		},
		{
			name:                      "Delete user by id with invalid id",
			id:                        "invalid-id",
			mockServiceExpectResponse: nil,
			mockServiceExpectError:    nil,
			expectedStatus:            400,
			expectedMessage:           "Bad request, invalid id",
		},
		{
			name: "Delete user by id with error",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal server error",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError: errors.New("Internal server error"),
			expectedStatus:         500,
			expectedMessage:        "Internal server error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			app := setupApp()
			app.Delete("/users/:id", func(c *fiber.Ctx) error {
				logger := utils.NewLogger()
				service := mocks.NewMockUserService(mockCtrl)
				validator := validator.NewUserValidator(*logger)
				if tc.mockServiceExpectResponse != nil {
					service.EXPECT().DeleteByUserId(gomock.Any()).Return(tc.mockServiceExpectResponse, tc.mockServiceExpectError)
				}
				handler := NewUserHandler(service, validator)
				return handler.DeleteByUserId(c)
			})
			request := httptest.NewRequest("DELETE", "/users/"+tc.id, nil)

			response, err := app.Test(request)
			if err != nil {
				t.Errorf("Error while making request %v", err)
			}

			if response.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, response.StatusCode)
			}

			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("Error while reading response body: %v", err)
			}

			var responseBody map[string]interface{}

			err = json.Unmarshal(bodyBytes, &responseBody)
			if err != nil {
				t.Errorf("Error while parsing response body: %v", err)
			}

			message := responseBody["message"].(string)
			if message != tc.expectedMessage {
				t.Errorf("Expected message %s, got %s", tc.expectedMessage, message)
			}
		})
	}
}

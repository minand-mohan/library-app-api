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
	servicemocks "github.com/minand-mohan/library-app-api/api/users/service/mocks"
	"github.com/minand-mohan/library-app-api/api/users/validator"
	validatormocks "github.com/minand-mohan/library-app-api/api/users/validator/mocks"
	"github.com/minand-mohan/library-app-api/utils"
)

func TestFindAllUsers(t *testing.T) {
	testCases := []struct {
		name                      string
		queryParams               string
		mockServiceExpectResponse *response.HTTPResponse
		mockServiceExpectError    error
		mockValidatorExpectError  error
		expectedStatus            int
		expectedMessage           string
	}{
		{
			name:        "Find all users with valid query params",
			queryParams: "name=test&phone=1234567890",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    200,
				Message: "Users found successfully",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   nil,
			mockValidatorExpectError: nil,
			expectedStatus:           200,
			expectedMessage:          "Users found successfully",
		},
		{
			name:        "Find users that do not exist",
			queryParams: "name=test&phone=123242345",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    404,
				Message: "No Users found",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   errors.New("No Users found"),
			mockValidatorExpectError: nil,
			expectedStatus:           404,
			expectedMessage:          "No Users found",
		},
		{
			name:                      "Find all users with invalid query params",
			queryParams:               "name=test&phone=invalid-phone",
			mockServiceExpectResponse: nil,
			mockServiceExpectError:    nil,
			mockValidatorExpectError:  errors.New("Invalid phone number"),
			expectedStatus:            400,
			expectedMessage:           "Bad request, invalid query params",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new fiber context for testing
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			app := setupApp()
			app.Get("/users", func(c *fiber.Ctx) error {
				validator := validatormocks.NewMockUserValidator(mockCtrl)
				service := servicemocks.NewMockUserService(mockCtrl)
				if tc.mockServiceExpectResponse != nil {
					service.EXPECT().FindAllUsers(gomock.Any()).Return(tc.mockServiceExpectResponse, tc.mockServiceExpectError)
				}
				validator.EXPECT().ValidateUserQueryParams(gomock.Any()).Return(tc.mockValidatorExpectError)
				handler := NewUserHandler(service, validator)
				return handler.FindAllUsers(c)
			})
			request := httptest.NewRequest("GET", "/users?"+tc.queryParams, nil)

			response, err := app.Test(request)
			if err != nil {
				t.Errorf("Error while making request %v", err)
			}

			if response.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, response.StatusCode)
			}

			// Read the entire response body
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("Error while reading response body: %v", err)
			}

			// Define a map or struct to hold the response body
			var responseBody map[string]interface{}

			// Parse the JSON response body
			err = json.Unmarshal(bodyBytes, &responseBody)
			if err != nil {
				t.Errorf("Error while parsing response body: %v", err)
			}

			// Now you can access the fields of the response body
			message := responseBody["message"].(string)
			if message != tc.expectedMessage {
				t.Errorf("Expected message %s, got %s", tc.expectedMessage, message)
			}
		})
	}
}
func TestFindByUserId(t *testing.T) {
	testCases := []struct {
		name                      string
		id                        string
		mockServiceExpectResponse *response.HTTPResponse
		mockServiceExpectError    error
		expectedStatus            int
		expectedMessage           string
	}{
		{
			name: "Find user by id with valid id",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User found successfully",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError: nil,
			expectedStatus:         200,
			expectedMessage:        "User found successfully",
		},
		{
			name:                      "Find user by id with invalid id",
			id:                        "invalid-id",
			mockServiceExpectResponse: nil,
			mockServiceExpectError:    nil,
			expectedStatus:            400,
			expectedMessage:           "Bad request, invalid id",
		},
		{
			name: "Find user by id that does not exist",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    404,
				Message: "User not found",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError: errors.New("User not found"),
			expectedStatus:         404,
			expectedMessage:        "User not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			app := setupApp()
			app.Get("/users/:id", func(c *fiber.Ctx) error {
				logger := utils.NewLogger()
				service := servicemocks.NewMockUserService(mockCtrl)
				validator := validator.NewUserValidator(*logger)
				if tc.mockServiceExpectResponse != nil {
					service.EXPECT().FindByUserId(gomock.Any()).Return(tc.mockServiceExpectResponse, tc.mockServiceExpectError)
				}
				handler := NewUserHandler(service, validator)
				return handler.FindByUserId(c)
			})
			request := httptest.NewRequest("GET", "/users/"+tc.id, nil)

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

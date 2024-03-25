package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
	"github.com/minand-mohan/library-app-api/api/response"
	servicemocks "github.com/minand-mohan/library-app-api/api/users/service/mocks"
	validatormocks "github.com/minand-mohan/library-app-api/api/users/validator/mocks"
)

func TestCreateUser(t *testing.T) {

	testCases := []struct {
		name                      string
		requestBody               map[string]interface{}
		mockServiceExpectResponse *response.HTTPResponse
		mockServiceExpectError    error
		mockValidatorExpectError  error
		expectedStatus            int
		expectedMessage           string
	}{
		{
			name: "Create user with valid request body",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User created successfully",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   nil,
			mockValidatorExpectError: nil,
			expectedStatus:           200,
			expectedMessage:          "User created successfully",
		},
		{
			name: "Create user with invalid request body - user already exists",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    400,
				Message: "Bad request, user already exists",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   errors.New("User already exists"),
			mockValidatorExpectError: nil,
			expectedStatus:           400,
			expectedMessage:          "Bad request, user already exists",
		},
		{
			name: "Create user with empty phone",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
			},
			mockServiceExpectResponse: nil,
			mockServiceExpectError:    nil,
			mockValidatorExpectError:  errors.New("Phone number is empty"),
			expectedStatus:            400,
			expectedMessage:           "Bad request, invalid request body",
		},
		{
			name: "Create user with service error",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal Server Error",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   errors.New("Internal Server Error"),
			mockValidatorExpectError: nil,
			expectedStatus:           500,
			expectedMessage:          "Internal Server Error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new fiber context for testing
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			app := setupApp()
			app.Post("/users", func(c *fiber.Ctx) error {
				// logger := utils.NewLogger()
				validator := validatormocks.NewMockUserValidator(mockCtrl)
				service := servicemocks.NewMockUserService(mockCtrl)
				if tc.mockServiceExpectResponse != nil {
					service.EXPECT().CreateUser(gomock.Any()).Return(tc.mockServiceExpectResponse, tc.mockServiceExpectError)
				}
				validator.EXPECT().ValidateUser(gomock.Any()).Return(tc.mockValidatorExpectError)
				handler := NewUserHandler(service, validator)
				t.Logf("Handler: %v", handler)
				return handler.CreateUser(c)
			})
			requestBody, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Errorf("Error while marshalling request body: %v", err)
			}
			t.Logf("Request body: %v", tc.requestBody)
			request := httptest.NewRequest("POST", "/users/", strings.NewReader(string(requestBody)))

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

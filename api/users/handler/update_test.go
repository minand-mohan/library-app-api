package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/service/service_tests"
	"github.com/minand-mohan/library-app-api/api/users/validator/validator_tests"
)

func TestUpdateByUserId(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Create a new fiber app for testing
	app := fiber.New()

	// Create a mock service and validator
	mockService := service_tests.NewMockUserService(mockCtrl)
	mockValidator := validator_tests.NewMockUserValidator(mockCtrl)

	// Create a user handler instance
	handler := NewUserHandler(mockService, mockValidator)

	// Register the route with the handler method
	app.Put("/users/:id", handler.UpdateByUserId)

	// Define test cases
	testCases := []struct {
		name                      string
		id                        string
		requestBody               map[string]interface{}
		mockServiceExpectResponse *response.HTTPResponse
		mockServiceExpectError    error
		mockValidatorExpectError  error
		expectedStatus            int
		expectedMessage           string
	}{
		{
			name: "Update user with valid request body",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    200,
				Message: "User updated successfully",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   nil,
			mockValidatorExpectError: nil,
			expectedStatus:           200,
			expectedMessage:          "User updated successfully",
		},
		{
			name: "Update user with invalid id",
			id:   "invalid-id",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: nil,
			mockServiceExpectError:    nil,
			mockValidatorExpectError:  nil,
			expectedStatus:            400,
			expectedMessage:           "Bad request, invalid id",
		},
		{
			name: "Update user with invalid request body",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
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
			name: "Update user with service error",
			id:   "d3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b",
			requestBody: map[string]interface{}{
				"username": "test",
				"email":    "test@example.com",
				"phone":    "1234567890",
			},
			mockServiceExpectResponse: &response.HTTPResponse{
				Code:    500,
				Message: "Internal server error",
				Content: map[string]interface{}{},
			},
			mockServiceExpectError:   errors.New("Service error"),
			mockValidatorExpectError: nil,
			expectedStatus:           500,
			expectedMessage:          "Internal server error",
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset the mock service expectations
			if tc.mockServiceExpectResponse != nil {
				mockService.EXPECT().UpdateByUserId(gomock.Any(), gomock.Any()).Return(tc.mockServiceExpectResponse, tc.mockServiceExpectError)
			}
			if tc.name != "Update user with invalid id" {
				mockValidator.EXPECT().ValidateUser(gomock.Any()).Return(tc.mockValidatorExpectError)
			}
			// Create a request body from the test case data
			requestBody, err := json.Marshal(tc.requestBody)
			if err != nil {
				t.Errorf("Error while marshalling request body: %v", err)
			}

			// Create a request with the test case data
			request := httptest.NewRequest("PUT", fmt.Sprintf("/users/%s", tc.id), strings.NewReader(string(requestBody)))

			// Perform the request
			response, err := app.Test(request)
			if err != nil {
				t.Errorf("Error while making request: %v", err)
			}

			// Check the response status code
			if response.StatusCode != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, response.StatusCode)
			}

			// Read the response body
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Errorf("Error while reading response body: %v", err)
			}

			// Parse the response body
			var responseBody map[string]interface{}
			err = json.Unmarshal(bodyBytes, &responseBody)
			if err != nil {
				t.Errorf("Error while parsing response body: %v", err)
			}

			// Check the response message
			message := responseBody["message"].(string)
			if message != tc.expectedMessage {
				t.Errorf("Expected message %s, got %s", tc.expectedMessage, message)
			}
		})
	}
}

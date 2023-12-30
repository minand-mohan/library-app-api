package response

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/minand-mohan/library-app-api/utils"
)

type HTTPResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Content interface{} `json:"content"`
}

type HTTPResponseContent struct {
	Count    int         `json:"count"`
	Previous *string     `json:"prev"`
	Next     *string     `json:"next"`
	Results  interface{} `json:"results"`
}

// Error details will be returned by a service function to the handler
type ErrorDetails struct {
	Code    int
	Message string
	Error   error
}

// StatusCode refer to Http Code
func WriteHTTPResponse(c *fiber.Ctx, statusCode int, responseBody *HTTPResponse) error {
	if statusCode < 100 || statusCode > 600 {
		return errors.New(fmt.Sprintf("Invalid status code for HTTP response: %v", statusCode))
	}
	c.Status(statusCode)
	err := c.JSON(responseBody)
	return err
}

// Code refer to Application Code
func GetErrorHTTPResponseBody(code int, message string) *HTTPResponse {
	return &HTTPResponse{
		Code:    code,
		Message: message,
		Content: map[string]interface{}{},
	}
}

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	log := utils.NewLogger()
	log.Error(fmt.Sprintf("Error thrown from DefaultErrorHandler %e", err))
	errorBody := GetErrorHTTPResponseBody(500, "Internal Server Error")
	return WriteHTTPResponse(c, 500, errorBody)
}

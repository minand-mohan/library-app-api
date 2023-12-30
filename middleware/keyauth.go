package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/minand-mohan/library-app-api/api/response"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {

	apiAuthToken, ok := os.LookupEnv("API_AUTH_TOKEN")
	if !ok {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}
	if key == apiAuthToken {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}
func errorHandler(c *fiber.Ctx, err error) error {
	errorBody := response.GetErrorHTTPResponseBody(401, "Missing or invalid Auth Token")
	return response.WriteHTTPResponse(c, 401, errorBody)
}

var KeyAuth = keyauth.New(keyauth.Config{
	Validator:    validateAPIKey,
	ErrorHandler: errorHandler,
})

package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/minand-mohan/library-app-api/api/response"
	userHandler "github.com/minand-mohan/library-app-api/api/users/handler"
	userRepository "github.com/minand-mohan/library-app-api/api/users/repository"
	userService "github.com/minand-mohan/library-app-api/api/users/service"
	userValidator "github.com/minand-mohan/library-app-api/api/users/validator"
	"github.com/minand-mohan/library-app-api/middleware"
	"github.com/minand-mohan/library-app-api/utils"
)

func getDefaultUserHandler(server *APIServer) *userHandler.UserHandler {
	logger := utils.NewLogger()
	repository := userRepository.NewUserRepository(server.dataSource.DB)
	validator := userValidator.NewUserValidator(*logger)
	service := userService.NewUserService(repository, *logger)
	return userHandler.NewUserHandler(service, validator)
}

func setUpDefaultRoutes(server *APIServer) {
	app := server.app
	app.All("/*", func(c *fiber.Ctx) error {
		body := response.GetErrorHTTPResponseBody(http.StatusMethodNotAllowed, "Method Not Allowed")
		response.WriteHTTPResponse(c, http.StatusMethodNotAllowed, body)
		return nil
	})
}

func SetupRoutes(server *APIServer) {

	app := server.app
	libraryv1 := app.Group("/library-app/api/v1")

	// User routes
	libraryv1.Post("/users", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := getDefaultUserHandler(server)
		return handler.CreateUser(c)
	})

	libraryv1.Get("/users", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := getDefaultUserHandler(server)
		return handler.FindAllUsers(c)
	})

	libraryv1.Get("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := getDefaultUserHandler(server)
		return handler.FindByUserId(c)
	})

	libraryv1.Put("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := getDefaultUserHandler(server)
		return handler.UpdateByUserId(c)
	})

	libraryv1.Delete("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := getDefaultUserHandler(server)
		return handler.DeleteByUserId(c)
	})

	setUpDefaultRoutes(server)

}

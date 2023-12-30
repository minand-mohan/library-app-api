package api

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/minand-mohan/library-app-api/api/users/handler"
	userRepository "github.com/minand-mohan/library-app-api/api/users/repository"
	userService "github.com/minand-mohan/library-app-api/api/users/service"
	userValidator "github.com/minand-mohan/library-app-api/api/users/validator"
	"github.com/minand-mohan/library-app-api/middleware"
	"github.com/minand-mohan/library-app-api/utils"
)

func GetDefaultUserHandler(server *APIServer) *userHandler.UserHandler {
	logger := utils.NewLogger()
	repository := userRepository.NewUserRepository(server.dataSource.DB)
	validator := userValidator.NewUserValidator(*logger)
	service := userService.NewUserService(repository, *logger)
	return userHandler.NewUserHandler(service, validator)
}

func SetupRoutes(server *APIServer) {

	app := server.app
	libraryv1 := app.Group("/library-app/api/v1")

	// User routes
	libraryv1.Post("/users", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := GetDefaultUserHandler(server)
		return handler.CreateUser(c)
	})

	libraryv1.Get("/users", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := GetDefaultUserHandler(server)
		return handler.FindAllUsers(c)
	})

	libraryv1.Get("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := GetDefaultUserHandler(server)
		return handler.FindByUserId(c)
	})

	libraryv1.Put("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := GetDefaultUserHandler(server)
		return handler.UpdateByUserId(c)
	})

	libraryv1.Delete("/users/:id", middleware.KeyAuth, func(c *fiber.Ctx) error {
		handler := GetDefaultUserHandler(server)
		return handler.DeleteByUserId(c)
	})

}

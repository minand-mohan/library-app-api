package api

func SetupRoutes(server *APIServer) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	app := server.app
	app.Group("/library-app/api/v1")

}

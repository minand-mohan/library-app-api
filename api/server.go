package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/system"
	"github.com/minand-mohan/library-app-api/utils"
)

type APIServer struct {
	// appConfig  *system.Config
	logger     *utils.AppLogger
	dataSource *system.DataSource
	app        *fiber.App
}

func NewServer() *APIServer {
	dataSource := system.NewDataSource()
	app := fiber.New(fiber.Config{
		CaseSensitive:         true,
		ServerHeader:          "minand-mohan/library-app-api",
		Concurrency:           1024,
		DisableStartupMessage: false,
		ErrorHandler:          response.DefaultErrorHandler,
	})
	appLogger := utils.NewLogger()
	return &APIServer{
		logger:     appLogger,
		dataSource: dataSource,
		app:        app,
	}
}

func (server *APIServer) StartServer() {
	log := server.logger
	SetupRoutes(server)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Start the server
		if err := server.app.Listen(":8080"); err != nil {
			log.Fatal(fmt.Sprintf("Error starting api server %e", err))
		}
	}()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Info("Terminating api server: context cancelled")
		server.app.Shutdown()
	case <-sigterm:
		log.Info("Terminating api server: via signal")
		server.app.Shutdown()
	}
	cancel()
}

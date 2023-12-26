package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/minand-mohan/library-app-api/api"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		apiServer := api.NewServer()
		print("Starting server..")
		apiServer.StartServer()
	}()
	// Wait for the server to shutdown
	cancel()
	wg.Wait()
	fmt.Println("Server shutdown gracefully")

}

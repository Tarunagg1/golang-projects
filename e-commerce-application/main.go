package main

import (
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api"
)

func main() {
	// Load application settings
	appConfig, err := config.SetupEnv()

	if err != nil {
		panic("Failed to load application settings: " + err.Error())
	}

	// Print the loaded configuration for debugging
	fmt.Printf("Loaded configuration: %+v\n", appConfig)

	// Start the server
	api.StartServer(appConfig)
}

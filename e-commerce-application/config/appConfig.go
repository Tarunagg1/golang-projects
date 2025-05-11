package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppCOnfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppCOnfig, err error) {
	// Load application settings from a configuration file or environment variables
	// This is a placeholder function. You can implement the actual loading logic here.
	// For example, you might use a library like viper to load settings from a .env file.
	fmt.Println("Loading application settings...")

	// if os.Getenv("APP_ENV") == "development" {
	godotenv.Load()
	// }

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppCOnfig{}, errors.New("env variables no found")
	}

	return AppCOnfig{
		ServerPort: httpPort,
	}, nil
}

package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppCOnfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	rh := &rest.RestHandler{
		App:    app,
		Config: config,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "I am breathing",
		"status":  "OK",
	})
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}

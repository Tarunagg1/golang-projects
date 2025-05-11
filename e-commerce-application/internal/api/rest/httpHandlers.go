package rest

import (
	"go-ecommerce-app/config"

	"github.com/gofiber/fiber/v2"
)

type RestHandler struct {
	App    *fiber.App
	Config config.AppCOnfig
	// DB     *gorm.DB
}

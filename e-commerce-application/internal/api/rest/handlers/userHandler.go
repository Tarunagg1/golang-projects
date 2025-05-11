package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usvc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	pubRoutes := app.Group("/api/v1/user")

	svc := service.UserService{}

	handler := UserHandler{
		usvc: svc,
	}

	// Public routes

	pubRoutes.Post("/register", handler.RegisterUser)
	pubRoutes.Post("/login", handler.LoginUser)

	// Private routes
	pvtRoutes := app.Group("/api/v1/user")

	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)

	pvtRoutes.Post("/profile", handler.CreateProfile)
	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Patch("/profile", handler.UpdateProfile)

	pvtRoutes.Post("/cart", handler.AddToCart)
	pvtRoutes.Get("/cart", handler.GetCart)

	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) RegisterUser(ctx *fiber.Ctx) error {

	user := dto.UserSignup{}

	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"status":  http.StatusBadRequest,
		})
	}

	token, error := h.usvc.Signup(user)

	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating user",
			"status":  http.StatusInternalServerError,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
		"token":   token,
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User GetVerificationCode in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User Verify in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User CreateProfile in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User GetProfile in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User UpdateProfile in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User AddToCart in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User GetCart in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User GetOrders in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User order in successfully",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User seller in successfully",
		"status":  http.StatusOK,
	})
}

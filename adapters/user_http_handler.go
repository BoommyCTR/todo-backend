package adapters

import (
	"time"

	"github.com/boomctr/todo-backend-go/auth"
	"github.com/boomctr/todo-backend-go/entities"
	"github.com/boomctr/todo-backend-go/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	service usecases.UserUseCase
}

func NewHttpUserHandler(service usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{
		service: service,
	}
}

func (h *HttpUserHandler) CreateUserHandler(c *fiber.Ctx) error {
	user := new(entities.Users)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created User Successfully"})
}

func (h *HttpUserHandler) LoginHandler(c *fiber.Ctx) error {
	user := new(entities.Users)

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := h.service.VerifyUser(user)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * 72),
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{"message": "Login Successfully"})
}

func (h *HttpUserHandler) WhoAmIHandler(c *fiber.Ctx) error {
	claims, err := auth.GetClaims(c)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	id := uint64((*claims)["user_id"].(float64))
	name, err := h.service.WhoAmIUser(id)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{"name": name})
}

func (h *HttpUserHandler) LogoutHandler(c *fiber.Ctx) error {
	c.ClearCookie("jwt")

	return c.JSON(fiber.Map{"message": "Logout Successfully"})
}

package adapters

import (
	"strconv"

	"github.com/boomctr/todo-backend-go/auth"
	"github.com/boomctr/todo-backend-go/entities"
	"github.com/boomctr/todo-backend-go/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpTodoHandler struct {
	service usecases.TodoUseCase
}

func NewHttpTodoHandler(service usecases.TodoUseCase) *HttpTodoHandler {
	return &HttpTodoHandler{
		service: service,
	}
}

func (h *HttpTodoHandler) CreateTodoHandler(c *fiber.Ctx) error {
	claims, err := auth.GetClaims(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	id := uint64((*claims)["user_id"].(float64))
	todo := new(entities.Todos)

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	todo.UsersID = uint(id)

	if err := h.service.CreateTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (h *HttpTodoHandler) GetTodosHandler(c *fiber.Ctx) error {
	claims, err := auth.GetClaims(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	id := uint64((*claims)["user_id"].(float64))
	Todos, err := h.service.GetAllTodos(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(Todos)
}

func (h *HttpTodoHandler) UpdateTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	TodoUpdate := new(entities.Todos)

	TodoUpdate, err = h.service.GetTodo(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&TodoUpdate); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = h.service.UpdateTodo(TodoUpdate)

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(TodoUpdate)
}

func (h *HttpTodoHandler) DeleteTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = h.service.DeleteTodo(id)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "Delete Successfully"})

}

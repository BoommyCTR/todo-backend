package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boomctr/todo-backend-go/adapters"
	"github.com/boomctr/todo-backend-go/auth"
	_ "github.com/boomctr/todo-backend-go/docs"
	"github.com/boomctr/todo-backend-go/entities"
	"github.com/boomctr/todo-backend-go/usecases"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func getHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	fmt.Println("Connecting to database...")

	// connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, databaseName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entities.Todos{}, &entities.Users{})

	fmt.Println("Connected to database")

	todoRepository := adapters.NewGormTodoRepository(db)
	todoService := usecases.NewTodoService(todoRepository)
	todoHandler := adapters.NewHttpTodoHandler(todoService)

	userRepository := adapters.NewGormUserRepository(db)
	userService := usecases.NewUserService(userRepository)
	userHandler := adapters.NewHttpUserHandler(userService)

	// fiber CRUD
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://127.0.0.1:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", getHello)
	app.Post("/register", userHandler.CreateUserHandler)
	app.Post("/login", userHandler.LoginHandler)

	app.Use(auth.CheckMiddleware)

	app.Get("/whoami", userHandler.WhoAmIHandler)
	app.Get("/todos", todoHandler.GetTodosHandler)
	app.Post("/todos", todoHandler.CreateTodoHandler)
	app.Put("/todos/:id", todoHandler.UpdateTodoHandler)
	app.Delete("/todos/:id", todoHandler.DeleteTodoHandler)
	app.Get("/logout", userHandler.LogoutHandler)

	app.Listen(":8080")
}

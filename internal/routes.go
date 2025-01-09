package internal

import (
	"ProjectsGo/internal/database"
	toDo "ProjectsGo/internal/entities/toDo"
	"ProjectsGo/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"os"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message":     "ToDo List",
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Use(middleware.ErrorHandler)
	app.Use(middleware.ResponseHandler)

	DatabaseRoutes(app)
	ToDoRoutes(app)

}

func DatabaseRoutes(app *fiber.App) {
	app.Post("/database/migrate", database.Migrate)
	//TODO fazer seed
}

func ToDoRoutes(app *fiber.App) {
	app.Post("/toDo/create", toDo.Create)
}

package internal

import (
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

	ToDoRoutes(app)

}

func ToDoRoutes(app *fiber.App) {
}

package database

import (
	"github.com/gofiber/fiber/v2"
)

func Migrate(c *fiber.Ctx) error {

	err := GetDB(c).AutoMigrate()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error migrating schemas": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "database successfully migrated"})
}

func Seed(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ok"}) //TODO fazer sistema de seed
}

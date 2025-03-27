package main

import (
	"ProjectsGo/internal"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		err = godotenv.Load("/env_secret/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	time.Local, _ = time.LoadLocation(os.Getenv("TIMEZONE_PROJECT")) //TODO ADICIONAR ENV
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 25 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowHeaders:  "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		ExposeHeaders: "Origin",
	}))

	app.Use(requestid.New())

	internal.RegisterRoutes(app)

	port := os.Getenv("PORT_SERVER")
	if len(port) == 0 {
		port = "7070"
	}

	/* Start Server */
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

}

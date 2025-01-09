package database

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var instance *gorm.DB

type UserKey string
type TraceIdKey string

func NewConnection() error {
	var dsn string

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	//sslmode := os.Getenv("DB_SSLMODE")
	//sslrootcert := os.Getenv("DB_SSLROOTCERT")
	//sslkey := os.Getenv("DB_SSLCLIENTKEY")
	//sslcert := os.Getenv("DB_SSLCLIENTCERT")

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, name, port)

	//if sslMode != "" {
	//	dsn += fmt.Sprintf(" sslmode=%s", sslMode)
	//}
	//
	//if rootCert != "" {
	//	dsn += fmt.Sprintf(" sslrootcert=%s", rootCert)
	//}
	//
	//if clientKey != "" {
	//	dsn += fmt.Sprintf(" sslkey=%s", clientKey)
	//}
	//
	//if clientCert != "" {
	//	dsn += fmt.Sprintf(" sslcert=%s", clientCert)
	//}
	//

	log.Printf("Connecting to database: %s:%s\n", host, port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent),
		PrepareStmt: true,
	})

	if err != nil {
		return err
	}

	if DB != nil {
		instance = DB
		log.Print("Database connected")
	}

	return nil
}

func GetDB(c *fiber.Ctx) *gorm.DB {
	if instance == nil {
		log.Println("Database connection not found!")
		maxAttempts := 3
		for attempt := 0; attempt < maxAttempts; attempt++ {
			log.Println("retrying connect... attempt: ", attempt)
			err := NewConnection()
			if instance != nil || err != nil {
				return instance
			}
		}
	}

	user := c.Locals("user")
	traceId := c.Locals("requestid")

	var userK UserKey
	var traceIdK TraceIdKey

	ctx := context.WithValue(c.Context(), userK, user)
	ctx = context.WithValue(ctx, traceIdK, traceId)

	return instance.WithContext(ctx)
}

package api

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmayola/mayola_bucket/handlers"
	"github.com/jmayola/mayola_bucket/middleware"
)

func StartApi() {
	// variables for app
	url := os.Getenv("APP_URL")
	if url == "" {
		url = ":3000"
	}

	name := os.Getenv("APP_NAME")
	if name == "" {
		name = "bucket"
	}
	//tls
	// pem := os.Getenv("PEM_FILE")
	// cert := os.Getenv("CERT_FILE")

	// app configuration
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "bucket",
		AppName:       name,
	})

	//middlewares
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(middleware.MiddleLogger())

	// handler
	app.Post("/upload", handlers.UploadFiles)

	// upload files to be accessed
	app.Static("/uploads", "/files", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 24 * time.Hour,
		MaxAge:        3600,
	})

	// app.Server().MaxConnsPerIP = 1

	log.Fatal(app.Listen(url))
}

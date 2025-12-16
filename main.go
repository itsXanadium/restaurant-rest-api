package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/itsXanadium/restaurant-rest-api/config"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	config.LoadEnv()
	config.DBConnect()
	log.Println("DB_USER", os.Getenv("DB_USER"))
	log.Println("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"{+}": "API IS RUNNING"})
	})
	app.Listen(":" + config.Appconfig.Appport)
}

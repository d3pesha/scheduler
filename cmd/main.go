package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"scheduler/internal/api"
	"scheduler/internal/service"
)

func main() {
	jobService := service.NewService()

	app := fiber.New()

	api.Register(app, jobService)
	if err := app.Listen(":8080"); err != nil {
		log.Fatalln("failed to listen:", err)
	}
}

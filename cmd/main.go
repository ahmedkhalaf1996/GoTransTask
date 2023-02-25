package main

import (
	_ "github.com/freddymac124/upwork-go/swegger"

	"github.com/gofiber/swagger"

	"github.com/freddymac124/upwork-go/handlers"
	"github.com/freddymac124/upwork-go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	models.ConnectDatabase() // New

	app := fiber.New(fiber.Config{
		AppName:      "Fiber with Planetscale",
		ServerHeader: "Fiber",
	})
	app.Use(logger.New())

	app.Post("/credit", handlers.AddNewCredit)
	app.Get("/Getrecoreds", handlers.GetLastRecored)
	app.Post("/debit", handlers.AddNewDebit)

	// swegger
	app.Get("/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}

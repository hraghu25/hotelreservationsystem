package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hraghu25/hotelreservationsystem/handler"
)

func Execute(listenAddress *string) {
	app := fiber.New()
	app.Use(logger.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/health", handler.HandlerHealth)

	api := app.Group("/api/v1")
	api.Get("/user", handler.HandleGetUser)
	api.Get("/user/:id", handler.HandleGetUserByID)

	app.Listen(*listenAddress)
}

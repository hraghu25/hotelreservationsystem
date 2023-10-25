package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hraghu25/hotelreservationsystem/database"
	"github.com/hraghu25/hotelreservationsystem/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://mongoadmin:secret@localhost:27017"

func Execute(listenAddress *string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/health", handler.HandlerHealth)

	userHandler := handler.NewUserHandler(database.NewMongoUserStore(client))
	api := app.Group("/api/v1")
	api.Get("/user", userHandler.HandleGetUser)
	api.Get("/user/:id", userHandler.HandleGetUserByID)

	app.Listen(*listenAddress)
}

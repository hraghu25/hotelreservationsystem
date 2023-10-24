package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hraghu25/hotelreservationsystem/handler"
	"github.com/hraghu25/hotelreservationsystem/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://mongoadmin:secret@localhost:27017"
const dbname = "hotel-reservation"
const userCollection = "users"

func Execute(listenAddress *string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userCollection)

	user := types.User{
		FirstName: "Raghavendra",
		LastName:  "Hiremath",
	}

	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	var usr types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&usr); err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr)

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

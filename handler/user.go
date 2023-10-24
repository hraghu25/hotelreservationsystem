package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hraghu25/hotelreservationsystem/types"
)

func HandleGetUser(c *fiber.Ctx) error {
	usr := types.User{
		FirstName: "Raghavendra",
		LastName:  "Hiremath",
	}
	return c.JSON(usr)
}

func HandleGetUserByID(c *fiber.Ctx) error {
	return c.JSON("James")
}

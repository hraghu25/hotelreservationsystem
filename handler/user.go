package handler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hraghu25/hotelreservationsystem/database"
	"github.com/hraghu25/hotelreservationsystem/types"
)

type UserHandler struct {
	userStore database.UserStore
}

func NewUserHandler(userDataStore database.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userDataStore,
	}
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	res, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(res)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	usr := types.User{
		FirstName: "Raghavendra",
		LastName:  "Hiremath",
	}
	return c.JSON(usr)
}

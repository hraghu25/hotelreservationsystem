package handler

import "github.com/gofiber/fiber/v2"

func HandlerHealth(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"Health": "ok",
	})
}

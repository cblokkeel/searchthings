package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(host string) {
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/index/:collection", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("collection"))
	})
	app.Get("/search/:collection", func(c *fiber.Ctx) error {
		return c.SendString(c.Query("q"))
	})
	err := app.Listen(host)
	if err != nil {
		panic(err)
	}
}

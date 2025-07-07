package api

import "github.com/gofiber/fiber/v2"

func ApplyRouter(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello! By stream-radar!")
	})
}

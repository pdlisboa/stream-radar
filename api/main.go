package api

import "github.com/gofiber/fiber/v2"

func InitApi() {
	app := fiber.New()
	ApplyRouter(app)
	app.Listen(":3000")
}

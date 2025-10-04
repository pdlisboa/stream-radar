package router

import (
	"github.com/gofiber/fiber/v2"
	"stream-radar/api/middleware"
	"stream-radar/internal/modules/auth"
	"stream-radar/internal/modules/streamer"
	"stream-radar/internal/modules/user"
)

func ApplyRouter(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello! By stream-radar!")
	})

	//Auth
	app.Post("/login", auth.Login)

	//User
	userGroup := app.Group("/user")
	userGroup.Post("/", user.CreateUser)
	userGroup.Get("/", middleware.JWTProtected, user.GetUser)

	//Streamer
	streamerGroup := app.Group("/streamer")
	streamerGroup.Post("/", middleware.JWTProtected, streamer.RegisterStreamers)

}

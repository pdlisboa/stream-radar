package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
	_ "stream-radar/api/docs"
	"stream-radar/api/router"
	"stream-radar/internal/config"
	"stream-radar/internal/database"
	"stream-radar/internal/environment"
	"stream-radar/internal/logger"
	"stream-radar/internal/utils"
)

func main() {
	environment.InitEnv()
	configs := config.InitConfigs()
	database.Connect(configs.Db)
	InitApi()
	//worker.TestTwitch("esacarry")
	//worker.TestKick("esa")
	//worker.TestYt("esacarry")

}

// @title Stream-Radar
// @version 1.0
// @description  Radar online streams
// @BasePath /
func InitApi() {
	log := logger.GetInstance()
	app := fiber.New()
	router.ApplyRouter(app)
	app.Get("/swagger/*", swagger.HandlerDefault)
	err := app.Listen(fmt.Sprintf(":%s", utils.GetEnv("SERVER_PORT", "3000")))

	if err != nil {
		msg := "Unable to start api"
		log.Fatal(msg, zap.Error(err))
		panic(msg)
	}

}

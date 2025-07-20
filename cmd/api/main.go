package main

import (
	"stream-radar/internal/config"
	"stream-radar/internal/database"
	"stream-radar/internal/environment"
	"stream-radar/worker"
)

func init() {
	environment.InitEnv()
}

func main() {
	configs := config.InitConfigs()

	database.Connect(configs.Db)
	//api.InitApi()
	worker.TestTwitch("esacarry")
	worker.TestKick("esa")
	worker.TestYt("esacarry")

}

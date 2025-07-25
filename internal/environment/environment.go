package environment

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
		panic("Erro on load environment")
	}
}

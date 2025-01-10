package enviroments

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	envLoad := godotenv.Load()
	if envLoad != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", envLoad)
	}
}

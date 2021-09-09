package env

import (
	"log"

	"github.com/enorith/enorith/internal/pkg/path"
	"github.com/joho/godotenv"
)

func LoadDotenv() {
	file := path.BasePath(".env")
	log.Printf("loading dotenv file %s", file)
	_ = godotenv.Load(file)
}

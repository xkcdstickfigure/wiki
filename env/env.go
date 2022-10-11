package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Domain string = os.Getenv("DOMAIN")
var DatabaseUrl string = os.Getenv("DATABASE_URL")
var StorageOrigin string = os.Getenv("STORAGE_ORIGIN")

package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var Domain = os.Getenv("DOMAIN")
var Origin = os.Getenv("ORIGIN")
var StorageOrigin = os.Getenv("STORAGE_ORIGIN")

var DatabaseUrl = os.Getenv("DATABASE_URL")

var DiscordClientId = os.Getenv("DISCORD_CLIENT_ID")
var DiscordClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")

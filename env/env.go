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
var DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")

var GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
var GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

var GiteaOrigin = os.Getenv("GITEA_ORIGIN")

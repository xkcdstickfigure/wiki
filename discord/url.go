package discord

import (
	"net/url"

	"alles/wiki/env"
)

const scope = "email guilds guilds.join identify"

func GenerateUrl() string {
	values := url.Values{}
	values.Set("client_id", env.DiscordClientId)
	values.Set("redirect_uri", env.Origin+"/discord/callback")
	values.Set("response_type", "code")
	values.Set("scope", scope)

	return "https://discord.com/api/oauth2/authorize?" + values.Encode()
}

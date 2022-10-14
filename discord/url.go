package discord

import (
	"net/url"

	"alles/wiki/env"
)

const scope = "email guilds guilds.join identify"

func GenerateUrl(state string) string {
	values := url.Values{}
	values.Set("client_id", env.DiscordClientId)
	values.Set("redirect_uri", env.Origin+"/discord/callback")
	values.Set("response_type", "code")
	values.Set("scope", scope)
	values.Set("state", state)

	return "https://discord.com/api/oauth2/authorize?" + values.Encode()
}

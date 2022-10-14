package hub

import (
	"net/http"

	"alles/wiki/discord"
)

func (h handlers) discordAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, discord.GenerateUrl(), http.StatusTemporaryRedirect)
}

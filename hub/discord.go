package hub

import (
	"net/http"

	"alles/wiki/discord"
)

// join
func (h handlers) discordJoin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, discord.GenerateUrl(), http.StatusTemporaryRedirect)
}

// callback
func (h handlers) discordCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	// get discord information
	profile, err := discord.GetProfile(code)
	if err != nil {
		http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
		return
	}

	// join guild
	err = discord.JoinGuild("1029151219819753512", profile.User.Id, profile.Token.AccessToken)
	if err != nil {
		http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
		return
	}

	// redirect
	http.Redirect(w, r, "/discord/success", http.StatusTemporaryRedirect)
}

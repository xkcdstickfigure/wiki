package hub

import (
	"fmt"
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
	profile, err := discord.GetProfile(code)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	fmt.Println(profile.User.Username)
}

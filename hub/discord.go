package hub

import (
	"net/http"

	"alles/wiki/discord"
	"alles/wiki/store"
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

	// create discord user
	discordUser, err := h.db.DiscordUserCreate(r.Context(), store.DiscordUser{
		Id:            profile.User.Id,
		Username:      profile.User.Username,
		Discriminator: profile.User.Discriminator,
		Avatar:        profile.User.Avatar,
		MfaEnabled:    profile.User.MfaEnabled,
		Locale:        profile.User.Locale,
		Email:         profile.User.Email,
		EmailVerified: profile.User.EmailVerified,
		AccessToken:   profile.Token.AccessToken,
		RefreshToken:  profile.Token.RefreshToken,
	})
	if err != nil {
		http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
		return
	}

	// create guilds
	for _, guild := range profile.Guilds {
		err := h.db.DiscordGuildCreate(r.Context(), store.DiscordGuild{
			Id:   guild.Id,
			Name: guild.Name,
			Icon: guild.Icon,
		}, discordUser.Id)

		if err != nil {
			http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
			return
		}
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

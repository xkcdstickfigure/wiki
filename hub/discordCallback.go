package hub

import (
	"net/http"
	"time"

	"alles/wiki/discord"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"
)

func (h handlers) discordCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token := r.URL.Query().Get("state")

	// get session
	session, err := sessionAuth.GetSession(h.db, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// get state
	state, err := h.db.DiscordStateUse(r.Context(), token)
	if err != nil || state.SessionId != session.Id || time.Since(state.CreatedAt).Seconds() > 180 {
		w.WriteHeader(400)
		return
	}

	// get site
	site, err := h.db.SiteGetByName(r.Context(), state.Site)
	if err != nil {
		w.WriteHeader(400)
		return
	}

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
	err = discord.JoinGuild(site.DiscordGuild, profile.User.Id, profile.Token.AccessToken)
	if err != nil {
		http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
		return
	}

	// set discord user for session
	err = h.db.SessionSetDiscord(r.Context(), session.Id, discordUser.Id)
	if err != nil {
		http.Redirect(w, r, "/discord/error", http.StatusTemporaryRedirect)
		return
	}

	// redirect
	http.Redirect(w, r, "/discord/success", http.StatusTemporaryRedirect)
}

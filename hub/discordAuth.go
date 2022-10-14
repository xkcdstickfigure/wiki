package hub

import (
	"net/http"

	"alles/wiki/discord"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"
)

func (h handlers) discordAuth(w http.ResponseWriter, r *http.Request) {
	site := r.URL.Query().Get("site")

	// get session
	session, err := sessionAuth.GetSession(h.db, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// create state
	state, err := h.db.DiscordStateCreate(r.Context(), store.DiscordState{
		SessionId: session.Id,
		Site:      site,
	})
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// redirect
	http.Redirect(w, r, discord.GenerateUrl(state.Token), http.StatusTemporaryRedirect)
}

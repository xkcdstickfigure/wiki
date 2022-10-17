package hub

import (
	"net/http"

	"alles/wiki/discord"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"
)

func (h handlers) discordAuth(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	value := r.URL.Query().Get("value")

	// get session
	session, err := sessionAuth.UseSession(h.db, w, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// create state
	state, err := h.db.DiscordStateCreate(r.Context(), store.DiscordState{
		SessionId: session.Id,
		Action:    action,
		Value:     value,
	})
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// redirect
	http.Redirect(w, r, discord.GenerateUrl(state.Token), http.StatusTemporaryRedirect)
}

package hub

import (
	"net/http"
	"time"

	"alles/wiki/google"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"
)

func (h handlers) googleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token := r.URL.Query().Get("state")

	// get session
	session, err := sessionAuth.GetSession(h.db, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// get state
	state, err := h.db.AuthStateUse(r.Context(), token)
	if err != nil || state.SessionId != session.Id || time.Since(state.CreatedAt).Seconds() > 300 {
		w.WriteHeader(400)
		return
	}

	// get profile
	profile, err := google.GetProfile(code)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// create account
	account, err := h.db.AccountCreate(r.Context(), store.Account{
		GoogleId:      profile.Id,
		Name:          profile.Name,
		Email:         profile.Email,
		EmailVerified: profile.EmailVerified,
		Avatar:        profile.Picture,
	})
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// set account for session
	err = h.db.SessionSetAccount(r.Context(), session.Id, account.Id)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// sync discord account
	if account.DiscordId.String == "" {
		if session.DiscordId.String != "" {
			// transfer discord id for session to account
			h.db.AccountSetDiscord(r.Context(), account.Id, session.DiscordId.String)
		}
	} else if session.DiscordId.String == "" {
		// transfer discord id for account to session
		h.db.SessionSetDiscord(r.Context(), session.Id, account.DiscordId.String)
	}

	// redirect
	http.Redirect(w, r, state.Redirect, http.StatusTemporaryRedirect)
}

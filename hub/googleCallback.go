package hub

import (
	"fmt"
	"net/http"
	"time"

	"alles/wiki/google"
	"alles/wiki/sessionAuth"
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
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	fmt.Println(profile.Email)

	// redirect
	http.Redirect(w, r, state.Redirect, http.StatusTemporaryRedirect)
}

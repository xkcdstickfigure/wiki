package hub

import (
	"net/http"

	"alles/wiki/google"
	"alles/wiki/sessionAuth"
	"alles/wiki/store"
)

func (h handlers) googleAuth(w http.ResponseWriter, r *http.Request) {
	redirect := r.URL.Query().Get("redirect")
	if redirect == "" {
		redirect = "/"
	}

	// get session
	session, err := sessionAuth.UseSession(h.db, w, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// create state
	state, err := h.db.AuthStateCreate(r.Context(), store.AuthState{
		SessionId: session.Id,
		Redirect:  redirect,
	})
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// redirect
	http.Redirect(w, r, google.GenerateUrl(state.Token), http.StatusTemporaryRedirect)
}

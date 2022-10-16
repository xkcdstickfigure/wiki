package hub

import (
	"net/http"
	"net/url"

	"alles/wiki/env"
	"alles/wiki/sessionAuth"
)

func (h handlers) giteaAuth(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")

	// get session
	session, err := sessionAuth.UseSession(h.db, w, r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// require account
	if session.AccountId.String == "" {
		http.Redirect(w, r, "/auth?redirect="+url.QueryEscape(r.URL.String()), http.StatusTemporaryRedirect)
		return
	}

	// create code
	code, err := h.db.GiteaCodeCreate(r.Context(), session.AccountId.String)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// redirect
	http.Redirect(w, r, env.GiteaCallback+"?code="+code.Code+"&state="+url.QueryEscape(state), http.StatusTemporaryRedirect)
}

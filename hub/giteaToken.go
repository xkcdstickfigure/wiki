package hub

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h handlers) giteaToken(w http.ResponseWriter, r *http.Request) {
	// parse body
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// get code
	code, err := h.db.GiteaCodeUse(r.Context(), r.PostForm.Get("code"))
	if err != nil || time.Since(code.CreatedAt).Seconds() > 30 {
		w.WriteHeader(400)
		return
	}

	// create token
	token, err := h.db.GiteaTokenCreate(r.Context(), code.AccountId)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// response
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  token.Token,
		TokenType:    "bearer",
		ExpiresIn:    30,
		RefreshToken: "",
	})
}

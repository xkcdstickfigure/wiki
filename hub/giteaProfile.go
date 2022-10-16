package hub

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (h handlers) giteaProfile(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("access_token")

	// get token
	token, err := h.db.GiteaTokenUse(r.Context(), t)
	if err != nil || time.Since(token.CreatedAt).Seconds() > 30 {
		w.WriteHeader(400)
		return
	}

	// get account
	account, err := h.db.AccountGet(r.Context(), token.AccountId)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// response
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Id     int    `json:"id"`
		Login  string `json:"login"`
		Email  string `json:"email"`
		Active bool   `json:"active"`
	}{
		Id:     account.Number,
		Login:  strconv.Itoa(account.Number),
		Email:  strconv.Itoa(account.Number) + "@glaffle.com",
		Active: true,
	})
}

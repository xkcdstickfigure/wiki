package hub

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	// response
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Active            bool   `json:"active"`
		AvatarUrl         string `json:"avatar_url"`
		Created           string `json:"created"`
		Description       string `json:"description"`
		Email             string `json:"email"`
		FollowersCount    int    `json:"followers_count"`
		FollowingCount    int    `json:"following_count"`
		FullName          string `json:"full_name"`
		Id                int    `json:"id"`
		IsAdmin           bool   `json:"is_admin"`
		Language          string `json:"language"`
		LastLogin         string `json:"last_login"`
		Location          string `json:"location"`
		Login             string `json:"login"`
		ProhibitLogin     bool   `json:"prohibit_login"`
		Restricted        bool   `json:"restricted"`
		StarredReposCount int    `json:"starred_repos_count"`
		Visibility        string `json:"visibility"`
		Website           string `json:"website"`
	}{
		Active:     true,
		Email:      "user123@glaffle.com",
		FullName:   account.Name,
		Id:         123,
		Login:      "user123",
		Visibility: "public",
	})
}

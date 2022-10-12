package discord

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"alles/wiki/env"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	Locale        string `json:"locale"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"verified"`
	Premium       int    `json:"premium_type"`
}

type Guild struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Profile struct {
	User   User
	Token  Token
	Guilds []Guild
}

func GetProfile(code string) (Profile, error) {
	// make token request
	values := url.Values{}
	values.Set("client_id", env.DiscordClientId)
	values.Set("client_secret", env.DiscordClientSecret)
	values.Set("redirect_uri", env.Origin+"/discord/callback")
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)

	resp, err := http.PostForm("https://discord.com/api/v10/oauth2/token", values)
	if err != nil {
		return Profile{}, err
	} else if resp.StatusCode != 200 {
		return Profile{}, errors.New("token request failed")
	}

	// parse token response
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return Profile{}, err
	}

	// compare scope
	tokenScope := strings.Split(token.Scope, " ")
	sort.Strings(tokenScope)
	if strings.Join(tokenScope, " ") != scope {
		return Profile{}, errors.New("invalid scope")
	}

	// make user request
	req, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
	if err != nil {
		return Profile{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return Profile{}, err
	} else if resp.StatusCode != 200 {
		return Profile{}, errors.New("user request failed")
	}

	// parse user response
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return Profile{}, err
	}

	// make guilds request
	req, err = http.NewRequest("GET", "https://discord.com/api/v10/users/@me/guilds", nil)
	if err != nil {
		return Profile{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return Profile{}, err
	} else if resp.StatusCode != 200 {
		return Profile{}, errors.New("guilds request failed")
	}

	// parse guilds response
	var guilds []Guild
	err = json.NewDecoder(resp.Body).Decode(&guilds)
	if err != nil {
		return Profile{}, err
	}

	// return
	return Profile{user, token, guilds}, nil
}

package gitea

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"alles/wiki/env"
)

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
	Sha      string `json:"sha"`
}

func GetArticleSource(site string, slug string) (string, error) {
	// make request
	resp, err := http.Get(env.GiteaOrigin + "/api/v1/repos/glaffle/" + url.QueryEscape(site) + "/contents/" + url.QueryEscape(slug) + "?token=" + url.QueryEscape(env.GiteaToken) + "&gitea_access=" + url.QueryEscape(env.GiteaAccess))
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", errors.New("request failed")
	}

	// parse response
	var file File
	err = json.NewDecoder(resp.Body).Decode(&file)
	if err != nil || file.Type != "file" || file.Encoding != "base64" {
		return "", err
	}

	// decode base64
	source, err := base64.StdEncoding.DecodeString(file.Content)
	if err != nil {
		return "", err
	}

	return string(source), nil
}

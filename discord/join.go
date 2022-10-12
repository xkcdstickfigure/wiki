package discord

import (
	"alles/wiki/env"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func JoinGuild(guildId string, userId string, accessToken string) error {
	body, err := json.Marshal(struct {
		AccessToken string `json:"access_token"`
	}{accessToken})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "https://discord.com/api/v10/guilds/"+guildId+"/members/"+userId, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bot "+env.DiscordBotToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 201 && resp.StatusCode != 204 {
		return errors.New("joining user to guild failed")
	}

	return nil
}

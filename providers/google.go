package provider

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func newGoogle(config *oauth2.Config) *Provider {
	if config == nil {
		config = &oauth2.Config{}
	}
	config.Endpoint = google.Endpoint
	config.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	return &Provider{
		config: config,
	}
}

func getGoogleUserInfo(ctx context.Context, config *oauth2.Config, accessToken *oauth2.Token) (*UserInfo, error) {
	client := config.Client(ctx, accessToken)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		Email:         userInfo.Email,
		EmailVerified: true,
		Name:          userInfo.Name,
		Picture:       userInfo.Picture,
		Locale:        "",
	}, nil
}

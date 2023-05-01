package oauth2providers

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func newGoogle(options *authOptions) *Provider {
	// google endpoint
	options.Endpoint = google.Endpoint
	if options.Scopes == nil {
		options.Scopes = []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		}
	}
	return &Provider{
		authOptions: options,
	}
}

func getGoogleUserInfo(ctx context.Context, authOptions *authOptions, accessToken *oauth2.Token) (*UserInfo, error) {
	client := authOptions.Client(ctx, accessToken)
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

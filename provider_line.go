package oauth2providers

import (
	"context"
	"encoding/json"

	"golang.org/x/oauth2"
)

func newLine(options *authOptions) *Provider {
	// line endpoint
	options.Endpoint = oauth2.Endpoint{
		AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
		TokenURL: "https://api.line.me/oauth2/v2.1/token",
	}
	if options.Scopes == nil {
		options.Scopes = []string{
			"profile",
			"openid",
			"email",
		}
	}

	return &Provider{
		authOptions: options,
	}
}

func getLineUserInfo(ctx context.Context, options *authOptions, accessToken *oauth2.Token) (*UserInfo, error) {
	client := options.Client(ctx, accessToken)
	resp, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID      string `json:"userId"`
		Email   string `json:"email"`
		Name    string `json:"displayName"`
		Picture string `json:"pictureUrl"`
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

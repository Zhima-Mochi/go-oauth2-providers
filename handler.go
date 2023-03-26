package goauthproviders

import (
	provider "github.com/Zhima-Mochi/go-auth-providers/providers"
)

type Oauth2Handler struct {
	provider *provider.Provider
}

func NewOauth2Handler(providerType provider.ProviderType, clientID, clientSecret, redirectURL string) *Oauth2Handler {
	return &Oauth2Handler{
		provider: provider.NewProvider(providerType, clientID, clientSecret, redirectURL),
	}
}

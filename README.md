# go-oauth2-providers
go-oauth2-providers is a Go library for OAuth2 authentication integration. It provides a unified interface to create OAuth2 provider instances and retrieve user information, and supports multiple providers, including Google, Facebook, and GitHub. It's extensively tested and can be used in projects of any size.

## Usage

```go
// Create an empty ProviderOption slice
providerOptions := []oauth2providers.ProviderOption{}

// Add Google OAuth2 Client ID to the ProviderOption slice
providerOptions = append(providerOptions, oauth2providers.WithProviderClientID("<YOUR_CLIENT_ID>"))

// Add Google OAuth2 Client Secret to the ProviderOption slice
providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret("<YOUR_CLIENT_SECRET>"))

// Add Google OAuth2 Redirect URL to the ProviderOption slice
providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL("<YOUR_REDIRECT_URL>"))

// Create a ProviderConfig object with the ProviderOption slice
providerConfig, _ := oauth2providers.NewProviderConfig(providerOptions...)

// Create an OAuth2Provider object with the GoogleOAuth2ProviderType and the ProviderConfig object
oauth2Provider, _ := oauth2providers.NewOAuth2Provider(oauth2providers.GoogleOAuth2ProviderType, providerConfig)

// Create a new Auth object with the OAuth2Provider object
googleAuth := oauth2providers.NewAuth(oauth2Provider, nil)
```
### Auth Interface
```go
type Auth interface {
	GetOAuth2AuthCodeURL(ctx context.Context, options ...AuthCodeOption) (url string)
	ExchangeOAuth2AuthCode(ctx context.Context, code string) (token Token, err error)
	RefreshOAuth2Token(ctx context.Context, refreshToken Token) (token Token, err error)
	GetOAuth2UserInfo(ctx context.Context, token Token) (userInfo UserInfo, err error)
}
```
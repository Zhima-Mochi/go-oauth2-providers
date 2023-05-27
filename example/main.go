package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	oauth2providers "github.com/Zhima-Mochi/go-oauth2-providers"
)

func main() {
	// example: go run main.go google
	argv := []string{}
	argv = append(argv, os.Args...)
	args := len(argv)
	if args < 2 {
		fmt.Println("error: missing provider name")
		return
	}
	var handler oauth2providers.Auth
	if argv[1] == "google" {
		handler = GetGoogleHandler()
	} else if argv[1] == "facebook" {
		handler = GetFacebookHandler()
	} else if argv[1] == "line" {
		handler = GetLineHandler()
	} else if argv[1] == "github" {
		handler = GetGithubHandler()
	} else {
		fmt.Println("error: provider name not found")
		return
	}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		opts := []oauth2providers.AuthCodeOption{}
		url := handler.GetOAuth2AuthCodeURL(r.Context(), opts...)
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body>
		<a href="/login">Login</a>
		</body></html>`)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "code cannot be empty", http.StatusBadRequest)
			return
		}

		token, err := handler.ExchangeOAuth2AuthCode(r.Context(), code)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get access token: %v", err), http.StatusInternalServerError)
			return
		}

		userInfo, err := handler.GetOAuth2UserInfo(r.Context(), token)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get user info: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Hello, %s!", userInfo.GetName())
	})

	server := &http.Server{
		Addr:        ":8080",
		ReadTimeout: 5 * time.Second,
	}
	fmt.Println("Server started on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func GetGoogleHandler() oauth2providers.Auth {
	providerOptions := []oauth2providers.ProviderOption{}
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientID("<client_id>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret("<client_secret>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL("https://example.com/callback"))
	providerConfig, _ := oauth2providers.NewProviderConfig(providerOptions...)
	oauth2Provider, _ := oauth2providers.NewOAuth2Provider(oauth2providers.GoogleOAuth2ProviderType, providerConfig)

	handler := oauth2providers.NewOAuth2Auth(oauth2Provider, nil)
	return handler
}

func GetFacebookHandler() oauth2providers.Auth {
	providerOptions := []oauth2providers.ProviderOption{}
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientID("<client_id>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret("<client_secret>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL("https://example.com/callback"))
	providerConfig, _ := oauth2providers.NewProviderConfig(providerOptions...)
	oauth2Provider, _ := oauth2providers.NewOAuth2Provider(oauth2providers.FacebookOAuth2ProviderType, providerConfig)

	handler := oauth2providers.NewOAuth2Auth(oauth2Provider, nil)

	return handler
}

func GetLineHandler() oauth2providers.Auth {
	providerOptions := []oauth2providers.ProviderOption{}
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientID("<client_id>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret("<client_secret>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL("https://example.com/callback"))
	providerConfig, _ := oauth2providers.NewProviderConfig(providerOptions...)
	oauth2Provider, _ := oauth2providers.NewOAuth2Provider(oauth2providers.LineOAuth2ProviderType, providerConfig)

	handler := oauth2providers.NewOAuth2Auth(oauth2Provider, nil)

	return handler
}

func GetGithubHandler() oauth2providers.Auth {
	providerOptions := []oauth2providers.ProviderOption{}
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientID("<client_id>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret("<client_secret>"))
	providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL("https://example.com/callback"))
	providerConfig, _ := oauth2providers.NewProviderConfig(providerOptions...)
	oauth2Provider, _ := oauth2providers.NewOAuth2Provider(oauth2providers.GithubOAuth2ProviderType, providerConfig)

	handler := oauth2providers.NewOAuth2Auth(oauth2Provider, nil)

	return handler
}

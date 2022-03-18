package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "1074900166822-7v6a3nf33a1ifi3tat37n4aa1is58tng.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8000/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

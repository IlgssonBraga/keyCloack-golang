package main

import (
	"context"
	oidc "github.com/coreos/go-oidc"
)

var (
	clientID = "app"
	clientSecret = "83858004-a187-449d-9647-ef5e486cafe9"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")

	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		Endpoint: provider.Endpoint(),
		RedirectURL: "http://localhost:8081/auth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "magica"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
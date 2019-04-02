package youtube

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi/transport"

	ytapi "google.golang.org/api/youtube/v3"
)

type Client struct {
	service *ytapi.Service
}

func NewClient(apiKey string, config *oauth2.Config) (*Client, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	})
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("can't scan code: %v", err)
	}
	token, err := config.Exchange(ctx, code)
	client := config.Client(ctx, token)
	service, err := ytapi.New(client)
	if err != nil {
		return nil, fmt.Errorf("can't create new service: %v", err)
	}
	return &Client{service: service}, nil
}

func Config(id, secret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Endpoint:     google.Endpoint,
		Scopes:       []string{ytapi.YoutubeScope},
	}
}

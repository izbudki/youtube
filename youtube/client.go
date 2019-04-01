package youtube

import (
	"net/http"

	ytapi "google.golang.org/api/youtube/v3"
)

type Client struct {
	service *ytapi.Service
}

func NewClient() (*Client, error) {
	client := http.Client{}
	service, err := ytapi.New(&client)
	if err != nil {
		return nil, err
	}
	return &Client{service: service}, nil
}

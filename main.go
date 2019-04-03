package main

import (
	"fmt"
	"log"
	"os"

	"github.com/izbudki/youtube/youtube"
)

const (
	envAPIKey       = "API_KEY"
	envClientID     = "CLIENT_ID"
	envClientSecret = "CLIENT_SECRET"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: youtube [broadcast name]")
	}
	name := os.Args[1]

	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		log.Fatalf("%v isn't provided", envAPIKey)
	}
	clientID := os.Getenv(envClientID)
	if clientID == "" {
		log.Fatalf("%v isn't provided", envClientID)
	}
	clientSecret := os.Getenv(envClientSecret)
	if clientSecret == "" {
		log.Fatalf("%v isn't provided", envClientSecret)
	}

	config := youtube.Config(clientID, clientSecret)
	client, err := youtube.NewClient(apiKey, config)
	if err != nil {
		log.Fatalf("can't create a new client: %v", err)
	}

	err = client.CreateBroadcast(name)
	if err != nil {
		log.Fatalf("can't create a new broadcast: %v", err)
	}
	ii, err := client.EncoderSetup()
	fmt.Println(ii.ServerURL, ii.StreamKey)
}

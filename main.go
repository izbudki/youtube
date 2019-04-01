package main

import (
	"log"
	"os"

	"github.com/izbudki/youtube/youtube"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Usage: youtube [playlist name]")
	}
	playlist := os.Args[0]

	client, err := youtube.NewClient()
	if err != nil {
		log.Fatalf("unable to create YouTube service: %v", err)
	}

	err = client.CreatePlaylist(playlist)
	if err != nil {
		log.Fatalf("can't create a new playlist: %v", err)
	}
}

package youtube

import (
	"fmt"

	ytapi "google.golang.org/api/youtube/v3"
)

func (c *Client) CreatePlaylist(name string) error {
	playlist := &ytapi.Playlist{
		Snippet: &ytapi.PlaylistSnippet{
			Title:       name,
			Description: "CREATED VIA YOUTUBE API",
			Tags:        []string{"test", "api", "playlist"},
		},
	}
	call := c.service.Playlists.Insert("snippet", playlist)
	resp, err := call.Do()
	if err != nil {
		return fmt.Errorf("can't send a request: %v", err)
	}
	fmt.Println(resp)
	return nil
}

package youtube

import (
	"fmt"

	ytapi "google.golang.org/api/youtube/v3"
)

func (c *Client) CreateBroadcast(name string) error {
	broadcast := &ytapi.LiveBroadcast{
		Snippet: &ytapi.LiveBroadcastSnippet{
			Title:       name,
			Description: "CREATED VIA YOUTUBE API",
		},
	}
	call := c.service.LiveBroadcasts.Insert("snippet", broadcast)
	response, err := call.Do()
	if err != nil {
		return fmt.Errorf("can't send a request: %v", err)
	}
	fmt.Println(response.Id)
	return nil
}

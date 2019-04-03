package youtube

import (
	"fmt"
	"time"

	ytapi "google.golang.org/api/youtube/v3"
)

type IngestionInfo struct {
	ServerURL string
	StreamKey string
}

func (c *Client) CreateBroadcast(name string) error {
	broadcast := &ytapi.LiveBroadcast{
		Snippet: &ytapi.LiveBroadcastSnippet{
			Title:              name,
			Description:        "CREATED VIA YOUTUBE API",
			ScheduledStartTime: time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		},
		Status: &ytapi.LiveBroadcastStatus{
			PrivacyStatus: "public",
		},
	}
	_, err := c.service.LiveBroadcasts.Insert("snippet,status,contentDetails", broadcast).Do()
	if err != nil {
		return fmt.Errorf("can't send a request: %v", err)
	}
	return nil
}

func (c *Client) EncoderSetup() (*IngestionInfo, error) {
	broadcastsResponse, err := c.service.LiveBroadcasts.List("id,snippet,contentDetails").BroadcastType("persistent").Mine(true).Do()
	if err != nil {
		return nil, fmt.Errorf("can't list broadcasts: %v", err)
	}
	boundStreamID := broadcastsResponse.Items[0].ContentDetails.BoundStreamId
	streamsResponse, err := c.service.LiveStreams.List("id,snippet,cdn").Id(boundStreamID).Do()
	if err != nil {
		return nil, fmt.Errorf("can't list live streams: %v", err)
	}
	ingestionInfo := streamsResponse.Items[0].Cdn.IngestionInfo
	return &IngestionInfo{
		ServerURL: ingestionInfo.IngestionAddress,
		StreamKey: ingestionInfo.StreamName,
	}, nil
}

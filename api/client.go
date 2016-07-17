package api

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// Client is a wrapper of slack.Client.
type Client struct {
	*slack.Client
}

// NewClient is a constructor.
func NewClient() *Client {
	client := slack.New(os.Getenv("KOMONJO_SLACK_API_TOKEN"))
	client.SetDebug(false)
	ret := Client{
		client,
	}
	return &ret
}

// GetChannel returns a slack.Channel searched by name.
func (c *Client) GetChannel(name string) (slack.Channel, error) {
	channels, err := c.GetChannels(false)
	if err != nil {
		return slack.Channel{}, err
	}
	for _, ch := range channels {
		if ch.Name == name {
			return ch, nil
		}
	}
	return slack.Channel{}, fmt.Errorf("No such channel: %s", name)
}

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

// GetUserMap returns User.ID -> User map.
func (c *Client) GetUserMap() (map[string]slack.User, error) {
	ret := make(map[string]slack.User)
	users, err := c.GetUsers()
	if err != nil {
		return ret, err
	}
	for _, u := range users {
		ret[u.ID] = u
	}
	return ret, nil
}

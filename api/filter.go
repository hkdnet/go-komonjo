package api

import (
	"strings"

	"github.com/nlopes/slack"
)

// ChannelFilterOption is a option for FilterChannel.
type ChannelFilterOption struct {
	Name string
	Mode int
}

// FilterChannel returns channel filtered by name.
func FilterChannel(channels []slack.Channel, name string) []slack.Channel {
	ret := make([]slack.Channel, 0, len(channels))
	for _, channel := range channels {
		if strings.Contains(channel.Name, name) {
			ret = append(ret, channel)
		}
	}
	return ret
}

package command

import (
	"fmt"
	"strings"

	"github.com/hkdnet/go-komonjo/api"
)

// ListCommand will show the history of a channel.
type ListCommand struct {
	Meta
}

// Run is a main method.
func (c *ListCommand) Run(args []string) int {
	client := api.NewClient()
	channels, err := client.GetChannels(false)
	if err != nil {
		return c.DealError(err)
	}
	for _, channel := range channels {
		var tmp string
		if channel.IsArchived {
			tmp = "-"
		} else {
			tmp = "*"
		}
		fmt.Printf("%s  %s\t#%s\n", tmp, channel.ID, channel.Name)
	}
	return 0
}

// Synopsis shows what this subcommand will do.
func (c *ListCommand) Synopsis() string {
	return "list something"
}

// Help shows how to use this subcommand.
func (c *ListCommand) Help() string {
	helpText := `
komonjo list TYPE
	TYPE is one of ...
		- channel
`
	return strings.TrimSpace(helpText)
}

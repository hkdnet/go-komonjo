package command

import (
	"fmt"
	"strings"
	"sync"

	"github.com/hkdnet/go-komonjo/api"
	"github.com/hkdnet/slack"
)

// ShowCommand will show the history of a channel.
type ShowCommand struct {
	Meta
}

// Run is a main method.
func (c *ShowCommand) Run(args []string) int {
	if len(args) < 1 {
		c.Ui.Warn(c.Help())
		return 1
	}
	client := api.NewClient()
	wg := new(sync.WaitGroup)
	ch := make(chan bool)
	var history *slack.History
	var userMap map[string]slack.User
	go func() {
		for {
			select {
			case <-ch:
				wg.Done()
			}
		}
	}()
	wg.Add(1)
	go func() {
		his, err := client.GetChannelHistoryByName(args[0], newHistoryParameters())
		if err != nil {
			panic(err)
		}
		history = his
		ch <- true
	}()
	wg.Add(1)
	go func() {
		ret, err := client.GetUserMap()
		if err != nil {
			panic(err)
		}
		userMap = ret
		ch <- true
	}()
	wg.Wait()
	for _, message := range history.Messages {
		fmt.Printf("[%s] %q\n", userMap[message.User].Name, message.Text)
	}
	return 0
}

func newHistoryParameters() slack.HistoryParameters {
	ret := slack.HistoryParameters{}
	ret.Count = 100
	/*
		Latest    string
		Oldest    string
		Count     int
		Inclusive bool
		Unreads   bool
	*/
	return ret
}

// Synopsis shows what this subcommand will do.
func (c *ShowCommand) Synopsis() string {
	return "show history of a channel"
}

// Help shows how to use this subcommand.
func (c *ShowCommand) Help() string {
	helpText := `
komonjo show [CHANNEL]
`
	return strings.TrimSpace(helpText)
}

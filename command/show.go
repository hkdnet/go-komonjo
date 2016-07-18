package command

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/hkdnet/go-komonjo/api"
	"github.com/nlopes/slack"
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
	go func() {
		for {
			select {
			case done := <-ch:
				if !done {
					panic(errors.New("?"))
				}
				wg.Done()
			}
		}
	}()
	wg.Add(1)
	go func() {
		channel, err := client.GetChannel(args[0])
		if err != nil {
			c.DealError(err)
			ch <- false
			return
		}
		his, err := client.GetChannelHistory(channel.ID, newHistoryParameters())
		if err != nil {
			c.DealError(err)
			ch <- false
			return
		}
		history = his
		ch <- true
	}()
	wg.Wait()
	for _, message := range history.Messages {
		fmt.Printf("[%s] %q\n", message.User, message.Text)
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

package command

import (
	"flag"
	"strings"

	"github.com/hkdnet/go-komonjo/server"
)

// ServerCommand will show the history of a channel.
type ServerCommand struct {
	Meta
}

var (
	serverPort int
)

// Run is a main method.
func (c *ServerCommand) Run(args []string) int {
	c.fs = c.newFlagSet()
	c.Parse(args)
	server.Up(server.Option{
		Port: serverPort,
	})
	return 0
}

func (c *ServerCommand) newFlagSet() *flag.FlagSet {
	ret := flag.NewFlagSet("server", flag.ContinueOnError)
	ret.IntVar(&serverPort, "p", 8080, "port to listen on")
	return ret
}

// Synopsis shows what this subcommand will do.
func (c *ServerCommand) Synopsis() string {
	return "run a http server"
}

// Help shows how to use this subcommand.
func (c *ServerCommand) Help() string {
	helpText := `
komonjo server [OPTIONS]
`
	return strings.TrimSpace(helpText)
}

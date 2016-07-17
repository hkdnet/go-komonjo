package command

import (
	"flag"

	"github.com/mitchellh/cli"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.Ui
	fs *flag.FlagSet
}

// SetFlagSet is a setter for flag.FlagSet
func (m *Meta) SetFlagSet(fs *flag.FlagSet) {
	m.fs = fs
}

// Parse parses all arguments. This allows option arguments after non-option arguments.
func (m *Meta) Parse(args []string) []string {
	ret := []string{}
	m.fs.Parse(args)
	for m.fs.NArg() != 0 {
		tmp := m.fs.Args()
		ret = append(ret, tmp[0])
		m.fs.Parse(tmp[1:])
	}
	return ret
}

// DealError shows error and return 1
func (m *Meta) DealError(err error) int {
	m.Ui.Error(err.Error())
	return 1
}

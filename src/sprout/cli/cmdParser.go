package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/docopt/docopt-go"
	"github.com/kataras/golog"
)

// SproutOptions represnts the different inputs that can be passed in
type SproutOptions struct {
	Command string
	Args    []string
	Help    bool
	Verbose bool
}

const usage = `usage:
  sprout [options] <command> [<args...>]

Options:
  -h --help      show this screen
  -X --verbose   show all logs when executing commands

Commands:
  create    open up a dialog to select which repos to clone in the folder structure specified in sprout_config.yml
  spread    execute the command following "spread" in all repos
`

// ParseCommand takes in the command line entry and parses it into a sproutOptions struct
func ParseCommand() (options SproutOptions) {
	parser := &docopt.Parser{OptionsFirst: true}
	args, err := parser.ParseArgs(usage, os.Args[1:], "")
	if err != nil {
		golog.Fatal(err)
	}

	options = SproutOptions{}
	options.Command, _ = args.String("<command>")
	options.Args, _ = args["<args>"].([]string)
	options.Help, _ = args.Bool("--help")
	options.Verbose, _ = args.Bool("--verbose")
	return
}

// RunCommand executes the commands given in sproutOptions
func RunCommand(config model.Config, options SproutOptions) {
	if options.Help {
		docopt.PrintHelpAndExit(nil, usage)
		return
	}
	switch options.Command {
	case "spread":
		SpreadCommand()
	case "create":
		CreateCommand(config)
	default:
		docopt.PrintHelpAndExit(errors.New(fmt.Sprint("Command not found")), usage)

	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/jessevdk/go-flags"
	"github.com/kataras/golog"
	"github.com/mitchellh/go-homedir"

	"gopkg.in/yaml.v2"
)

var opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Spread  []bool `short:"s" long:"spread" description:"Execute command across all repos"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		golog.Fatal(err)
	}
	if len(opts.Verbose) > 0 && opts.Verbose[0] {
		golog.SetLevel("debug")
	}
	golog.Debug("Checking for config")
	ParseConfig()
	// RunCommand(config)
}

// func RunCommand(config model.Config) {

// }

func ParseConfig() model.Config {
	dir, err := homedir.Dir()
	if err != nil {
		golog.Fatal(err)
	}
	path := fmt.Sprintf("%s%s%s", dir, utils.PathSeparator, utils.ConfigFileName)
	golog.Info(path)
	if !utils.FileExists(path) {
		golog.Fatal(fmt.Sprintf(`No %s found in "%s", terminating`, utils.ConfigFileName, dir))
	}

	configYaml, err := ioutil.ReadFile(path)
	if err != nil {
		golog.Fatal(err)
	}

	config := model.Config{}
	if err := yaml.Unmarshal(configYaml, &config); err != nil {
		golog.Fatal(err)
	}

	return config
}

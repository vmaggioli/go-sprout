package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Fair2Dare/sprout/src/sprout/cli"
	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"github.com/mitchellh/go-homedir"

	"gopkg.in/yaml.v2"
)

func main() {
	golog.SetLevel("debug")
	config := ParseConfig()
	options := cli.ParseCommand()
	if config.Verbose || options.Verbose {
		golog.SetLevel("debug")
		golog.Debug("Debug logging enabled")
	}
	cli.RunCommand(options)
}

// ParseConfig reads from sprout_config.yml to generate the project layout
func ParseConfig() model.Config {
	dir, err := homedir.Dir()
	if err != nil {
		golog.Fatal(err)
	}
	path := fmt.Sprintf("%s%s%s", dir, utils.PathSeparator, utils.ConfigFileName)
	if !utils.FileExists(path) {
		golog.Fatal(fmt.Sprintf(`No %s found in "%s"`, utils.ConfigFileName, dir))
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

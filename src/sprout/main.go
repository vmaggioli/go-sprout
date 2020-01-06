package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"github.com/mitchellh/go-homedir"

	"gopkg.in/yaml.v2"
)

func main() {
	golog.SetLevel("debug")
	golog.Debug("Checking for config")
	ParseConfig()
}

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
		golog.Fatal("Error reading file, terminating")
	}

	config := model.Config{}
	if err := yaml.Unmarshal(configYaml, &config); err != nil {
		golog.Fatal("Error parsing config, terminating")
	}

	golog.Info(config.Structure.Projects[0].Projects[0].Name)
	return config
}

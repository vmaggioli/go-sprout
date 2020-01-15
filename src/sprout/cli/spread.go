package cli

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v2"
)

// SpreadCommand executes the provided command across all cloned repositories
func SpreadCommand(command string) {
	golog.Info("Spreading \"%s\"", command)
	splitCommand := strings.Fields(command)
	currDir, _ := os.Getwd()
	sproutRootPath := currDir + utils.PathSeparator + ".sprout_root.yml"
	if !utils.FileExists(sproutRootPath) {
		golog.Error("No sprouted project root exists in current directory")
		return
	}

	sproutRoot := model.SproutRoot{Repos: []model.RepoWithPath{}}
	file, err := ioutil.ReadFile(sproutRootPath)
	if err != nil {
		golog.Fatal(err)
	}
	err = yaml.Unmarshal(file, &sproutRoot)
	if err != nil {
		golog.Fatal(err)
	}
	for _, repo := range sproutRoot.Repos {
		golog.Infof("Inside \"%s\"", repo.Name)
		cmd := exec.Command(splitCommand[0], splitCommand[1:]...)
		cmd.Dir = repo.Path
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			golog.Error(err)
		}
	}
}

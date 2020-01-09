package cli

import (
	"os"
	"os/exec"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
)

// CreateCommand opens a prompt to select which repositories to clone
func CreateCommand(config model.Config) {
	createStructure(config.Projects, config.Repos)
}

func createStructure(projects []model.Project, repos []string) {
	for _, repoURL := range repos {
		golog.Debugf("Cloning %s", repoURL)
		cmd := exec.Command("git", "clone", repoURL)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			golog.Error(err)
		}
	}

	for _, project := range projects {
		golog.Debugf("Creating directory for project \"%s\"", project.Name)
		err := utils.Mkdir(project.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		os.Chdir(project.Name)
		createStructure(project.Projects, project.Repos)
		os.Chdir("..")
	}
}

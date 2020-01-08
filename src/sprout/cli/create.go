package cli

import (
	"fmt"
	"os"

	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"gopkg.in/src-d/go-git.v4"
)

// CreateCommand opens a prompt to select which repositories to clone
func CreateCommand(config model.Config) {
	createStructure(config.Projects, config.Repos)
}

func createStructure(projects []model.Project, repos []string) {
	fmt.Println(len(projects))
	for _, repoURL := range repos {
		fmt.Println(repos)
		golog.Debug("Cloning %s", repoURL)
		git.PlainClone(".", false, &git.CloneOptions{URL: repoURL, Progress: os.Stdout})
	}

	for _, project := range projects {
		golog.Debug("Creating directory for project %s", project)
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

package cli

import (
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
)

// CreateCommand opens a prompt to select which repositories to clone
func CreateCommand(config *model.Config) {
	options := config.GetAllRepos()
	selected := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select repos to clone",
		Options: options,
	}
	survey.AskOne(prompt, &selected)
	createStructure(config.Projects, config.Repos, selected)
}

func createStructure(projects []model.Project, repos []model.Repo, reposToClone []string) {
	for _, project := range projects {
		golog.Debugf("Creating directory for project \"%s\"", project.Name)
		err := utils.Mkdir(project.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		os.Chdir(project.Name)
		createStructure(project.Projects, project.Repos, reposToClone)
		os.Chdir("..")
	}

	for _, repo := range repos {
		if !utils.Contains(reposToClone, repo.Name) {
			continue
		}
		golog.Debugf("Cloning %s", repo.Name)
		cmd := exec.Command("git", "clone", repo.URL)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			golog.Error(err)
		}
	}
}

package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v2"
)

// CreateCommand opens a prompt to select which repositories to clone
func CreateCommand(config *model.Config) {
	currDir, _ := os.Getwd()
	if utils.FileExists(fmt.Sprintf("%s/%s", currDir, ".sprout_root.yml")) {
		golog.Error("Sprouted project already exsists, please run \"sprout slash\" to remove project and start a new one")
		return
	}
	options := config.GetAllRepos()
	selected := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select repos to clone",
		Options: options,
	}
	survey.AskOne(prompt, &selected)
	createStructure(config.Projects, config.Repos, selected, true)
}

func createStructure(projects []model.Project, repos []model.Repo, reposToClone []string, isRoot bool) {
	branch := model.SproutBranch{Folders: []string{}, Repos: []string{}}
	for _, project := range projects {
		golog.Debugf("Creating directory for project \"%s\"", project.Name)
		branch.Folders = append(branch.Folders, project.Name)
		err := utils.Mkdir(project.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		os.Chdir(project.Name)
		createStructure(project.Projects, project.Repos, reposToClone, false)
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

	golog.Debug("Generating branch file")
	yamlFile, err := yaml.Marshal(&branch)
	if err != nil {
		golog.Fatal(err)
	}
	var file *os.File

	if isRoot {
		file, err = os.Create(".sprout_root.yml")
	} else {
		file, err = os.Create(".sprout_branch.yml")
	}

	if err != nil {
		golog.Fatal(err)
	}
	file.Write(yamlFile)
	defer file.Close()

}

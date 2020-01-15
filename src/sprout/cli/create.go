package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Fair2Dare/sprout/src/sprout/model"
	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v2"
)

// CreateCommand opens a prompt to select which repositories to clone
func CreateCommand(config *model.Config) {
	currDir, _ := os.Getwd()
	if utils.FileExists(fmt.Sprintf("%s%s%s", currDir, utils.PathSeparator, model.SproutRootFileName)) {
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
	file, err := os.Create(model.SproutRootFileName)
	if err != nil {
		golog.Fatal(err)
	}
	root := &model.SproutRoot{Repos: []model.RepoWithPath{}}
	createStructure(config.Projects, config.Repos, selected, root)
	text, err := yaml.Marshal(root)
	if err != nil {
		golog.Fatal(err)
	}
	file.Write(text)
	defer file.Close()
}

func createStructure(projects []model.Project, repos []model.Repo, reposToClone []string, sproutRoot *model.SproutRoot) {
	for _, project := range projects {
		golog.Debugf("Creating directory for project \"%s\"", project.Name)
		err := utils.Mkdir(project.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		os.Chdir(project.Name)
		createStructure(project.Projects, project.Repos, reposToClone, sproutRoot)
		os.Chdir("..")
	}

	for _, repo := range repos {
		if !utils.Contains(reposToClone, repo.Name) {
			continue
		}
		currPath, _ := os.Getwd()
		absCurrPath, _ := filepath.Abs(currPath)
		urlRepoName := getRepoNameFromUrl(repo.URL)
		newRepo := model.RepoWithPath{Name: urlRepoName, Path: fmt.Sprintf("%s%s%s", absCurrPath, utils.PathSeparator, urlRepoName)}
		sproutRoot.Repos = append(sproutRoot.Repos, newRepo)
		golog.Debugf("Cloning %s", repo.Name)
		cmd := exec.Command("git", "clone", repo.URL)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			golog.Error(err)
		}
	}
}

func getRepoNameFromUrl(url string) string {
	gitIndex := strings.Index(url, ".git")
	if gitIndex == -1 {
		return ""
	}
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		return ""
	}
	runes := []rune(url)
	return string(runes[lastSlashIndex+1 : gitIndex])
}

package model

// SproutRootFileName sets the file name that sprout generates when creating the directory structure
const SproutRootFileName string = ".sprout_root.yml"

// Config represents the layout of the sprout_config.yml file and all available flag options
type Config struct {
	Verbose  bool      `yaml:"verbose"`
	Projects []Project `yaml:"projects"`
	Repos    []Repo    `yaml:"repos"`
}

// SproutRoot represents the structure of o
type SproutRoot struct {
	Repos []RepoWithPath `yaml:"repos"`
}

// GetAllRepos retrieves the names of all repositories in the specified Config struct
func (config *Config) GetAllRepos() (repos []string) {
	options := []string{}
	for _, repo := range config.Repos {
		options = append(options, repo.Name)
	}
	projectsRepos := getProjectsRepos(config.Projects)
	for _, repo := range projectsRepos {
		options = append(options, repo.Name)
	}
	return options
}

func getProjectsRepos(projects []Project) []Repo {
	if projects == nil {
		return nil
	}
	options := []Repo{}
	for _, project := range projects {
		options = append(options, project.Repos...)
		options = append(options, getProjectsRepos(project.Projects)...)
	}
	return options
}

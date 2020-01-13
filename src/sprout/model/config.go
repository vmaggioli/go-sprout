package model

// Config represents the layout of the sprout_config.yml file and all available flag options
type Config struct {
	Verbose  bool      `yaml:"verbose"`
	Projects []Project `yaml:"projects"`
	Repos    []Repo    `yaml:"repos"`
}

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

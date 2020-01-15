package model

// Repo represents a code repository with a name and a git URL to clone from
type Repo struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// RepoWithPath is used to write to our hidden .sprout_root.yml file when generating the project layout
type RepoWithPath struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

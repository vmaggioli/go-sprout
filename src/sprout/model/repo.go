package model

// Repo represents a code repository with a name and a git URL to clone from
type Repo struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

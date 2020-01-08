package model

// Project represents a grouping of child Projects and Repos
type Project struct {
	Name     string    `yaml:"name"`
	Projects []Project `yaml:"projects"`
	Repos    []string  `yaml:"repos"`
}

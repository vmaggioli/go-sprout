package model

type Config struct {
	Verbose   bool          `yaml:"verbose"`
	Structure ProjectLayout `yaml:"structure"`
	Repos     []Repo        `yaml:"repos"`
}

type ProjectLayout struct {
	Projects []Project `yaml:"projects"`
	Repos    []Repo    `yaml:"config"`
}

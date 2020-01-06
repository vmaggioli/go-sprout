package model

type Config struct {
	Debug     bool          `yaml:"debug"`
	Structure ProjectLayout `yaml:"structure"`
	Repos     []Repo        `yaml:"repos"`
}

type ProjectLayout struct {
	Projects []Project `yaml:"projects"`
	Repos    []Repo    `yaml:"config"`
}

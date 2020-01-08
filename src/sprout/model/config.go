package model

// Config represents the layout of the sprout_config.yml file and all available flag options
type Config struct {
	Verbose  bool      `yaml:"verbose"`
	Projects []Project `yaml:"projects"`
	Repos    []string  `yaml:"repos"`
}

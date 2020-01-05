package model

// Config represents a parsed sprout_config.yml file
type Config struct {
	Verbose     bool
	MainProject Project
}

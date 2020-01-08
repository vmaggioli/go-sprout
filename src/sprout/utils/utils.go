package utils

import (
	"os"
)

const PathSeparator = string(os.PathSeparator)
const ConfigFileName = "sprout_config.yml"

// FileExists checks the path provided and returns whether or not it exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Mkdir is a helper function for the "mkdir" command, creating a new directory
func Mkdir(name string) error {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

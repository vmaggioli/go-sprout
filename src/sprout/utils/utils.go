package utils

import (
	"os"
)

const PathSeparator = string(os.PathSeparator)
const ConfigFileName = "sprout_config.yml"

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

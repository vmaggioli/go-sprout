package cli

import (
	"fmt"
	"os"

	"github.com/Fair2Dare/sprout/src/sprout/utils"
	"github.com/kataras/golog"
)

// SpreadCommand executes the provided command across all cloned repositories
func SpreadCommand() {
	currDir, _ := os.Getwd()
	if !utils.FileExists(fmt.Sprintf("%s/%s", currDir, ".sprout_root.yml")) {
		golog.Error("No sprouted project exists in current directory")
		return
	}
	if utils.FileExists(fmt.Sprintf("%s/%s", currDir, ".sprout_branch.yml")) {
		golog.Error("Please execute spread commands in project root")
		return
	}
}

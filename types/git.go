package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runGit(execstruct structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"pull", "push"}) {
		logger.StepVerboseExec(execstruct)
		if execstruct.Command == "pull" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "git pull ")
		} else if execstruct.Command == "reset" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "git pull ")
		}
	} else {
		logger.StepVerboseError(execstruct)
	}
}

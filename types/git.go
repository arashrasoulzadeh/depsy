package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"log"
	"strings"
)

func runGit(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"pull", "push"}) {
		logger.StepVerboseExec(execstruct)
		if execstruct.Command == "pull" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "git pull ",s)
		} else if execstruct.Command == "reset" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "git reset --hard ", s)
		}
	} else {
		logger.StepVerboseError(execstruct,s)
	}
}

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

func runComposer(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"install"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct,execstruct.Command)
		if execstruct.Command == "install" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "composer install ", s)
		}
	} else {
		logger.StepVerboseError(execstruct.Command, execstruct, s, supportedCommands)
	}
}

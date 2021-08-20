package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"github.com/briandowns/spinner"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runFolder(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"create", "delete", "force_delete"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct)
		if execstruct.Command == "create" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "mkdir "+execstruct.Path, s)
		} else if execstruct.Command == "create" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "rmdir "+execstruct.Path, s)
		} else if execstruct.Command == "force_delete" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "rm -rf "+execstruct.Path, s)
		}
	} else {
		logger.StepVerboseError(execstruct, s, supportedCommands)
	}
}

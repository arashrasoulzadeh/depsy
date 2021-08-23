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

func runPm2(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"restart_all"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct,execstruct.Command)
		if execstruct.Command == "restart_all" {
			utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "pm2 restart all ", s)
		}
	} else {
		logger.StepVerboseError(execstruct.Command, execstruct, s, supportedCommands)
	}
}

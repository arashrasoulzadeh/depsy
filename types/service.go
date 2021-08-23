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

func runService(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"restart", "reload", "stop", "start"}
	for _, i := range execstruct.Args {
		if utils.StringInSlice(i.Value, supportedCommands) {
			logger.StepVerboseExec(execstruct)
			//utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" nginx", s)
		} else {
			logger.StepVerboseError(execstruct.Command, execstruct, s, supportedCommands)
		}
	}


}

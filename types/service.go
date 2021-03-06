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
		if utils.StringInSlice(i.Key, supportedCommands) {
			logger.StepVerboseExec(execstruct, i.Value+" "+i.Key)
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+i.Key+" "+i.Value, s)
		} else {
			logger.StepVerboseError(i.Key, execstruct, s, supportedCommands)
		}
	}

}

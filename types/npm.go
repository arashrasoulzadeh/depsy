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

func runNpm(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"install"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct)

		if execstruct.Command == "install" {
			for _, i := range execstruct.Args {
				if i.Key == "package" {
					log.Println("\t"+ strings.ToUpper(i.Value))
					utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "npm i "+i.Value, s)
				}
			}
		}
	} else {
		logger.StepVerboseError(execstruct.Command, execstruct, s, supportedCommands)
	}
}

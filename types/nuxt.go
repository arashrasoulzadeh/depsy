package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"log"
	"strings"

	"github.com/briandowns/spinner"

	"github.com/fatih/color"
)

func runNuxt(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"build"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct, execstruct.Command)
		utils.RunCustomBashCommand(execstruct.Path, execstruct.PassOnError, "nuxt "+execstruct.Command, s)
	} else {
		logger.StepVerboseError(execstruct.Command, execstruct, s, supportedCommands)
	}
}

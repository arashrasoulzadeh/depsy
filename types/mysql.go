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

func runMysql(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))
	supportedCommands := []string{"restart", "stop", "start"}
	if utils.StringInSlice(execstruct.Command, supportedCommands) {
		logger.StepVerboseExec(execstruct)
		utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" mysql", s)
	} else {
		logger.StepVerboseError(execstruct, s, supportedCommands)
	}
}

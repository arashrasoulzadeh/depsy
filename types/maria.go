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

func runMaria(execstruct structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"restart", "stop", "start"}) {
		logger.StepVerboseExec(execstruct)
		utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" mariadb", s)
	} else {
		logger.StepVerboseError(execstruct, s)
	}
}

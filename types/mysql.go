package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runMysql(execstruct structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"restart", "stop", "start"}) {
		log.Printf("\t%s", green(strings.ToUpper(execstruct.Command)))
		utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" mysql")
	} else {
		logger.StepError()
	}
}

package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"github.com/fatih/color"
	"log"
	"strings"
)

func RunMaria(execstruct structs.ExecStruct){
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"restart", "stop", "start"}) {
		log.Printf("\t%s", green(strings.ToUpper(execstruct.Command)))
		utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" mariadb")
	} else {
		logger.StepError()
	}
}

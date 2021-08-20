package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runFolder(execstruct structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"create", "delete", "force_delete"}) {
		log.Printf("\t%s", green(strings.ToUpper(execstruct.Command)))
		if execstruct.Command == "create" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "mkdir "+execstruct.Path)
		} else if execstruct.Command == "create" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "rmdir "+execstruct.Path)
		} else if execstruct.Command == "force_delete" {
			utils.RunCustomBashCommand("/", execstruct.PassOnError, "rm -rf "+execstruct.Path)
		}
	} else {
		logger.StepError()
	}
}

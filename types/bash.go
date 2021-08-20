package types

import (
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runBash(execstep structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstep.Name)))
	utils.RunBashCommand(execstep)
}

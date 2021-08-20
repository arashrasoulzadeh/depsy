package types

import (
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"
	"github.com/briandowns/spinner"
	"log"
	"strings"

	"github.com/fatih/color"
)

func runBash(execstep structs.ExecStruct, s *spinner.Spinner) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstep.Name)))
	utils.RunBashCommand(execstep,s)
}

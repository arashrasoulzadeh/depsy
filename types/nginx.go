package types

import (
	"arashrasoulzadeh/deepzy/structs"
	. "arashrasoulzadeh/deepzy/utils"

	"log"
	"strings"

	"github.com/fatih/color"
)

func RunNginx(execstruct structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()

	if StringInSlice(execstruct.Command, []string{"restart", "reload", "stop", "start"}) {
		log.Printf("\t%s", green(strings.ToUpper("NGINX "+execstruct.Command)))
	}

}

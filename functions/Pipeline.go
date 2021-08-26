package functions

import (
	. "arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/utils"

	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/types"

	"log"
	"strings"

	"github.com/fatih/color"
)

func Execute(config structs.Config) {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for i := 0; i < len(config); i++ {
		go Hook(config[i].Hook, config[i].Name, "start", 1)
		log.Printf("STEP %s : \n", blue(strings.ToUpper(config[i].Name)))
		if config[i].Become {
			log.Printf("\t%s\t", green(strings.ToUpper("ROOT CHECK")))
			if strings.Compare(utils.GetProcessOwner(), "root") == 0 {
				StepPass()
				go Hook(config[i].Hook, "check_root", "pass", 1)
			} else {
				go Hook(config[i].Hook, "check_root", "error", 0)
				StepError("SUDO")
				StepBreak("YOU ARE NOT SUDOER")
			}
			log.Println()
		}
		for s := 0; s < len(config[i].Exec); s++ {
			Log.Printf("%s: %s ", config[i].Exec[s].Type, config[i].Exec[s].Command)
			go Hook(config[i].Hook, config[i].Exec[s].Command, "start", 1)
			types.Run(config[i].Exec[s])
		}
	}
	log.Println("")
	log.Println("DONE!")
}

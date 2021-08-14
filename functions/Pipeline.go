package functions

import (
	"arashrasoulzadeh/deepzy/structs"
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Execute(config structs.Config) {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for i := 0; i < len(config); i++ {

		log.Printf("STEP %s : \n", blue(strings.ToUpper(config[i].Name)))
		if config[i].Become {
			log.Printf("\t%s\t", green(strings.ToUpper("ROOT CHECK")))
			if strings.Compare(getProcessOwner(), "root") == 0 {
				StepPass()
			} else {
				StepError()
				StepBreak("YOU ARE NOT SUDOER")
			}
			log.Println()
		}
		for s := 0; s < len(config[i].Exec); s++ {
			Runner(config[i].Exec[s])
		}
	}
	log.Println("")
	log.Println("DONE!")
}

func RenderCommand(cmd string, args []structs.ExecArgs) string {

	for i := 0; i < len(args); i++ {
		cmd = strings.ReplaceAll(cmd, "{"+args[i].Key+"}", args[i].Value)
	}
	return cmd
}

func Runner(execstep structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()

	log.Printf("\t%s", green(strings.ToUpper(execstep.Name)))
	cmd := exec.Command("sh", "-c", RenderCommand(execstep.Command, execstep.Args))
	cmd.Dir = execstep.Path

	err := cmd.Run()

	if err != nil {
		if !execstep.PassOnError {
			StepError()
			StepBreak(err)
		} else {
			StepError()
		}
	} else {
		StepPass()
	}
}

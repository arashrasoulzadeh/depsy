package functions

import (
	"arashrasoulzadeh/deepzy/structs"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func Execute(config structs.Config) {
	blue := color.New(color.FgBlue).SprintFunc()

	for i := 0; i < len(config); i++ {
		fmt.Printf("STEP %s : \n", blue(strings.ToUpper(config[i].Name)))
		for s := 0; s < len(config[i].Exec); s++ {
			Runner(config[i].Exec[s])
		}
	}
	fmt.Println("")
	fmt.Println("DONE!")
}

func RenderCommand(cmd string, args []structs.ExecArgs) string {

	for i := 0; i < len(args); i++ {
		cmd = strings.ReplaceAll(cmd, "{"+args[i].Key+"}", args[i].Value)
	}
	return cmd
}

func Runner(execstep structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("\t%s:", green(strings.ToUpper(execstep.Name)))
	cmd := exec.Command(execstep.Command)
	cmd.Dir = execstep.Path
	fmt.Println(RenderCommand(execstep.Command, execstep.Args))

	err := cmd.Run()

	if err != nil {
		if !execstep.PassOnError {
			fmt.Println("")
			log.Fatal(err)
			fmt.Println("")
		} else {
			StepError()
		}
	} else {
		StepPass()
	}
}

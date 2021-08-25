package logger

import (
	"arashrasoulzadeh/deepzy/structs"
	"fmt"
	"github.com/briandowns/spinner"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	Log *log.Logger
)

func StepVerboseError(command string, reason structs.ExecStruct, s *spinner.Spinner, commands []string) {
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	log.Printf("\t⬆️  %s INVALID %s COMMAND %s {%s}", red("ERROR => "), blue(strings.ToUpper(reason.Type)), yellow(strings.ToUpper(command)), strings.ToUpper(strings.Join(commands, ",")))
}

func StepVerboseExec(reason structs.ExecStruct,command string) {
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	log.Printf("\t⬆  ️%s %s ", blue(strings.ToUpper(reason.Type)), yellow(strings.ToUpper(command)))
}

func StepError(reason string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Printf("\t⬆  ️%s", red("ERROR => "+reason))
}
func StepPass() {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t⬆  ️%s", green("PASS"))
}
func StepBreak(v ...interface{}) {
	defer fmt.Print("logger not initialized!")
	if v != nil {
		log.Println()
		log.Fatal(fmt.Sprintf("%s", v))
		log.Println("")
	}
}

func StepInfo(message string) {
 	if message != "" {
		log.Println(fmt.Sprintf("%s", message))
		log.Println("")
	}
}


func init() {
	// set location of log file
	var logpath = "error.log"

	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

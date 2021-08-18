package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)
var (
	Log      *log.Logger
)

func StepError() {
	red := color.New(color.FgRed).SprintFunc()
	log.Printf("\t⬆️%s", red("FAILED"))
}
func StepPass() {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t⬆️%s", green("PASS"))
}
func StepBreak(v ...interface{}) {
	log.Println()
	log.Fatal(v)
	log.Println("")
}

func init() {
	// set location of log file
	var logpath =   "error.log"

 	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

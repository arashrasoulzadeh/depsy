package functions

import (
	"log"

	"github.com/fatih/color"
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

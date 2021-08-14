package functions

import (
	"fmt"

	"github.com/fatih/color"
)

func StepError() {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("\t%s", red("FAILED"))
}
func StepPass() {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("\t%s", green("PASS"))
}

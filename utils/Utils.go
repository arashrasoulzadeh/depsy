package utils

import (
	. "arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/briandowns/spinner"
)

func GetProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		StepBreak(err)
	}
	user := ""
	if runtime.GOOS == "windows" {
		user = strings.TrimRight(string(stdout), "\r\n")
	} else {
		user = strings.TrimRight(string(stdout), "\n")
	}

	return user
}
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func RenderCommand(cmd string, args []structs.ExecArgs) string {

	for i := 0; i < len(args); i++ {
		cmd = strings.ReplaceAll(cmd, "{"+args[i].Key+"}", args[i].Value)
	}
	return cmd
}
func ExecuteSystemCall(command string, path string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = path
	return cmd.CombinedOutput()

}
func RunCustomBashCommand(path string, pass_on_error bool, command string, s *spinner.Spinner) {
	output, err := ExecuteSystemCall(command, path)

	s.Stop()
	fmt.Printf("\r")
	if err != nil {
		if !pass_on_error {
			go Hook(command, command, "error", 0)
			StepError("pass on error is false")
			Log.Printf("%s => %s", fmt.Sprint(err), output)
			StepBreak(err)
		} else {
			go Hook(command, command, "error_but_passed", 3)
			StepError("pass on error is on, passing")
		}
	} else {
		go Hook(command, command, "pass", 1)
		StepPass()
	}
}
func RunBashCommand(execstep structs.ExecStruct, s *spinner.Spinner) {
	RunCustomBashCommand(execstep.Path, execstep.PassOnError, RenderCommand(execstep.Command, execstep.Args), s)
}

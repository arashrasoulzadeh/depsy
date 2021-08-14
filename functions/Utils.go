package functions

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func getProcessOwner() string {
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

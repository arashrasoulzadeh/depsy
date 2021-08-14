package utils

import (
	. "arashrasoulzadeh/deepzy/logger"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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

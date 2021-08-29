package test

import (
	"arashrasoulzadeh/deepzy/utils"
	"testing"
)

func TestRunCustomBashCommand(t *testing.T) {
	_, err := utils.ExecuteSystemCall("bash -c ls", "")
	if err != nil {
		t.Error(err)
	}
}
func TestRunCustomInvalidBashCommand(t *testing.T) {
	_, err := utils.ExecuteSystemCall("bash -c someinvalidcommand", "")
	if err == nil {
		t.Error(err)
	}
}
func TestRunBashUnameCommand(t *testing.T) {
	_, err := utils.ExecuteSystemCall("bash -c uname -a", "")
	if err != nil {
		t.Error(err)
	}
}

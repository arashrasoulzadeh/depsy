package test

import (
	"arashrasoulzadeh/deepzy/functions"
	"testing"
)

func TestRunCustomBashCommand(t *testing.T) {
	_, err := functions.ReadFileToString("sample.yaml")
	if err != nil {
		t.Error(err)
	}
}

package test

import (
	"arashrasoulzadeh/deepzy/functions"
	"arashrasoulzadeh/deepzy/structs"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestLoadYamlFile(t *testing.T) {
	_, err := functions.ReadFileToString("sample.yaml")
	if err != nil {
		t.Error(err)
	}
}
func TestParseLoadedYamlFile(t *testing.T) {
	str, err := functions.ReadFileToString("sample.yaml")
	if err != nil {
		t.Error(err)
	}
	config := structs.Config{}
	errOnUnmarshal := yaml.Unmarshal([]byte(str), &config)
	if errOnUnmarshal != nil {
		t.Error(errOnUnmarshal)
	}
}

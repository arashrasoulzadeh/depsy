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
func TestParseLoadedYamlFileItems(t *testing.T) {
	str, err := functions.ReadFileToString("sample.yaml")
	if err != nil {
		t.Error(err)
	}
	config := structs.Config{}
	errOnUnmarshal := yaml.Unmarshal([]byte(str), &config)
	if errOnUnmarshal != nil {
		t.Error(errOnUnmarshal)
	}
	if config[0].Name!="api"{
		t.Error("invalid parser for Name" )
	}

	if config[0].Become!=true{
		t.Error("invalid parser for Become" )
	}

	if config[0].Exec[0].Name!="test"{
		t.Error("invalid parser for Exec Name" )
	}
	if config[0].Exec[0].Command!="create"{
		t.Error("invalid parser for Exec Command" )
	}
	if config[0].Exec[0].Type!="folder"{
		t.Error("invalid parser for Exec Type" )
	}
	if config[0].Exec[0].Path!="/var/www/html/test"{
		t.Error("invalid parser for Exec Path" )
	}
	if config[0].Exec[0].PassOnError!=true{
		t.Error("invalid parser for Exec PassOnError" )
	}

}

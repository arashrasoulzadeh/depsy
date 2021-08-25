package functions

import (
	"arashrasoulzadeh/deepzy/logger"
	"io/ioutil"
)

func ReadFileToString(path string) (string, error) {
	logger.StepInfo("loading config from " + path)
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
 		return "nil", err
	}
	str := string(b) // convert content to a 'string'

	return str, err
}

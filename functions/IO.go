package functions

import (
	"fmt"
	"io/ioutil"
)

func ReadFileToString(path string) (string, error) {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return "nil", err
	}
	str := string(b) // convert content to a 'string'

	return str, err
}

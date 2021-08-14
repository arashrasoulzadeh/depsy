package functions

import (
	"fmt"
	"io/ioutil"
)

func ReadFileToString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'

	return str
}

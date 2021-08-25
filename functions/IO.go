package functions

import (
	"arashrasoulzadeh/deepzy/logger"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ReadFileToString(path string) (string, error) {
	logger.StepInfo("loading config from " + path)
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		logger.StepInfo("Downloading... " + path)
		dlPath := "down_" + RandStringBytes(10) + ".yaml"
		err := DownloadFile(dlPath, path)
		if err != nil {
			return "", err
		}
		path = dlPath
		logger.StepInfo("Done Downloading => " + path)

	}
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return "nil", err
	}
	str := string(b) // convert content to a 'string'

	return str, err
}
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

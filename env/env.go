package env

import (
	"errors"
	"os"

	"github.com/rabadiw/onemark-nmh/config"

	"bufio"
	"strings"
)

var envFilename string

func init() {
	envFilename = config.GetEnvPath()
}

// GetLogLevel gets the Log_Level set in .env file
func GetLogLevel() string {
	v, _ := GetEnvValue("LOG_LEVEL")
	return v
}

// GetEnvPath gets the ENV_PATH set in .env file
func GetEnvPath() string {
	v, _ := GetEnvValue("ENV_PATH")
	return v
}

// GetEnvValue returns the value of a given key from the .env file
// with default location of the executing app
func GetEnvValue(key string) (string, error) {
	//envFilename := ".env"
	f, err := os.OpenFile(envFilename, os.O_RDONLY, 0600)
	if err != nil {
		return "", errors.New("Failed to open file " + envFilename)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	var v = ""
	for reader.Scan() {
		token := strings.Split(reader.Text(), "=")
		if strings.Compare(token[0], key) == 0 {
			v = token[1]
			break
		}
	}
	return v, nil
}

package nmh

import (
	"os"

	"bufio"
	"strings"

	"github.com/rabadiw/onemark-nmh/logger"
)

// GetEnvValue returns the value of a given key from the .env file
// with default location of the executing app
func GetEnvValue(key string) string {
	envFilename := ".env"
	f, err := os.OpenFile(envFilename, os.O_RDONLY, 0600)
	if err != nil {
		logger.LogInfo("Failed to open file " + envFilename)
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
	return v
}

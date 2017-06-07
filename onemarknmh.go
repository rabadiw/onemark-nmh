package main

import (
	"flag"
	"fmt"

	"github.com/rabadiw/onemark-nmh/config"
	"github.com/rabadiw/onemark-nmh/logger"
	"github.com/rabadiw/onemark-nmh/nmh"
)

var envPath string

func init() {
	flag.StringVar(&envPath, "env", ".env", "Environment file")
	flag.Parse()
	config.SetEnvPath(envPath)
	logger.LogInfo(fmt.Sprintf("Filepath %s", envPath))
}

func main() {
	logger.LogInfo("Application started")
	nmh.Receive()
}

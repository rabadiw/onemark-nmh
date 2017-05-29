package main

import (
	"github.com/rabadiw/onemark-nmh/logger"
	"github.com/rabadiw/onemark-nmh/nmh"
)

func main() {
	logger.LogInfo("Application started")
	nmh.Receive()
}

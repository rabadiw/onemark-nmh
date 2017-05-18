package main

import (
	"fmt"

	"github.com/rabadiw/onemark-nmh/logger"
	"github.com/rabadiw/onemark-nmh/nmh"
)

func main() {
	logger.LogInfo("Application started")
	nmh.Receive()
}

func sendMessage(msg ...interface{}) {
	_, err := fmt.Println(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
}

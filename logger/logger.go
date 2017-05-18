// Package logger contains log functions supporting log levels
package logger

import (
	"fmt"
	"os"
	"time"
)

// Logs a given msg into a file.
func LogInfo(msg string) {
	f, _ := os.OpenFile("onemarknhm_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	s := fmt.Sprintf("%v: [info] %v\n", time.Now().Format(time.RFC850), msg)
	f.WriteString(s)
}

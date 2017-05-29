// Package logger contains log functions supporting log levels
package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rabadiw/onemark-nmh/env"
)

// LogLevel type of log
type LogLevel int

const (
	// INFO log level, info as minimum level
	INFO LogLevel = 1 + iota
	// WARN log level, warn as minimum level
	WARN
	// ERROR log level, errors as minimum level
	ERROR
	// NONE log level, disables all logging
	NONE
)

var (
	logLevel LogLevel
)

var logNames = [...]string{
	"INFO",
	"WARN",
	"ERROR",
	"NONE",
}

func (l LogLevel) String() string { return logNames[l-1] }

func init() {
	level, err := env.GetEnvValue("LOG_LEVEL")
	if err != nil {
		LogError(err.Error())
		return
	}

	switch strings.ToUpper(level) {
	case INFO.String():
		logLevel = INFO
	case WARN.String():
		logLevel = WARN
	case ERROR.String():
		logLevel = ERROR
	default:
		logLevel = NONE
	}

	fmt.Println(level)
}

// LogInfo logs a given msg into a file for LOG_LEVEL of Info+
func LogInfo(msg string) {
	log(msg, INFO)
}

// LogWarn logs a given msg for LOG_LEVEL of Warn+
func LogWarn(msg string) {
	log(msg, WARN)
}

// LogError logs a given msg for LOG_LEVEL of Error+
func LogError(msg string) {
	log(msg, ERROR)
}

func log(msg string, level LogLevel) {
	f, _ := os.OpenFile("onemarknhm_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	s := fmt.Sprintf("%v: [%v] %v\n", time.Now().Format(time.RFC850), strings.ToLower(level.String()), msg)
	f.WriteString(s)
}

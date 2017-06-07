// Package logger contains log functions supporting log levels
package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rabadiw/onemark-nmh/config"
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
	level := config.GetLogLevel()

	switch strings.ToUpper(level) {
	case WARN.String():
		logLevel = WARN
	case ERROR.String():
		logLevel = ERROR
	case NONE.String():
		logLevel = NONE
	default:
		logLevel = INFO
	}
}

// LogInfo logs a given msg into a file for LOG_LEVEL of Info+
func LogInfo(msg string) {
	if logLevel != INFO || logLevel == NONE {
		return
	}
	log(msg, INFO)
}

// LogWarn logs a given msg for LOG_LEVEL of Warn+
func LogWarn(msg string) {
	if logLevel != WARN || logLevel == NONE {
		return
	}
	log(msg, WARN)
}

// LogError logs a given msg for LOG_LEVEL of Error+
func LogError(msg string) {
	if logLevel != ERROR || logLevel == NONE {
		return
	}
	log(msg, ERROR)
}

func log(msg string, level LogLevel) {
	f, _ := os.OpenFile("onemarknhm_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	s := fmt.Sprintf("%v: [%v] %v\n", time.Now().Format(time.RFC850), strings.ToLower(level.String()), msg)
	f.WriteString(s)
}

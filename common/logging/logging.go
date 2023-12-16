package logging

import (
	"fmt"
	"log"
	"os"
)

type logLevel string

const (
	DEBUG logLevel = "DEBUG"
	INFO  logLevel = "INFO"
	WARN  logLevel = "WARN"
	ERROR logLevel = "ERROR"
)

var loggerDebug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
var loggerInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var loggerWarn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
var loggerError = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func LogInfo(message string) {
	loggerInfo.Println(message)
}
func logMessage(loggerName string, prefix logLevel, message string) {
	switch prefix {
	case DEBUG:
		loggerDebug.Println(fmt.Sprintf("%v - %v", loggerName, message))
	case INFO:
		loggerInfo.Println(fmt.Sprintf("%v - %v", loggerName, message))
	case WARN:
		loggerWarn.Println(fmt.Sprintf("%v - %v", loggerName, message))
	case ERROR:
		loggerError.Println(fmt.Sprintf("%v - %v", loggerName, message))
	}
}

type customLogger struct {
	loggerName string
}

func NewLogger(loggerName string) *customLogger {
	return &customLogger{loggerName}
}

func (logger *customLogger) Debug(message string) {
	logMessage(logger.loggerName, DEBUG, message)
}
func (logger *customLogger) Info(message string) {
	logMessage(logger.loggerName, INFO, message)
}
func (logger *customLogger) Warn(message string) {
	logMessage(logger.loggerName, WARN, message)
}
func (logger *customLogger) Error(message string) {
	logMessage(logger.loggerName, ERROR, message)
}

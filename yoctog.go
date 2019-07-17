package yoctog

import (
	"fmt"
	"log"
)

// Level is the level of logging
type Level int8

// Log levels
const (
	LevelDebug Level = 1 + iota
	LevelInfo
	LevelWarn
	LevelError
	LevelAllSuppress
)

var logLevel = LevelDebug

// Debug are output debug level message
func Debug(v ...interface{}) {
	if logLevel <= LevelDebug {
		logger(append([]interface{}{"[DEBUG]"}, v...)...)
	}
}

// Info are output information level message
func Info(v ...interface{}) {
	if logLevel <= LevelInfo {
		logger(append([]interface{}{"[INFO]"}, v...)...)
	}
}

// Warn are output warning level message
func Warn(v ...interface{}) {
	if logLevel <= LevelWarn {
		logger(append([]interface{}{"[WARN]"}, v...)...)
	}
}

// Error are output error level message
func Error(v ...interface{}) {
	if logLevel <= LevelError {
		logger(append([]interface{}{"[ERROR]"}, v...)...)
	}
}

// logger are Output wrapper
func logger(v ...interface{}) {
	_ = log.Output(3, fmt.Sprintln(v...))
}

// SetLogLevel is set the log output level
func SetLogLevel(level Level) {
	logLevel = level
}

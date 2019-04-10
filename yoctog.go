package yoctog

import (
	"fmt"
	"log"
)

// logger are Output wrapper
func logger(v ...interface{}){
	_ = log.Output(3, fmt.Sprintln(v...))
}

// Debug are output debug level message
func Debug(v ...interface{}){
	logger(append([]interface{}{"[DEBUG]"}, v...)...)
}

// Info are output information level message
func Info(v ...interface{}){
	logger(append([]interface{}{"[INFO]"}, v...)...)
}

// Warn are output warning level message
func Warn(v ...interface{}){
	logger(append([]interface{}{"[WARN]"}, v...)...)
}

// Error are output error level message
func Error(v ...interface{}){
	logger(append([]interface{}{"[ERROR]"}, v...)...)
}

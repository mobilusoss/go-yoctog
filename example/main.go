package main

import (
	"fmt"
	"github.com/mobilusoss/go-yoctog"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	fmt.Println("Set level to debug")
	yoctog.SetLogLevel(yoctog.LevelDebug)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")

	fmt.Println("Set level to info")
	yoctog.SetLogLevel(yoctog.LevelInfo)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")

	fmt.Println("Set level to warn")
	yoctog.SetLogLevel(yoctog.LevelWarn)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")

	fmt.Println("Set level to error")
	yoctog.SetLogLevel(yoctog.LevelError)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")

	fmt.Println("Set level to all suppress")
	yoctog.SetLogLevel(yoctog.LevelAllSuppress)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")
}

package main

import (
	"github.com/mobilusoss/go-yoctog"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	yoctog.Debug("This debug log")
	yoctog.Info("This information log")
	yoctog.Warn("This warning log")
	yoctog.Error("This error log")
}

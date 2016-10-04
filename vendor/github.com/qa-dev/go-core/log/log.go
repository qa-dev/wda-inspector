package log

import (
	"github.com/qa-dev/go-core/color"
	"log"
)

func Info(format string, a ...interface{}) {
	format = color.Color(color.Blue, "INFO: ")+format
	log.Printf(format, a...)
}

func Fatal(format string, a ...interface{}) {
	format = color.Color(color.Red, "FATAL: ")+format
	log.Fatalf(format, a...)
}

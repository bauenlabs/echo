// Package log provides a standard set of logging tools.
package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	TraceLog = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLog = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WarningLog = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLog = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Trace prints a trace log.
func Trace(logString string) {
	TraceLog.Println(logString)
}

// Info prints a info log.
func Info(logString string) {
	InfoLog.Println(logString)
}

// Warning prints a warning log.
func Warning(logString string) {
	WarningLog.Println(logString)
}

// Error prints a error log.
func Error(logString string) {
	ErrorLog.Println(logString)
}
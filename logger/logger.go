package logger

import (
	"io"
	"log"
)

var (
	// Trace requires a comment to not trigger the linter!
	Trace *log.Logger
	// Info shut up linter
	Info *log.Logger
	// Warning shut up linter
	Warning *log.Logger
	// Error shut up linter
	Error *log.Logger
)

// Init logger handlers
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

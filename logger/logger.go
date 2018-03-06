package logger

import (
	"io"
	"log"
)

var (
	// Debug requires a comment to not trigger the linter!
	Debug *log.Logger
	// Info shut up linter
	Info *log.Logger
	// Warning shut up linter
	Warning *log.Logger
	// Error shut up linter
	Error *log.Logger
)

// Init logger handlers
func Init(
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Debug = log.New(debugHandle,
		"DEBUG: ",
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

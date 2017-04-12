package command

// UI is an interface for interacting with the terminal
type UI interface {
	// Ask asks the user for input using the given query. The response is
	// returned as the given string, or an error.
	Ask(string) (string, error)

	// Output is called for normal standard output.
	Output(string)

	// Info is called for information related to the previous output.
	// In general this may be the exact same as Output, but this gives
	// Ui implementors some flexibility with output formats.
	Info(string)

	// Error is used for any error messages that might appear on standard
	// error.
	Error(string)
}
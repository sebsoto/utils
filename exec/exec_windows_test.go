// +build windows

package exec

var shell = "powershell"

// outputWriterCommand returns a command and args which will write the given string to the output stream
func outputWriterCommand(input string) []string {
	return []string{"cmd", "/Q", "/C", "echo", input}
}

// errorWriterCommand returns a command and args which will write the given string to the error stream
func errorWriterCommand(input string) []string {
	return []string{"cmd", "/Q", "/C", "echo", input, "1>&2"}
}

// errorAndOutputWriterCommand returns a command and args which will write the given strings to the output and error
// streams
func errorAndOutputWriterCommand(output, error string) []string {
	return []string{"cmd", "/Q", "/C", "echo", output, "&&", "echo", error, "1>&2"}
}

// sleepCommand returns a command and args which will sleep for the time specified. parameter should be in numeric form
// "1", "2", etc
func sleepCommand(seconds string) []string {
	return []string{"powershell", "Start-Sleep", "-s", seconds}
}

// printEnvVarCommand retunrs a command and args which writes the contents of the given env variable to the output stream
func printEnvVarCommand(envVar string) []string {
	return []string{"cmd", "/Q", "/C", "echo", "%" + envVar + "%"}
}

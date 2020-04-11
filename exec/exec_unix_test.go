// +build linux darwin

package exec

var shell = "sh"

// outputWriterCommand returns a command and args which will write the given string to the output stream
func outputWriterCommand(output string) []string {
	return []string{"echo", output}
}

// errorWriterCommand returns a command and args which will write the given string to the error stream
func errorWriterCommand(error string) []string {
	return []string{"/bin/sh", "-c", "echo " + error + " > /dev/stderr"}
}

// errorAndOutputWriterCommand returns a command and args which will write the given strings to the output and error
// streams
func errorAndOutputWriterCommand(output, error string) []string {
	return []string{"/bin/sh", "-c", "echo '" + output + "'>&1; echo '" + error + "'>&2"}
}

// sleepCommand returns a command and args which will sleep for the time specified. parameter should be in numeric form
// "1", "2", etc
func sleepCommand(seconds string) []string {
	return []string{"sleep", seconds}
}

// printEnvVarCommand retunrs a command and args which writes the contents of the given env variable to the output stream
func printEnvVarCommand(envVar string) []string {
	return []string{"/bin/sh", "-c", "echo $" + envVar}
}

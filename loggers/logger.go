package logger

import (
	"fmt"

	"github.com/fatih/color"
)

// LogRedError print red error detail
func LogRedError(err error) {
	redOutput := color.New(color.FgRed)
	errorOutput := redOutput.Add(color.Bold)
	errorOutput.Println(fmt.Errorf("%s", err))
}

// LogRedMessage print red string message
func LogRedMessage(format string, a ...interface{}) {
	redOutput := color.New(color.FgRed)
	errorOutput := redOutput.Add(color.Bold)
	errorOutput.Println(fmt.Sprintf(format, a...))
}

// LogGreenMessage print green string message
func LogGreenMessage(format string, a ...interface{}) {
	greenOutput := color.New(color.FgGreen)
	errorOutput := greenOutput.Add(color.Bold)
	errorOutput.Println(fmt.Sprintf(format, a...))
}

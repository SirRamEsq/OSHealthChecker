package logger

import (
	"fmt"
	"healthChecker/cmd"
	"math"
	"strconv"

	"github.com/fatih/color"
)

type Logger struct {
	fails  int
	passes int
}

// Get total number of Fails and Passes
func (l *Logger) ChecksCount() int {
	return l.fails + l.passes
}
func (l *Logger) Fail(message string) {
	color.Set(color.FgRed, color.Bold)
	cmd.PrintlnLeftRight(message, "FAIL")
	color.Unset()
	l.fails += 1
}

func (l *Logger) Pass(message string) {
	color.Set(color.FgGreen, color.Bold)
	cmd.PrintlnLeftRight(message, "PASS")
	color.Unset()
	l.passes += 1
}

func (l *Logger) Error(message string) {
	color.Set(color.BgRed, color.FgWhite, color.Bold)
	cmd.PrintlnRightAlign(message)
	color.Unset()
	l.passes += 1
}

func (l *Logger) PrintStatus() {
	failString := "FAILED: " + strconv.Itoa(l.fails)
	passString := "PASSED: " + strconv.Itoa(l.passes)
	color.Set(color.FgGreen, color.Bold)
	cmd.PrintlnLeftAlign(passString)

	color.Set(color.FgRed, color.Bold)
	cmd.PrintlnLeftAlign(failString)
	color.Unset()

	termWidth := cmd.GetTerminalWidth()
	percent := float64(float64(l.passes) / float64(l.ChecksCount()))
	passedCharacters := int(math.Floor(float64(termWidth) * percent))
	failedCharacters := termWidth - passedCharacters

	passedString := ""
	failedString := ""
	for i := 0; i < passedCharacters; i++ {
		passedString += "="
	}
	for i := 0; i < failedCharacters; i++ {
		failedString += "="
	}
	color.Set(color.FgGreen, color.Bold)
	fmt.Print(passedString)
	color.Set(color.FgRed, color.Bold)
	fmt.Print(failedString)
	color.Unset()
	fmt.Println()
}

func New() Logger {
	return Logger{}
}

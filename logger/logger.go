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
}

func (l *Logger) PrintStatus() {
	color.Set(color.FgWhite, color.Bold)
	finalStr := "Total Checks: " + strconv.Itoa(l.ChecksCount())
	fmt.Println(finalStr)
	color.Unset()

	failString := "FAILED: " + strconv.Itoa(l.fails)
	passString := "PASSED: " + strconv.Itoa(l.passes)
	color.Set(color.FgGreen, color.Bold)
	cmd.PrintlnLeftAlign(passString)

	color.Set(color.FgRed, color.Bold)
	cmd.PrintlnLeftAlign(failString)
	color.Unset()

	if l.ChecksCount() == 0 {
		cmd.PrintlnCenter("NO CHECKS!")
		return
	}

	termWidth := cmd.GetTerminalWidth()
	percent := float64(float64(l.passes) / float64(l.ChecksCount()))
	passedCharacters := int(math.Floor(float64(termWidth) * percent))
	failedCharacters := termWidth - passedCharacters

	percentInt := int(percent * 100)
	percentString := strconv.Itoa(percentInt)

	if percentInt == 100 {
		color.Set(color.FgGreen, color.Bold)
	} else if percentInt > 60 {
		color.Set(color.FgYellow, color.Bold)

	} else {
		color.Set(color.FgRed, color.Bold)
	}
	cmd.PrintlnCenter(percentString + "%")
	color.Unset()

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

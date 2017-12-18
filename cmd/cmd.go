package cmd

import (
	"fmt"
	"lengine/logger"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode/utf8"
)

var MAX_WIDTH int

func init() {
	MAX_WIDTH = 104
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func GetTerminalDimensions() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		logger.Critical("Cannot get size of terminal!")
	}
	/*
		out contains the width and height of the terminal in the format
		"hhh www"
	*/

	//need to separate string at white space
	fields := strings.Fields(string(out))
	width, err := StrToInt(fields[1])
	height, err := StrToInt(fields[0])

	return width, height
}
func GetTerminalWidth() int {
	w, _ := GetTerminalDimensions()
	if w > MAX_WIDTH {
		return MAX_WIDTH
	}
	return w
}
func GetTerminalHeight() int {
	_, h := GetTerminalDimensions()
	return h
}
func Separator(width int) string {
	separator := ""
	for i := 0; i < width; i++ {
		separator += "─"
	}
	return separator
}

func PrintlnLeftAlign(str string) {
	fmt.Println(str)
}
func PrintlnRightAlign(str string) {
	tWidth := GetTerminalWidth()
	subStrings := strings.Split(str, "\n")
	for i := 0; i != len(subStrings); i++ {
		sub := subStrings[i]
		strLength := utf8.RuneCountInString(sub)
		if strLength <= tWidth {
			bufferWidth := tWidth - strLength
			finalString := ""
			for i := 0; i < bufferWidth; i++ {
				finalString += " "
			}
			finalString += sub
			fmt.Println(finalString)
		} else {
			PrintlnLeftAlign(sub)
		}
	}
}
func PrintlnCenter(str string) {
	tWidth := GetTerminalWidth()
	strLength := utf8.RuneCountInString(str)
	if strLength <= tWidth {
		bufferWidth := (tWidth / 2) - (strLength / 2)
		finalString := ""
		for i := 0; i < bufferWidth; i++ {
			finalString += " "
		}
		finalString += str
		fmt.Println(finalString)
	} else {
		PrintlnLeftAlign(str)
	}
}
func PrintlnLeftRight(left string, right string) {
	totalWidth := utf8.RuneCountInString(left) + utf8.RuneCountInString(right)
	terminalWidth := GetTerminalWidth()
	if totalWidth <= terminalWidth {
		separatorWidth := terminalWidth - totalWidth
		finalString := left
		for i := 0; i < separatorWidth; i++ {
			finalString += " "
		}
		finalString += right
		fmt.Println(finalString)
	} else {
		PrintlnLeftAlign(left)
		PrintlnRightAlign(right)
	}
}

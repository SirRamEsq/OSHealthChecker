package main

import (
	"fmt"
	"healthChecker/cmd"
	"healthChecker/jsonParser"
	"healthChecker/logger"
	"io/ioutil"

	"github.com/fatih/color"
)

func PrintBanner() {
	tWidth := cmd.GetTerminalWidth()
	separator := cmd.Separator(tWidth)
	color.Set(color.FgWhite, color.Bold)
	fmt.Println(separator)
	cmd.PrintlnCenter("Health checker!")
	cmd.PrintlnLeftRight("--~--~", "~--~--")
	fmt.Println(separator)
	fmt.Println()
	color.Unset()
}

func main() {
	PrintBanner()
	tWidth := cmd.GetTerminalWidth()
	separator := cmd.Separator(tWidth)
	log := logger.New()

	configFileName := "test.json"
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Error("Error opening config file" + err.Error())
	}

	executables, err := jsonParser.ParseExecutables(configFile)
	if err != nil {
		log.Error("Error Parsing config file" + err.Error())
	}

	for i := 0; i < len(executables); i++ {
		exe := executables[i]
		passed, err := exe.Execute()
		if !passed {
			log.Fail(exe.String(passed))
		} else {
			log.Pass(exe.String(passed))
		}

		if err != nil {
			log.Error(err.Error())
		}

	}
	fmt.Println(separator)

	log.PrintStatus()
}

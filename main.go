package main

import (
	"fmt"
	"healthChecker/cmd"
	"healthChecker/jsonParser"
	"healthChecker/logger"
	"io/ioutil"

	"github.com/fatih/color"
)

func main() {
	tWidth := cmd.GetTerminalWidth()
	separator := cmd.Separator(tWidth)
	color.Set(color.FgWhite, color.Bold)
	fmt.Println(separator)
	cmd.PrintlnCenter("Health checker!")
	cmd.PrintlnLeftRight("--~--~", "~--~--")
	fmt.Println(separator)
	fmt.Println()
	log := logger.New()

	configFileName := "test.json"
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Error("opening config file" + err.Error())
	}

	executables, err := jsonParser.ParseExecutables(configFile)
	if err != nil {
		log.Error("parsing config file" + err.Error())
	}

	for i := 0; i < len(executables); i++ {
		exe := executables[i]
		result, err := exe.Execute()
		if err != nil {
			log.Error(err.Error())
			continue
		}

		if !result {
			log.Fail(exe.String())
		} else {
			log.Pass(exe.String())
		}
	}
	fmt.Println(separator)

	log.PrintStatus()
}

package main

import (
	"encoding/json"
	"fmt"
	"healthChecker/cmd"
	"healthChecker/fileExists"
	"healthChecker/logger"
	"os"
	"strconv"

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

	configFile, err := os.Open("test.json")
	if err != nil {
		log.Error("opening config file" + err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	newFE := []fileExists.FileExists{}
	if err = jsonParser.Decode(&newFE); err != nil {
		log.Error("parsing config file" + err.Error())
	}

	for i := 0; i < len(newFE); i++ {
		fileExists := newFE[i]
		result := fileExists.Execute()

		if !result {
			log.Fail(fileExists.String())
		} else {
			log.Pass(fileExists.String())
		}
	}
	fmt.Println(separator)
	color.Set(color.FgWhite, color.Bold)
	finalStr := "Total Checks: " + strconv.Itoa(log.ChecksCount())
	fmt.Println(finalStr)
	color.Unset()

	log.PrintStatus()
}

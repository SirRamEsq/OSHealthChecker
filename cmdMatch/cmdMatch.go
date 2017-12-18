package cmdMatch

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type CmdMatch struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
	//Can use either output or file output
	//FileOutputPath overrides Output
	Output         string `json:"output"`
	FileOutputPath string `json:"fileOutput"`

	Invert bool `json:"invert"`

	ExpectedOutput string
	ActualOutput   string
}

func (run *CmdMatch) checkCmdAgainstFile(cmdOutput *bytes.Buffer) (bool, error) {
	//If output file cannot be found
	if _, err := os.Stat(run.FileOutputPath); err != nil {
		return false, err
	}

	outputFile, err := ioutil.ReadFile(run.FileOutputPath)
	if err != nil {
		return false, err
	}

	run.ExpectedOutput = string(outputFile)
	run.ActualOutput = cmdOutput.String()

	return bytes.Equal(outputFile, cmdOutput.Bytes()), nil
}

func (run *CmdMatch) checkCmdAgainstOutput(cmdOutput *bytes.Buffer) (bool, error) {
	run.ActualOutput = cmdOutput.String()
	run.ExpectedOutput = run.Output
	return (cmdOutput.String() == run.Output), nil
}

// Check if file exists
func (run *CmdMatch) Execute() (bool, error) {
	returnValue := false

	//Run command
	cmd := exec.Command(run.Cmd)
	cmd.Args = run.Args
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	useFileOutput := (run.FileOutputPath != "")
	if useFileOutput {
		returnValue, err = run.checkCmdAgainstFile(&out)
	} else {
		returnValue, err = run.checkCmdAgainstOutput(&out)
	}

	if run.Invert {
		returnValue = !returnValue
	}

	return returnValue, err
}

func (run CmdMatch) String(passed bool) string {
	str := "Command \"" + run.Cmd + "\""
	if run.Invert {
		str = "(NOT) " + str
	}
	if !passed {
		//indent
		str += "\nExpected Output:\n    " + strings.Replace(run.ExpectedOutput, "\n", "\n    ", -1)
		str += "\nActual Output:\n    " + strings.Replace(run.ActualOutput, "\n", "\n    ", -1)
	}
	return str
}

func New(cmd string, file string, output string, invert bool) CmdMatch {
	return CmdMatch{Cmd: cmd, FileOutputPath: file, Output: output, Invert: invert}
}

package cmdMatch

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func RemoveWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

type CmdMatch struct {
	Cmd         string   `json:"cmd"`
	Args        []string `json:"args"`
	Environment []string `json:"env"`
	//Can use either output or file output
	//FileOutputPath overrides Output
	Output         string `json:"output"`
	FileOutputPath string `json:"fileOutput"`

	Invert     bool `json:"invert"`
	Whitespace bool `json:"whitespace"`

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

	run.ActualOutput = string(norm.NFC.Bytes([]byte(run.ActualOutput)))
	run.ExpectedOutput = string(norm.NFC.Bytes([]byte(run.ExpectedOutput)))

	if !run.Whitespace {
		run.ActualOutput = RemoveWhitespace(run.ActualOutput)
		run.ExpectedOutput = RemoveWhitespace(run.ExpectedOutput)
	}

	return (run.ActualOutput == run.ExpectedOutput), nil
}

func (run *CmdMatch) checkCmdAgainstOutput(cmdOutput *bytes.Buffer) (bool, error) {
	run.ActualOutput = cmdOutput.String()
	run.ExpectedOutput = run.Output
	run.ActualOutput = string(norm.NFC.Bytes([]byte(run.ActualOutput)))
	run.ExpectedOutput = string(norm.NFC.Bytes([]byte(run.ExpectedOutput)))

	if !run.Whitespace {
		run.ActualOutput = RemoveWhitespace(run.ActualOutput)
		run.ExpectedOutput = RemoveWhitespace(run.ExpectedOutput)
	}

	return (run.ActualOutput == run.ExpectedOutput), nil
}

// Check if Command matches expected output
func (run *CmdMatch) Execute() (bool, error) {
	returnValue := false

	//Run command
	cmd := exec.Command(run.Cmd)
	for i := 0; i != len(run.Args); i++ {
		cmd.Args = append(cmd.Args, run.Args[0])
	}
	cmd.Env = os.Environ()
	for i := 0; i != len(run.Environment); i++ {
		cmd.Env = append(cmd.Env, run.Environment[i])
	}
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
	//only show verbose output if not inverted
	if (!passed) && (!run.Invert) {
		//indent
		str += "\nExpected Output:\n    " + strings.Replace(run.ExpectedOutput, "\n", "\n    ", -1)
		str += "\nActual Output:\n    " + strings.Replace(run.ActualOutput, "\n", "\n    ", -1)
	}
	return str
}

func New(cmd string, file string, output string, invert bool) CmdMatch {
	return CmdMatch{Cmd: cmd, FileOutputPath: file, Output: output, Invert: invert}
}

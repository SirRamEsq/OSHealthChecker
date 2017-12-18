package jsonParser

import (
	"encoding/json"
	"errors"
	"healthChecker/cmdExist"
	"healthChecker/cmdMatch"
	"healthChecker/fileExists"
)

type Executable interface {
	Execute() (bool, error)
	String(bool) string
}

func ParseExecutables(data []byte) ([]Executable, error) {
	type Check struct {
		CheckType string          `json:"check"`
		Args      json.RawMessage `json:"args"`
	}

	var checks []Check
	err := json.Unmarshal(data, &checks)
	if err != nil {
		return nil, err
	}

	var executables []Executable
	aggregatedError := errors.New("")

	for _, c := range checks {
		var exe Executable
		var err error
		switch c.CheckType {
		case "cmdMatch":
			cmd := new(cmdMatch.CmdMatch)
			err = json.Unmarshal(c.Args, &cmd)
			exe = cmd
		case "cmdExist":
			cmd := new(cmdExist.CmdExist)
			err = json.Unmarshal(c.Args, &cmd)
			exe = cmd
		case "fileExists":
			fe := new(fileExists.FileExists)
			err = json.Unmarshal(c.Args, &fe)
			exe = fe
		default:
			err = errors.New("\"" + c.CheckType + "\" is not a valid check")
		}
		if err != nil {
			aggregatedError = errors.New(aggregatedError.Error() + "\n" + err.Error())
		} else {
			executables = append(executables, exe)
		}
	}

	if aggregatedError.Error() == "" {
		return executables, nil
	}
	return executables, aggregatedError
}

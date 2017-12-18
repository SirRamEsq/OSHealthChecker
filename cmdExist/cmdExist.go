package cmdExist

import (
	"os/exec"
)

type CmdExist struct {
	Cmd    string `json:"cmd"`
	Invert bool   `json:"invert"`
}

// Check if command exists
func (exist *CmdExist) Execute() (bool, error) {

	//Run command
	cmd := exec.Command(exist.Cmd)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		if !exist.Invert {
			return false, err
		}
	}

	if exist.Invert {
		if err == nil {
			return false, err
		}
	}

	return true, nil
}

func (exist CmdExist) String(passed bool) string {
	str := "Command \"" + exist.Cmd + "\"" + " Exists"
	if exist.Invert {
		str = "(NOT) " + str
	}
	return str
}

func New(cmd string, invert bool) CmdExist {
	return CmdExist{Cmd: cmd, Invert: invert}
}

package fileExists

import (
	"os"
)

type FileExists struct {
	FilePath string `json:"path"`
	Invert   bool   `json:"invert"`
}

// Check if file exists
func (fe FileExists) Execute() (bool, error) {
	returnValue := false
	var err error
	if _, err = os.Stat(fe.FilePath); err == nil {
		returnValue = true
	}
	if fe.Invert {
		returnValue = !returnValue
	}
	return returnValue, err
}

func (fe FileExists) String(passed bool) string {
	str := "File Exists at \"" + fe.FilePath + "\""
	if fe.Invert {
		str = "(NOT) " + str
	}
	return str
}

func New(path string, invert bool) FileExists {
	return FileExists{FilePath: path, Invert: invert}
}

package fileExists

import (
	"os"
)

type FileExists struct {
	FilePath string `json:"path"`
	Invert   bool   `json:"invert"`
}

// Check if file exists
func (fe FileExists) Execute() bool {
	returnValue := false
	if _, err := os.Stat(fe.FilePath); err == nil {
		returnValue = true
	}
	if fe.Invert {
		returnValue = !returnValue
	}
	return returnValue
}

func (fe FileExists) String() string {
	str := "File Exists at \"" + fe.FilePath + "\""
	if fe.Invert {
		str = "(NOT) " + str
	}
	return str
}

func New(path string, invert bool) FileExists {
	return FileExists{FilePath: path, Invert: invert}
}

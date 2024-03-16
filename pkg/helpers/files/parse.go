package files

import (
	"os"

	. "github.com/Kapparina/Dotty/pkg/errors"
)

func GetContents(filePath string) ([]byte, error) {
	contents, readErr := os.ReadFile(filePath)
	if readErr != nil {
		return nil, NewFileError("reading file contents", filePath, Read, readErr.Error())
	}
	return contents, nil
}

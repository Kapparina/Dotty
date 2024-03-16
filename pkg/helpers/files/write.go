package files

import (
	"os"

	. "github.com/Kapparina/Dotty/pkg/errors"
)

func WriteToFile(destination string, data []byte) error {
	writeErr := os.WriteFile(destination, data, 0644)
	if writeErr != nil {
		return NewFileError("writing to file", destination, Write, writeErr.Error())
	}
	return nil
}

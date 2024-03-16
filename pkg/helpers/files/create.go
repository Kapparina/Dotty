package files

import (
	"errors"
	"fmt"
	"os"

	derr "github.com/Kapparina/Dotty/pkg/errors"
)

func CreateFile(filePath string) (retErr error) {
	if err := ValidatePath(filePath); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()

	return
}

func CreateDir(dirPath string) error {
	err := ValidatePath(dirPath)
	if err != nil {
		if !errors.As(err, &derr.NoExistError{}) {
			return err
		}
		// If Path doesn't exist, try to create the directory
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			// Added more specific error message
			return fmt.Errorf("error while creating directory: %w", err)
		}
	}
	return nil
}

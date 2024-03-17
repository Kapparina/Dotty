package files

import (
	"errors"
	"os"
	"path/filepath"

	derr "github.com/Kapparina/Dotty/pkg/errors"
)

const (
	noAccess        = "the path provided exists, but isn't accessible to you"
	noExist         = "the path doesn't exist; the root is accessible however"
	partialAccess   = "can't access the path, but access to the root appears permitted"
	invalidPath     = "the entire path is invalid"
	partialAccessOp = "Accessing file along filepath"
	fullAccessOp    = "Accessing any part of filepath"
)

func ValidatePath(path string) error {
	accessible, accessErr := pathAccessible(path)
	if _, statErr := os.Stat(path); statErr != nil || !accessible {
		if errors.Is(statErr, os.ErrNotExist) {
			if !filePathRootAccessible(path) {
				return derr.NewFileError(fullAccessOp, path, derr.InvalidPath, invalidPath)
			} else {
				return derr.NewFileError(partialAccessOp, path, derr.NoExist, noExist)
			}
		} else if errors.Is(statErr, os.ErrPermission) || errors.Is(accessErr, os.ErrPermission) {
			if !filePathRootAccessible(path) {
				return derr.NewFileError(fullAccessOp, path, derr.NoAccess, noAccess)
			} else {
				return derr.NewFileError(partialAccessOp, path, derr.PartialAccess, partialAccess)
			}
		}
	}
	return nil
}

func filePathRootAccessible(filePath string) bool {
	if _, err := os.Stat(filepath.VolumeName(filePath)); err != nil {
		return false
	} else {
		return true
	}
}

func pathAccessible(path string) (bool, error) {
	handle, err := os.Open(path)
	defer func(handle *os.File) {
		err := handle.Close()
		if err != nil {
			return
		}
	}(handle)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

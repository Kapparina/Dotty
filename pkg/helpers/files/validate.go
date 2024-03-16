package files

import (
	"errors"
	"os"
	"path/filepath"

	derr "github.com/Kapparina/Dotty/pkg/errors"
)

const (
	noAccess        = "the path provided exists, but isn't accessible to you"
	noExist         = "the path doesn't exist; the root is accessible though (could use somewhere along `%s`)"
	partialAccess   = "can't access the path, but access to root (`%s`) appears permitted"
	invalidPath     = "the entire path (incl. `%s`) is invalid"
	partialAccessOp = "Accessing file along filepath"
	fullAccessOp    = "Accessing any part of filepath"
)

func ValidatePath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if !filePathRootAccessible(path) {
				return derr.NewFileError(fullAccessOp, path, derr.InvalidPath, invalidPath)
			} else {
				return derr.NewFileError(partialAccessOp, path, derr.NoExist, noExist)
			}
		} else if errors.Is(err, os.ErrPermission) {
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

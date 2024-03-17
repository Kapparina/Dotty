package files

import (
	"path/filepath"
	"runtime"
	"testing"

	derr "github.com/Kapparina/Dotty/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidatePath(t *testing.T) {
	assertion := assert.New(t)

	completelyInvalidPath := "/path/does/not/exist"
	nonExistingRelativePath, _ := filepath.Abs("./path/does/not/exist")
	determinePaths := func() (accessDeniedPath string, validPath string) {
		if runtime.GOOS == "windows" {
			return "C:\\Windows\\ServiceState", "C:\\Windows"
		} else {
			return "/etc/sudoers", "/etc"
		}
	}
	deniedPath, validPath := determinePaths()

	tests := []struct {
		name     string
		path     string
		expected error
	}{
		{
			name: "InvalidPath",
			path: completelyInvalidPath,
			expected: derr.NewFileError(
				fullAccessOp,
				completelyInvalidPath,
				derr.InvalidPath,
				invalidPath,
			),
		},
		{
			name: "DoesNotExist",
			path: nonExistingRelativePath,
			expected: derr.NewFileError(
				partialAccessOp,
				nonExistingRelativePath,
				derr.NoExist,
				noExist,
			),
		},
		{
			name: "PermissionDenied",
			path: deniedPath,
			expected: derr.NewFileError(
				partialAccessOp,
				deniedPath,
				derr.PartialAccess,
				partialAccess,
			),
		},
		{
			name:     "PathValid",
			path:     validPath,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePath(tt.path)
			assertion.IsType(tt.expected, err)
			if err != nil {
				assertion.Equal(tt.expected.Error(), err.Error())
			}
		})
	}
}

package errors

import (
	"errors"
	"fmt"
	"path/filepath"
)

type FileOperationError interface {
	Error() string
	GetOperation() string
	GetPath() string
}

type FileError struct {
	Operation string
	Path      string
	Err       error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s: %s: %v", e.Operation, e.Path, e.Err)
}

func (e *FileError) GetOperation() string {
	return e.Operation
}

func (e *FileError) GetPath() string {
	return filepath.Clean(e.Path)
}

type PartialAccessError struct {
	*FileError
}

type NoAccessError struct {
	*FileError
}

type NoExistError struct {
	*FileError
}

type InvalidPathError struct {
	*FileError
}

type ReadError struct {
	*FileError
}

type WriteError struct {
	*FileError
}

type ErrorType int

type ErrorCreator func(*FileError) FileOperationError

const (
	PartialAccess = iota
	NoAccess
	NoExist
	InvalidPath
	Read
	Write
)

var errorCreators = map[ErrorType]ErrorCreator{
	PartialAccess: func(e *FileError) FileOperationError { return &PartialAccessError{e} },
	NoAccess:      func(e *FileError) FileOperationError { return &NoAccessError{e} },
	NoExist:       func(e *FileError) FileOperationError { return &NoExistError{e} },
	InvalidPath:   func(e *FileError) FileOperationError { return &InvalidPathError{e} },
	Read:          func(e *FileError) FileOperationError { return &ReadError{e} },
	Write:         func(e *FileError) FileOperationError { return &WriteError{e} },
}

func NewFileError(op string, path string, errType ErrorType, msg string) error {
	creator, ok := errorCreators[errType]
	if !ok {
		return fmt.Errorf("error creating error type: %v", errType)
	}
	newFileErr := &FileError{Operation: op, Path: path, Err: errors.New(msg)}
	return creator(newFileErr)
}

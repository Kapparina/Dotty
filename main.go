package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/akamensky/argparse"
)

type ExitCode int

const (
	NoArgsErr ExitCode = iota
	Success
	ScanDirErr
	CloneRepoErr
	CopyFilesErr
)

type ExitHandler struct {
	ExitCode
	error
}

func (e *ExitHandler) Error() string {
	return e.error.Error()
}

func (e *ExitHandler) Exit() {
	switch e.ExitCode {
	case NoArgsErr:
		fmt.Println("No arguments were provided.")
	default:
		if e.error != nil {
			fmt.Println(e.error)
		}
	}
	os.Exit(int(e.ExitCode))
}

const (
	dotFilesRepo string = "https://github.com/Kapparina/dotfiles"
)

type dotFile struct {
	name          string
	repoLocation  string
	localLocation string
}

func getHomePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}

// ScanDir recursively scans the given directory path
func ScanDir(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Print the path of the file
		fmt.Println(path)
		return nil
	})
}

func main() {
	var runtimeTerminator *ExitHandler = &ExitHandler{Success, nil}
	defer runtimeTerminator.Exit()

	parser := argparse.NewParser("gocode", "Go code with argparse")
	stringArg := parser.String("s", "someArg", &argparse.Options{Required: true, Help: "Some argument"})

	if err := parser.Parse(os.Args); err != nil {
		runtimeTerminator.ExitCode = NoArgsErr
		fmt.Println(parser.Usage(err))
		runtimeTerminator.Exit()
	}

	fmt.Println("Input Arg: ", *stringArg)

	homePath, pathErr := getHomePath()
	if pathErr != nil {
		runtimeTerminator.ExitCode = ScanDirErr
		runtimeTerminator.error = pathErr
		return
	}
	var dotFilesPath string = homePath + "/dots"

	scanErr := ScanDir(dotFilesPath)
	if scanErr != nil {
		runtimeTerminator.ExitCode = ScanDirErr
		runtimeTerminator.error = scanErr
		return
	}
}

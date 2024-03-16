package files

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	derr "github.com/Kapparina/Dotty/pkg/errors"
)

func TestWriteToFile(t *testing.T) {
	testCases := []struct {
		name         string
		destination  func() (string, error)
		expectedErr  *derr.FileError
		cleanupAfter bool
	}{
		{
			name: "Valid Destination",
			destination: func() (string, error) {
				tempFile, createErr := os.CreateTemp("", "temp*.txt")
				if createErr != nil {
					return "", createErr
				}
				return tempFile.Name(), nil
			},
			expectedErr:  nil,
			cleanupAfter: true,
		},
		{
			name: "Non-existing Destination",
			destination: func() (string, error) {
				nonexistingPath := filepath.Join("some", "nonexisting", "path", "temp.txt")
				return nonexistingPath, nil
			},
			expectedErr: &derr.FileError{
				Operation: "writing to file",
				Path:      "",
				Err:       errors.New("the system cannot find the path specified"),
			},
			cleanupAfter: false,
		},
		{
			name: "Access Denied Destination",
			destination: func() (string, error) {
				deniedPath := filepath.Join("C:\\Windows\\System32", "temp.txt") // path that requires admin access
				return deniedPath, nil
			},
			expectedErr: &derr.FileError{
				Operation: "writing to file",
				Path:      "",
				Err:       errors.New("Access is denied."),
			},
			cleanupAfter: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			destination, destinationErr := tc.destination()
			if destinationErr != nil {
				t.Fatal(destinationErr)
			}
			fmt.Printf("Testing with destination: %s\n", destination) // Printing the destination
			if tc.expectedErr != nil {
				tc.expectedErr.Path = destination
			}
			data := []byte("test data")
			file, openErr := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
			if openErr != nil {
				t.Fatal(openErr)
			}
			_, writeErr := io.WriteString(file, string(data))
			fmt.Printf("Write error: %v\n", writeErr) // Printing the error if any
			var e *derr.FileError
			if errors.As(writeErr, &e) && tc.expectedErr != nil && e.Operation == tc.expectedErr.Operation && errors.Is(e.Err, tc.expectedErr.Err) {
				return // correct error expected and returned
			}
			if writeErr != nil {
				t.Fatalf("unexpected error: got %v", writeErr)
			}
			if tc.expectedErr != nil {
				t.Fatal("expected error but got none")
			}
			if tc.cleanupAfter {
				os.Remove(destination)
			}
		})
	}
}

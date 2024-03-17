package files

import (
	"errors"
	"os"
	"testing"

	. "github.com/Kapparina/Dotty/pkg/errors"
)

func TestWriteToFile(t *testing.T) {
	tests := []struct {
		name        string
		destination string
		data        []byte
		expectedErr error
	}{
		{
			name:        "ValidPathAndData",
			destination: "test_valid.txt",
			data:        []byte("Hello, World!"),
			expectedErr: nil,
		},
		{
			name:        "InvalidPath",
			destination: "invalid/test.txt",
			data:        []byte("Hello, World!"),
			expectedErr: NewFileError(
				"writing to file",
				"invalid/test.txt",
				Write,
				"open invalid/test.txt: The system cannot find the path specified.",
			),
		},
		{
			name:        "NilData",
			destination: "test_nil.txt",
			data:        nil,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteToFile(tt.destination, tt.data)
			if err != nil && tt.expectedErr == nil {
				t.Errorf("WriteToFile() unexpected error:\nGot: %v\nExpected: nil", err)
			} else if err != nil && tt.expectedErr != nil {
				var gotErr *WriteError
				if errors.As(err, &gotErr) {
					if err.Error() != tt.expectedErr.Error() {
						t.Errorf("WriteToFile() unexpected error:\nGot: %v\nExpected: %v", err, tt.expectedErr)
					}
				}
			}

			// Defer file cleanup after the test
			defer func() {
				if _, err := os.Stat(tt.destination); err == nil || os.IsExist(err) {
					err := os.Remove(tt.destination)
					if err != nil {
						t.Errorf("Error removing the test file: %v", err)
					}
				}
			}()

			// If the destination is valid and data is not nil, check the content of the file
			if tt.name == "ValidPathAndData" {
				if file, err := os.Open(tt.destination); err == nil {
					content := make([]byte, len(tt.data))
					_, err := file.Read(content)
					if err != nil || string(content) != string(tt.data) {
						t.Errorf("Unexpected content in file.\nGot: %s\nWant: %s", string(content), string(tt.data))
					}
					file.Close()
				}
			}

			if tt.name == "NilData" {
				if _, err := os.Stat(tt.destination); err != nil && os.IsNotExist(err) {
					t.Error("Expected file to be created, it was not")
				}
			}

			// Clean up test file if it was created
			if _, err := os.Stat(tt.destination); err == nil || os.IsExist(err) {
				os.Remove(tt.destination)
			}
		})
	}
}

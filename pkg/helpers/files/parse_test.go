package files

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetContents(t *testing.T) {
	// setup
	content := []byte("test file content\n")
	testFile, err := os.CreateTemp("", "temp*.txt")
	if err != nil {
		t.Fatalf("TestGetContents() error creating test file: %v", err)
	}
	_, err = testFile.Write(content)
	if err != nil {
		t.Fatalf("TestGetContents() error writing to test file: %v", err)
	}
	defer func() {
		_ = os.Remove(testFile.Name()) // clean up
	}()

	// define test cases
	testCases := []struct {
		name         string
		filePath     string
		wantData     []byte
		wantErrIsNil bool
	}{
		{
			name:         "file exists",
			filePath:     testFile.Name(),
			wantData:     content,
			wantErrIsNil: true,
		},
		{
			name:         "file does not exist",
			filePath:     "testdata/nonexistent_file.txt",
			wantErrIsNil: false,
		},
	}

	// run the tests
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotErr := GetContents(tt.filePath)
			if (gotErr == nil) != tt.wantErrIsNil {
				t.Errorf("\nGetContents() error = %v\nwantErrIsNil %v", gotErr, tt.wantErrIsNil)
				return
			}
			if !cmp.Equal(gotData, tt.wantData) {
				t.Errorf("\nGetContents() = %v\nwant %v", gotData, tt.wantData)
			}
		})
	}
}

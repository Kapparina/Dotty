package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileError(t *testing.T) {
	assertion := assert.New(t)

	testCases := []struct {
		name     string
		op       string
		path     string
		errType  ErrorType
		msg      string
		expected string
	}{
		{
			name:     "invalid error type",
			op:       "read",
			path:     "/test/path",
			errType:  ErrorType(9999), // assuming 9999 is not a valid ErrorType
			msg:      "unexpected error type",
			expected: "error creating error type: 9999",
		},
		// using a known existing ErrorType with a valid creator
		{
			name:     "valid error type",
			op:       "write",
			path:     "/another/test/path",
			errType:  Write,
			msg:      "unable to write to file",
			expected: "write: /another/test/path: unable to write to file",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newFileErr := NewFileError(testCase.op, testCase.path, testCase.errType, testCase.msg)
			actualErrorMsg := newFileErr.Error()
			assertion.Equal(testCase.expected, actualErrorMsg)
		})
	}
}

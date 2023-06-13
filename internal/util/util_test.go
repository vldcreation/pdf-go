package util_test

import (
	"errors"
	"testing"

	"github.com/vldcreation/privy-pdf-go/internal/util"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		filename string
		ext      string
		err      error
	}{
		{
			name:     "valid input",
			input:    "test.pdf",
			filename: "test",
			ext:      "pdf",
			err:      nil,
		}, {
			name:     "valid input with path",
			input:    "test/test.pdf",
			filename: "test",
			ext:      "pdf",
			err:      nil,
		},
		{
			name:     "invalid input",
			input:    "test",
			filename: "",
			ext:      "",
			err:      errors.New("file input not valid"),
		},
		{
			name:     "invalid input with path",
			input:    "test/test",
			filename: "",
			ext:      "",
			err:      errors.New("file input not valid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename, ext, err := util.ParseFile(tt.input)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("ParseInput() error = %v, wantErr %v", err, tt.err)
				}
				return
			}

			if filename != tt.filename {
				t.Errorf("ParseInput() filename = %v, want %v", filename, tt.filename)
			}

			if ext != tt.ext {
				t.Errorf("ParseInput() ext = %v, want %v", ext, tt.ext)
			}
		})
	}
}

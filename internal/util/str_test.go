package util_test

import (
	"testing"

	"github.com/vldcreation/privy-pdf-go/internal/util"
)

func TestSetUnlockFilename(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "Invalid filename",
			filename: "file1",
			want:     "file1",
		}, {
			name:     "Valid filename",
			filename: "file1.pdf",
			want:     "file1_unlock.pdf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := util.SetUnlockFilename(&tt.filename)
			if err != nil {
				if err != util.ErrFileInputNotValid {
					t.Errorf("SetUnlockFilename() error = %v", err)
					return
				}
			}

			if tt.filename != tt.want {
				t.Errorf("SetUnlockFilename() got = %v, want %v", tt.filename, tt.want)
			}
		})
	}
}

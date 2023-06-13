package util_test

import (
	"testing"

	"github.com/vldcreation/privy-pdf-go/internal/util"
)

func TestGenerateRandoomUUID(t *testing.T) {
	u := util.GenerateRandoomUUID()

	// UUID length is 36
	// 8-4-4-4-12
	// based on https://en.wikipedia.org/wiki/Universally_unique_identifier
	if len(u) != 36 {
		t.Fatalf("UUID length is not 36")
	}
}

package config_test

import (
	"testing"

	"github.com/vldcreation/privy-pdf-go/internal/config"
	"github.com/vldcreation/privy-pdf-go/internal/constants"
)

func TestNewConfig(t *testing.T) {
	c := config.NewConfig(constants.ROOT_FROM_CONFIG)

	if c == nil {
		t.Fatal("Something error")
	}

	if c.UniDocLicenseKey == "" {
		t.Fatal("License Key expected to be not empty")
	}

}

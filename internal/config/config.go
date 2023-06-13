package config

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/vldcreation/privy-pdf-go/internal/util"
)

type Config struct {
	UniDocLicenseKey     string `mapstructure:"UNIDOC_LICENSE_API_KEY"`
	AppPrivyVerifyDomain string `mapstructure:"APP_PRIVY_VERIFY_DOMAIN"`
}

var (
	o sync.Once
	c Config
)

func NewConfig(path string) *Config {
	log.Info().Msg("Initializing config")
	err := util.LoadConfig(path, &c)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while load config")
	}
	log.Info().Msg("Config initialized")

	return &c
}

func (c *Config) GetUniDOcLicenseKey() string {
	return c.UniDocLicenseKey
}

package config

import (
	"time"
)

type ConfigJWTEnv struct {
	JwtSecret        string
	DurationTRefresh time.Duration
}

func (c *ConfigJWTEnv) GetSecretVal() string {
	return c.JwtSecret
}

func (c *ConfigJWTEnv) GetDurationRefreshVal() time.Duration {
	return c.DurationTRefresh
}

func NewJwtConfig(config Config) *ConfigJWTEnv {
	secret := config.Get("JWT_SECRET")
	tokenDuration := 15 * time.Minute
	cfgEnv := ConfigJWTEnv{
		JwtSecret:        secret,
		DurationTRefresh: tokenDuration,
	}
	return &cfgEnv
}

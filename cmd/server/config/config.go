// package config contains the configuration of the burl server.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	dbURL    = "DB_URL"
	httpPort = "HTTP_PORT"
)

// Config represents the configuration of the burl server.
// It uses viper to read the configuration from the environment.
// All the configuration keys are prefixed with "burl".
// For example, the key "DB_URL" is read from the environment as "BURL_DB_URL".
type Config struct {
	viper *viper.Viper
}

// New returns a new Config.
func New() *Config {
	v := viper.New()
	v.SetEnvPrefix("burl")
	v.AutomaticEnv()
	return &Config{
		viper: v,
	}
}

// int returns the value of the key as an integer.
func (c *Config) int(key string) (int, error) {
	if !c.viper.IsSet(key) {
		return 0, fmt.Errorf("%s is not set", key)
	}

	return c.viper.GetInt(key), nil
}

// str returns the value of the key as a string.
func (c *Config) str(key string) (string, error) {
	if !c.viper.IsSet(key) {
		return "", fmt.Errorf("%s is not set", key)
	}

	return c.viper.GetString(key), nil
}

// DBURL returns the database URL.
func (c *Config) DBURL() (string, error) { return c.str(dbURL) }

// HTTPPort returns the HTTP port.
func (c *Config) HTTPPort() (int, error) { return c.int(httpPort) }

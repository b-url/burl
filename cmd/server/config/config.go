// Package config contains the configuration of the burl server.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	EnvironmentPrefix = "burlserver"

	FlagLogLevel    = "log-level"
	FlagLogType     = "log-type"
	FlagDatabaseURL = "db-url"
	FlagHTTPPort    = "http-port"
)

// Config represents the configuration of the burl server.
// It uses viper to read the configuration from the environment.
//
// All the configuration keys are prefixed with "burlserver".
// For example, the key "DB_URL" is read from the environment as "BURLSERVER_DB_URL".
type Config struct {
	viper *viper.Viper
}

// New returns a new Config.
func New() *Config {
	v := viper.New()

	v.SetEnvPrefix(EnvironmentPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	return &Config{
		viper: v,
	}
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable).
func (c *Config) BindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name

		if !f.Changed && c.viper.IsSet(configName) {
			val := c.viper.Get(configName)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
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
func (c *Config) DBURL() (string, error) { return c.str(FlagDatabaseURL) }

// HTTPPort returns the HTTP port.
func (c *Config) HTTPPort() (int, error) { return c.int(FlagHTTPPort) }

// LogLevel returns the log level that was configured for the server.
func (c *Config) LogLevel() (slog.Level, error) {
	ll, err := c.str(FlagLogLevel)
	if err != nil {
		return slog.LevelInfo, err
	}

	var level slog.Level
	err = level.UnmarshalText([]byte(ll))
	if err != nil {
		return slog.LevelInfo, fmt.Errorf("invalid log level '%s': %w", ll, err)
	}

	return level, nil
}

// LogType returns the log type that was configured for the server.
func (c *Config) LogType() (LogType, error) {
	lt, err := c.str(FlagLogType)
	if err != nil {
		return LogTypeText, err
	}

	var logType LogType
	err = logType.UnmarshalText([]byte(lt))
	if err != nil {
		return LogTypeText, fmt.Errorf("invalid log type '%s': %w", lt, err)
	}

	return logType, nil
}

// NewLogger creates a new *slog.Logger based on the configuration.
func (c *Config) NewLogger() *slog.Logger {
	logLevel, err := c.LogLevel()
	if err != nil {
		fmt.Println(err)
	}

	var logType LogType
	logType, err = c.LogType()
	if err != nil {
		fmt.Println(err)
	}

	switch logType {
	case LogTypeText:
		return slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				Level: logLevel,
			}),
		)
	case LogTypeJSON:
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: logLevel,
			}),
		)
	default:
		return nil
	}
}

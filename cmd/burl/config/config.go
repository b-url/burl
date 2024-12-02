package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ConfigFileName    = "config.yaml"
	ConfigFolder      = ".config/burl"
	EnvironmentPrefix = "BURL"

	APIURLName = "api_url"
)

// Config represents the configuration of the burl command.
type Config struct {
	APIURL string `yaml:"api_url"`
	DBURL  string `yaml:"db_url"`
	DBPASS string `yaml:"db_pass"`
}

func New() (Config, error) {
	c := Config{}

	err := viper.GetViper().Unmarshal(&c, viper.DecoderConfigOption(
		func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "yaml"
		}))

	return c, err
}

func (c Config) Write() (string, error) {
	dir := fileDir()
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	path := viper.ConfigFileUsed()
	return path, viper.WriteConfig()
}

func fileDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return filepath.Join(home, ConfigFolder)
}

func Filepath() string {
	configPath := fileDir()
	burlConfigPath := filepath.Join(configPath, ConfigFileName)

	return burlConfigPath
}

func Init() {
	configFile := Filepath()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.SetConfigPermissions(os.FileMode(0600))

	viper.SetEnvPrefix(EnvironmentPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// TODO: use debug logger and make log level configurable.
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

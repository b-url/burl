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
)

// Config represents the configuration of the burl command.
type Config struct {
	APIURL     string `yaml:"api-url"`
	DeviceName string `yaml:"deviceName"`
}

func New() (Config, error) {
	c := Config{}
	err := viper.GetViper().Unmarshal(&c, viper.DecoderConfigOption(func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "yaml"
	}))

	return c, err
}

func Filepath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configPath := filepath.Join(home, ConfigFolder)
	burlConfigPath := filepath.Join(configPath, ConfigFileName)

	return burlConfigPath
}

func Init() {
	fmt.Println("Initializing config")
	configFile := Filepath()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.SetConfigPermissions(os.FileMode(0600))

	viper.SetEnvPrefix(EnvironmentPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/b-url/burl/cmd/burl/config"
)

func TestDefaultConfigFolder(t *testing.T) {
	t.Run("should return the default config folder path", func(t *testing.T) {
		home, err := os.UserHomeDir()
		if err != nil {
			t.Fatalf("os.UserHomeDir() error = %v; want nil", err)
		}

		expected := filepath.Join(home, config.ConfigFolder, config.ConfigFileName)
		got := config.Filepath()

		if got != expected {
			t.Errorf("DefaultConfigFolder() = %s; want %s", got, expected)
		}
	})
}

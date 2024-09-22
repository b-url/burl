package main

import (
	"testing"
)

func TestNewRootCMD(t *testing.T) {
	t.Run("should return a new root command", func(t *testing.T) {
		if NewRootCMD() == nil {
			t.Error("NewRootCMD() = nil; want a new root command")
		}
	})
}

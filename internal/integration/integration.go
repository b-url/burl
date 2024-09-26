// Package integration provides helpers for integration testing.
package integration

import (
	"fmt"
	"os"
	"testing"
)

// RequireIntegration checks if integration tests should run.
// Skips the test if in short mode or "INTEGRATION" env var is unset.
func RequireIntegration(t *testing.T) bool {
	t.Helper()

	if testing.Short() || os.Getenv("INTEGRATION") == "" {
		fmt.Println("Skipping integration test")
		t.Skip()
		return false
	}

	return true
}

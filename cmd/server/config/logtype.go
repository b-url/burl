package config

import (
	"fmt"
	"strings"
)

// LogType determines what kind of output is written by the logger.
type LogType int

// Common log types.
//
//go:generate stringer -type LogType -trimprefix LogType
const (
	LogTypeText LogType = iota
	LogTypeJSON
)

// UnmarshalText implements [encoding.TextUnmarshaler].
func (lt *LogType) UnmarshalText(data []byte) error {
	// Convert the input data to uppercase
	str := strings.ToUpper(string(data))

	switch str {
	case "TEXT":
		*lt = LogTypeText
	case "JSON":
		*lt = LogTypeJSON
	default:
		return fmt.Errorf("invalid LogType: %s", data)
	}

	return nil
}

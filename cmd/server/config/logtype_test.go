package config_test

import (
	"testing"

	"github.com/b-url/burl/cmd/server/config"
)

func TestLogType_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    config.LogType
		wantErr bool
	}{
		{"valid TEXT", []byte("TEXT"), config.LogTypeText, false},
		{"valid text", []byte("text"), config.LogTypeText, false},
		{"valid JSON", []byte("JSON"), config.LogTypeJSON, false},
		{"valid json", []byte("json"), config.LogTypeJSON, false},
		{"invalid type", []byte("xml"), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lt config.LogType
			err := lt.UnmarshalText(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if lt != tt.want {
				t.Errorf("UnmarshalText() = %v, want %v", lt, tt.want)
			}
		})
	}
}

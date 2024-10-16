package otp_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/b-url/burl/cmd/server/login/otp"
)

// MockWriter is a mock implementation of the io.Writer interface.
type MockWriter struct {
	err error
}

// Write writes the token and email to the writer.
func (m *MockWriter) Write(p []byte) (int, error) {
	return len(p), m.err
}

func TestNewWriter(t *testing.T) {
	if otp.NewNotifierWriter(nil) == nil {
		t.Error("NewWriter returned nil")
	}
}

func TestWriter_Notify(t *testing.T) {
	t.Run("writer should write the token and email", func(t *testing.T) {
		w := &strings.Builder{}
		writer := otp.NewNotifierWriter(w)
		d := otp.Notification{Password: "Password", Email: "email"}
		err := writer.Notify(d)
		if err != nil {
			t.Errorf("Writer.Notify() error = %v", err)
		}
		if got := w.String(); got != fmt.Sprintf("Password: %q Email: %q\n", d.Password, d.Email) {
			t.Errorf("Writer.Notify() = %v, want %v", got, fmt.Sprintf("Password: %q Email: %q\n", d.Password, d.Email))
		}
	})

	t.Run("writer should return error if write fails", func(t *testing.T) {
		var expectedErr = errors.New("write error")
		w := &MockWriter{err: expectedErr}
		writer := otp.NewNotifierWriter(w)
		d := otp.Notification{Password: "Password", Email: "email"}
		if err := writer.Notify(d); !errors.Is(err, expectedErr) {
			t.Errorf("Writer.Notify() error = %v, want %v", err, expectedErr)
		}
	})
}

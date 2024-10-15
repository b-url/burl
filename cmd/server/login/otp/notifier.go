package otp

import (
	"fmt"
	"io"
)

// Notification is a struct that contains the token and contact details of the user.
type Notification struct {
	Password string // Password is the password of the user
	Email    string // Email is the email address of the user
}

// Notifier notifies the user of the token.
type Notifier interface {
	Notify(Notification) error
}

// NotifierWriter implements the Notifier interface and writes the token to the writer.
type NotifierWriter struct {
	w io.Writer
}

// NewWriter returns a new Writer.
func NewNotifierWriter(w io.Writer) *NotifierWriter {
	return &NotifierWriter{w: w}
}

func (n *NotifierWriter) Notify(d Notification) error {
	_, err := n.w.Write([]byte(fmt.Sprintf("Password: %q Email: %q\n", d.Password, d.Email)))
	if err != nil {
		return fmt.Errorf("failed to write token: %w", err)
	}
	return err
}

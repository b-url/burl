package otp_test

import (
	"errors"
	"testing"

	"github.com/b-url/burl/cmd/server/login/otp"
)

func TestNewOTP(t *testing.T) {
	if otp.New(nil, nil) == nil {
		t.Error("NewOTP returned nil")
	}
}

func TestOTP_Create(t *testing.T) {
	t.Run("should return error if generator fails", func(t *testing.T) {
		var genErr = errors.New("err")
		o := otp.New(nil, otp.GeneratorFunc(func(length int) (string, error) {
			return "", genErr
		}))
		if err := o.Send(""); !errors.Is(err, genErr) {
			t.Errorf("OTP.Create() error = %v, want %v", err, genErr)
		}
	})

	t.Run("should return error if notifier fails", func(t *testing.T) {
		var notifierErr = errors.New("err")
		w := &MockWriter{err: nil}
		o := otp.New(otp.NewNotifierWriter(w), otp.GeneratorFunc(func(length int) (string, error) {
			return "token", notifierErr
		}))
		if err := o.Send(""); !errors.Is(err, notifierErr) {
			t.Errorf("OTP.Create() error = %v, want %v", err, notifierErr)
		}
	})
}

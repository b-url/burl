package otp_test

import (
	"testing"

	"github.com/b-url/burl/cmd/server/login/otp"
)

func TestGeneratorFunc_Generate(t *testing.T) {
	f := func(_ int) (string, error) {
		return "123456", nil
	}

	g := otp.GeneratorFunc(f)
	got, err := g.Generate(6)
	if err != nil {
		t.Errorf("GeneratorFunc.Generate() error = %v", err)
	}
	if got != "123456" {
		t.Errorf("GeneratorFunc.Generate() = %v, want %v", got, "123456")
	}
}

func TestNewNumericOTPGenerator(t *testing.T) {
	g := otp.NewNumericOTPGenerator()
	if g == nil {
		t.Error("NewNumericOTPGenerator returned nil")
	}
	got, err := g.Generate(6)
	if err != nil {
		t.Errorf("NewNumericOTPGenerator.Generate() error = %v", err)
	}
	if len(got) != 6 {
		t.Errorf("NewNumericOTPGenerator.Generate() = %v, want a length of %v", got, 6)
	}
}

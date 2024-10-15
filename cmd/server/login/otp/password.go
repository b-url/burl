package otp

import (
	"crypto/rand"
	"math/big"
)

// Generator is an interface that generates a random string.
type Generator interface {
	Generate(length int) (string, error)
}

type GeneratorFunc func(length int) (string, error)

func (g GeneratorFunc) Generate(length int) (string, error) {
	return g(length)
}

// generateNumericOTP generates a numeric one-time password of the given length.
func generateNumericOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}

// NewNumericOTPGenerator returns a new numeric OTP generator.
func NewNumericOTPGenerator() Generator {
	return GeneratorFunc(generateNumericOTP)
}

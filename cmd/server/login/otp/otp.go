// Package otp provides the One-Time Password authentication mechanism.
package otp

const passwordLength = 6 // passwordLength is the length of the OTP password.

// OTP is One-Time Password authentication mechanism that is used to authenticate the user.
type OTP struct {
	notifier  Notifier
	generator Generator
}

// New returns a new OTP.
func New(notifier Notifier, generator Generator) *OTP {
	return &OTP{notifier: notifier, generator: generator}
}

func (o *OTP) Send(email string) error {
	token, err := o.generator.Generate(passwordLength)
	if err != nil {
		return err
	}
	// todo: save token to the database
	return o.notifier.Notify(Notification{Password: token, Email: email})
}

package entities

type MailType string

const (
	MailTypeOTP                MailType = "otp"
	MailTypeVerifyLink         MailType = "verifylink"
)
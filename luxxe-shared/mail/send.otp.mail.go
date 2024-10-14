package mail

import (
	"strconv"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func SendOtpMail(email, otp, timeLeft string) {
	mailContext := map[string]interface{}{
		"mailType": entities.MailTypeOTP,
	}
	appURL := config.EnvConfig.FRONTEND_PUBLIC_URL

	SendMail(
		&MailInfoStruct{
			To:      email,
			Subject: "Luxxe Account OTP",
			Text: 	 "Please proceed with secure otp " + otp + ". You have " + timeLeft,
			TemplateID: "e3116348-ffd7-4099-8ee0-3fec7b39ec36",
			TemplateData: map[string]string{
				"otp": 					otp,
				"appURL": 			appURL,
				"current_year": strconv.Itoa(shared.GetCurrentYear()),
			},
		},
		mailContext,
		MailTrapMailClient,
	)
}

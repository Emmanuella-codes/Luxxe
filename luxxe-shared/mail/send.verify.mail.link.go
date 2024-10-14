package mail

import (
	"fmt"
	"strconv"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func createVerifyEmailRoute(userID string) string {
	return fmt.Sprintf("%s/auth/verify-email?userID=%s", config.EnvConfig.BACKEND_ROUTE, userID)
}

func SendVerifyMailLink(email, userID, firstName string) {
	inviteLink := createVerifyEmailRoute(userID)
	mailContext := map[string]interface{}{
		"mailType": entities.MailTypeVerifyLink,
	}
	appURL := config.EnvConfig.FRONTEND_PUBLIC_URL

	SendMail(
		&MailInfoStruct{
			To:      email,
			Subject: "Luxxe Account Email Verification",
			Text: 	 "Click on this link to verify your email:  " + inviteLink,
			TemplateID: "a3d86160-f34f-43e0-a9b4-246aa1f02b10",
			TemplateData: map[string]string{
				"firstName":    firstName,
				"inviteLink":   inviteLink,
				"appURL": 			appURL,
				"current_year": strconv.Itoa(shared.GetCurrentYear()),
			},
		},
		mailContext,
		MailTrapMailClient,
	)
}

package mail

import (
	"fmt"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

type SendPulseEmailRes struct {
	Result    bool   `json:"result"`
	Id        string `json:"id"`
	Message   string `json:"message"`
	ErrorCode int32  `json:"error_code"`
}

func SendMailBySendPulse(mailInfo *MailInfoStruct) bool {
	senderEmail, subject, recipientEmail, body, templateID, templateData := mailInfo.From,
		mailInfo.Subject, mailInfo.To, mailInfo.Text, mailInfo.TemplateID, mailInfo.TemplateData

	if senderEmail == "" {
		senderEmail = config.EnvConfig.MAILDATASENDER
	}

	if subject == "" {
		subject = config.EnvConfig.MAILDATASUBJECT
	}

	if recipientEmail == "" {
		recipientEmail = config.EnvConfig.MAILDATARECIPIENT
	}

	if body == "" {
		body = "Welcome to Weddn"
	}

	mailData := map[string]interface{}{
		"email": map[string]interface{}{
			"subject": subject,
			"template": map[string]interface{}{
				"id":        templateID,
				"variables": templateData,
			},
			"from": map[string]string{
				"name":  senderEmail,
				"email": senderEmail,
			},
			"to": []map[string]string{
				{
					"email": recipientEmail,
					"name":  recipientEmail,
				},
			},
		},
	}

	var success bool
	res, err := shared.POST[map[string]interface{}](
		shared.HttpUtilsReq{
			BaseRoute: config.EnvConfig.SENDPULSE_BASE_URL,
			Ext:       "/smtp/emails",
			Body:      mailData,
			Token:     GlobalSendpulseKeyCache.Data.AccessToken,
		},
	)
	if err != nil {
		fmt.Println("Err: ", err)
	} else {
		fmt.Println("Res: ", res)

		// Check if "result" key exists and is a bool
		if result, ok := res["result"].(bool); ok {
			success = result
		} else {
			// Handle the case where "result" key does not exist or is not a bool
			fmt.Println("Result key not found or is not a bool")
			success = false
		}

		// print error details if they exist
		if message, ok := res["message"].(string); ok {
			fmt.Println("Error message:", message)
		}
		if errorCode, ok := res["error_code"].(float64); ok { // JSON numbers are parsed as float64
			fmt.Println("Error code:", int(errorCode))
		}
	}

	return success
}

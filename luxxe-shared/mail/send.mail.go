package mail

type MailClient string

const SendPulseMailClient MailClient = "sendpulse"

func SendMail(
	mailInfo *MailInfoStruct,
	mailContext map[string]interface{},
	mailClients ...MailClient,
) {
	if len(mailClients) == 0 {
		SendMailBySendPulse(mailInfo)
	}  
}

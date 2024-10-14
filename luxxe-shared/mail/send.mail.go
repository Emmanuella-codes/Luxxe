package mail

type MailClient string

const MailTrapMailClient MailClient = "mailtrap"

func SendMail(
	mailInfo *MailInfoStruct,
	mailContext map[string]interface{},
	mailClients ...MailClient,
) {
	if len(mailClients) == 0 {
		SendMailByMailTrap(mailInfo)
	}  
}

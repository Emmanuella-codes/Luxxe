package mail

func SendMail(
	mailInfo *MailInfoStruct,
	mailContext map[string]interface{},
) {
		SendMailByGmail(mailInfo)
}

package mail

import (
	"fmt"
  "net/smtp"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
)

func SendMailByGmail(mailInfo *MailInfoStruct) {
  // set up auth info
  auth := smtp.PlainAuth("", config.EnvConfig.GMAIL_USERNAME, config.EnvConfig.GMAIL_PASSWORD, "smtp.gmail.com")

  // Create the email message.
	to := []string{mailInfo.To}
	msg := []byte(
		"From: " + mailInfo.From + "\r\n" +
			"To: " + mailInfo.To + "\r\n" +
			"Subject: " + mailInfo.Subject + "\r\n" +
			"\r\n" + mailInfo.Text + "\r\n",
	)

  // Send the email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, mailInfo.From, to, msg)
	if err != nil {
		fmt.Println("error sending email: %v", err)
    panic(err)
	}

	fmt.Println("Email sent successfully!")
}

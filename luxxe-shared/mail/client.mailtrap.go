package mail

import (
	"fmt"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	gomail "gopkg.in/mail.v2"
)

func SendMailByMailTrap(mailInfo *MailInfoStruct) {
  // create a new message
  message := gomail.NewMessage()

  // Set email headers using data from MailInfoStruct
	message.SetHeader("From", mailInfo.From)
	message.SetHeader("To", mailInfo.To)
	message.SetHeader("Subject", mailInfo.Subject)

  // set email body (text content)
  message.SetBody("text/plain", mailInfo.Text)

  // setup the Mailtrap SMTP dialer
  dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, config.EnvConfig.MAILTRAP_USERNAME, config.EnvConfig.MAILTRAP_PASSWORD)

  // send the email
  if err := dialer.DialAndSend(message); err != nil {
    fmt.Println("Error while sending the email:", err)
		panic(err)
  } else {
    fmt.Println("Email sent successfully!")
  }
}

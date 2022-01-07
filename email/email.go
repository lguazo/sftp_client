package email

import (
	"fmt"
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendEmail() {

	var smtpPort int

	emailPort := os.Getenv("SMTP_PORT")
	_, err := fmt.Sscan(emailPort, &smtpPort)

	server := mail.NewSMTPClient()
	server.Host = os.Getenv("SMTP_HOST")
	server.Port = smtpPort
	server.Username = os.Getenv("SMTP_USER")
	server.Password = os.Getenv("SMTP_PASSWORD")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom(os.Getenv("EMAIL_FROM"))
	email.AddTo(os.Getenv("EMAIL_TO"))
	email.AddCc(os.Getenv("EMAIL_CC"))
	email.SetSubject(os.Getenv("EMAIL_SUBJECT"))

	email.SetBody(mail.TextHTML, os.Getenv("EMAIL_BODY"))
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}

package messaging

import (
	"fmt"
	"log"
	"net/smtp"
)



var (
	conf Config = ReadConfig()
)


func SendMessage(subject string, message string) {
	msg := []byte("To: " + conf.To[0] + "\r\n" +
	"From: " + conf.From + "\r\n" +
	"Subject: " + subject + "\r\n" +
	message)

	auth := smtp.PlainAuth("", conf.From, conf.Password, conf.SmtpHost)
	err := smtp.SendMail(conf.SmtpHost+":"+conf.SmtpPort, auth, conf.From, conf.To, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message Sent...")
}
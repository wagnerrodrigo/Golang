package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	// Config Mailtrap
	host := "sandbox.smtp.mailtrap.io"
	port := 587
	user := ""
	password := ""

	msg := gomail.NewMessage()

	msg.SetHeader("From", "teste@gmail.com")
	msg.SetHeader("To", "oi@gmail.com")
	msg.SetHeader("Subject", "Ola Mundo do Email")
	msg.SetBody("Text/html", "<h1>Meu primeiro envio de Email. ola mundo</h1>")

	dialer := gomail.NewDialer(host, port, user, password)

	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Println("nao deu certo")
	} else {
		fmt.Println("AEEEEEEEEEEEEEEEEEEEEE")
	}

}

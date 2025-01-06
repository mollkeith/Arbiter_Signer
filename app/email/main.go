// Copyright (c) 2025 The bel2 developers

package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

var from = "mollkeith.elastos@gmail.com"
var pass = "ziyueyejzh12"

func sendEmail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	return d.DialAndSend(m)
}

func main() {

	to := []string{from}
	subject := "Test Email from GoFrame"
	body := "Hello, this is a test email sent from a GoFrame service!"

	if err := sendEmail(to, subject, body); err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}
}

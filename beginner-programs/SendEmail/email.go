package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "barbershopwale@gmail.com", "harshit65@", "smtp.gmail.com")

	// Email details
	from := "barbershopwale@gmail.com"
	to := "harshitdixit69@gmail.com"
	subject := "Hello there!"
	body := "Hello there my friend!"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587", auth, "barbershopwale@gmail.com", []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	// Create email template variables
	outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}

	// New Outgoing Email information
	newOutgoing := SimpleMailer.Outgoing{
		Email: "info@example.com",
		Subject: "SimpleMailer in Golang",
		Template: "welcome.html",
		Variables: outVars }

	// Send the email
	sendSuccess := SimpleMailer.SendSingle(newOutgoing)

	if sendSuccess {
		fmt.Println("Email was successfully sent to ",newOutgoing.Email)
	} else {
		fmt.Println("The email was not sent")
	}

}
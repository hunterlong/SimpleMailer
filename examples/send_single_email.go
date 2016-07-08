package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "Sender Name", "from@domain.com", "./emails/")

	outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
	newOutgoing := SimpleMailer.Outgoing{
		Email: "info@example.com",
		Subject: "SimpleMailer in Golang",
		Template: "welcome.html",
		Variables: outVars }
	sendSuccess := SimpleMailer.SendSingle(newOutgoing)

	if sendSuccess {
		fmt.Println("Email was successfully sent to ",newOutgoing.Email)
	} else {
		fmt.Println("The email was not sent")
	}

}
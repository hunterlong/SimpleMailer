package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	var newOutgoing []SimpleMailer.Outgoing

	outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"billy", "DIFFICULTY": "simple"}}
	newOutgoingMessage := SimpleMailer.Outgoing{Email: "firstuser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: outVars}
	newOutgoing = append(newOutgoing, newOutgoingMessage)

	secondoutVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gopher", "DIFFICULTY": "super easy"}}
	secondnewOutgoingMessage := SimpleMailer.Outgoing{Email: "seconduser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: secondoutVars}
	newOutgoing = append(newOutgoing, secondnewOutgoingMessage)

	// here is where you send multiple emails from a Outgoing array
	responses := SimpleMailer.SendMultiple(newOutgoing)

	for _,successSend := range responses {
		userEmail := successSend["email"]
		sentStatus := successSend["status"].(bool)

		if sentStatus {
			fmt.Println("Email to "+userEmail.(string)+" was successfully sent")
		} else {
			fmt.Println("ERROR - Email to "+userEmail.(string)+" was not able to send")
		}
	}
}
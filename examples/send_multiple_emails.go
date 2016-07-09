// +build ignore

package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	// New Outgoing Email array
	var newOutgoing []SimpleMailer.Outgoing

	// Create the first email with variables
	outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"billy", "DIFFICULTY": "simple"}}
	newOutgoingMessage := SimpleMailer.Outgoing{Email: "firstuser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: outVars}
	newOutgoing = append(newOutgoing, newOutgoingMessage)

	// Create the second email with variables
	secondoutVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gopher", "DIFFICULTY": "super easy"}}
	secondnewOutgoingMessage := SimpleMailer.Outgoing{Email: "seconduser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: secondoutVars}
	newOutgoing = append(newOutgoing, secondnewOutgoingMessage)

	// Send Multiple Outgoing Emails with variables
	responses := SimpleMailer.SendMultiple(newOutgoing)

	// for each email sent in this process
	for _,successSend := range responses {
		userEmail := successSend["email"]
		// was it successful?
		sentStatus := successSend["status"].(bool)

		if sentStatus {
			fmt.Println("Email to "+userEmail.(string)+" was successfully sent")
		} else {
			fmt.Println("ERROR - Email to "+userEmail.(string)+" was not able to send")
		}
	}
}
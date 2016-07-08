package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	// Create a slice of email addresses
	allEmails := []string{"info@emaildomain.com", "noreply@nodomain.com", "hey@gmail.com"}

	// Setup the subject, template html file, and a slice of email addresses
	bulkSend := SimpleMailer.BulkSend{Emails: allEmails, Subject: "Hello Bulk Sender", Template: "bulk.html"}

	// Send Static HTML Template to all email addresses
	response := SimpleMailer.SendBulkEmails(bulkSend)

	if response!=nil {
		fmt.Println("Multiple Static Emails where sent.")
		fmt.Println(response)
	}


}
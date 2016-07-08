package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	allEmails := []string{"info@emaildomain.com", "noreply@nodomain.com", "hey@gmail.com"}

	bulkSend := SimpleMailer.BulkSend{Emails: allEmails, Subject: "Hello Bulk Sender", Template: "bulk.html"}

	response := SimpleMailer.SendBulkEmails(bulkSend)

	if response!=nil {
		fmt.Println("Multiple Static Emails where sent.")
		fmt.Println(response)
	}


}
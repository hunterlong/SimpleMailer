![alt tag](http://pichoster.net/images/2016/05/30/simplemailer.jpg)

# SimpleMailer    [![Build Status](https://travis-ci.org/hunterlong/SimpleMailer.svg?branch=master)](https://travis-ci.org/Hunterlong/SimpleMailer)  [![Code Climate](https://codeclimate.com/github/hunterlong/SimpleMailer/badges/gpa.svg)](https://codeclimate.com/github/hunterlong/SimpleMailer)  [![Coverage Status](https://coveralls.io/repos/github/Hunterlong/SimpleMailer/badge.svg?branch=master)](https://coveralls.io/github/Hunterlong/SimpleMailer?branch=master)  [![GoDoc](https://godoc.org/github.com/Hunterlong/SimpleMailer?status.svg)](https://godoc.org/github.com/hunterlong/SimpleMailer)
#### KISS principle SMTP emailing solution that allows variables and template HTML emails.

An extremely simple SMTP emailing solution for your Go Language application. This golang package allows you to create a HTML email templates, and have variables inside of it by including {{USERNAME}} to be replaced by the actual username in the array.
Simple Mailer can also send bulk emails without variables, checkout this package and try it out!

## 1. Import the package
```
go get github.com/hunterlong/simplemailer
```
```go
import "github.com/hunterlong/simplemailer"
```

## 2. Create 'emails' folder and create a new file called 'welcome.html'
```
You just got an email from SimpleMailer {{USERNAME}}. I hope you enjoyed how {{DIFFICULTY}} it was to do this.
```
Notice the {{USERNAME}} variable inside the HTML template. You'll be able to insert many of these variables in a simple array.

## 3. Setup SMTP Information
```go
// SMTP host, port, username, password, name, send from address, email directory
SimpleMailer.SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "Sender Name", "from@domain.com", "./emails/")
```
##### Be sure to set the email directory for the last parameter (./emails/) ** Keep slash on end! **

![alt tag](http://pichoster.net/images/2016/05/30/githubbreakerJKAya.jpg)


### Send a Single Email with Variables
```go
outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
newOutgoing := SimpleMailer.Outgoing{
                    Email: "info@domain.com", 
                    Subject: "SimpleMailer in Golang", 
                    Template: "welcome.html", 
                    Variables: outVars }
sendSuccess := SimpleMailer.SendSingle(newOutgoing)
fmt.Println(sendSuccess)
// outputs true or false
```
####### This will replace {{USERNAME}} in your HTML template to 'gophers'

### Send Multiple Emails with Variables
```go
var newOutgoing []SimpleMailer.Outgoing

outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"billy", "DIFFICULTY": "simple"}}
newOutgoingMessage := SimpleMailer.Outgoing{Email: "firstuser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: outVars}
newOutgoing = append(newOutgoing, newOutgoingMessage)

secondoutVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"l33tguy", "DIFFICULTY": "super easy"}}
secondnewOutgoingMessage := SimpleMailer.Outgoing{Email: "firstuser@domain.com", Subject: "Welcome Email", Template: "welcome.html", Variables: secondoutVars}
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
```


### Send Multiple Emails without Variables
```go
allEmails := []string{"info@emaildomain.com", "noreply@nodomain.com", "hey@gmail.com"}
bulkSend := SimpleMailer.BulkSend{Emails: allEmails, Subject: "Hello Bulk Sender", Template: "welcome.html"}

response := SimpleMailer.SendBulkEmails(bulkSend)
```
##### This function will send an email without variables to an array of emails


### Full Example for Sending Single Email
```go
package main

import (
	"github.com/hunterlong/simplemailer"
	"fmt"
)

func main() {

	SimpleMailer.SetSMTPInfo("emailserver-hidden.com", "465", "stmpusername", "passwordhere", "Sender Name", "info@sendfrom.com", "./emails/")

	outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
	newOutgoing := SimpleMailer.Outgoing{
		Email: "info@sendtoemail.com",
		Subject: "SimpleMailer in Golang",
		Template: "welcome.html",
		Variables: outVars }
	sendSuccess := SimpleMailer.SendSingle(newOutgoing)

	fmt.Println(sendSuccess)
	// outputs true if it was sent, false if error

}
```

### Google Gmail SMTP info
```go
SimpleMailer.SetSMTPInfo("smtp.gmail.com", "465", "<gmail email>", "<gmail password>", "Sender Name", "<gmail email>", "./emails/")
```

### Yahoo Mail SMTP info
```go
SimpleMailer.SetSMTPInfo("smtp.mail.yahoo.com", "465", "<yahoo email>", "<yahoo password>", "Sender Name", "<yahoo email>", "./emails/")
```

### Outlook Office 365 SMTP info
```go
SimpleMailer.SetSMTPInfo("smtp.office365.com", "587", "<outlook email>", "<outlook password>", "Sender Name", "<outlook email>", "./emails/")
```


### Awesome HTML Email Templates
https://github.com/leemunroe/responsive-html-email-template

https://github.com/mailgun/transactional-email-templates

https://github.com/mailchimp/email-blueprints

https://github.com/InterNations/antwort

https://github.com/g13nn/Email-Framework

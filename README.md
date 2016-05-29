# SimpleMailer    [![Build Status](https://travis-ci.org/Hunterlong/SimpleMailer.svg?branch=master)](https://travis-ci.org/Hunterlong/SimpleMailer)  [![Code Climate](https://codeclimate.com/github/Hunterlong/SimpleMailer/badges/gpa.svg)](https://codeclimate.com/github/Hunterlong/SimpleMailer)  [![Coverage Status](https://coveralls.io/repos/github/Hunterlong/SimpleMailer/badge.svg?branch=master)](https://coveralls.io/github/Hunterlong/SimpleMailer?branch=master)  [![GoDoc](https://godoc.org/github.com/Hunterlong/SimpleMailer?status.svg)](https://godoc.org/github.com/Hunterlong/SimpleMailer)
#### KISS principle SMTP emailing solution that allows variables and template HTML emails.

An extremely simple SMTP emailing solution for your Go Language application.

## 1. Import the package
```
go get github.com/hunterlong/simplemailer
```
```go
import "github.com/hunterlong/simplemailer"
```

## 2. Setup SMTP Information
```go
// SMTP host, port, username, password, send from address, email directory
SimpleMailer.SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "from@domain.com", "./emails")
```
##### Be sure to set the email directory for the last parameter (./emails)

## 3. Create 'emails' folder and create a new file called 'welcome.html'
```
You just got an email from SimpleMailer {{USERNAME}}. I hope you enjoyed how {{DIFFICULTY}} it was to do this.
```
Notice the {{USERNAME}} variable inside the HTML template. You'll be able to insert many of these variables in a simple array.

### Send a Single Email with Variables
```go
outVars := SimpleMailer.Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
newOutgoing := SimpleMailer.Outgoing{
                    Email: "info@domain.com", 
                    Subject: "SimpleMailer in Golang", 
                    Template: "welcome.html", 
                    Variables: outVars }
sendSuccess := SimpleMailer.SendSingle(newOutgoing)
// outputs true or false
```


### Send a Multiple Emails with Variables
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
		t.Log("Email to "+userEmail.(string)+" was successfully sent")
	} else {
		t.Log("ERROR - Email to "+userEmail.(string)+" was not able to send")
	}
}
```


# SimpleMailer
#### KISS principle SMTP emailing solution that allows variables and template HTML emails.

[![Build Status](https://travis-ci.org/Hunterlong/SimpleMailer.svg?branch=master)](https://travis-ci.org/Hunterlong/SimpleMailer)  [![Code Climate](https://codeclimate.com/github/Hunterlong/SimpleMailer/badges/gpa.svg)](https://codeclimate.com/github/Hunterlong/SimpleMailer)  [![Coverage Status](https://coveralls.io/repos/github/Hunterlong/SimpleMailer/badge.svg?branch=master)](https://coveralls.io/github/Hunterlong/SimpleMailer?branch=master)

An extremely simple SMTP emailing solution for your Go Language application.

## 1. Import the package
```go
import "github.com/hunterlong/simplemailer"
```

## 2. Setup SMTP Information
```go
// SMTP host, port, username, password, send from address, email directory
SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "from@domain.com", "./emails")
```

## 3. Create 'emails' folder and create a new file called 'welcome.html'
```
You just got an email from SimpleMailer {{USERNAME}}. I hope you enjoyed how {{DIFFICULTY}} it was to do this.
```
Notice the {{USERNAME}} variable inside the HTML template. You'll be able to insert many of these variables in a simple array.

### Send a Single Email with Variables
```go
outVars := Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
newOutgoing := Outgoing{
                  Email: "info@domain.com", 
                  Subject: "SimpleMailer in Golang", 
                  Template: "welcome.html", 
                  Variables: outVars }
sendSuccess := SendSingle(newOutgoing)
// outputs true or false
```

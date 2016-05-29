# SimpleMailer
#### KISS principle SMTP emailing solution that allows variables and template HTML emails.

An extremely simple SMTP emailing solution for your Go Language application.

## Import the package
```
import "github.com/hunterlong/simplemailer"
```

## Setup SMTP Information
```
SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "from@domain.com", "./emails")
```

### Send a Single Email with Variables
```
outVars := Variables{map[string]interface{}{"USERNAME":"gophers", "AWESOMEKEY": "j7sN8qo7xb1Xy0"}}
newOutgoing := Outgoing{
                  Email: "info@domain.com", 
                  Subject: "SimpleMailer in Golang", 
                  Template: "welcome.html", 
                  Variables: outVars }
sendSuccess := SendSingle(newOutgoing)
// outputs true or false
```

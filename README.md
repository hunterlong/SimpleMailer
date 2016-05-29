# SimpleMailer
#### KISS principle SMTP emailing solution that allows variables and template HTML emails.

An extremely simple SMTP emailing solution for your Go Language application.

## 1. Import the package
```
import "github.com/hunterlong/simplemailer"
```

## 2. Setup SMTP Information
```
SetSMTPInfo("emailserveraddress.com", "465", "info@domain.com", "passwordhere", "from@domain.com", "./emails")
```

## 3. Create 'emails' folder and create a new file called 'welcome.html'
```
You just got an email from SimpleMailer {{USERNAME}}. I hope you enjoyed how simple it was to do this.
```
Notice the {{USERNAME}} variable inside the HTML template. You'll be able to insert many of these variables in a simple array.

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

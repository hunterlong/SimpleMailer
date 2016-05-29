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

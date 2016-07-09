package SimpleMailer

import (

)

var configs Config

// configs for SMTP login information and email template directory
type Config struct {
	SMTPuser string
	SMTPpass string
	SMTPhost string
	SMTPname string
	SMTPfrom string
	SMTPport string
	EmailsDir string
}

// variables in an array that will be replaced inside template
// USERNAME with text: {{USERNAME}} in the .html template
type Variables struct {
	Inputs map[string]interface{}
}

// Array of email addresses, subject for email, and .html template
type BulkSend struct {
	Emails []string
	Subject string
	Template string
}

// A single email address, subject, .html template and array of Variables
type Outgoing struct {
	Email string
	Subject string
	Template string
	Variables Variables
}

// function to set the SMTP login information and email directory
// be sure to leave forward slash on end of email directory
func SetSMTPInfo(host string, port string, user string, password string, fromName string, fromAddress string, emailsDir string){
	 configs = Config{
		 SMTPhost: host,
		 SMTPuser: user,
		 SMTPpass: password,
		 SMTPname: fromName,
		 SMTPfrom: fromAddress,
		 SMTPport: port,
		 EmailsDir: emailsDir }
}


// function to send a single email from Outgoing struct
func SendSingle(outgoing Outgoing) bool {
	chk := SendEmail(outgoing)
	return chk
}

// function to send multiple emails without any variables.
// good for sending a flat html email to many. uses BulkSend struct
func SendBulkEmails(outgoing BulkSend) []map[string]interface{} {
	var response []map[string]interface{}
	for _,sendEmail := range outgoing.Emails {
		sendSingle := Outgoing{Email: sendEmail, Subject: outgoing.Subject, Template: outgoing.Template, Variables: Variables{}}
		success := SendEmail(sendSingle)
		thisResponse := map[string]interface{}{"email": sendEmail, "status": success}
		response = append(response, thisResponse)
	}
	return response
}

// function to send multiple emails with variables.
// uses and array of Outgoing structs
func SendMultiple(outgoing []Outgoing) []map[string]interface{} {

	var response []map[string]interface{}
	for _,sendEmail := range outgoing {
		success := SendEmail(sendEmail)
		thisResponse := map[string]interface{}{"email": sendEmail.Email, "status": success}
		response = append(response, thisResponse)
	}
	return response
}

// simple check error panic function
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
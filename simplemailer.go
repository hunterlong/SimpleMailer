package SimpleMailer

import (

)

var configs Config

type Config struct {
	SMTPuser string
	SMTPpass string
	SMTPhost string
	SMTPfrom string
	SMTPport string
	EmailsDir string
}

type Variables struct {
	Inputs map[string]interface{}
}

type BulkSend struct {
	Emails []string
	Subject string
	Template string
}

type Outgoing struct {
	Email string
	Subject string
	Template string
	Variables Variables
}

 func SetSMTPInfo(host string, port string, user string, password string, fromAddress string, emailsDir string){
	 configs = Config{
		 SMTPhost: host,
		 SMTPuser: user,
		 SMTPpass: password,
		 SMTPfrom: fromAddress,
		 SMTPport: port,
		 EmailsDir: emailsDir }
 }



func SendSingle(outgoing Outgoing) bool {
	chk := SendEmail(outgoing)
	return chk
}


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


func SendMultiple(outgoing []Outgoing) []map[string]interface{} {

	var response []map[string]interface{}
	for _,sendEmail := range outgoing {
		success := SendEmail(sendEmail)
		thisResponse := map[string]interface{}{"email": sendEmail.Email, "status": success}
		response = append(response, thisResponse)
	}
	return response
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
package SimpleMailer

import (
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"crypto/tls"
	"io/ioutil"
	"strings"
)


// function to send an email with TLS and PlainAuth
// this function is ran multiple times in some scripts
func SendEmail(outgoingEmail Outgoing) bool {
		email_html, err := ioutil.ReadFile(configs.EmailsDir + outgoingEmail.Template)
		from := mail.Address{"", configs.SMTPfrom}
		to := mail.Address{"", outgoingEmail.Email}
		subj := outgoingEmail.Subject
		body := string(email_html)
		headers := make(map[string]string)
		headers["From"] = from.String()
		headers["To"] = to.String()
		headers["Subject"] = subj
		message := ""
		for k, v := range headers {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		newbody := ReplaceContentText(outgoingEmail.Variables.Inputs, body)
		message += newbody
		servername := configs.SMTPhost+":"+configs.SMTPport
		host, _, _ := net.SplitHostPort(servername)
		auth := smtp.PlainAuth("", configs.SMTPuser, configs.SMTPpass, host)
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName: host,
		}
		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			return false
			log.Panic(err)
		}
		c, err := smtp.NewClient(conn, host)
		if err != nil {
			return false
			log.Panic(err)
		}
		if err = c.Auth(auth); err != nil {
			return false
			log.Panic(err)
		}
		if err = c.Mail(from.Address); err != nil {
			return false
			log.Panic(err)
		}
		if err = c.Rcpt(to.Address); err != nil {
			return false
			log.Panic(err)
		}
		w, err := c.Data()
		if err != nil {
			return false
			log.Panic(err)
		}
		_, err = w.Write([]byte(message))
		if err != nil {
			return false
			log.Panic(err)
		}
		err = w.Close()
		if err != nil {
			return false
			log.Panic(err)
		}
		c.Quit()
	return true
}


// function to replace variables as: {{USERNAME}} to USERNAME
// inside the html template. input as array of
// {"USERNAME": "gopher"}
func ReplaceContentText(array map[string]interface{}, content string) string {
	var newmeailstring string
	newmeailstring = content
	for k, v := range array {
		toreplace := "{{"+k+"}}"
		newmeailstring = strings.Replace(newmeailstring, toreplace, v.(string), -1)
	}
	return newmeailstring
}

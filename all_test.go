package SimpleMailer


import (
	"testing"
	"os"
)

func TestSetSMTPInfo(t *testing.T){
	SetSMTPInfo(os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), "SimpleMailer", os.Getenv("from"), "./emails/")
	t.Log("Set SMTP information and Email directory \n")
}

func TestSendSingleEmail(t *testing.T) {
	outVars := Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
	newOutgoing := Outgoing{Email: "info@socialeck.com", Subject: "Welcome Email", Template: "welcome.html", Variables: outVars}
	success := SendSingle(newOutgoing)
	if success {
		t.Log("Sending a single email! \n")
	} else {
		t.Log("There was an issue sending that email! \n")
		t.Fail()
	}
}

func TestMultipleEmails(t *testing.T){

	var newOutgoing []Outgoing

	outVars := Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "simple"}}
	newOutgoingMessage := Outgoing{Email: "info@socialeck.com", Subject: "Welcome Email", Template: "welcome.html", Variables: outVars}
	newOutgoing = append(newOutgoing, newOutgoingMessage)

	secondoutVars := Variables{map[string]interface{}{"USERNAME":"gophers", "DIFFICULTY": "super easy"}}
	secondnewOutgoingMessage := Outgoing{Email: "djzebular@gmail.com", Subject: "Welcome Email", Template: "welcome.html", Variables: secondoutVars}
	newOutgoing = append(newOutgoing, secondnewOutgoingMessage)

	responses := SendMultiple(newOutgoing)

	for _,successSend := range responses {
		userEmail := successSend["email"]
		sentStatus := successSend["status"].(bool)

		if sentStatus {
			t.Log("Email to "+userEmail.(string)+" was successfully sent")
		} else {
			t.Log("ERROR - Email to "+userEmail.(string)+" was not able to send")
		}
	}

}


func TestSendMultipleNoVars(t *testing.T){
	allEmails := []string{"info@socialeck.com", "djzebular@gmail.com", "hey@gmail.com"}
	output := BulkSend{Emails: allEmails, Subject: "Hello Bulk Sender", Template: "bulk.html"}

	response := SendBulkEmails(output)

	for _,email := range response {
		t.Log(email)
	}

}


func TestCheckErrors(t *testing.T){
	error := ""
	checkErr(error)
}
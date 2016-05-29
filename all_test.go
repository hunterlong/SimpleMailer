package simplemailer


import (
	"testing"
	"os"
)

func TestSetSMTPInfo(t *testing.T){
	SetSMTPInfo(os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("from"), "./emails")
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

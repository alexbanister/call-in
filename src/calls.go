package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/spf13/viper"
)

type twilioRequest struct {
	endpoint string
	method   string
	values   map[string]string
}

func callData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}
	transcription := new(Transcription)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err = decoder.Decode(transcription, r.Form)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", transcription)
	params := twilioRequest{
		endpoint: "/Messages.json",
		method:   "POST",
		values: map[string]string{
			"To":   r.URL.Query().Get("sms"),
			"From": viper.GetString("numbers.from"),
			"Body": transcription.TranscriptionText,
		},
	}
	sendSMS(params)
	w.WriteHeader(204)
	return
}

func twiml(w http.ResponseWriter, r *http.Request) {
	host := viper.GetString("host")
	twiml := TwiML{
		Pause: Pause{
			Length: 5,
		},
		Play: Play{
			Digits: viper.GetString("numbers.digits"),
		},
		Record: Record{
			Action:             host + "stub",
			Timeout:            5,
			MaxLength:          10,
			TranscribeCallback: host + "callData?sms=" + r.URL.Query().Get("sms"),
			PlayBeep:           true,
		},
	}
	log.Printf("[TWIML] %+v", twiml)
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func call(params twilioRequest) {
	//get user details
	//make call
	response, err := twilioHTTPRequest(params)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(response)
	//log call & id

	return
}

func startCall() {
	params := twilioRequest{
		endpoint: "/Calls.json",
		method:   "POST",
		values: map[string]string{
			"To":   viper.GetString("numbers.call"),
			"From": viper.GetString("numbers.from"),
			"Url":  viper.GetString("host") + "twiml?sms=" + viper.GetString("numbers.sms"),
		},
	}
	// Create Client
	call(params)
	return
}

func testCall(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key != viper.GetString("test.key") {
		fmt.Println(time.Now().Format("[2006-06-05 15:13:11]"), "Unauthorized request to /test")
		http.Error(w, "Unauthorized Request", http.StatusUnauthorized)
		return
	}

	params := twilioRequest{
		endpoint: "/Calls.json",
		method:   "POST",
		values: map[string]string{
			"To":   viper.GetString("test.call"),
			"From": viper.GetString("numbers.from"),
			"Url":  viper.GetString("host") + "twiml?sms=" + r.URL.Query().Get("sms"),
		},
	}
	// Create Client
	call(params)
	fmt.Fprintf(w, "Thank You")
}

func incomingCall(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{
		Pause: Pause{
			Length: 1,
		},
		Say: Say{
			Voice: "woman",
			Text:  "We're sorry, this number does not receive incoming calls",
		},
	}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

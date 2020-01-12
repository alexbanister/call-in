package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

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
			"To":   "8177165097",
			"From": "9405399177",
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
			Digits: "1w86475259wwwwwwwwwwwwwwwwwwwwww1",
		},
		Record: Record{
			Action:             host + "stub",
			Timeout:            5,
			MaxLength:          10,
			TranscribeCallback: host + "callData",
			PlayBeep:           true,
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

func call(params twilioRequest) {
	//get user details
	//make call
	response, err := twilioHTTPRequest(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	//log call & id

	return
}

func testCall(w http.ResponseWriter, r *http.Request) {
	params := twilioRequest{
		endpoint: "/Calls.json",
		method:   "POST",
		values: map[string]string{
			"To":   "7202218948",
			"From": "9405399177",
			"Url":  viper.GetString("host") + "twiml",
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

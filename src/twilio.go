package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause,omitempty"`
	Say     Say      `xml:"Say,omitempty"`
	Play    Play     `xml:"Play,omitempty"`
	Record  Record   `xml:"Record,omitempty"`
	Hangup  string   `xml:"Hangup"`
}

type Pause struct {
	Length int `xml:"length,attr,omitempty"`
}

type Say struct {
	Voice string `xml:"voice,attr,omitempty"`
	Text  string `xml:",innerxml"`
}

type Play struct {
	XMLName xml.Name `xml:"Play"`
	Digits  string   `xml:"digits,attr,omitempty"`
	Say     string   `xml:",innerxml"`
}

type Record struct {
	Action             string `xml:"action,attr,omitempty"`
	Timeout            int    `xml:"timeout,attr,omitempty"`
	MaxLength          int    `xml:"maxLength,attr,omitempty"`
	TranscribeCallback string `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool   `xml:"playBeep,attr,omitempty"`
}

type Transcription struct {
	TranscriptionSid    string
	TranscriptionText   string
	TranscriptionStatus string
	TranscriptionUrl    string
	RecordingSid        string
	RecordingUrl        string
	CallSid             string
	From                string
}

func twilioHTTPRequest(params twilioRequest) (map[string]interface{}, error) {
	accountSid := viper.GetString("twilio.prod.accountSid")
	authToken := viper.GetString("twilio.prod.authToken")

	// Let's set some initial default variables
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + params.endpoint
	v := url.Values{}
	for key, value := range params.values {
		v.Set(key, value)
	}
	rb := *strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	// make request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var data map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		e := errors.New(resp.Status)
		return nil, e
	}
	if er := json.Unmarshal(bodyBytes, &data); er != nil {
		fmt.Println(er)
		return nil, er
	}
	return data, nil
}

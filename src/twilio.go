package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}

func twiml(w http.ResponseWriter, r *http.Request) {
	twiml := TwiML{Say: "Hello World!"}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func call(w http.ResponseWriter, r *http.Request) {
	accountSid := viper.GetString("twilio.prod.accountSid")
	authToken := viper.GetString("twilio.prod.authToken")
	host := viper.GetString("host")

	// Let's set some initial default variables
	hostURL := host + "twiml"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Calls.json"
	v := url.Values{}
	v.Set("To", "3035123030")
	v.Set("From", "9405399177")
	v.Set("Url", hostURL)
	v.Set("sendDigits", "ww1234ww545")
	rb := *strings.NewReader(v.Encode())
	fmt.Println("POST ", urlStr)
	fmt.Println("URL ", hostURL)

	// Create Client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	resp, _ := client.Do(req)

	var data map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
		fmt.Println(string(bodyBytes))
		w.Write([]byte("Go Royals!"))
	}
}

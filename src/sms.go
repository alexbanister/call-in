package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func sendSMS(params twilioRequest) {
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

func testSMS() {
	body := "[TEST] The time is " + time.Now().Format("2006-05-06 15:15")
	params := twilioRequest{
		endpoint: "/Messages.json",
		method:   "POST",
		values: map[string]string{
			"To":   viper.GetString("test.call"),
			"From": viper.GetString("numbers.from"),
			"Body": body,
		},
	}
	fmt.Printf("TEXT %+v", params)
	sendSMS(params)
}

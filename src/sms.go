package main

import (
	"fmt"
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

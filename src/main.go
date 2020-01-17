package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	host = viper.GetString("host")
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.BindEnv("host", "HOST_URL")
	viper.BindEnv("bind-port", "PORT")
	viper.BindEnv("test.key", "TEST_KEY")
	viper.BindEnv("test.call", "TEST_CALL")

	viper.BindEnv("twilio.prod.accountSid", "TWILIO_ACCOUNT_ID")
	viper.BindEnv("twilio.prod.authToken", "TWILIO_TOKEN")

	viper.BindEnv("numbers.from", "NUMBERS_FROM")
	viper.BindEnv("numbers.to", "NUMBERS_TO")
	viper.BindEnv("numbers.sms", "NUMBERS_SMS")
}

func main() {
	http.HandleFunc("/test", testCall)
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/callData", callData)
	http.HandleFunc("/incoming", incomingCall)
	http.HandleFunc("/stub", stub)

	port := viper.GetString("bind-port")
	fmt.Println("Starting Server on port " + port + "...")
	log.Fatal(http.ListenAndServe(port, nil))
}

func stub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank You")
	return
}

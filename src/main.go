package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
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
	location, err := time.LoadLocation("America/Denver")
	if err != nil {
		log.Println("Unfortunately can't load a location")
		log.Println(err)
	} else {
		gocron.ChangeLoc(location)
	}
}

func scheduler() {
	gocron.Every(1).Day().At("6:45").Do(startCall)
	gocron.Every(1).Day().At("22:27").Do(task)
	<-gocron.Start()
}
func main() {

	viper.BindEnv("host", "HOST_URL")
	viper.BindEnv("test.key", "TEST_KEY")
	viper.BindEnv("test.call", "TEST_CALL")

	viper.BindEnv("twilio.prod.accountSid", "TWILIO_ACCOUNT_ID")
	viper.BindEnv("twilio.prod.authToken", "TWILIO_TOKEN")

	viper.BindEnv("numbers.from", "NUMBERS_FROM")
	viper.BindEnv("numbers.to", "NUMBERS_TO")
	viper.BindEnv("numbers.sms", "NUMBERS_SMS")

	port := os.Getenv("PORT")

	if port == "" {
		port = viper.GetString("bind-port")
	}

	go scheduler()

	http.HandleFunc("/test", testCall)
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/callData", callData)
	http.HandleFunc("/incoming", incomingCall)
	http.HandleFunc("/stub", stub)

	fmt.Println("HOST:::: ", viper.GetString("host")+":"+port)

	fmt.Println("Starting Server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func stub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank You")
	return
}

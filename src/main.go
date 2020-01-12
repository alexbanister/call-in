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
}

func main() {
	http.HandleFunc("/test", testCall)
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/callData", callData)
	http.HandleFunc("/incoming", incomingCall)
	http.HandleFunc("/stub", stub)

	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func stub(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank You")
	return
}

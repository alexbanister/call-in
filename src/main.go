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

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/call", call)
	http.HandleFunc("/twiml", twiml)

	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

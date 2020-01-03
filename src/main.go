package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/call", nil)

	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

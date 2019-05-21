package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

var recipients = []string{"+61437200015"}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	client := messagebird.New(os.Getenv("MESSAGEBIRD_KEY"))
	msg, err := sms.Create(
		client,
		os.Getenv("SMS_SENDER"),
		recipients,
		"Hello World, I am a text message and I was hatched by Go code!",
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// You can log the msg variable for development, or discard it by assigning it to '_'
	log.Println(msg)
}

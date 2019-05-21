package main

import (
	"log"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

var sender = "+61437200015"
var recipients = []string{"+61437200015"}

func main() {
	client := messagebird.New("8uF0snGf07eB4cj6ZnbjI6Mzw")
	msg, err := sms.Create(
		client,
		sender,
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

package main

import (
	"log"

	"github.com/AleksandrAkhapkin/go-fcm"
)

func main() {
	// Create the message to be sent.
	msg := &fcm.Message{
		To: "sample_device_token",
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "title",
			Body:  "body",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("sample_api_key")
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
}

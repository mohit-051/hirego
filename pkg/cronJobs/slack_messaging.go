package cronjobs

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendSlackMessage(webhook_url string, message string) {

	// Send a slack message
	fmt.Println("Sending slack message")
    
	// Define the message to be sent.
	body := []byte(`{"text": "` + message + `"}`)
    
	// Send the message to the slack channel.
	response, err := http.Post(webhook_url, "application/json", bytes.NewBuffer(body))

	// Check if there is an error.
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

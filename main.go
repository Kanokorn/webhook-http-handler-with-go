package main

import (
	"log"
	"net/http"

	"github.com/omise/omise-go"
)

func ProcessWebhook(w http.ResponseWriter, r *http.Request, event *omise.Event) {
	switch event.Key {
	case "charge.create":
		charge := event.Data.(*omise.Charge)
		log.Printf("ID: %s", charge.ID)
		log.Printf("Amount: %d", charge.Amount)
	case "charge.complete":
		charge := event.Data.(*omise.Charge)
		log.Printf("ID: %s", charge.ID)
		log.Printf("Amount: %d", charge.Amount)
	default:
		// Do sometinng
	}
}

func main() {
	http.Handle("/", omise.WebhookHTTPHandler(omise.EventHandlerFunc(ProcessWebhook)))
	http.ListenAndServe(":3001", nil)
}

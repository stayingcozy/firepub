package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func addStream(ctx context.Context, client *firestore.Client, newDoc string, url string, user string, pass string, serverIP string) error {
	_, err := client.Collection("streams").Doc(newDoc).Set(ctx, map[string]interface{}{
			"url":    url,
			"user":   user,
			"pass":   pass,
			"server": serverIP,
			"status": "readyToStream",
	})
	if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
	}

	return err
}
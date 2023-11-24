package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func updateAsStreaming(ctx context.Context, client *firestore.Client, newDoc string) error {
	_, err := client.Collection("streams").Doc(newDoc).Update(ctx, []firestore.Update{
		{
			Path:  "status",
			Value: "streaming",
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error while updating has occurred: %s", err)
	}

	return err
}


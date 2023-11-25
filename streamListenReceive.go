package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

func streamListenReceive(ctx context.Context, client *firestore.Client, newDoc string) (bool, error) {

	maxAttempts := 10
	attempts := 1

	for {
		dsnap, err := client.Collection("streams").Doc(newDoc).Get(ctx)
		if err != nil {
			return true, err
		}

		data := dsnap.Data()
		if data["status"] == "readyToReceive" {
			return false, nil
		}

		if attempts >= maxAttempts {
			return true, nil
		}

		// if ready to receive doesn't come for 5 sec, write as broken and rerun ?
		attempts++
		time.Sleep(1 * time.Second)
	}
}

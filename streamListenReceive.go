package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

func streamListenReceive(ctx context.Context, client *firestore.Client, newDoc string) (bool, error) {

	var err error
	var doc *firestore.DocumentSnapshot

	// set timer duration
	duration := 10 * time.Second
	timer := time.NewTimer(duration)

	iter := client.Collection("streams").Where("status", "==", "readyToReceive").Documents(ctx)
	for {
		select {
		case <-timer.C:
			return true, nil
		default:
			doc, err = iter.Next()
			if err != nil {
				return true, err
			}
			streamDocName := doc.Ref.ID

			if streamDocName == newDoc {
				return false, nil
			}
		}
	}
}

	//

	// maxAttempts := 10
	// attempts := 1

	// for {
	// 	dsnap, err := client.Collection("streams").Doc(newDoc).Get(ctx)
	// 	if err != nil {
	// 		// Handle any errors in an appropriate way, such as returning them.
	// 		log.Printf("An error has occurred: %s", err)
	// 		return true, err
	// 	}

	// 	data := dsnap.Data()
	// 	if data["status"] == "readyToReceive" {
	// 		return false, nil
	// 	}

	// 	if attempts >= maxAttempts {
	// 		return true, nil
	// 	}

	// 	// if ready to receive doesn't come for 5 sec, write as broken and rerun ?
	// 	attempts++
	// 	time.Sleep(500 * time.Millisecond)
	// }

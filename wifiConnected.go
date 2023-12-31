package main

import(
	"math/rand"
	"time"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func wifiConnected(ctx context.Context, client *firestore.Client, uid string) {
	//// Update account with wifi connected ////
	// assumption - wifi checks from runPetCamera allow this code to run + wanted to keep firebase code in one spot
	// Initialize the random number generator with a seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(1000001)
	_, errU := client.Collection("users").Doc(uid).Update(ctx, []firestore.Update{
		{
			Path:  "wifiStatus",
			Value: randomNumber,
		},
	})
	if errU != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error while updating has occurred: %s", errU)
	}

}